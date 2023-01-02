dev:
	DATABASE_URL=postgresql://postgres:postgres@localhost:5432/gokbtg?sslmode=disable PORT=:2565 go run server.go
	# DATABASE_URL=postgresql://postgres:postgres@localhost:5432/gokbtg?connect_timeout=10 PORT=:2565 go run server.go

host-postgres: 
	docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgresdb

up:
	docker compose -f postgres-compose.yaml up -d

down: 
	docker compose -f postgres-compose.yaml down