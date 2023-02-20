#! /usr/bin/bash

set -e

cd _docker/postgres
ENV=${ENV:-local} bash ./start.sh
