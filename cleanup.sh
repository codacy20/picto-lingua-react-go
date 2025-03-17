#!/bin/bash

# This script cleans up the directory structure after reorganizing backend files

echo "Starting cleanup process..."

# Verify that backend files are properly copied
if [ ! -f "backend/main.go" ]; then
  echo "Error: backend/main.go not found. Did you run the setup script first?"
  exit 1
fi

# Remove files from the root that have been moved to backend
echo "Removing files from root that have been moved to backend..."
rm -f main.go
rm -rf api
rm -rf config

echo "Cleanup complete. The directory structure is now properly organized."
echo ""
echo "To start the backend, run: cd backend && go run main.go"
echo "To start the frontend (once implemented), run: cd frontend && npm start" 