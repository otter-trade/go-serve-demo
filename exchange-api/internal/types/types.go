// Code generated by goctl. DO NOT EDIT.
package types

type AdminCommonResp struct {
}

type Admin struct {
	Id         int64  `json:"id"`          //id
	RoleId     int64  `json:"role_id"`     //角色组
	Account    string `json:"account"`     //用户名
	Nickname   string `json:"nickname"`    //昵称
	Mobile     string `json:"mobile"`      //手机号
	Password   string `json:"password"`    //密码
	Salt       string `json:"salt"`        //密码盐
	Avatar     string `json:"avatar"`      //头像
	Email      string `json:"email"`       //邮箱
	LoginFail  int64  `json:"login_fail"`  //登录失败次数
	LoginTime  int64  `json:"login_time"`  //登录时间
	LoginIp    string `json:"login_ip"`    //登录 ip
	Status     string `json:"status"`      //状态
	StatusText string `json:"status_text"` //状态描述
	CreateTime int64  `json:"create_time"` //创建时间
	UpdateTime int64  `json:"update_time"` //更新时间
}

type AdminSearchReq struct {
	Page       int64  `json:"page,optional"`        //第几页
	PageSize   int64  `json:"page_size,optional"`   //每页多少条
	OrderField string `json:"order_field,optional"` //排序字段
	OrderParam string `json:"order_param,optional"` //排序方法
	StartTime  int64  `json:"start_time,optional"`  //开始时间
	EndTime    int64  `json:"end_time,optional"`    //结束时间
}

type AdminSearchResp struct {
	Total int64    `json:"total"` // 总数
	List  []*Admin `json:"list"`  // 数据
}

type AdminUpdateReq struct {
	Id        int64  `json:"id"`         //id
	RoleId    int64  `json:"role_id"`    //角色组
	Account   string `json:"account"`    //用户名
	Nickname  string `json:"nickname"`   //昵称
	Mobile    string `json:"mobile"`     //手机号
	Password  string `json:"password"`   //密码
	Avatar    string `json:"avatar"`     //头像
	Email     string `json:"email"`      //邮箱
	LoginFail int64  `json:"login_fail"` //登录失败次数
	LoginTime int64  `json:"login_time"` //登录时间
	LoginIp   string `json:"login_ip"`   //登录 ip
	Status    string `json:"status"`     //状态
}

type AdminDeleteReq struct {
	Id int64 `json:"id" validate:"required"` //id
}

type AdminAddReq struct {
	RoleId    int64  `json:"role_id"`             //角色组
	Account   string `json:"account"`             //用户名
	Nickname  string `json:"nickname"`            //昵称
	Mobile    string `json:"mobile"`              //手机号
	Password  string `json:"password"`            //密码
	Salt      string `json:"salt,optional"`       //密码盐
	Avatar    string `json:"avatar"`              //头像
	Email     string `json:"email"`               //邮箱
	LoginFail int64  `json:"login_fail,optional"` //登录失败次数
	LoginTime int64  `json:"login_time,optional"` //登录时间
	LoginIp   string `json:"login_ip,optional"`   //登录 ip
	Status    string `json:"status"`              //状态
}

type AdminDetailReq struct {
	Id int64 `form:"id" validate:"required"` //id
}

type AdminDetailResp struct {
	Id         int64  `json:"id"`          //id
	RoleId     int64  `json:"role_id"`     //角色组
	Account    string `json:"account"`     //用户名
	Nickname   string `json:"nickname"`    //昵称
	Mobile     string `json:"mobile"`      //手机号
	Password   string `json:"password"`    //密码
	Salt       string `json:"salt"`        //密码盐
	Avatar     string `json:"avatar"`      //头像
	Email      string `json:"email"`       //邮箱
	LoginFail  int64  `json:"login_fail"`  //登录失败次数
	LoginTime  int64  `json:"login_time"`  //登录时间
	LoginIp    string `json:"login_ip"`    //登录 ip
	Status     string `json:"status"`      //状态
	StatusText string `json:"status_text"` //状态描述
	CreateTime int64  `json:"create_time"` //创建时间
	UpdateTime int64  `json:"update_time"` //更新时间
}
