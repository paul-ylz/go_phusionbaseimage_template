#!/bin/bash
ssh -t -t user@host << EOF
cd /projectdir
git pull
git checkout staging
docker-machine use machinename
docker-machine ls
docker-compose -f staging.yml build
docker-compose stop
docker-compose rm -f
docker-compose -f staging.yml up -d --remove-orphans
docker exec containername ./migrations/up
exit
EOF
