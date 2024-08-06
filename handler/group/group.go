package group

import (
	"github.com/sslime336/bot-486/bot"
	"github.com/sslime336/bot-486/bot/command"
	"github.com/sslime336/bot-486/logging"
	"github.com/sslime336/bot-486/service/mhw"
	"github.com/tencent-connect/botgo/dto"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
	log = logging.Logger().Named("handler.group")
}

func Handler(event *dto.WSPayload, data []byte) error {
	do := bot.ExtractGroupMessage(data)
	log.Debug("received atGroupMessage", zap.Any("model.GroupAtMessage", *do))

	msgId := do.MsgId()
	groupOpenId := do.GroupOpenId()
	memberOpenId := do.D.Author.MemberOpenid

	if cmd, ok := bot.Subaru.ParseCommand(do.RawContent()); ok {
		switch cmd {
		case command.Help:
			return mhw.Help(groupOpenId, msgId)
		case command.SignUp:
			return mhw.SignUp(groupOpenId, memberOpenId, msgId, do.Content())
		case command.PartyCode:
			return mhw.PartyCode(groupOpenId,memberOpenId, msgId, do.Content())
		case command.CreatePartyCode:
			return mhw.CreatePartyCode(groupOpenId, msgId, do.Content())
		case command.DeletePartyCode:
			return mhw.DeletePartyCode(groupOpenId, msgId, do.Content())
		case command.Prefile:
			return mhw.Prefile(groupOpenId, msgId)
		}
	}
	return nil
}
