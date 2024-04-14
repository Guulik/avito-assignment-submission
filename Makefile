.PHONY: build up integration_tests fake_banners lint

#up - stop - start because volume and tables in the database are created at the first startup,
#so the database may not be ready to accept the connection and the application will fall down :(
first-run: build .up .stop .start integration_tests

test: .up integration_tests

rebuild: build .up

build:
	docker-compose  build

.up:
	docker-compose up -d

.start:
	docker-compose start

integration_tests:
	go test ./tests/integration/...

fake_banners:
	go test -run TestCreate_Happy ./tests/integration/admincreate_test.go

.stop:
	docker-compose stop

down:
	docker-compose down

lint:
	golangci-lint run -c ./config/.golangci.yml

