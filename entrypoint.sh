#!/bin/sh

set -e

echo "⏳ Waiting for PostgreSQL to be ready..."
until nc -z "$DB_HOST" "$DB_PORT"; do
  sleep 1
done

echo "🚀 Running Goose migrations..."
goose -dir ./migrations postgres \
  "host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASS dbname=$DB_NAME sslmode=disable" up

echo "✅ Migrations done. Starting app..."
exec ./main
