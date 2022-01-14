# Money Processing Service
This is a challenge.

Clients of the bank who keep money there.
Each Client can have multiple Accounts.
Each Account balance can be denominated in one currency.
For example Client1 can have three accounts: one USD account and two COP accounts.
The following currencies should be supported: USD, COP, MXN.

Transaction is an action to update an Account or Accounts balance.
A transaction should belong to one of following types:
- Deposit money
- Withdraw money
- Money transfer between Clients (Currency conversion is not needed. Only transfer
between accounts with the same currency is allowed.)

The following methods should be implemented:
- Create Client
- Create Account for Client
- Get Client
- Get Account
- Get Transactions - return list of transactions for account
- Create Transaction - create transaction of needed type

Examples to make request [Check details](./documents/client)

API - JSON via REST
Programming language - Golang
Database - PostgreSQL

# Money processing service -  Example usage
Steps:
1. Create the bellow docker-compose file

```
## moneyprocessing.yml

version: '3.1'

services:

  postgres:
    image: postgres
    container_name: postgres
    restart: always
    environment:
      POSTGRES_DB: moneyprocessing
      POSTGRES_USER: crpostgres
      POSTGRES_PASSWORD: mypass
    ports:
      - "5432:5432"

  moneyprocessing:
    image: crca13/moneyprocessing:latest
    container_name: moneyprocessing
    environment:
      HOST: moneyprocessing.com
      DBHOST: postgres
      DBNAME: moneyprocessing
      DBUSER: crpostgres
      DBPASS: mypass
    ports:
      - "8080:8080"

networks:
  default:
    external:
      name: dev-net
```

Image Repository [Check details](https://hub.docker.com/r/crca13/moneyprocessing)

2. Run [up.sh](./documents/docker/up.sh)
```
## Start app using Makefile
$ make deploylocally

## Stop app using Makefile
$ make stoplocaldeployment
```
