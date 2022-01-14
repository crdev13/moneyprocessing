rundevenvwithmemory:
	go run ./cmd/memory/app.go
rundevenvwithpostgresdb:
	go run ./cmd/postgres/app.go


# DB Local deployment
runlocalpostgresdb:
	cd ./documents/database/local/ && bash up.sh
stoplocalpostgresdb:
	cd ./documents/database/local/ && bash down.sh

# App Local deployment
deploylocally:
	cd ./documents/docker/ && bash up.sh
stoplocaldeployment:
	cd ./documents/docker/ && bash down.sh
