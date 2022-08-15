#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail
[[ "${DEBUG:=false}" == 'true' ]] && set -o xtrace

export PGPASSFILE="./.pgpass"

until pg_isready --host=localhost --port=5432 --dbname=postgres --username=user; do
  sleep 1;
done

psql --host "localhost" --username "user" --dbname "postgres" -lqt | cut -d \| -f 1 | grep -qw "${1:-postgres}"
