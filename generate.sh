#!/bin/bash
declare -a PROTO_DIRS=("util" "services/user_data" "services/game" "services/game/teen_patti" "services/bot" "services/daily_bonus" "services/matchmaking" "services/matchmaking_system" "services/user_accounts" "services/admin" "services/kafka_websocket")

for dir in ${PROTO_DIRS[@]} ; do \
    mkdir -p generated/$dir ; \
	protoc -I=protodefs/$dir --go_out=plugins=grpc:generated/$dir protodefs/$dir/*.proto ; \
done