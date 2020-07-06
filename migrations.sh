#!/usr/bin/env sh

set -e

ENV=${ENV:-local}
echo "ENV: ${ENV}"

sql-migrate up -env="${ENV}"
sql-migrate status -env="${ENV}"
