#!/bin/bash

# Variables
DB_NAME="$MYSQL_DATABASE"
DB_USER="$MYSQL_USER"
DB_PASS="$MYSQL_PASSWORD"
BACKUP_DIR="/db_dump/tables_backup"

# Create backup directory if it doesn't exist
mkdir -p "$BACKUP_DIR"

# Get the list of tables in the database
TABLES=$(mysql -u "$DB_USER" -p"$DB_PASS" -D "$DB_NAME" -e "SHOW TABLES;" | awk '{ print $1}' | grep -v '^Tables' )

# Loop through each table and back up the first 1000 rows
for TABLE in $TABLES; do
  echo "Backing up table: $TABLE"
  mysqldump -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" "$TABLE" --where="1 LIMIT 1000" > "$BACKUP_DIR/${TABLE}.sql"
done

echo "Backup completed. Files are stored in $BACKUP_DIR."
