package enum

// Status
type UserStatus int

const (
	StatusActive    UserStatus = 1 // 活跃
	StatusSuspended UserStatus = 2 // 暂停
	StatusBanned    UserStatus = 3 // 已封禁
)

var UserStatusMap = map[UserStatus]string{
	StatusActive:    "活跃",
	StatusSuspended: "暂停",
	StatusBanned:    "已封禁",
}

func GetUserStatusDesc(code UserStatus) string {
	if desc, ok := UserStatusMap[code]; ok {
		return desc
	}
	return "未知状态"
}

func GetUserStatusCode(desc string) UserStatus {
	for code, d := range UserStatusMap {
		if d == desc {
			return code
		}
	}
	return 0
}

// Gender
type UserGender int

const (
	GenderMale   UserGender = 1 // 男
	GenderFemale UserGender = 2 // 女
	GenderOther  UserGender = 3 // 其他
)

var UserGenderMap = map[UserGender]string{
	GenderMale:   "男",
	GenderFemale: "女",
	GenderOther:  "其他",
}

func GetUserGenderDesc(code UserGender) string {
	if desc, ok := UserGenderMap[code]; ok {
		return desc
	}
	return "未知"
}

func GetUserGenderCode(desc string) UserGender {
	for code, d := range UserGenderMap {
		if d == desc {
			return code
		}
	}
	return 0
}

// Role
type UserRole int

const (
	RoleAdmin UserRole = 1 // 管理员
	RoleUser  UserRole = 2 // 普通用户
)

var UserRoleMap = map[UserRole]string{
	RoleAdmin: "管理员",
	RoleUser:  "用户",
}

func GetUserRoleDesc(code UserRole) string {
	if desc, ok := UserRoleMap[code]; ok {
		return desc
	}
	return "无角色"
}

func GetUserRoleCode(desc string) UserRole {
	for code, d := range UserRoleMap {
		if d == desc {
			return code
		}
	}
	return 0
}
