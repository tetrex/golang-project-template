#!/bin/bash

# Set the database, dump directory, and authentication details
DUMP_DIR="/db_dump/lykStage"
DB_NAME="db"
MONGO_USER="root"
MONGO_PASS="pass"
AUTH_DB="admin"  # Change if your authentication database is different

# Loop through all .bson files in the dump directory
for file in "$DUMP_DIR"/*.bson; do
    # Extract the collection name from the filename
    collection=$(basename "$file" .bson)
    
    echo "Restoring collection: $collection from file: $file"
 
    # Create collection
    mongosh --eval "db.createCollection('$collection')"
    
    # Import data
    # mongoimport --db "$DB_NAME" --collection "$collection" \
    #             --username "$MONGO_USER" --password "$MONGO_PASS" \
    #             --authenticationDatabase "$AUTH_DB" --file "$file" --jsonArray
    # Restore the collection using mongorestore with authentication
    mongorestore --db "$DB_NAME" --collection "$collection" \
                 --username "$MONGO_USER" --password "$MONGO_PASS" \
                 --authenticationDatabase "$AUTH_DB" "$file"
done

echo "Restore process completed."
