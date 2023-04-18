test:
	go test ./... -v -count=1

cover:
	go test ./... -v -count=1 -coverprofile=cover.out

cover-html:
	go tool cover -html=cover.out

build:
	go build -o gosync cmd/app/main.go
