package bot

import (
	"encoding/json"

	"github.com/sslime336/bot-486/logging"
	"github.com/sslime336/bot-486/model"
	"go.uber.org/zap"
)

func ExtractPrivateChatMessage(data []byte) *model.PrivateChatMessage {
	var msg model.PrivateChatMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		logging.Error("failed to unmarshal PrivateChatMessage", zap.Error(err))
		return nil
	}
	return &msg
}

func ExtractGroupMessage(data []byte) *model.GroupAtMessage {
	var msg model.GroupAtMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		logging.Error("failed to unmarshal GroupAtMessage", zap.Error(err))
		return nil
	}
	return &msg
}
