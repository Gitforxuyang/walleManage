package dao


type Admin struct {
	Id int32 `xorm:"id"`
	Name string `xorm:"name"`
	UserName string `xorm:"user_name"`
	Password string `xorm:"password"`
	Token string `xorm:"token"`
}
