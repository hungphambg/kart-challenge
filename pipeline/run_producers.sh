#!/bin/bash

# This script runs the Kafka producer for each couponbase file.

# Assuming this script is run inside a container that has access to the producer binary and the .gz files

for i in 1 2 3; do
  FILE="/app/producer/couponbase${i}" # Path inside the container
  echo "Running producer for $FILE..."
  /app/producer/main $FILE # Execute the producer binary directly
  echo "Finished producer for $FILE."
  sleep 5 # Optional: wait a bit between runs
done

echo "All producers finished."
