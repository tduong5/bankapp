#!/bin/sh

set -e # exit if return non-zero

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up # similar one found in Makefile

echo "start the app"
exec "$@" # takes all parameters passed to the script and run it
