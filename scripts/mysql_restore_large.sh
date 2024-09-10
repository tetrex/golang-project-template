#!/bin/bash

# Variables
DB_NAME="$MYSQL_DATABASE"
DB_USER="$MYSQL_USER"
DB_PASS="$MYSQL_PASSWORD"
SQL_FILE="/db_dump/mysql_backup.sql"

# Check if SQL file exists
if [ ! -f "$SQL_FILE" ]; then
  echo "SQL file not found!"
  exit 1
fi

# Restore the database
# pv "$SQL_FILE" | mysql -u $DB_USER -p$DB_PASS $DB_NAME
mysql -u $DB_USER -p$DB_PASS $DB_NAME < $SQL_FILE


# Check if the restore was successful
if [ $? -eq 0 ]; then
  echo "Database restored successfully."
else
  echo "Failed to restore the database."
fi