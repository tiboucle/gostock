#!/bin/bash

database="gostock.db"

if [ ! -f $database ]; then
	echo "Creating Database..."
	exec `sqlite3 gostock.db < migration/ddl.sql`

	echo "Migrating Data..."
	exec `sqlite3 gostock.db < migration/dml.sql`

	if [ -f $database ]; then
		echo "Database is created"
	fi
else
	echo "Database is ready"
fi

echo "Starting web server GoStock on http://localhost:8090"
exec `go run application.go`