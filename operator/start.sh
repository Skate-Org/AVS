#!/bin/bash

trap 'kill $PID1 $PID2 $PID3; exit' SIGINT SIGTERM

cd ..

go run operator/main.go monitor --signer-config 1 &
PID1=$!

go run operator/main.go monitor --signer-config 2 &
PID2=$!

go run operator/main.go monitor --signer-config 3 &
PID3=$!

wait $PID1 $PID2 $PID3
