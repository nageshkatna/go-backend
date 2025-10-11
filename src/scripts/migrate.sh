#!/bin/sh

# Database connection string
SCRIPT_DIR="$(dirname "$0")"
DB_URL="postgres://postgres:postgres@user_db:5432/postgres?sslmode=disable"
MIGRATIONS_PATH="/root/database/migrations"

# Check arguments
if [ $# -eq 0 ]; then
  echo "Usage:"
  echo "  $0 -up"
  echo "  $0 -force <version_number>"
  exit 1
fi

case "$1" in
  -up)
    migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" up
    ;;
  -force)
    if [ -z "$2" ]; then
      echo "Error: You must provide a version number with -force"
      echo "Example: $0 -force 1"
      exit 1
    fi
    migrate -database "$DB_URL" -path "$MIGRATIONS_PATH" force "$2"
    ;;
  *)
    echo "Unknown option: $1"
    echo "Usage:"
    echo "  $0 -up"
    echo "  $0 -force <version_number>"
    exit 1
    ;;
esac