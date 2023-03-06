package pipe

import (
	"bufio"
	"context"
	"log"
	"os"
	"syscall"
)

func Connect(
	ctx context.Context,
	filename string,
	updateMillisec int,
	ticker chan []byte) error {
	var (
		f   *os.File
		err error
	)
	defer f.Close()

	f, err = os.OpenFile(filename, os.O_RDONLY|syscall.O_NONBLOCK, os.ModeNamedPipe)
	if err != nil {
		log.Println("[ERROR] Open named pipe file error:", err)
		return err
	}

	reader := bufio.NewReader(f)

	go func() {
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				continue
			}
			ticker <- line
		}
	}()

	<-ctx.Done()

	return nil
}
