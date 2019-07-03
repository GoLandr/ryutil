package ryutil

import (
	"runtime/debug"
	"log"
	"context"
)

func RecoverPanic(f func()) (err interface{}) {
	defer func() {
		err = recover()
		if err != nil {
			debug.PrintStack()
			log.Printf("panic: %v", err)
		}
	}()

	f()
	return
}

func RunForever(ctx context.Context, f func()) {
	for {
		if ctx.Err() != nil {
			break
		}

		err := RecoverPanic(f)
		log.Printf("RunForever: panic: %v", err)
	}
}
