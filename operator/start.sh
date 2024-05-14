#!/bin/bash

# Array to hold PIDs
declare -a PIDs

# Function to kill all processes when the script receives a SIGINT or SIGTERM
trap 'kill "${PIDs[@]}"; exit' SIGINT SIGTERM

# Move up one directory
cd ..

# Start each operator in the background
for i in {1..3}; do
    ./bin/operator monitor --signer-config $i --verbose=false &
    # Store PID of the last background process
    PIDs+=($!)
done

# Wait for all processes to exit
wait "${PIDs[@]}"
