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
	    <head>
	      <link rel="stylesheet" href="index.min.css" />
	    </head>
	    <body>
	      <h1 class="bg-[url(/img/grid.svg)] text-sky-500 hover:text-sky-600">Hello Tailwind!</h1>
	    </body>
	  </html>
	`)
	fmt.Println(css)
}
```

This results in a CSS file:

```css
/*! tailwindcss v3.2.4 | MIT License | https://tailwindcss.com*/*,:after,:before{border:0 solid #e5e7eb;box-sizing:border-box}:after,:before{--tw-content:""}html{-webkit-text-size-adjust:100%;font-feature-settings:normal;font-family:ui-sans-serif,system-ui,-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Helvetica Neue,Arial,Noto Sans,sans-serif,Apple Color Emoji,Segoe UI Emoji,Segoe UI Symbol,Noto Color Emoji;line-height:1.5;-moz-tab-size:4;-o-tab-size:4;tab-size:4}
/* ... */
```

For production, I recommend building this CSS file and embedding it alongside the processed HTML file. 

For development, you may want to run this processor dynamically. Keep in mind that this processor currently depends on CGO. 

I originally built this project for [Bud](https://github.com/livebud/bud), which handles the environment differences for you.

## CLI Usage

We've also included a rudimentary CLI for quickly testing. If you already have the official CLI installed, stick with that.

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
