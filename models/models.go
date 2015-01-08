package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	//"log"
	"os"
	"path"
	"strconv"
	//"strings"
	"time"
)

const (
	_DB_NAME        = "data/party.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//党员类别表
type Category struct {
	Id          int64
	Title       string `orm:"index"` //党员类别
	Created     time.Time
	PeopleCount int64 //人数

}

//权限表
type Authority struct {
	Id       int64
	Rolename string //角色名
	Username string
	Comment  string `orm:"size(1000)"` //备注
}

//党费缴纳情况表
type Pay struct {
	Id            int64
	Year          int64  //年度
	Firstquarter  int64  //第一季度
	Secondquarter int64  //第二季度
	Thirdquarter  int64  //第三季度
	Fourthquarter int64  //第四季度
	Total         int64  //全年总计
	Comment       string //备注

}

//参加活动情况表
type Activity struct {
	Id          int64
	Atime       time.Time //活动时间
	Aname       string    //活动名称
	Adepartment string    //组织单位
	Acontent    string    //活动内容

}

//流入情况表
type Inflow struct {
	Id      int64
	Iadress string    //流入地详细地址
	Itime   time.Time //流入时间
	Icard   string    //持证情况
	Comment string    //备注

}

// 流出情况表
type Outflow struct {
	Id      int64
	Oadress string    //流出地详细地址
	Otime   time.Time //流出时间
	Ocard   string    //持证情况
	Comment string    //备注
}

func RegisterDB() {
	// 检查数据库文件
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	// 注册模型
	orm.RegisterModel(new(User), new(Authority), new(Pay), new(Activity), new(Inflow), new(Outflow), new(Category))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

// 注册新用户
func AddUser(nickname, password string) error {
	o := orm.NewOrm()
	user1 := &User{
		Nickname: nickname,
		Password: password,
	}

	qs := o.QueryTable("user")
	// 如果没有找到，则err不为nil，为nil时表示已经找到
	err := qs.Filter("nickname", nickname).One(user1)
	if err == nil {
		return err
	}
	//插入数据
	_, err = o.Insert(user1)
	if err != nil {
		return err
	}
	return nil

}

func AddInformation(uid, username, sex, identity, native, nation, number, education, phone, email, address, resident, company, position, group, branch, category, linkman, introducer, feature, comment, photo, birthday, jointime string) error {

	uidNum, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	user := &User{Id: uidNum}
	if o.Read(user) == nil {

		user.Username = username
		user.Sex = sex
		user.Identity = identity
		user.Native = native
		user.Nation = nation
		user.Number = number
		user.Education = education
		user.Phone = phone
		user.Email = email
		user.Address = address
		user.Resident = resident
		user.Birthday = birthday
		user.Jointime = jointime
		user.Company = company
		user.Position = position
		user.Group = group
		user.Branch = branch
		user.Category = category
		user.Linkman = linkman
		user.Introducer = introducer
		user.Feature = feature
		user.Comment = comment
		user.Photo = photo
		_, err = o.Update(user)
		if err != nil {
			return err
		}
	}

	return nil

}
func GetUser(uid string) (*User, error) {
	uidNum, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	user := new(User)

	qs := o.QueryTable("user")
	err = qs.Filter("id", uidNum).One(user)
	if err != nil {
		return nil, err
	}
	return user, err
}
func GetAllUsers(identy, cate string) ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	if len(identy) > 0 {
		qs = qs.Filter("identity", identy)
	}
	if len(cate) > 0 {
		qs = qs.Filter("category", cate)
	}
	_, err := qs.All(&users)
	return users, err

}

func ModifyUser(uid, username, sex, identity, native, nation, number, education, phone, email, address, resident, company, position, group, branch, category, linkman, introducer, feature, comment, photo, birthday, jointime string) error {

	uidNum, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return err
	}

	var oldPhoto string
	o := orm.NewOrm()
	user := &User{Id: uidNum}
	if o.Read(user) == nil {
		oldPhoto = user.Photo
		user.Username = username
		user.Sex = sex
		user.Identity = identity
		user.Native = native
		user.Nation = nation
		user.Number = number
		user.Education = education
		user.Phone = phone
		user.Email = email
		user.Address = address
		user.Resident = resident
		user.Birthday = birthday
		user.Jointime = jointime
		user.Company = company
		user.Position = position
		user.Group = group
		user.Branch = branch
		user.Category = category
		user.Linkman = linkman
		user.Introducer = introducer
		user.Feature = feature
		user.Comment = comment
		user.Photo = photo
		_, err = o.Update(user)
		if err != nil {
			return err
		}
	}
	// 删除旧照片
	if len(oldPhoto) > 0 {
		os.Remove(path.Join("photo", oldPhoto))
	}
	return nil

}

func DeleteUser(uid string) error {
	// TODO:删除后应该更新人数统计
	uidNum, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	user := &User{Id: uidNum}
	if o.Read(user) == nil {
		_, err = o.Delete(user)
		if err != nil {
			return err
		}
	}
	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{
		Title:   name,
		Created: time.Now(),
	}

	// 查询数据
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
