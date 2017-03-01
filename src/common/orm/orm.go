package orm

import (
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

func NewOrm(server string,username string,password string,db_name string,charset string,port string) *xorm.Engine  {
	log.Printf("db initializing...")
	var err error
	orm, err := xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":"+port+")/"+db_name+"?charset="+charset+"&parseTime=true")
	//Orm, err = xorm.NewEngine("mysql", "root:@/chat?charset=utf8&parseTime=true")
	//fmt.Print(username+":"+password+"@tcp("+server+")/"+db_name+"?charset="+charset+"&parseTime=true")
	if err != nil {
		log.Println(err)
	}
	orm.TZLocation = time.Local
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "gosys_")
	orm.SetTableMapper(tbMapper)
	return orm
}



