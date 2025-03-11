#!/bin/bash

# Load environment variables from .env file (for ash shell of Alpine Linux)
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo "Error: .env file not found!"
    exit 1
fi

# Check if SQL file is provided as argument
if [ -z "$1" ]; then
    echo "Usage: $0 <sql_file>"
    exit 1
fi

SQL_FILE=$1

# Check if SQL file exists
if [ ! -f "$SQL_FILE" ]; then
    echo "Error: File '$SQL_FILE' not found!"
    exit 1
fi

# Executed SQL file using psql command
PGPASSWORD="root" psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f $SQL_FILE

# Check psql command status
if [ $? -eq 0 ]; then
    # Add or update LAST_EXECUTED_SQL variable in .env file
    if grep -q '^LAST_EXECUTED_SQL=' .env; then
        # If already exists, update the value
        sed -i "s|^LAST_EXECUTED_SQL=.*|LAST_EXECUTED_SQL=$SQL_FILE|" .env
    else
        # If not exists, add the value to the end of the file
        echo "LAST_EXECUTED_SQL=$SQL_FILE" >> .env
    fi
    echo "Successfully executed: $SQL_FILE"
else
    echo "Error: Failed to execute SQL file."
    exit 1
fi
