package tgbot

//go:generate ./objects.py

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	APIURLBase        string        = "https://api.telegram.org/bot"
	CooldownStart     time.Duration = time.Second
	CooldownScale     int           = 2
	GetUpdatesTimeout int           = 3 //60
)

type UpdateHandler func(*Basic, *Update)

type Basic struct {
	Name          string
	APIURL        string
	UpdateHandler UpdateHandler
	Cooldown      time.Duration
}

var (
	ErrNotOk = errors.New("reply has a false 'ok' field")
)

type Reply struct {
	OK          bool        `json:"ok"`
	Description string      `json:"description"`
	ErrorCode   int         `json:"error_code"`
	Result      interface{} `json:"result"`
}

type ReplyNotOk Reply

func (r *ReplyNotOk) Error() string {
	if r.OK {
		return "no error"
	}
	return fmt.Sprintf("reply not ok: %s (code = %d)", r.Description, r.ErrorCode)
}

func (b *Basic) Get(loc string, q Query, result interface{}) error {
	urltail := loc + "?" + q.Encode()
	log.Println("GET", urltail)
	url := b.APIURL + urltail
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reply := &Reply{}
	reply.Result = result

	err = json.NewDecoder(resp.Body).Decode(reply)
	if err != nil {
		return err
	}
	if !reply.OK {
		return (*ReplyNotOk)(reply)
	}
	// log.Println("result:", reply.Result)
	return nil
}

func incCooldown(cd time.Duration) time.Duration {
	if cd == 0 {
		return CooldownStart
	}
	return cd * time.Duration(CooldownScale)
}

// Main is the main loop.
func (b *Basic) Main() {
	offset := 0
	for {
		var lastCooldown time.Duration
		if b.Cooldown > 0 {
			log.Println("cooldown for", b.Cooldown)
			<-time.After(b.Cooldown)
			lastCooldown = b.Cooldown
			b.Cooldown = 0
		}
		var updates []Update
		err := b.Get("/getUpdates", Query{"offset": offset, "timeout": GetUpdatesTimeout}, &updates)
		if err != nil {
			log.Println("error with /getUpdates:", err)
			b.Cooldown = incCooldown(lastCooldown)
			continue
		}
		log.Printf("%d update(s)\n", len(updates))
		for _, update := range updates {
			offset = update.UpdateID + 1
			if b.UpdateHandler != nil {
				b.UpdateHandler(b, &update)
			}
		}
	}
}

func Main(name, token string, mh UpdateHandler) {
	b := &Basic{name, APIURLBase + token, mh, 0}
	b.Main()
}

func MainWithTokenFile(name, fname string, mh UpdateHandler) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("cannot read token file: %s", err)
	}
	token := strings.TrimSpace(string(buf))
	Main(name, token, mh)
}
