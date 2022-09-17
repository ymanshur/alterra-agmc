run:
	clear && air
test:
	go test ./... -coverprofile=cover.out && go tool cover -func=cover.out
test-result:
	go tool cover -html=cover.out