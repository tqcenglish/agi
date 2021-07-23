package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/tqcenglish/agi"
)

const Host = "192.168.11.243"
const Port = "2700"
const buffsize = 80000

type Message struct {
	Result []struct {
		Conf  float64
		End   float64
		Start float64
		Word  string
	}
	Text string
}

var m Message

func main() {
	fmt.Printf("listen 18081\n")
	agi.Listen(":18081", handler)
}

func handler(a *agi.AGI) {
	defer a.Close()
	fmt.Print("record demo start\n")
	a.Answer()
	//保存路径 /var/lib/asterisk/sounds
	//按#号结束
	err := a.Record("tqcenglish", &agi.RecordOptions{})
	if err != nil {
		fmt.Printf("%+v", err)
	}
	a.StreamFile("/var/lib/asterisk/sounds/tqcenglish", "", 0)
	fmt.Print("record demo end\n")
	kaldi("/var/lib/asterisk/sounds/tqcenglish.wav")
	fmt.Print("end all\n")
	a.Hangup()
}

func kaldi(path string) {
	u := url.URL{Scheme: "ws", Host: Host + ":" + Port, Path: ""}
	log.Info("connecting to ", u.String())

	// Opening websocket connection
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	check(err)
	defer c.Close()

	f, err := os.Open(path)
	check(err)

	for {
		buf := make([]byte, buffsize)
		dat, err := f.Read(buf)

		if dat == 0 && err == io.EOF {
			err = c.WriteMessage(websocket.TextMessage, []byte("{\"eof\" : 1}"))
			check(err)
			break
		}
		check(err)

		err = c.WriteMessage(websocket.BinaryMessage, buf)
		check(err)

		// Read message from server
		_, _, err = c.ReadMessage()
		check(err)
	}

	// Read final message from server
	_, msg, err := c.ReadMessage()
	check(err)

	// Closing websocket connection
	c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

	// Unmarshalling received message
	err = json.Unmarshal(msg, &m)
	check(err)
	log.Info(m.Text)
}

func check(err error) {

	if err != nil {
		log.Error(err)
	}
}
