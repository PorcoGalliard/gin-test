POSTGRES_USER = ultrasur
POSTGRES_PASSWORD = ultrasur123
POSTGRES_DB = simple-forum
POSTGRES_HOST = localhost
POSTGRES_PORT = 5432

export POSTGRES_URL='postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database $(POSTGRES_URL) -path scripts/migrations up

migrate-down:
	@ migrate -database $(POSTGRES_URL) -path scripts/migrations down

migrate-force:
	@ migrate -database $(POSTGRES_URL) -path scripts/migrations force $(version)