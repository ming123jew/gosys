package model

import (
	. "common"
)


func init()  {
	err := Orm.Sync2(new(Group))
	if err != nil{

	}
}
type Group struct {
	Id	int	`xorm:"int(11) pk autoincr"`
	Name	string	`xorm:"varchar(25) unique  notnull"`
	Num 	int	`xorm:"int(11) notnull"`
	Ctime   int	`xorm:"created notnull"`

}
