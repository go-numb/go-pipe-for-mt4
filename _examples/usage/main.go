package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-numb/go-pipe-for-mt4"
)

const (
	FILENAME = `C:\Users\<USER>\AppData\Roaming\MetaQuotes\Terminal\<UUID>\MQL4\Files\pipe\EURUSD.csv`
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	ch := make(chan []byte, 512)
	go pipe.Connect(ctx, FILENAME, 10, ch)

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
