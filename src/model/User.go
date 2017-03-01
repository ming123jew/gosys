package model

import (
	"log"
	. "common"
)



type User struct {
	Id       int    ` int(10) NOT NULL`
	Username string ` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL`
	Password string ` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL`
	Regtime  int    ` int(10) NOT NULL`
	Status   int    ` tinyint(2) NOT NULL DEFAULT 1`
	Groupid  int    ` int(5) NOT NULL DEFAULT 0`
	Email    string ` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL`
	Roleid   string ` varchar(20) NOT NULL DEFAULT 0`
}

//查询所有列表，不建议使用
func (self *User) SelectAll() ([]User, error) {

	var chatuser []User
	err := Orm.Find(&chatuser)
	return chatuser, err
}

func (self *User) Count() (int64,error) {
	var chatuser User
	return Orm.Count(chatuser)
}

func (self *User) GetUser() (*User,error)  {
	chatuser := &User{}
	_, err := Orm.Id(self.Id).Get(chatuser)
	return chatuser,err
}

func (self *User) GetUserById(id int) (*User, error) {
	chatuser := &User{Id: id}
	_, err := Orm.Get(chatuser)
	return chatuser, err
}

func (self *User) Exist() (bool, error) {

	return Orm.Get(self)

}

func (self *User) ExistEmail() (bool, error) {
	log.Print(self.Email)
	return Orm.Get(&User{Email: self.Email})

}

type UserSession struct {
	Username string ` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL`
}
