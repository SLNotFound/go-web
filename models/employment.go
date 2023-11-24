package models

type User struct {
	// 登录账号
	UserLogin string `json:"userLogin"`
	// 员工姓名
	Username string `json:"username"`
	// 登录密码
	Password string `json:"password"`
	// 员工聘用形式，1正式, 2非正式
	FormOfEmployment int64 `json:"formOfEmployment"`
	// 员工工号
	WorkNumber string `json:"workNumber"`
	// 员工转正日期
	CorrectionTime string `json:"correctionTime"`
	// 员工部门id
	DepartmentID int64 `json:"departmentId"`
	// 头像信息
	StaffPhoto *string `json:"staffPhoto,omitempty"`
	// 员工入职日期，格式: 2020-01-01
	TimeOfEntry string `json:"timeOfEntry"`
}
