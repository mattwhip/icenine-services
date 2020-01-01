package dailybonus

import (
	"context"
	"fmt"
	"sync"
	"time"

	core "github.com/mattwhip/icenine-services"
	cfg "github.com/mattwhip/icenine-services/config"
	pb "github.com/mattwhip/icenine-services/generated/services/daily_bonus"
	"github.com/pkg/errors"
	"github.com/vitessio/vitess/go/pools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// Mediator provides an interface to interact with the Daily Bonus service
type Mediator struct {
	poolMutex  *sync.Mutex
	clientPool *pools.ResourcePool
	config     *cfg.Config
}

// ResourceConn wraps a grpc client connection for use in a resource pool
type ResourceConn struct {
	*grpc.ClientConn
	pb.DailyBonusClient
}

// Close a Vitess Resource connection
func (r *ResourceConn) Close() {
	r.ClientConn.Close()
}

// mutex protect the following variables
var defaultMediatorMutex *sync.Mutex
var defaultMediator *Mediator

func init() {
	defaultMediatorMutex = &sync.Mutex{}
}

// Get returns the default core service library UserData mediator
func Get() *Mediator {
	defaultMediatorMutex.Lock()
	defer defaultMediatorMutex.Unlock()
	if defaultMediator == nil {
		defaultMediator = New()
		core.RegisterCleanup(func() error {
			err := defaultMediator.Close()
			defaultMediator = nil
			return err
		})
	}
	return defaultMediator
}

// New returns a new Mediator instance
func New() *Mediator {
	config := cfg.New()
	address := fmt.Sprintf("%s:%s", config.RPCHost, config.RPCPort)
	pool := pools.NewResourcePool(func() (pools.Resource, error) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		client := pb.NewDailyBonusClient(conn)
		return &ResourceConn{
			conn,
			client,
		}, err
	}, config.RPCConnectionPool.InitialCapacity, config.RPCConnectionPool.MaxCapacity, config.RPCConnectionPool.IdleTimeout)
	return &Mediator{
		clientPool: pool,
		config:     &config,
		poolMutex:  &sync.Mutex{},
	}
}

// Close cleans up resources used by the Mediator instance
func (m *Mediator) Close() error {
	m.poolMutex.Lock()
	defer m.poolMutex.Unlock()
	if m.clientPool != nil {
		m.clientPool.Close()
		m.clientPool = nil
	}
	return nil
}

func (m *Mediator) getResourceConn(ctx context.Context) (*ResourceConn, error) {
	resourceConn, err := m.clientPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	rc := resourceConn.(*ResourceConn)
	if rc.ClientConn.GetState() == connectivity.TransientFailure || rc.ClientConn.GetState() == connectivity.Shutdown {
		// Close existing broken connection and attempt to generate a new one
		rc.ClientConn.Close()
		address := fmt.Sprintf("%s:%s", m.config.RPCHost, m.config.RPCPort)
		conn, dialErr := grpc.Dial(address, grpc.WithInsecure())
		if dialErr != nil {
			m.clientPool.Put(rc)
			return nil, errors.Wrap(dialErr, "failed to dial new grpc comnection")
		}
		client := pb.NewDailyBonusClient(conn)
		rc.ClientConn = conn
		rc.DailyBonusClient = client
	}
	return rc, nil
}

func (m *Mediator) putResourceConn(rc *ResourceConn) {
	if m.clientPool == nil {
		return
	}
	m.clientPool.Put(rc)
}

// InitNewUser initializes a new user in the Daily Bonus system
func (m *Mediator) InitNewUser(userID string) (*Status, error) {
	// Get client connection from resource pool
	resourceConn, err := m.getResourceConn(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get grpc client connection from pool")
	}
	defer m.putResourceConn(resourceConn)
	// Execute grpc request
	ctx, cancel := m.newContext()
	defer cancel()
	req := &pb.UserRequest{
		UserID: userID,
	}
	resp, err := resourceConn.DailyBonusClient.InitNewUser(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute grpc InitNewUser to daily bonus service")
	}
	// Craft response
	return &Status{
		IsAvailable:           resp.Status.IsAvailable,
		SecondsUntilAvailable: resp.Status.SecondsUntilAvailable,
		Streak:                resp.Status.Streak,
		WheelValues:           resp.Wheel.Values,
	}, nil
}

// GetStatus retrieves the current status
func (m *Mediator) GetStatus(userID string) (*Status, error) {
	// Get client connection from resource pool
	resourceConn, err := m.getResourceConn(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get grpc client connection from pool")
	}
	defer m.putResourceConn(resourceConn)
	// Execute grpc request
	ctx, cancel := m.newContext()
	defer cancel()
	req := &pb.UserRequest{
		UserID: userID,
	}
	resp, err := resourceConn.DailyBonusClient.GetStatus(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute grpc GetStatus to daily bonus service")
	}
	// Craft response
	return &Status{
		IsAvailable:           resp.Status.IsAvailable,
		SecondsUntilAvailable: resp.Status.SecondsUntilAvailable,
		Streak:                resp.Status.Streak,
		WheelValues:           resp.Wheel.Values,
	}, nil
}

func (m *Mediator) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(m.config.RPCTimeoutSeconds)*time.Second)
}
