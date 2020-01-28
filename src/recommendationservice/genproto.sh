#!/bin/bash -eu

set -e

python3 -m grpc_tools.protoc -I../../pb --python_out=. --grpc_python_out=. ../../pb/goloco.proto
