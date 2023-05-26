package main

import (
	"os"
	"time"

	xlogs "github.com/mt1976/appFrame/logs"
	xpush "github.com/mt1976/appFrame/pushover"
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
	xlogs.Info("Sending Push Notification to Pushover")

	env := CatchEnvironment()

	Debug(env)

	// Pushover API Token
	//	apiToken := "azjcvqy7ajf9i875ndinapzg7focya"
	//	userToken := "uyosdopsu9wxxo7b264bmnnhbfz8nj"
	title := "Torrent Complete"
	body := "Torrent Completed"
	if env.TR_TORRENT_NAME != "" {
		body = body + env.TR_TORRENT_NAME
	}
	body = body + " at " + env.TR_TIME_LOCALTIME.Format("15:04:05 on Mon Jan 2 2006")
	//var fields xlogs.Fields
	xlogs.WithFields(xlogs.Fields{"SUBJECT": title, "BODY": body}).Info("Sending Message...")
	xpush.Normal(title, body)
	xlogs.Info("Message Sent")
	//apiPushover := pushover.New(apiToken)
	//msgRecipient := pushover.NewRecipient(userToken)

	//msgText := pushover.NewMessage("Hello from Go!")

	//response, err := apiPushover.SendMessage(msgText, msgRecipient)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//log.Println(response)
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
			xlogs.Println("Error", err)
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
		//xlogs.Println("Environment Variable", v, "is not set")
		return ""
	}
	return r
}

func Debug(t TransmissionEnvironment) {
	dbgtxt := "Env"
	dbgFields := xlogs.Fields{
		"VERSION":          t.TR_APP_VERSION,
		"TIME":             t.TR_TIME_LOCALTIME.Format("Mon Jan 2 15:04:05 2006"),
		"BYTES_DOWNLOADED": t.TR_TORRENT_BYTES_DOWNLOADED,
		"DIR":              t.TR_TORRENT_DIR,
		"HASH":             t.TR_TORRENT_HASH,
		"ID":               t.TR_TORRENT_ID,
		"LABELS":           t.TR_TORRENT_LABELS,
		"NAME":             t.TR_TORRENT_NAME,
		"TRACKERS":         t.TR_TORRENT_TRACKERS,
	}

	xlogs.WithFields(dbgFields).Info(dbgtxt)

}
