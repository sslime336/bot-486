package mhw

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sslime336/bot-486/bot"
	"github.com/sslime336/bot-486/bot/command"
	"github.com/sslime336/bot-486/db"
	"github.com/sslime336/bot-486/db/orm"
	"github.com/sslime336/bot-486/logging"
	"github.com/sslime336/bot-486/service/prefile"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	log              *zap.Logger
	persist          *gorm.DB
	DefaultAliveTime = 6 * time.Hour
)

func Init() {
	log = logging.Logger().Named("service.mhw")
	persist = db.Sqlite
}

func Help(groupOpenId, msgId string) error {
	return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, command.Brief)
}

func SignUp(groupOpenId, memberOpenId, msgId, mhwId string) error {
	var hunter orm.Hunter
	if err := persist.Where(&orm.Hunter{GroupOpenId: groupOpenId, MemberOpenId: memberOpenId}).First(&hunter).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			persist.Create(&orm.Hunter{
				GroupOpenId:  groupOpenId,
				MemberOpenId: memberOpenId,
				MHWId:        mhwId,
			})
			return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, fmt.Sprintf("MHW ID: %s 登记成功", mhwId))
		}
		log.Error("hunter sign up failed", zap.Error(err))
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "更新/登录失败")
	}
	if err := persist.Model(&hunter).Update("MHWId", mhwId).Error; err != nil {
		log.Error("hunter info update failed", zap.Error(err))
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "更新/登录失败")
	}
	return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, fmt.Sprintf("MHW ID: %s 已更新", mhwId))
}

func Prefile(groupOpenId, msgId string) error {
	return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, prefile.Info())
}

func PartyCode(groupOpenId, memberOpenId, msgId, targetPartyCode string) error {
	log := log.Named("joinImpart")

	var groupImpart orm.GroupImpart
	if err := persist.Where(&orm.GroupImpart{GroupOpenId: groupOpenId}).Preload("Imparts").First(&groupImpart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "当前群组未创建集会码")
		} else {
			log.Error("query group impart error", zap.Error(err))
			return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "查询集会码失败")
		}
	}

	if groupImpart.Imparts == nil || len(groupImpart.Imparts) == 0 {
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "当前群组未创建集会码")
	}

	var b strings.Builder
	found := -1
	for idx, impart := range groupImpart.Imparts {
		b.WriteString(fmt.Sprintf("%d. %s\n", idx+1, impart.Code))
		persist.Preload("Hunters").First(&impart)
		for _, hunter := range impart.Hunters {
			b.WriteString(fmt.Sprintf("\t- %s\n", hunter.MHWId))
		}
		if targetPartyCode == impart.Code {
			found = idx
		}
	}
	if targetPartyCode == "" {
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, fmt.Sprintf("当前群组中的集会:\n%s", b.String()))
	}

	if found == -1 {
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, fmt.Sprintf("要加入的集会 %s 不存在，当前群组中的集会:\n%s", targetPartyCode, b.String()))
	}

	var hunter orm.Hunter
	if err := persist.Where(&orm.Hunter{GroupOpenId: groupOpenId, MemberOpenId: memberOpenId}).First(&hunter).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "请先登记MHW ID")
		}
		log.Error("query hunter error", zap.Error(err))
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "Hunter信息查询失败")
	}
	if err := persist.Model(&hunter).Update("ImpartId", groupImpart.Imparts[found].Id).Error; err != nil {
		log.Error("update hunter impart id error", zap.Error(err))
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "加入集会失败")
	}
	return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, fmt.Sprintf("%s 已加入集会 %s", hunter.MHWId, groupImpart.Imparts[found].Code))
}

func CreatePartyCode(groupOpenId, msgId, partyCode string) error {
	var groupImpart orm.GroupImpart
	if err := persist.Where(&orm.GroupImpart{GroupOpenId: groupOpenId}).Preload("Imparts").First(&groupImpart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := persist.Create(&orm.GroupImpart{
				GroupOpenId: groupOpenId,
				Imparts:     []orm.Impart{{Code: partyCode}},
			}).Error; err != nil {
				log.Error("create group impart error", zap.Error(err))
				return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码创建失败")
			}
			return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码创建成功")
		} else {
			log.Error("query group impart error", zap.Error(err))
			return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码创建失败")
		}
	}

	if err := persist.Model(&groupImpart).Association("Imparts").Append(&orm.Impart{Code: partyCode}); err != nil {
		log.Error("update group impart error", zap.Error(err))
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码创建失败")
	}
	return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码创建成功")
}

func DeletePartyCode(groupOpenId, msgId, partyCode string) error {
	if partyCode == "" {
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "要删除的集会信息不能为空，通过 /help 查看命令帮助")
	}

	var groupImpart orm.GroupImpart
	if err := persist.Where(&orm.GroupImpart{GroupOpenId: groupOpenId}).Preload("Imparts").First(&groupImpart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "当前群组不存在集会码")
		}
		log.Error("query group impart error", zap.Error(err))
		return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码查询失败")
	}

	for _, impart := range groupImpart.Imparts {
		log.Debug("check pivot", zap.Bool("equal", impart.Code == partyCode))
		if impart.Code == partyCode {
			if err := persist.Model(&groupImpart).Association("Imparts").Delete(impart); err != nil {
				log.Error("delete impart association error", zap.Error(err))
				return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码删除失败")
			}
			if err := persist.Delete(&impart).Error; err != nil {
				log.Error("delete impart error", zap.Error(err))
				return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码删除失败")
			}
		}
	}
	return bot.Subaru.ToGroup(groupOpenId).Reply(msgId, "集会码 "+partyCode+" 删除成功")
}
