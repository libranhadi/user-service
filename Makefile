.PHONY: migrateup, migratedown

migrateup:
	docker run -v ./migrations:/migrations --network user-service_user-networks migrate/migrate -path=/migrations/ -database "postgres://postgres:postgres@postgresql-service:5432/service_user?sslmode=disable" up

migratedown: 
	docker run -v ./migrations:/migrations --network user-service_user-networks migrate/migrate -path=/migrations/ -database "postgres://postgres:postgres@postgresql-service:5432/service_user?sslmode=disable" down 1

migrateforce:
	docker run -v ./migrations:/migrations --network user-service_user-networks migrate/migrate -path=/migrations/ -database "postgres://postgres:postgres@postgresql-service:5432/service_user?sslmode=disable" force 6