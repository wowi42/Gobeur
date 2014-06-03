/*

This prototype is a simple tcp server
When a message is received, it replies with the MD5 sum of this message

*/

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			buf := make([]byte, 4096) // Buffer used to store received data
			for {
				/* Receive the data */
				n, err := c.Read(buf)
				if err != nil || n == 0 {
					c.Close()
					break
				}

				/* Computing the MD5 sum */
				s := string(buf[0:n-2]) // Received string

				hash := md5.New()
				io.Copy(hash, strings.NewReader(s))
				sum := hash.Sum(nil)

				/* Send the MD5 sum */
				n, err = c.Write([]byte(fmt.Sprintf("%s\n", hex.EncodeToString(sum))))
				if err != nil {
					c.Close()
					break
				}
			}

			// Shut down the connection.
			c.Close()
		}(conn)
	}
	os.Exit(0)
	io.Copy(nil, nil)
}
