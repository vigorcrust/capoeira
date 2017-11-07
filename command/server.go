package command

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/urfave/cli"
	"github.com/vigorcrust/capoeira/util"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//CmdServer is the main entrypoint for the server
func CmdServer(c *cli.Context) error {
	log.SetPrefix("[" + util.Name + "-server] ")
	if c.NumFlags() < 2 {
		util.PrintErrorAndUsage("", c)
		return nil
	}

	if c.GlobalBool("verbose") {
		log.Println(c.String("port"))
	}

	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServeTLS(c.String("port"), "server.crt", "server.key", nil))

	return nil
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		currenttime := time.Now().Format("2006-01-02 15:04:05.9999")
		log.Printf("recv: %s", message)
		log.Printf("send: %s", currenttime)
		err = c.WriteMessage(mt, []byte(currenttime))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
