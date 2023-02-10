package main

import (
	"context"
	"fmt"
	"os"
	"path"

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
		return fmt.Errorf("usage: tailwind <file>")
	}
	ctx := context.Background()
	fpath := os.Args[1]
	vm, err := v8.Load(&js.Console{
		Log:   os.Stdout,
		Error: os.Stderr,
	})
	if err != nil {
		return fmt.Errorf("unable to load v8: %w", err)
	}
	compiler := tailwind.New(vm)
	data, err := os.ReadFile(fpath)
	if err != nil {
		return fmt.Errorf("unable to read %s file: %w", path.Ext(fpath), err)
	}
	css, err := compiler.Process(ctx, fpath, string(data))
	if err != nil {
		return fmt.Errorf("unable to compile %s file: %w", path.Ext(fpath), err)
	}
	fmt.Println(css)
	return nil
}
