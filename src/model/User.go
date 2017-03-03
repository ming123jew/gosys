package model

import (
	. "common"
	"log"

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

	var user []User
	err := Orm.Find(&user)
	return user, err
}

//根据分页查询
func (self *User)SelectByPage(pageSize int,start int,cache bool)([]User,error)  {
	var user []User
	var err error
	key := "User_SelectByPage_"+string(start)

	//读取缓存
	data := CacheGetList(key)
	if data != nil && cache==true{
		user = data.([]User)
		log.Println("beforeok:",data)
	}

	//缓存不存在则读取数据库
	if data==nil{

		after := func() {
			if cache==true  {
				//写入缓存
				CachePutList(key,user,CACHE_TIME_OUT)
				log.Println("afterok")
			}
		}
		err = Orm.Limit(pageSize,start).Find(&user)
		after()
	}

	return user,err
}

func (self *User) Count() (int64,error) {
	var user User
	return Orm.Count(user)
}

func (self *User) GetUser() (*User,error)  {
	user := &User{}
	_, err := Orm.Id(self.Id).Get(user)
	return user,err
}

func (self *User) GetUserById(id int) (*User, error) {
	user := &User{Id: id}
	_, err := Orm.Get(user)
	return user, err
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
