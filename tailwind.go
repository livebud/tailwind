package tailwind

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/livebud/js"
)

//go:generate go run github.com/evanw/esbuild/cmd/esbuild tailwind.ts --platform=neutral --format=iife --global-name=tailwind --bundle --inject:inject.ts --outfile=asset/tailwind.js --log-level=warning --main-fields=browser,module,main --loader:.css=text --minify

//go:embed asset/tailwind.js
var compiler string

func New(vm js.VM) *Processor {
	_, err := vm.Evaluate(context.Background(), "asset/tailwind.js", compiler)
	if err != nil {
		panic(err)
	}
	return &Processor{vm}
}

type Processor struct {
	vm js.VM
}

func (p *Processor) Process(ctx context.Context, name, src string) (string, error) {
	expr := fmt.Sprintf(`tailwind.process({
		content: [{
			raw: %q,
			extension: "svelte",
		}],
	})`, src)
	return p.vm.Evaluate(ctx, name, expr)
}
