create_migration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate_up:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migrate_down:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

.PHONY: migrate migrate_up migrate_down