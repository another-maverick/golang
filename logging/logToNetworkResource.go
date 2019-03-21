package main

// create a server by executing nc -lk 1234
import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")

	if err != nil {
		panic("failed to connect on port 1234")
	}

	defer conn.Close()

	logger := log.New(conn, "testGoLog  ", log.LstdFlags|log.Lshortfile)
	logger.Println("my first log message via logger")
	logger.Panicln("panicing now. Future logs from this program wont get printed")
}
