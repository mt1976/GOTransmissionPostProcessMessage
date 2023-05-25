package main

import (
	"log"
	"os"
	"time"

	pushover "github.com/gregdel/pushover"
)

type TransmissionEnvironment struct {
	TR_APP_VERSION              string
	TR_TIME_LOCALTIME           time.Time
	TR_TORRENT_BYTES_DOWNLOADED string
	TR_TORRENT_DIR              string
	TR_TORRENT_HASH             string
	TR_TORRENT_ID               string
	TR_TORRENT_LABELS           string
	TR_TORRENT_NAME             string
	TR_TORRENT_TRACKERS         string
}

func main() {
	log.Println("Sending Push Notification to Pushover")

	env := CatchEnvironment()

	Debug(env)

	log.Println("POO", env.TR_TORRENT_NAME, "WEE")

	// Pushover API Token
	apiToken := "azjcvqy7ajf9i875ndinapzg7focya"
	userToken := "uyosdopsu9wxxo7b264bmnnhbfz8nj"

	apiPushover := pushover.New(apiToken)
	msgRecipient := pushover.NewRecipient(userToken)

	msgText := pushover.NewMessage("Hello from Go!")

	response, err := apiPushover.SendMessage(msgText, msgRecipient)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(response)
}

func CatchEnvironment() TransmissionEnvironment {
	t := TransmissionEnvironment{}
	t.TR_APP_VERSION = Get("TR_APP_VERSION")
	lt := Get("TR_TIME_LOCALTIME")

	if lt == "" {
		t.TR_TIME_LOCALTIME = time.Now()
	} else {
		// Time Example : "Tue Dec 13 06:24:36 2016"
		nt, err := time.Parse("Mon Jan 2 15:04:05 2006", lt)
		if err != nil {
			log.Println("Error", err)
			t.TR_TIME_LOCALTIME = time.Now()
		}
		t.TR_TIME_LOCALTIME = nt
	}

	t.TR_TORRENT_BYTES_DOWNLOADED = Get("TR_TORRENT_BYTES_DOWNLOADED")
	t.TR_TORRENT_DIR = Get("TR_TORRENT_DIR")
	t.TR_TORRENT_HASH = Get("TR_TORRENT_HASH")
	t.TR_TORRENT_ID = Get("TR_TORRENT_ID")
	t.TR_TORRENT_LABELS = Get("TR_TORRENT_LABELS")
	t.TR_TORRENT_NAME = Get("TR_TORRENT_NAME")
	t.TR_TORRENT_TRACKERS = Get("TR_TORRENT_TRACKERS")
	return t
}

func Get(v string) string {
	r := os.Getenv(v)
	//log.Println(v, r)
	if r == "" {
		log.Println("Environment Variable", v, "is not set")
		return ""
	}
	return r
}

func Debug(t TransmissionEnvironment) {
	log.Println("TR_APP_VERSION", t.TR_APP_VERSION)
	log.Println("TR_TIME_LOCALTIME", t.TR_TIME_LOCALTIME.Format("Mon Jan 2 15:04:05 2006"))
	log.Println("TR_TORRENT_BYTES_DOWNLOADED", t.TR_TORRENT_BYTES_DOWNLOADED)
	log.Println("TR_TORRENT_DIR", t.TR_TORRENT_DIR)
	log.Println("TR_TORRENT_HASH", t.TR_TORRENT_HASH)
	log.Println("TR_TORRENT_ID", t.TR_TORRENT_ID)
	log.Println("TR_TORRENT_LABELS", t.TR_TORRENT_LABELS)
	log.Println("TR_TORRENT_NAME", t.TR_TORRENT_NAME)
	log.Println("TR_TORRENT_TRACKERS", t.TR_TORRENT_TRACKERS)

}
