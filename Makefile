init_db:
	docker run \
		--name postgres-db \
		--rm -e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=0000 \
		-e POSTGRES_DB=db-go-backend-se101 \
		-p 5432:5432 -it \
		-d postgres:latest

# short command with make local
local: 
	go run cmd/main.go -config ./config/env -env=local -upgrade=${m}
# change upgrade false to true -> create a new user table
# change m=false or true