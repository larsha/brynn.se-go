.PHONY: run

run:
	export $$(cat .env | grep -v ^\# | xargs) && \
	go run main.go
