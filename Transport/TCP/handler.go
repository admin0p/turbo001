package TCP

import (
	"net"
	"realNow/logger"
)

/**
* all message handling happens here
*
 */

func HandleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)

	// Read data from the connection
	_, err := conn.Read(buffer)
	if err != nil {
		logger.Error("Error reading from connection: ", err.Error())
		return
	}
	logger.Info("Received: ", string(buffer))
	// Process the data (this is where you would handle your business logic)

	conn.Write([]byte("Hello from server"))
	logger.Info("Sent: Hello from server")
	// Close the connection
	conn.Close()
	logger.Info("Connection closed")

}
