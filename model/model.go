package model

import (
	"strings"
	"time"
)

type WsDtoBase struct {
	Op int64  `json:"op"`
	S  int64  `json:"s"`
	T  string `json:"t"`
	Id string `json:"id"`
}

type General struct {
	WsDtoBase
	D any `json:"d"`
}

// PrivateChatMessage 群聊中被 @
type PrivateChatMessage struct {
	WsDtoBase
	D struct {
		Author struct {
			Id         string `json:"id"`
			UserOpenid string `json:"user_openid"`
		} `json:"author"`
		Content string `json:"content"`
		ID      string `json:"id"`
	} `json:"d"`
}

func (pm *PrivateChatMessage) UserOpenId() string {
	return pm.D.Author.UserOpenid
}

func (pm *PrivateChatMessage) MsgId() string {
	return pm.D.ID
}

func (pm *PrivateChatMessage) Content() string {
	return pm.D.Content
}

type GroupAtMessage struct {
	WsDtoBase
	D struct {
		Author struct {
			ID           string `json:"id"`
			MemberOpenid string `json:"member_openid"`
		} `json:"author"`
		Content     string    `json:"content"`
		GroupID     string    `json:"group_id"`
		GroupOpenid string    `json:"group_openid"`
		ID          string    `json:"id"`
		Timestamp   time.Time `json:"timestamp"`
	} `json:"d"`
}

func (gm *GroupAtMessage) GroupOpenId() string {
	return gm.D.GroupOpenid
}

func (gm *GroupAtMessage) MsgId() string {
	return gm.D.ID
}

// RawContent 返回已经移除前后空白字符的消息内容
func (gm *GroupAtMessage) RawContent() string {
	return strings.TrimSpace(gm.D.Content)
}

// Content 返回移除指令的消息
func (gm *GroupAtMessage) Content() string {
	txt := strings.Split(gm.RawContent(), " ")
	if len(txt) < 2 {
		return ""
	}
	return txt[1]
}
