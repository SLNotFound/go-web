package models

type Department struct {
	ID int `gorm:"primary_key" json:"id"`

	PID         int    `json:"pid"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	ManagerId   int    `json:"managerId"`
	ManagerName string `json:"managerName"`
	Introduce   string `json:"introduce"`
	CreateTime  string `json:"createTime"`
}

func GetDepts(maps interface{}) (depts []Department) {
	db.Where(maps).Find(&depts)
	return
}

func AddDept(deptCode, deptName, introduce string, deptPID, managerId int) bool {
	db.Create(&Department{
		Code:      deptCode,
		Name:      deptName,
		ManagerId: managerId,
		Introduce: introduce,
		PID:       deptPID,
	})
	return true
}
