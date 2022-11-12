package tailwind_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/livebud/js"
	v8 "github.com/livebud/js/v8"
	"github.com/livebud/tailwind"
	"github.com/matryer/is"
	"github.com/matthewmueller/diff"
)

type Test struct {
	Name   string
	Input  string
	Expect string
	Error  string
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
	processor := tailwind.New(vm)
	code, err := processor.Process(ctx, test.Name, test.Input)
	if err != nil {
		if err.Error() == test.Error {
			return nil
		}
		return fmt.Errorf("tailwind: process error: %w", err)
	}
	if code != test.Expect {
		if delta := diff.String(test.Expect, code); delta != "" {
			return fmt.Errorf("tailwind: unexpected result\n\nExpected:\n%s\n\nActual:\n%s\n\nDiff:\n%s\n\n", test.Expect, code, delta)
		}
	}
	return nil
}

func TestSimple(t *testing.T) {
	is := is.New(t)
	is.NoErr(runTest(Test{
		Name: t.Name(),
		Input: `
			<script lang="typescript">
				export let name: string = "Mark";
			</script>
			<h1 class="text-gray-600">Hello {name}!</h1>
		`,
		Expect: `<h1>Hello Matt!</h1>`,
	}))
}
