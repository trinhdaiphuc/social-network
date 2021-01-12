build-server:
	go build -o bin/social-network cmd/main.go

run-server: build-server
	./bin/social-network

build-client:
	cd web && yarn build

run-client:
	cd web && yarn start