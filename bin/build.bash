#!/bin/bash
cd ../src/cmd
go build -o ../../bin/IM ./main.go
cd ../../bin
./scfy-im