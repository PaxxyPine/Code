#!/bin/sh

protoc -I=. --go_out=. gameapi.proto
