#!/bin/sh
# TODO: run "chmod +x start.sh" for one time after creating this file

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
