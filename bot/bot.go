package bot

import (
	"os"
	"strings"

	"github.com/sslime336/bot-486/bot/command"
	"github.com/sslime336/bot-486/logging"
	"github.com/tencent-connect/botgo/openapi"
	"go.uber.org/zap"
)

type SubaruBot struct {
	openapi.OpenAPI
	hostUrl string
}

var hostUrlTemplate struct {
	User  string
	Group string
}

var Subaru *SubaruBot

const (
	hostUser         = "https://api.sgroup.qq.com/v2/users/%s/messages"
	hostGroup        = "https://api.sgroup.qq.com/v2/groups/%s/messages"
	hostUserSandbox  = "https://sandbox.api.sgroup.qq.com/v2/users/%s/messages"
	hostGroupSandbox = "https://sandbox.api.sgroup.qq.com/v2/groups/%s/messages"
)

func BuildClient(api openapi.OpenAPI) {
	Subaru = new(SubaruBot)
	Subaru.OpenAPI = api

	hostUrlTemplate.User = hostUserSandbox
	hostUrlTemplate.Group = hostGroupSandbox

	if os.Getenv("SUBARU_MODE") == "release" {
		hostUrlTemplate.User = hostUser
		hostUrlTemplate.Group = hostGroup
	}
}

func (b *SubaruBot) ParseCommand(content string) (command.Command, bool) {
	ctnt := strings.TrimSpace(content)
	logging.Debug("bot received message", zap.String("trimed-content", content))
	fields := strings.Split(ctnt, " ")
	logging.Debug("bot received message fields", zap.Strings("message-fields", fields))

	cmd, ok := command.CommandMap[fields[0]]
	return cmd, ok
}
