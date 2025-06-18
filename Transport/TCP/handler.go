package TCP

import (
	"bufio"
	"net"
	"realNow/logger"
)

/**
* all message handling happens here
*
 */

func HandleConnection(conn net.Conn) {
	for {

		bufferReader := bufio.NewReader(conn)
		payload, err := bufferReader.ReadString('\n')
		if err != nil {
			logger.Error("Error reading from connection: ", err)
			conn.Close()
			return
		}
		if payload == "exit\n" {
			logger.Info("Client requested to close the connection")
			conn.Close()
			return
		}
		logger.Info("Received payload: ", payload)
		// Here you would typically unmarshal the payload into a data structure
		// and process it accordingly. For now, we just log it.
		conn.Write([]byte("Message received: " + payload + "\n"))

	}
}
