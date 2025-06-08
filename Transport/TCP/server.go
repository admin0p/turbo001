package TCP

import (
	"net"
	"realNow/logger"
)

func Serve() {

	// listen for incoming connections
	listener, err := net.Listen("tcp", ":1229")
	if err != nil {
		logger.Error("Error starting server: ", err.Error())
		return
	}

	defer listener.Close()

	logger.Info("Server started on port 1229")

	for {
		// accept connection from client
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Error accepting connection: ", err.Error())
			return
		}

		logger.Info("Client connected: ", conn.RemoteAddr().String())
		// handle client connection in a separate goroutine
		go HandleConnection(conn)
		logger.Info("Handling client connection in a separate goroutine")
	}
}

//  function to handle client connection
