run:
	go run cmd/main.go

migrate:
	go run cmd/migrations/main.go

up:
	docker-compose up 

down:
	docker-compose down 

exec:
	docker exec -it http_crud_pgsql_1 bash