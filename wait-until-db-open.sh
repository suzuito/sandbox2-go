#!/bin/sh

while [ true ]; do
    mysql -u root -h $DB_HOST -P $DB_PORT -e 'show databases' > /dev/null 2>&1
    if [ $? -eq 0 ]; then
        echo 'DB connection is OK'
        break
    fi
    echo 'Waiting until DB connection is OK'
    sleep 1
done