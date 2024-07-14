package orm

// GatewayUser TODO: 分离?完善鉴权?
type GatewayUser struct {
	// 网关用户ID, 全局唯一
	AccountID uint32 `gorm:"primary_key;not_null;uniqueIndex;auto_increment"`
	// 用户名
	Username string
}
