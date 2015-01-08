package models

import (
	//"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//"os"
	//"path"
	//"strconv"
	"strings"
	//"time"
)

//用户基本信息表
type User struct {
	Id       int64
	Username string
	Nickname string
	Password string
	// Authority string
	Sex       string
	Identity  string //身份
	Native    string //籍贯
	Nation    string //民族
	Number    string //身份证号
	Education string //学历
	Phone     string
	Email     string
	Address   string //家庭住址
	Resident  string //户口所在地
	// Birthday   time.Time //出生日期
	// Jointime   time.Time `orm:"index"` //入党时间
	Birthday   string
	Jointime   string
	Company    string //原工作单位
	Position   string //党内职务
	Group      string //隶属党小组
	Branch     string //隶属党支部
	Category   string //党员类别
	Linkman    string //培养联系人
	Introducer string //入党介绍人
	Feature    string `orm:"size(1000)"` //有何特长
	Comment    string `orm:"size(1000)"` //备注
	Photo      string //照片

}

// 登录时检查密码是否正确
// 正确返回true
// 错误返回false
func CheckAccount(name, pwd1 string) bool {
	// 生成orm对象
	o := orm.NewOrm()
	// 生成一个新User对象，并将Username赋值为传入的name
	// User{}，就是一个新的User类型变量
	// i := 0       0 叫做 整型字面量
	// "hello world"  叫做 字符串字面量
	// 1 + 1i         叫做 复数字面量
	// User{}         叫做 User类型的字面量，User类型的0值
	// User{Username: name}
	// user的类型是User的指针类型
	user := &User{Nickname: name}

	// 执行了：
	// rset = select * from user where username=name
	// rset[0].id => user.Id
	// rset[0].username => user.Username
	// .....

	err := o.Read(user, "Nickname")
	log.Println(user)
	if err != nil {
		log.Print(err)
		return false
	}
	pwd := user.Password
	ok := strings.EqualFold(pwd, pwd1)
	return ok
}
