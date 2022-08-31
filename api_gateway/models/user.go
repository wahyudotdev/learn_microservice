package models

type User struct {
	Id       string `json:"id" form:"id" gorm:"primaryKey,column:id"`
	Email    string `json:"email" form:"email" gorm:"column:email"`
	Name     string `json:"name" form:"name" gorm:"column:name"`
	Phone    string `json:"phone" form:"phone" gorm:"column:phone"`
	Password string `json:"-" form:"password" gorm:"column:password"`
	Role     string `json:"role" form:"role" gorm:"column:role"`
	Photo    string `json:"photo,omitempty" form:"photo" gorm:"column:photo"`
	Token    string `json:"token,omitempty" form:"token" gorm:"-"`
	IsActive bool   `json:"is_active" gorm:"column:is_active"`
	Imei     string `json:"imei" gorm:"column:imei"`
}

func (User) TableName() string {
	return "user"
}

type NewUser struct {
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Imei     string `json:"imei" form:"imei"`
}

func (NewUser) TableName() string {
	return "user"
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserUpdate struct {
	Name  string `json:"name" form:"name" gorm:"column:name"`
	Phone string `json:"phone" form:"phone" gorm:"column:phone"`
}

type UserUpdatePassword struct {
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
}

type SearchUserQuery struct {
	Page  int64  `json:"page" query:"page"`
	Query string `json:"q" query:"q"`
}

type EnableOrDisableUser struct {
	UserId   string `json:"user_id" form:"user_id"`
	IsActive bool   `json:"is_active" form:"is_active"`
}

type UpdateImei struct {
	UserId string `json:"user_id" form:"user_id"`
	Imei   string `json:"imei" form:"imei"`
}
