#!/bin/bash

echo "Creating databases..."

mysql -u root -p"$MYSQL_ROOT_PASSWORD" -e "
CREATE DATABASE IF NOT EXISTS \`$DB_NAME\`;
CREATE DATABASE IF NOT EXISTS \`$GUAC_DB_NAME\`;
GRANT ALL PRIVILEGES ON \`$DB_NAME\`.* TO 'root'@'%';
GRANT ALL PRIVILEGES ON \`$GUAC_DB_NAME\`.* TO 'root'@'%';
FLUSH PRIVILEGES;
"

echo "Databases created successfully!"
