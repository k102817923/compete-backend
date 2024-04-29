package models

import (
	"database/sql/driver"
	"encoding/json"
	"paracraft-compete-backend/middleware/database/mysql"
	"time"
)

type scoreRange struct {
	Lable string
	Value int
}

type option struct {
	Key      string
	Title    string
	Required bool
}

type previewFile struct {
	Url  string
	Name string
}

type rule struct {
	Type   int
	Amount int
}

type extra map[string]interface{}
type showPlatform []int

type Compete struct {
	IModel
	OrgId           uint      `gorm:"column:orgId;not null;default:0;comment:'比赛单位id, 单一赛事填写'" json:"orgId"`
	ClassId         uint      `gorm:"column:classId;not null;default:0;comment:'比赛班级id, 单一赛事填写'" json:"classId"`
	Name            string    `gorm:"column:name;type:varchar(50);not null;comment:'名称'" json:"name"`
	Link            string    `gorm:"column:link;type:varchar(50);not null;comment:'地址'" json:"link"`
	Intro           string    `gorm:"column:intro;type:text;not null;comment:'简介'" json:"intro"`
	Icon            string    `gorm:"column:icon;type:varchar(512);not null;comment:'首页图标'" json:"icon"`
	Sponsor         string    `gorm:"column:sponsor;type:varchar(50);not null;comment:'主办方'" json:"sponsor"`
	Type            uint8     `gorm:"column:type;type:tinyint;not null;default:1;comment:'赛事类型: 默认1单一赛事, 2大赛组';check:type IN(1, 2)" json:"type"`
	ApplyStartAt    time.Time `gorm:"column:applyStartAt;not null;comment:'报名开始时间'" json:"applyStartAt"`
	ApplyEndAt      time.Time `gorm:"column:applyEndAt;not null;comment:'报名结束时间'" json:"applyEndAt"`
	ReviewStartAt   time.Time `gorm:"column:reviewStartAt;comment:'评审开始时间'" json:"reviewStartAt"`
	ReviewEndAt     time.Time `gorm:"column:reviewEndAt;comment:'评审结束时间'" json:"reviewEndAt"`
	GradeResultType uint8     `gorm:"column:gradeResultType;type:tinyint;not null;default:0;comment:'成绩结果显示类型, 默认0显示总分, 1总分转义';check gradeResultType IN(0, 1)" json:"gradeResultType"`
	// ScoreRanges     []scoreRange `gorm:"column:scoreRanges;type:json;not null;default:'[]';comment:'总分转义区间'" json:"scoreRanges"`
	StartAt time.Time `gorm:"column:startAt;not null;comment:'比赛开始时间'" json:"startAt"`
	EndAt   time.Time `gorm:"column:endAt;not null;comment:'比赛结束时间'" json:"endAt"`
	// RegionIds            []int         `gorm:"column:regionIds;type:json;not null;default:'[]';comment:'比赛区域'" json:"regionIds"`
	// Options              []option      `gorm:"column:options;type:json;not null;default:'[]';comment:'赛事报名项'" json:"options"`
	IsOpenRegistry uint8 `gorm:"column:isOpenRegistry;type:tinyint;not null;default:0;comment:'是否开启赛事报名';check isOpenRegistry IN(0, 1)" json:"isOpenRegistry"`
	// PreviewFiles         []previewFile `gorm:"column:previewFiles;type:json;not null;default:'{}';comment:'预览文件'" json:"previewFiles"`
	IsPreviewFile uint8  `gorm:"column:isPreviewFile;type:tinyint;not null;default:0;comment:'是否展示预览文件';check:isPreviewFile IN (0, 1)" json:"isPreviewFile"`
	Description   string `gorm:"column:description;type:text;comment:'测评说明'" json:"description"`
	// Rules                rule          `gorm:"column:rules;type:json;not null;default:'{}';comment:'评分规则'" json:"rules"`
	Status               uint8        `gorm:"column:status;type:tinyint;not null;default:0;comment:'赛事状态';check:status IN (0, 1, 2, 3)" json:"status"`
	IsPublished          uint8        `gorm:"column:isPublished;type:tinyint;not null;default:0;comment:'状态 0.下架, 1.上架';check:isPublished IN (0, 1)" json:"isPublished"`
	IsShow               uint8        `gorm:"column:isShow;type:tinyint;not null;default:0;comment:'是否在竞赛列表页展示';check:isShow IN (0, 1)" json:"isShow"`
	IsAssociated         uint8        `gorm:"column:isAssociated;type:tinyint;not null;default:0;commetn:'是否关联了大赛组: 默认0没关联, 1已关联大赛组';check:isAssociated IN (0, 1)" json:"isAssociated"`
	Sn                   int          `gorm:"column:sn;type:integer;not null;default:0;comment:'自定义顺序'" json:"sn"`
	Extra                extra        `gorm:"column:extra;type:json;default:'{}';comment:'额外字段'" json:"extra"`
	IsMemberJoin         uint8        `gorm:"column:isMemberJoin;type:tinyint;not null;default:0;comment:'队员是否参赛';check:isMemberJoin IN (0, 1)" json:"isMemberJoin"`
	OperatorId           uint         `gorm:"column:operatorId;type:bigint;not null;default:0;comment:'最后操作的管理员id'" json:"operatorId"`
	CreatorId            uint         `gorm:"column:creatorId;type:bigint;not null;default:0;comment:'创建人id'" json:"creatorId"`
	ShowPlatform         showPlatform `gorm:"column:showPlatform;type:json;not null;default:'[]';comment:'展示平台: 1.edu'" json:"showPlatform"`
	JoinType             uint8        `gorm:"column:joinType;type:tinyint;not null;default:1;comment:'参赛形式';check:joinType IN (1, 2, 3)" json:"joinType"`
	IsUploadFile         uint8        `gorm:"column:isUploadFile;type:tinyint;not null;default:0;comment:'是否上传附件';check:isUploadFile IN (0, 1)" json:"isUploadFile"`
	IsUploadFileRequired uint8        `gorm:"column:isUploadFileRequired;type:tinyint;not null;default:0;comment:'是否必须上传附件';check:isUploadFileRequired IN (0, 1)" json:"isUploadFileRequired"`
}

func (Compete) TableName() string {
	return "competes"
}

func (s showPlatform) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}

func (s *showPlatform) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), s)
}

func (e extra) Value() (driver.Value, error) {
	b, err := json.Marshal(e)
	return string(b), err
}

func (e *extra) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), e)
}

func GetCompeteList() (competes []Compete) {
	unDeletedTime := time.Date(1970, 1, 1, 8, 0, 0, 0, time.UTC)
	mysql.CompeteDB.Unscoped().Where("deletedAt = ? or deletedAt is null", unDeletedTime).Find(&competes)
	return
}
