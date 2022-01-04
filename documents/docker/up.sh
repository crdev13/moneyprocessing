#! /bin/bash

docker-compose -f docker-compose.local.yml up -d postgres
sleep 5 && cat ../database/moneyprocessing.sql | docker exec -i postgres psql -U crpostgres -d moneyprocessing
sleep 3 && docker-compose -f docker-compose.local.yml up -d moneyprocessing