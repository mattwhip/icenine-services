package userdata

import (
	"context"
	"fmt"
	"sync"
	"time"

	core "github.com/mattwhip/icenine-services"
	cfg "github.com/mattwhip/icenine-services/config"
	pbud "github.com/mattwhip/icenine-services/generated/services/user_data"
	"github.com/pkg/errors"
	"github.com/vitessio/vitess/go/pools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// UserData provides access to User Data
type UserData struct {
	poolMutex  *sync.Mutex
	clientPool *pools.ResourcePool
	config     *cfg.Config
}

// ResourceConn wraps a grpc client connection for use in a resource pool
type ResourceConn struct {
	*grpc.ClientConn
	pbud.UserDataClient
}

// Close a Vitess Resource connection
func (r *ResourceConn) Close() {
	r.ClientConn.Close()
}

// mutex protext the following variables
var defaultUserDataMutex *sync.Mutex
var defaultUserData *UserData

func init() {
	defaultUserDataMutex = &sync.Mutex{}
}

// Get returns the default core service library UserData
func Get() *UserData {
	if defaultUserData == nil {
		defaultUserDataMutex.Lock()
		defer defaultUserDataMutex.Unlock()
		if defaultUserData == nil {
			defaultUserData = New()
			core.RegisterCleanup(func() error {
				err := defaultUserData.Close()
				defaultUserData = nil
				return err
			})
		}
	}
	return defaultUserData
}

// New returns a new UserData instance
func New() *UserData {
	config := cfg.New()
	address := fmt.Sprintf("%s:%s", config.RPCHost, config.RPCPort)
	pool := pools.NewResourcePool(func() (pools.Resource, error) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		client := pbud.NewUserDataClient(conn)
		return &ResourceConn{
			conn,
			client,
		}, err
	}, config.RPCConnectionPool.InitialCapacity, config.RPCConnectionPool.MaxCapacity, config.RPCConnectionPool.IdleTimeout)
	return &UserData{
		clientPool: pool,
		config:     &config,
		poolMutex:  &sync.Mutex{},
	}
}

// Close cleans up resources used by the UserData instance
func (u *UserData) Close() error {
	u.poolMutex.Lock()
	defer u.poolMutex.Unlock()
	if u.clientPool != nil {
		u.clientPool.Close()
		u.clientPool = nil
	}
	return nil
}

func (u *UserData) getResourceConn(ctx context.Context) (*ResourceConn, error) {
	resourceConn, err := u.clientPool.Get(ctx)
	if err != nil {
		return nil, err
	}
	rc := resourceConn.(*ResourceConn)
	if rc.ClientConn.GetState() == connectivity.TransientFailure || rc.ClientConn.GetState() == connectivity.Shutdown {
		// Close existing broken connection and attempt to generate a new one
		rc.ClientConn.Close()
		address := fmt.Sprintf("%s:%s", u.config.RPCHost, u.config.RPCPort)
		conn, dialErr := grpc.Dial(address, grpc.WithInsecure())
		if dialErr != nil {
			u.clientPool.Put(rc)
			return nil, errors.Wrap(dialErr, "failed to dial new grpc comnection")
		}
		client := pbud.NewUserDataClient(conn)
		rc.ClientConn = conn
		rc.UserDataClient = client
	}
	return rc, nil
}

func (u *UserData) putResourceConn(rc *ResourceConn) {
	if u.clientPool == nil {
		return
	}
	u.clientPool.Put(rc)
}

// DoCoinTransaction executes a coin transaction for one or more users. Takes a map of userID to desired coin delta as
// input and returns a map of userID to final balance.
func (u *UserData) DoCoinTransaction(transaction map[string]int64) (map[string]int64, error) {
	// Get client connection from resource pool
	resourceConn, err := u.getResourceConn(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get grpc client connection from pool")
	}
	defer u.putResourceConn(resourceConn)
	// Execute grpc request
	ctx, cancel := u.newContext()
	defer cancel()
	req := &pbud.CoinTransactionRequest{
		Transactions: transaction,
	}
	resp, err := resourceConn.UserDataClient.DoCoinTransaction(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute grpc coin transaction to user_data service")
	}
	// Craft response
	return resp.Balances, nil
}

