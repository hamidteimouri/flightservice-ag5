#!/bin/bash

protoc --proto_path=. --go_out=../pbs --go-grpc_out=require_unimplemented_servers=false:../pbs *.proto