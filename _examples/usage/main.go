package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-numb/go-pipe-for-mt4"
)

const (
	// for MQ4 pipe
	//   _sym = StringSubstr(Symbol(),0,6);
	//   pipe_name = _sym;
	//   Print("connected file to ", pipe_name);
	//   pipe = FileOpen("\\\\.\\pipe\\" + _sym, FILE_WRITE | FILE_BIN | FILE_ANSI);
	FILENAME = `\\.\pipe\USDJPY`
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	ch := make(chan []byte, 512)
	go pipe.Pipe(ctx, FILENAME, 10, ch)

L:
	for {
		select {
		case <-ctx.Done():
			break L

		case v := <-ch:
			fmt.Println(pipe.ByteToTicker(v))
		}
	}
}
