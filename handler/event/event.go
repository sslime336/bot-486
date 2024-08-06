package event

const (
	GuildCreate           = "GUILD_CREATE"
	GuildUpdate           = "GUILD_UPDATE"
	GuildDelete           = "GUILD_DELETE"
	ChannelCreate         = "CHANNEL_CREATE"
	ChannelUpdate         = "CHANNEL_UPDATE"
	ChannelDelete         = "CHANNEL_DELETE"
	GuildMemberAdd        = "GUILD_MEMBER_ADD"
	GuildMemberUpdate     = "GUILD_MEMBER_UPDATE"
	GuildMemberRemove     = "GUILD_MEMBER_REMOVE"
	MessageCreate         = "MESSAGE_CREATE"
	MessageReactionAdd    = "MESSAGE_REACTION_ADD"
	MessageReactionRemove = "MESSAGE_REACTION_REMOVE"
	AtMessageCreate       = "AT_MESSAGE_CREATE"
	PublicMessageDelete   = "PUBLIC_MESSAGE_DELETE"
	DirectMessageCreate   = "DIRECT_MESSAGE_CREATE"
	DirectMessageDelete   = "DIRECT_MESSAGE_DELETE"
	AudioStart            = "AUDIO_START"
	AudioFinish           = "AUDIO_FINISH"
	AudioOnMic            = "AUDIO_ON_MIC"
	AudioOffMic           = "AUDIO_OFF_MIC"
	MessageAuditPass      = "MESSAGE_AUDIT_PASS"
	MessageAuditReject    = "MESSAGE_AUDIT_REJECT"
	MessageDelete         = "MESSAGE_DELETE"
	ForumThreadCreate     = "FORUM_THREAD_CREATE"
	ForumThreadUpdate     = "FORUM_THREAD_UPDATE"
	ForumThreadDelete     = "FORUM_THREAD_DELETE"
	ForumPostCreate       = "FORUM_POST_CREATE"
	ForumPostDelete       = "FORUM_POST_DELETE"
	ForumReplyCreate      = "FORUM_REPLY_CREATE"
	ForumReplyDelete      = "FORUM_REPLY_DELETE"
	ForumAuditResult      = "FORUM_PUBLISH_AUDIT_RESULT"
	InteractionCreate     = "INTERACTION_CREATE"

	C2CMessageCreate     = "C2C_MESSAGE_CREATE"      // 用户单聊发消息给机器人时候
	GroupAtMessageCreate = "GROUP_AT_MESSAGE_CREATE" // 用户在群里@机器人时收到的消息
)
