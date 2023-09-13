createmigration:
	migrate create -ext=sql -dir=internal/infra/database/migrations -seq init
	

migrateup:
	migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" --verbose up

migratedown:
	migrate -path=internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" --verbose down


.PHONY: migrateup migratedown createmigration