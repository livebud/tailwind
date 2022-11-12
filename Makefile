run:
	@ go generate
	@ go run cmd/tailwind/main.go example/play.html

test:
	@ go generate
	@ go test ./...

time:
	@ go generate
	@ go build -o tailwind cmd/tailwind/main.go
	@ time ./tailwind example/play.html
	@ rm tailwind
