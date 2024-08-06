package command

var CommandMap = map[string]Command{
	"/help":    Help,
	"/signup":  SignUp,
	"/p":       PartyCode,
	"/cp":      CreatePartyCode,
	"/dp":      DeletePartyCode,
	"/prefile": Prefile,
}

type Command int

const (
	Unknown Command = iota

	// Help 查看帮助信息
	Help

	// SignUp 登记 MHW 游戏昵称
	SignUp

	// PartyCode 查看当前QQ群中的集会码信息，或者加入已存在的集会
	PartyCode

	// CreatePartyCode 创建集会码，若已经存在则更新集会码创建时间
	CreatePartyCode

	// DeletePartyCode 删除集会码
	DeletePartyCode

	// Prefile 查看机器人运行环境信息
	Prefile
)

const Brief = `bot-486 可用指令：
/help    查看指令详细信息
/signup  绑定 MHW 中的昵称到当前账号，如: /signup Kokona
/p       查看群中已登记的集会码
/cp      创建新的群集会码，如: /cp v7A3eDYPk7=t 
/dp      删除群集会码，如: /dp e4n27?G7BQcR
/prefile 查看机器人运行状态
`