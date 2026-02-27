package wazero_atomic

import (
	"context"
	_ "embed"
	"os"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed test\.wasm
var testwasm []byte

func TestModule(t *testing.T) {
	var (
		ctx = context.Background()
	)
	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig())
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	hostModule := New()
	hostModule.Register(ctx, r)

	compiled, err := r.CompileModule(ctx, testwasm)
	if err != nil {
		panic(err)
	}
	cfg := wazero.NewModuleConfig().WithStdout(os.Stdout).WithName(`a`)
	mod1, err := r.InstantiateModule(ctx, compiled, cfg)
	if err != nil {
		t.Errorf(`%v`, err)
		return
	}
	cfg = wazero.NewModuleConfig().WithStdout(os.Stdout).WithName(`b`)
	mod2, err := r.InstantiateModule(ctx, compiled, cfg)
	if err != nil {
		t.Errorf(`%v`, err)
		return
	}
	ctx, err = hostModule.InitContext(ctx, mod1)
	if err != nil {
		t.Fatalf(`%v`, err)
	}

	ctx = ContextCopy(ctx, ctx)

	t.Run(`uint64`, func(t *testing.T) {
		var n int
		t.Run(`add`, func(t *testing.T) {
			for n = range 10 {
				stack, err := mod1.ExportedFunction(`testUint64Add`).Call(ctx, uint64(1))
				if err != nil {
					t.Fatalf("%v", err)
				}
				if stack[0] != n {
					t.Fatalf("expected %d, got %d", n, stack[0])
				}
			}
		})
		t.Run(`load`, func(t *testing.T) {
			stack, err := mod2.ExportedFunction(`testUint64Load`).Call(ctx)
			if err != nil {
				t.Fatalf("%v", err)
			}
			if stack[0] != uint64(n+1) {
				t.Fatalf("expected %d, got %d", n, stack[0])
			}
		})
		t.Run(`store`, func(t *testing.T) {
			n = 100
			stack, err := mod1.ExportedFunction(`testUint64Store`).Call(ctx, uint64(n))
			if err != nil {
				t.Fatalf("%v", err)
			}
			stack, err = mod2.ExportedFunction(`testUint64Load`).Call(ctx)
			if err != nil {
				t.Fatalf("%v", err)
			}
			if stack[0] != uint64(n) {
				t.Fatalf("expected %d, got %d", n, stack[0])
			}
		})
	})

	hostModule.Stop()
}
