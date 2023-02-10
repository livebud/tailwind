install:
	@ go mod tidy
	@ npm install

generate:
	@ go generate

run: generate
	@ go run cmd/tailwind/main.go example/play.html

test: generate
	@ go test ./...

time: generate
	@ go generate
	@ go build -o tailwind cmd/tailwind/main.go
	@ time ./tailwind example/play.html
	@ rm tailwind
