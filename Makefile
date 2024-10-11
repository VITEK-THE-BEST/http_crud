run:
	go run cmd/app/main.go

migrate:
	go run cmd/migrations/main.go

up:
	docker-compose up 

down:
	docker-compose down 

exec:
	docker exec -it http_crud_pgsql_1 bash