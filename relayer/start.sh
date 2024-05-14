#!/bin/bash

trap 'kill $PID1 $PID2; exit' SIGINT SIGTERM

cd ..

./bin/relayer retrieve &
PID1=$!

./bin/relayer publish --signer-config 1 &
PID2=$!

wait $PID1 $PID2
