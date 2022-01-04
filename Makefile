rundevenvwithmemory:
	go run ./cmd/memory/app.go
rundevenvwithpostgresdb:
	go run ./cmd/postgres/app.go
runlocalpostgresdb:
	cd ./documents/database/local/ && bash up.sh
stoplocalpostgresdb:
	cd ./documents/database/local/ && bash down.sh
deploylocally:
	cd ./documents/docker/ && bash up.sh
