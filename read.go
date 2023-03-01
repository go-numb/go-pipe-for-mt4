package pipe

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func Connect(
	ctx context.Context,
	filename string,
	updateMillisec int,
	ticker chan []byte) error {
	var (
		f   *os.File
		err error

		buf = make([]byte, 1028)
		n   int
	)
	defer f.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			// now := time.Now()
			f, err = os.OpenFile(filename, os.O_RDONLY|syscall.O_NONBLOCK, os.ModeNamedPipe)
			if err != nil {
				log.Println("[ERROR] Open named pipe file error:", err)
				return err
			}

			for {
				n, err = f.Read(buf)
				if err == io.EOF {
					break
				} else if err != nil {
					log.Println("[ERROR] ", err)
					return err
				}

				ticker <- buf[:n]
			}
			f.Close()
			// fmt.Println(time.Since(now))
			time.Sleep(time.Duration(updateMillisec) * time.Millisecond)
		}
	}
}

type Ticker struct {
	Ltp       float64
	Ask       float64
	Bid       float64
	Volume    float64
	Timestamp time.Time
}

func ByteToTicker(buf []byte) *Ticker {
	str := strings.Split(string(buf), ",")
	last, _ := strconv.ParseFloat(str[0], 64)
	ask, _ := strconv.ParseFloat(str[1], 64)
	bid, _ := strconv.ParseFloat(str[2], 64)
	volume, _ := strconv.ParseFloat(str[3], 64)
	t, _ := strconv.ParseInt(str[4], 10, 64)
	timestamp := time.UnixMilli(t)
	return &Ticker{
		Ltp:       last,
		Ask:       ask,
		Bid:       bid,
		Volume:    volume,
		Timestamp: timestamp,
	}
}

func (t *Ticker) String() string {
	return fmt.Sprintf(
		"ltp: %f, ask: %f, bid: %f, volume: %f, timestamp: %d:%d:%d.%d",
		t.Ltp,
		t.Ask,
		t.Bid,
		t.Volume,
		t.Timestamp.Hour(),
		t.Timestamp.Minute(),
		t.Timestamp.Second(),
		t.Timestamp.Nanosecond())
}
