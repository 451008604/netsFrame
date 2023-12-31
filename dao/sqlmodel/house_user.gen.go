// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sqlmodel

const TableNameHouseUser = "house_user"

// HouseUser mapped from table <house_user>
type HouseUser struct {
	ID           int32  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:自增ID" json:"id"`             // 自增ID
	UniID        int64  `gorm:"column:uni_id;type:bigint;not null;comment:唯一ID" json:"uni_id"`                       // 唯一ID
	Nickname     string `gorm:"column:nickname;type:varchar(18);not null;comment:昵称" json:"nickname"`                // 昵称
	HeadImage    string `gorm:"column:head_image;type:varchar(128);not null;default:0;comment:头像" json:"head_image"` // 头像
	RegisterTime int32  `gorm:"column:register_time;type:int;not null;comment:注册时间" json:"register_time"`            // 注册时间
}

// TableName HouseUser's table name
func (*HouseUser) TableName() string {
	return TableNameHouseUser
}
