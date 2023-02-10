# Tailwind.go

Use Tailwind in Go.

Unlike other options, this package embeds tailwind directly and doesn't spawn a subprocess, so it's faster.

## Example Usage

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/livebud/js"
	v8 "github.com/livebud/js/v8"
	"github.com/livebud/tailwind"
)

func main() {
	// Initialize a V8 isolate
	vm, _ := v8.Load(&js.Console{
		Log:   os.Stdout,
		Error: os.Stderr,
	})
  defer vm.Close()
	// Generate the CSS for this index.html file
	processor := tailwind.New(vm)
	ctx := context.Background()
	css, _ := processor.Process(ctx, "index.html", `
    <html>
      <head></head>
      <body>
        <h1 class="bg-[url(/img/grid.svg)] text-sky-500 hover:text-sky-600">Hello Tailwind!</h1>
      </body>
    </html>
  `)
	fmt.Println(css)
}
```

## CLI Usage

```
go install github.com/livebud/tailwind/cmd/tailwind@latest
tailwind example/play.html
```

## Development

You can install the dependencies with make install:

```sh
make install
```

This command expects that you have `go` and `npm` in your toolchain. After installing the dependencies, you can run tests with:

```sh
make test
```

## License

MIT
