#!/bin/bash
cd ../src/cmd
go build -o ../../bin/scfy-im ./main.go
cd ../../bin
./scfy-im