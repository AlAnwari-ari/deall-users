package model

type CreateUser struct {
	RoleID   uint64 `json:"role_id" binding:"required,numeric"`
	Username string `json:"username" binding:"required,alphanum"`
	Fullname string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required,decryptiontext"`
	Email    string `json:"email" binding:"email"`
}

type UpdateUser struct {
	UserID   uint64 `json:"user_id" binding:"required,numeric"`
	RoleID   uint64 `json:"role_id" binding:"omitempty,uuid4"`
	Username string `json:"username" binding:"omitempty,required"`
	Fullname string `json:"fullname" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email"`
}

type User struct {
	UserID   uint64 `json:"user_id" uri:"user_id" binding:"numeric" gorm:"primaryKey;autoIncrement:true"`
	RoleID   uint64 `json:"role_id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	RoleName string `json:"role_name" gorm:"->"`
	Password string `json:"-"`
	TimeAt
}

func (User) TableName() string {
	return "users"
}
