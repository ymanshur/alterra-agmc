run:
	clear && air
test:
	clear
	go test ./... -coverprofile=cover.out
	go tool cover -func=cover.out
ifeq ($(mode), html)
	go tool cover -html=cover.out
endif
	