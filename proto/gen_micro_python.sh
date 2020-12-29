#!/bin/bash
#protoc --proto_path=. --micro_out=. --go_out=. --python_out=. --grpc_python_out=. termite.proto
python2.7 -m grpc_tools.protoc --proto_path=. --python_out=. --grpc_python_out=. termite.proto
