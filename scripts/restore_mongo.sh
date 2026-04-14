#!/bin/bash
set -e

# Usage: ./restore_mongo.sh 2025-01-10

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
cd "$PROJECT_ROOT"

DATE="$1"
BACKUP_DIR="backups/mongodb/$DATE"
CONTAINER="mongodb"

if [ -z "$DATE" ]; then
    echo "Usage: $0 <YYYY-MM-DD>"
    exit 1
fi

if [ ! -d "$BACKUP_DIR" ]; then
    echo "Backup directory not found: $BACKUP_DIR"
    exit 1
fi

echo "Copying backup '$DATE' into container..."
docker cp "$BACKUP_DIR" "$CONTAINER":/tmp/restore

echo "Restoring MongoDB backup..."

source .env

docker exec -it "$CONTAINER" mongorestore \
  --username ${DB_USERNAME} \
  --password ${DB_PASSWORD} \
  --authenticationDatabase admin \
  --drop \
  /tmp/restore

echo "Cleaning up temp restore directory inside container..."
docker exec "$CONTAINER" rm -rf /tmp/restore

echo "Restore completed successfully for backup: $DATE"
