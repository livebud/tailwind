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
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: svelte2go <file>")
	}
	ctx := context.Background()
	vm, err := v8.Load(&js.Console{
		Log:   os.Stdout,
		Error: os.Stderr,
	})
	if err != nil {
		return fmt.Errorf("unable to load v8: %w", err)
	}
	compiler := tailwind.New(vm)
	svelte, err := os.ReadFile(os.Args[1])
	if err != nil {
		return fmt.Errorf("unable to read svelte file: %w", err)
	}
	css, err := compiler.Process(ctx, os.Args[1], string(svelte))
	if err != nil {
		return fmt.Errorf("unable to compile svelte file: %w", err)
	}
	fmt.Println(css)
	return nil
}
