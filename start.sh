#!/bin/bash

export GOPATH=$(go env GOPATH)
sudo docker network prune -f
sudo docker volume prune -f
cd fixtures && docker-compose up -d
cd ..
rm education
go build
./education