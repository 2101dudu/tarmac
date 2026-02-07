#!/bin/bash
set -e

CONTAINER="mongodb"
BACKUP_DIR="backups/mongodb"
TIMESTAMP=$(date +%F)

# Create backup directory if it doesn't exist
mkdir -p "$BACKUP_DIR"

if [ -d "$BACKUP_DIR/$TIMESTAMP" ]; then
    read -p "Backup for today exists, overwrite? [y/N]: " yn
    case $yn in
        [Yy]*) rm -rf "$BACKUP_DIR/$TIMESTAMP" ;;
        *) echo "Aborting"; exit 1 ;;
    esac
fi

source .env 

# Run mongodump inside the container
docker exec "$CONTAINER" \
    mongodump \
      --username ${DB_USERNAME} \
      --password ${DB_PASSWORD} \
      --authenticationDatabase admin \
      --out /tmp/mongo_backup_$TIMESTAMP

# Copy results out of the container
docker cp "$CONTAINER":/tmp/mongo_backup_$TIMESTAMP "$BACKUP_DIR/$TIMESTAMP"

# Cleanup temporary folder inside container
docker exec "$CONTAINER" rm -rf /tmp/mongo_backup_$TIMESTAMP

# Delete backups older than 14 days
find "$BACKUP_DIR" -maxdepth 1 -type d -mtime +14 -exec rm -rf {} +

echo "Backup complete: $BACKUP_DIR/$TIMESTAMP"
