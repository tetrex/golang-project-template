#!/bin/bash

# Set the database, dump directory, and authentication details
DUMP_DIR="/db_dump/mongo_backup"
DB_NAME="db"
MONGO_USER="root"
MONGO_PASS="pass"
AUTH_DB="admin"  # Change if your authentication database is different

# Loop through all .json files in the dump directory
for file in "$DUMP_DIR"/*.json; do
    # Extract the collection name from the filename
    collection=$(basename "$file" .json)
    
    echo "Restoring collection: $collection from file: $file"
    
    # Create collection (optional)
    mongosh --eval "db.getSiblingDB('$DB_NAME').createCollection('$collection')" \
            --username "$MONGO_USER" --password "$MONGO_PASS" \
            --authenticationDatabase "$AUTH_DB"
    
    # Import data (line-delimited JSON format, no --jsonArray flag)
    mongoimport --db "$DB_NAME" --collection "$collection" \
                --username "$MONGO_USER" --password "$MONGO_PASS" \
                --authenticationDatabase "$AUTH_DB" --file "$file"
done

echo "Restore process completed."
