# Tailwind for Go (WIP)

Call Tailwind programmatically from Go. Unlike other options, this package embeds tailwind directly and doesn't spawn a subprocess.

## CLI Usage

```
go run cmd/tailwind/main.go example/play.html
```

## API Usage

```go
css, err := tailwind.Process(ctx, "index.html", "some css")
```
