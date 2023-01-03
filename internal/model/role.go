package model

type Role struct {
	RoleID    uint64 `json:"role_id" uri:"role_id" binding:"numeric" gorm:"primaryKey;autoIncrement:true"`
	RoleName  string `json:"role_name"`
	CanAdd    bool   `json:"-"`
	CanUpdate bool   `json:"-"`
	CanRead   bool   `json:"-"`
	CanDelete bool   `json:"-"`
	TimeAt
}

func (Role) TableName() string {
	return "roles"
}
