#!/bin/bash

protoc $1/$1pb/$1.proto --go_out=plugins=grpc:.