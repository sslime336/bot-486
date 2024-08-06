package bot

type Message struct {
	Content string `json:"content"`
	MsgType int    `json:"msg_type"`
	MsgId   string `json:"msg_id"` // 带上该字段后发送消息变成被动消息，限制比主动消息少
}

func TextMessage(msg, msgId string) Message {
	return Message{
		Content: msg,
		MsgType: 0,
		MsgId:   msgId,
	}
}
