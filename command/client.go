package command

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/urfave/cli"
	"github.com/vigorcrust/capoeira/util"
)

//CmdClient is the main entrypoint for the client
func CmdClient(c *cli.Context) error {
	log.SetPrefix("[" + util.Name + "-client] ")
	if c.NumFlags() < 2 {
		util.PrintErrorAndUsage("", c)
		return nil
	}

	if c.GlobalBool("verbose") {
		log.Println(c.String("url"))
	}

	dialer := &websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	conn, _, err := dialer.Dial(c.String("url"), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	go timeWriter(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return nil
		}

		fmt.Printf("received: %s\n", message)
	}
}

func timeWriter(conn *websocket.Conn) {
	for {
		fmt.Println()

		time.Sleep(time.Second * 2)
		currenttime := time.Now().Format("2006-01-02 15:04:05.9999")
		fmt.Println("send: " + currenttime)
		conn.WriteMessage(websocket.TextMessage, []byte(currenttime))
	}
}
