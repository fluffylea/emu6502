#!/usr/bin/env bash

# Build hello.rom
cd Test || exit
if ! ./Ophis/bin/ophis -l helloWorld.l -m helloWorld.m helloWorld.oph; then
    echo "Failed to build helloWorld.oph"
    cd ..
    exit 1
fi
cd ..

go run . \
	-loglevel info \
	-listing Test/helloWorld.l \
	-mapping Test/helloWorld.m
