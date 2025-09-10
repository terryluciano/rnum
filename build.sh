#! /bin/bash

set -e

go build -o rnum main.go

sudo mv rnum /usr/local/bin/

echo "rnum installed to /usr/local/bin/rnum"