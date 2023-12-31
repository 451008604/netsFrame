// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sqlmodel

const TableNameHouseAccount = "house_account"

// HouseAccount mapped from table <house_account>
type HouseAccount struct {
	ID       int32  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:唯一ID" json:"id"`      // 唯一ID
	UserID   int32  `gorm:"column:user_id;type:int;not null;comment:用户ID" json:"user_id"`                 // 用户ID
	Account  string `gorm:"column:account;type:varchar(64);not null;comment:账号" json:"account"`           // 账号
	Password string `gorm:"column:password;type:varchar(128);not null;comment:MD5加密后的密码" json:"password"` // MD5加密后的密码
}

// TableName HouseAccount's table name
func (*HouseAccount) TableName() string {
	return TableNameHouseAccount
}
