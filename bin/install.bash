#!/bin/bash
cd ../cmd
go build -o ../bin/livechat ./livechat
go build -o ../bin/account ./account
cd ../bin
./account
./livechat
