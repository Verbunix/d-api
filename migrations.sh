#!/usr/bin/env sh

set -e

ENV=${ENV:-local}
echo "ENV: ${ENV}"

#go get -v github.com/rubenv/sql-migrate/...
sql-migrate up -env="${ENV}"
sql-migrate status -env="${ENV}"
