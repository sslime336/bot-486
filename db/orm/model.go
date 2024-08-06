package orm

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GroupImpart struct {
	Model

	// GroupOpenId QQ群OpenId
	GroupOpenId string

	Imparts []Impart
}

// Impart QQ群内怪猎集会
type Impart struct {
	Model

	Code    string
	Hunters []Hunter

	GroupImpartId uint
}

type Hunter struct {
	Model

	ImpartId uint

	// GroupOpenId QQ群OpenId
	GroupOpenId string

	// MemberOpenId QQ群成员OpenId
	MemberOpenId string

	// MHWId 怪物猎人ID
	MHWId string
}
