#!/usr/bin/env bash

# Build hello.rom
cd Test || exit
if ! ./Ophis/bin/ophis -l helloWorld.l helloWorld.oph; then
    echo "Failed to build helloWorld.oph"
    cd ..
    exit 1
fi
cd ..

go run .
