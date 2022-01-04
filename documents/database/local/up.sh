#! /bin/bash

docker-compose -f docker-compose.local.yml up -d
sleep 5 && cat ../moneyprocessing.sql | docker exec -i postgres psql -U crpostgres -d moneyprocessing