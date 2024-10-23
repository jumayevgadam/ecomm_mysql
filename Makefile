.PHONY: create
create:
	migrate create -ext sql -dir ./internal/migrations/ -seq db_schema

.PHONY: migr_with_docker
migr_with_docker:
	docker run -it --rm --network host --volume "$(pwd)/internal:/internal" migrate/migrate:v4.17.0 create -ext sql -dir ./internal/migrations/ init_schema	

.PHONY: migration_up
migration_up:
	docker run -it --rm --network host --volume ./internal:/internal migrate/migrate:v4.17.0 -path=/internal/migrations -database 'mysql://root:NewPassword!123@tcp(localhost:3306)/ecomm_mysql' up

.PHONY: docker-run
docker-run:
	docker run --name ecomm-mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=NewPassword!123 -d mysql:8.4

.PHONY: docker-exec
docker-exec:
	docker exec -i ecomm_mysql mysql -uroot -password <<< "CREATE DATABASE ecomm_mysql;"