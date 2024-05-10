#!/bin/bash

trap 'kill $PID1 $PID2; exit' SIGINT SIGTERM

cd ..

go run relayer/main.go retrieve &
PID1=$!

go run relayer/main.go publish --signer-config 1 &
PID2=$!

wait $PID1 $PID2
