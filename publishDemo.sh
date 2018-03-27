#!/bin/bash

docker build -t big-earl-cli -f Dockerfile.cli .

docker tag big-earl-cli micahlee/big-earl-cli

docker push micahlee/big-earl-cli
