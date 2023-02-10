package tailwind

import (
	"context"
	_ "embed"
	"fmt"
	"path"

	"github.com/livebud/js"
)

//go:generate go run github.com/evanw/esbuild/cmd/esbuild tailwind.ts --platform=neutral --format=iife --global-name=tailwind --bundle --inject:inject.ts --outfile=asset/tailwind.js --log-level=warning --main-fields=browser,module,main --loader:.css=text --minify

//go:embed asset/tailwind.js
var tailwindjs string

func New(vm js.VM) *Processor {
	_, err := vm.Evaluate(context.Background(), "asset/tailwind.js", tailwindjs)
	if err != nil {
		panic(err)
	}
	return &Processor{vm}
}

type Processor struct {
	vm js.VM
}

func (p *Processor) Process(ctx context.Context, name, src string) (css string, err error) {
	expr := fmt.Sprintf(`tailwind.process({
		content: [{
			extension: %q,
			raw: %q,
		}],
	})`, path.Ext(name), src)
	css, err = p.vm.Evaluate(ctx, name, expr)
	if err != nil {
		return "", fmt.Errorf("tailwind: %w", err)
	}
	return css, nil
}
