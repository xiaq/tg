package tg

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	APIURLBase        string        = "https://api.telegram.org/bot"
	CooldownStart     time.Duration = time.Second
	CooldownScale     int           = 2
	GetUpdatesTimeout int           = 3600
)

type UpdateHandler func(*Bot, *Update)

type Bot struct {
	Name           string
	APIURL         string
	UpdateHandlers []UpdateHandler
	Cooldown       time.Duration
}

var (
	ErrNotOk = errors.New("reply has a false 'ok' field")
)

type Reply struct {
	OK          bool        `json:"ok"`
	Description string      `json:"description"`
	ErrorCode   int64       `json:"error_code"`
	Result      interface{} `json:"result"`
}

type ReplyNotOk Reply

func (r *ReplyNotOk) Error() string {
	if r.OK {
		return "no error"
	}
	return fmt.Sprintf("reply not ok: %s (code = %d)", r.Description, r.ErrorCode)
}

func (b *Bot) Get(loc string, q Query, result interface{}) error {
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

func (b *Bot) OnUpdate(uh UpdateHandler) {
	b.UpdateHandlers = append(b.UpdateHandlers, uh)
}

// Main is the main loop.
func (b *Bot) Main() {
	var offset int64
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
		if len(updates) > 0 {
			offset = updates[len(updates)-1].UpdateID + 1
		}
		if len(b.UpdateHandlers) > 0 {
			for _, update := range updates {
				for _, uh := range b.UpdateHandlers {
					uh(b, &update)
				}
			}
		}
	}
}

func (b *Bot) getUsername() string {
	var me User
	err := b.Get("/getMe", Query{}, &me)
	if err != nil {
		log.Println("error when retrieving username:", err)
		os.Exit(1)
	}
	return *me.Username
}

func NewBot(token string) *Bot {
	bot := Bot{"", APIURLBase + token, nil, 0}
	bot.Name = bot.getUsername()
	return &bot
}
