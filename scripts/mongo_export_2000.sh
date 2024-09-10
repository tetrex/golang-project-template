#!/bin/bash

# Set MongoDB credentials and database name
MONGO_HOST="local_mongodb"
MONGO_PORT="27017"
MONGO_USER="root"
MONGO_PASS="pass"
AUTH_DB="admin"
TARGET_DB="db"
OUTPUT_DIR="/db_dump/mongo_backup"

# Ensure output directory exists
mkdir -p "$OUTPUT_DIR"

collections=$(mongosh --quiet --host $MONGO_HOST --port $MONGO_PORT -u $MONGO_USER -p $MONGO_PASS --authenticationDatabase $AUTH_DB --eval "
use $TARGET_DB;")

# Get a list of collections from the target database
collections=$(mongosh --quiet --host $MONGO_HOST --port $MONGO_PORT -u $MONGO_USER -p $MONGO_PASS --authenticationDatabase $AUTH_DB --eval "
db.getSiblingDB('$TARGET_DB').getCollectionNames().join('\n');
")

# Print collections for debugging (optional)
echo "Collections: $collections"

# Loop through each collection and export the first 2000 documents
for collection in $collections; do
  echo "Exporting first 2000 documents from collection: $collection"

  # Export the first 2000 documents to a JSON file
  mongosh --quiet --host $MONGO_HOST --port $MONGO_PORT -u $MONGO_USER -p $MONGO_PASS --authenticationDatabase $AUTH_DB --eval "
    db.getSiblingDB('$TARGET_DB').$collection.find().limit(2000).forEach(function(doc) {
      print(JSON.stringify(doc));
    });
  " > "$OUTPUT_DIR/$collection.json"
done

echo "Backup completed to $OUTPUT_DIR"
