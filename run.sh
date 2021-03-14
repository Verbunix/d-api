#!/usr/bin/env sh

rm ./app
clear && go build -o app . && ./app
