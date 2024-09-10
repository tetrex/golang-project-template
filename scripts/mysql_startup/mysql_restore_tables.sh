#!/bin/bash

# Variables
DB_NAME="$MYSQL_DATABASE"
DB_USER="$MYSQL_USER"
DB_PASS="$MYSQL_PASSWORD"
BACKUP_DIR="/db_dump/tables_backup"


# Wait for MySQL to be ready
until mysqladmin ping -h"localhost" --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 3
done


# Import each SQL file
for file_name in $BACKUP_DIR/*.sql; do
    echo "Importing $file_name..."
    mysql -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" < "$file_name"
done

echo "All SQL files imported."