// FindUser retrieves a previously initialized user by userID. If the user is not found, nil will be returned.
func (u *UserData) FindUser(userID string) (*User, error) {
	// Get client connection from resource pool
	resourceConn, err := u.getResourceConn(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get grpc client connection from pool")
	}
	defer u.putResourceConn(resourceConn)
	// Execute grpc request
	ctx, cancel := u.newContext()
	defer cancel()
	req := &pbud.UserRequest{
		UID: userID,
	}
	resp, err := resourceConn.UserDataClient.GetUser(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute grpc GetUser to user_data service")
	}
	// Craft response
	if !resp.Exists {
		return nil, nil
	}
	return &User{
		Coins: resp.Coins,
		ID:    userID,
		Rating: &Rating{
			Value:      resp.Rating.Value,
			Deviation:  resp.Rating.Deviation,
			Volatility: resp.Rating.Volatility,
		},
	}, nil
}

// InitNewUser initializes a new user in the UserData system
func (u *UserData) InitNewUser(userID string) (*User, error) {
	// Get client connection from resource pool
	resourceConn, err := u.getResourceConn(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get grpc client connection from pool")
	}
	defer u.putResourceConn(resourceConn)
	// Execute grpc request
	ctx, cancel := u.newContext()
	defer cancel()
	req := &pbud.UserRequest{
		UID: userID,
	}
	resp, err := resourceConn.UserDataClient.InitNewUser(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute grpc InitNewUser to user_data service")
	}
	// Craft response
	return &User{
		Coins: resp.Coins,
		ID:    userID,
		Rating: &Rating{
			Value:      resp.Rating.Value,
			Deviation:  resp.Rating.Deviation,
			Volatility: resp.Rating.Volatility,
		},
	}, nil
}

// UpdateSkillLevels updates skill levels for a group of users involved in a collection of matches
func (u *UserData) UpdateSkillLevels(matchResults []*MatchResult) (map[string]*Rating, error) {
	// Get client connection from resource pool
	resourceConn, err := u.getResourceConn(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get grpc client connection from pool")
	}
	defer u.putResourceConn(resourceConn)
	// Execute grcp request
	ctx, cancel := u.newContext()
	defer cancel()
	mrs := make([]*pbud.MatchResult, 0, len(matchResults))
	for _, matchResult := range matchResults {
		mrs = append(mrs, &pbud.MatchResult{
			Player1: matchResult.UserID1,
			Player2: matchResult.UserID2,
		})
	}
	req := &pbud.UpdateSkillRequest{
		MatchResults: mrs,
	}
	resp, err := resourceConn.UserDataClient.UpdateSkillLevels(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute grpc UpdateSkillLevels to user_data service")
	}
	// Craft response
	results := map[string]*Rating{}
	for userID, rating := range resp.Ratings {
		results[userID] = &Rating{
			Value:      rating.Value,
			Deviation:  rating.Deviation,
			Volatility: rating.Volatility,
		}
	}
	return results, nil
}

// GetBalances retrieves coin balances for one or more users
func (u *UserData) GetBalances(userIDs []string) (map[string]int64, error) {
	// Get client connection from resource pool
	resourceConn, err := u.getResourceConn(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get grpc client connection from pool")
	}
	defer u.putResourceConn(resourceConn)
	// Execute grpc request
	ctx, cancel := u.newContext()
	defer cancel()
	req := &pbud.BalancesRequest{
		UserIDs: userIDs,
	}
	resp, err := resourceConn.UserDataClient.GetBalances(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute grpc balances request to user_data service")
	}
	// Craft response
	return resp.Balances, nil
}

func (u *UserData) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(u.config.RPCTimeoutSeconds)*time.Second)
}
