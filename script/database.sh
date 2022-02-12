#! /bin/bash

### Setup PostgreSQL server in container for development.
docker run --name postgres-container -p 5432:5432 -v ~/postgresql/data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=postgres -d postgres:bullseye

### PostgreSQL files:
# - /etc/postgresql : configuration files
# - /var/log/postgresql : log files
# - /var/lib/postgresql/data : database files

### Using PGAdmin to interact with database:
docker run --name pgadmin-container -p 5050:80 -e PGADMIN_DEFAULT_EMAIL=marvin@42.fr -e PGADMIN_DEFAULT_PASSWORD=pgadmin -d dpage/pgadmin4

### or using PostgreSQL client, psql:
#psql -h hostname -p port -U postgres -W