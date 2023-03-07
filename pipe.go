package pipe

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/Microsoft/go-winio"
)

func Pipe(
	ctx context.Context,
	filename string,
	updateMillisec int,
	ticker chan []byte) error {

	l, err := winio.ListenPipe(filename, nil)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}

		go handler(conn, ticker)
	}

	<-ctx.Done()

	return nil
}

func handler(c net.Conn, ch chan []byte) {
	defer c.Close()

	var (
		buf = make([]byte, 1028)
		n   int
		err error
	)

	for {
		n, err = c.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("[ERROR]read error: %v\n", err)
			}
			break
		}
		ch <- buf[:n]
	}
	log.Println("Client disconnected")
}
