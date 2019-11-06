package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int `PK`
	Name string `orm:"size(100)"`
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
	Post []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Userinfo struct {
	Id     int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username    string
	Departname  string
	Created     time.Time
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default","mysql","root:root@tcp(localhost:3306)/beedbdemo?charset=utf8",30)
	orm.SetMaxIdleConns("default",30)
	orm.SetMaxOpenConns("default",100)
	orm.RegisterModel(new(Userinfo),new(User), new(Profile), new(Tag),new(Post))
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	profile := Profile{Age:35}
	pid,err := o.Insert(&profile)
	fmt.Println(pid,err)

	user := User{Name:"bruce1299999",Profile:&profile}
	id, err := o.Insert(&user)
	fmt.Printf("ID:%d, ERR: %v\n",id ,err)

	q_user := User{Id:3}
	err = o.Read(&q_user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到!")
	}
	fmt.Println(err)
	if err == nil{
		fmt.Println(q_user)
		fmt.Println(q_user.Id)
		fmt.Println(q_user.Name)
		q_user.Name = "bruce_update"
		num,err := o.Update(&q_user)
		if err == nil{
			fmt.Println("num ",num)
		}else{
			fmt.Println("error ", err)
		}
	}

	//var rs orm.RawSeter
	sql := `
select * 
from user 
where id > 2
order by id desc
`
	rs := o.Raw(sql)

	var users []User
	num,err :=rs.QueryRows(&users)
	fmt.Println(num,err)

	for _,f_user := range users {
		fmt.Println(f_user.Id,f_user.Name)
	}

	//users := []User{
	//	{Name:"slene"},
	//	{Name:"bruce1"},
	//	{Name:"lily"},
	//}
	//successNums, err := o.InsertMulti(100,users)
	//fmt.Println(successNums,err)
	//user.Name = "bruce1"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//u := User{Id:user.Id}
	//err = o.Read(&u)
	//fmt.Printf("ERROR: %v\n",err)
	//
	//user.Name="bruce2"
	//
	//fmt.Println(u)
	//fmt.Println("------")
	//fmt.Println(user)

	//num,err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

}
