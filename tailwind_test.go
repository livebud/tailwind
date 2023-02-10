package tailwind_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/livebud/js"
	v8 "github.com/livebud/js/v8"
	"github.com/livebud/tailwind"
	"github.com/matryer/is"
)

type Test struct {
	Path     string
	Input    string
	Contains []string
	Error    string
}

func runTest(test Test) error {
	ctx := context.Background()
	vm, err := v8.Load(&js.Console{
		Log:   os.Stdout,
		Error: os.Stderr,
	})
	if err != nil {
		if err.Error() == test.Error {
			return nil
		}
		return fmt.Errorf("v8: load error: %w", err)
	}
	defer vm.Close()
	processor := tailwind.New(vm)
	css, err := processor.Process(ctx, test.Path, test.Input)
	if err != nil {
		if err.Error() == test.Error {
			return nil
		}
		return fmt.Errorf("tailwind: process error: %w", err)
	}
	for _, contain := range test.Contains {
		if !strings.Contains(css, contain) {
			return fmt.Errorf("tailwind: expected result to contain %q.\n\n%q", contain, css)
		}
	}
	return nil
}

func TestHTML(t *testing.T) {
	is := is.New(t)
	is.NoErr(runTest(Test{
		Path: "index.html",
		Input: `
			<h1 class="bg-[url(/img/grid.svg)] text-sky-500 hover:text-sky-600">Hello {name}!</h1>
		`,
		Contains: []string{
			// ".bg-\\[url\\(\\/img\\/grid\\.svg\\)\\]{background-image:url(/img/grid.svg)}",
			// ".text-sky-500{--tw-text-opacity:1;color:rgb(14 165 233/var(--tw-text-opacity))}",
			// ".hover\\:text-sky-600:hover{--tw-text-opacity:1;color:rgb(2 132 199/var(--tw-text-opacity))}",
		},
	}))
}

func TestSvelte(t *testing.T) {
	is := is.New(t)
	is.NoErr(runTest(Test{
		Path: "index.svelte",
		Input: `
			<script lang="typescript">
				export let name: string = "Mark";
			</script>
			<h1 class="bg-[url(/img/grid.svg)] text-sky-500 hover:text-sky-600">Hello {name}!</h1>
		`,
		Contains: []string{
			// ".bg-\\[url\\(\\/img\\/grid\\.svg\\)\\]{background-image:url(/img/grid.svg)}",
			// ".text-sky-500{--tw-text-opacity:1;color:rgb(14 165 233/var(--tw-text-opacity))}",
			// ".hover\\:text-sky-600:hover{--tw-text-opacity:1;color:rgb(2 132 199/var(--tw-text-opacity))}",
		},
	}))
}
