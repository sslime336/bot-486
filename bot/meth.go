package bot

import (
	"context"
	"fmt"
)

func (bot *SubaruBot) Message(msg Message) error {
	_, err := bot.Transport(context.Background(), "POST", bot.hostUrl, msg)
	return err
}

// Send 发送文本消息
func (bot *SubaruBot) Send(msg string) error {
	return bot.Message(TextMessage(msg, ""))
}

// Reply 发送被动文本消息
func (bot *SubaruBot) Reply(msgId, msg string) error {
	return bot.Message(TextMessage(msg, msgId))
}

func (bot *SubaruBot) ToUser(userOpenId string) *SubaruBot {
	bot.hostUrl = fmt.Sprintf(hostUrlTemplate.User, userOpenId)
	return bot
}

func (bot *SubaruBot) ToGroup(groupOpenId string) *SubaruBot {
	bot.hostUrl = fmt.Sprintf(hostUrlTemplate.Group, groupOpenId)
	return bot
}
