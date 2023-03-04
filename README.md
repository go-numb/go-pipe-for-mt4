# go-pipe-for-mt4

This module reads the Ticker retrieved by Meta Trader4 from the Golang program. An mq4 sample file is included for saving.

## USAGE
```go 
package main

import (
	"context"
	"fmt"
	"github.com/go-numb/go-pipe-for-mt4"
	"time"
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
            // -> ltp: 0.000000, ask: 1.065350, bid: 1.065190, volume: 0.000000, timestamp: 2:20:53.0
		}
	}
}

```