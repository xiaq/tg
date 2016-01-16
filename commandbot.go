package tg

type CommandHandler func(b *CommandBot, args string, msg *Message)
type AnyCommandHandler func(b *CommandBot, cmd, args string, msg *Message)

// CommandBot only handles commands.
type CommandBot struct {
	*Bot
	Handlers    map[string][]CommandHandler
	AnyHandlers []AnyCommandHandler
}

func Split(text string, sep byte) (string, string) {
	var i, j int
	for i = 0; i < len(text) && text[i] != sep; i++ {
	}
	for j = i; j < len(text) && text[j] == sep; j++ {
	}
	return text[0:i], text[j:len(text)]
}

func (b *CommandBot) HandleUpdate(_ *Bot, up *Update) {
	if up.Message == nil {
		return
	}
	if up.Message.Text == nil {
		return
	}
	text := *up.Message.Text
	if len(text) == 0 || text[0] != '/' {
		// Not a command
		return
	}
	cmd, args := Split(text[1:len(text)], ' ')
	cmd, to := Split(cmd, '@')
	if to != "" && to != b.Name {
		// Not addressed to me
		return
	}
	for _, ach := range b.AnyHandlers {
		ach(b, cmd, args, up.Message)
	}
	chs, ok := b.Handlers[cmd]
	if !ok {
		return
	}
	for _, ch := range chs {
		ch(b, args, up.Message)
	}
}

func (b *CommandBot) OnCommand(cmd string, ch CommandHandler) {
	b.Handlers[cmd] = append(b.Handlers[cmd], ch)
}

func (b *CommandBot) OnAnyCommand(ach AnyCommandHandler) {
	b.AnyHandlers = append(b.AnyHandlers, ach)
}

func NewCommandBot(name, token string) *CommandBot {
	b := &CommandBot{NewBot(name, token),
		make(map[string][]CommandHandler), nil}
	b.OnUpdate(b.HandleUpdate)
	return b
}
