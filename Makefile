# Protobuf generation config
GC_PROJECT=indian-game-system
BASE_SRC_PROTO_DIR=${CURDIR}/protodefs
BASE_DEST_PROTO_DIR=${CURDIR}/generated
GOBUILD_DIR=/root/gobuffalo
REQUIRED_PROTO_DIRS="util" "services/user_data" "services/game" "services/game/teen_patti" "services/bot" "services/daily_bonus" "services/matchmaking" "services/matchmaking_system" "services/user_accounts" "services/admin" "services/kafka_websocket"

# Tests
TEST_REPORTS="bot" "config" "daily_bonus" "instance/game" "matchmaking" "math/sampling" "middleware" "rand" "redis" "session" "stream" "user_accounts" "user_data" "uuid"

###########
# Protobuf
###########
generate:
	rm -rf generated/*
	docker run -v ${CURDIR}/generate.sh:${GOBUILD_DIR}/generate.sh \
	 -v ${BASE_SRC_PROTO_DIR}:${GOBUILD_DIR}/protodefs \
	 -v ${BASE_DEST_PROTO_DIR}:${GOBUILD_DIR}/generated \
	 --rm gcr.io/${GC_PROJECT}/gobuild:latest \
	 bash generate.sh

test:
	mkdir -p ./test-reports
	ginkgo -r

	for testDir in ${TEST_REPORTS} ; do \
		mv ./$$testDir/test*.xml ./test-reports ; \
	done	