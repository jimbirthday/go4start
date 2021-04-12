package httpup

import (
	"io"
	"log"
	"net"
	"testing"
	"time"
)

func TestT1(T *testing.T) {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(accept)
	}

}

func handleConn(accept net.Conn) {
	defer accept.Close()
	for {
		_, err := io.WriteString(accept, time.Now().Format(time.RFC3339))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
