package consts

// 数据库相关
const ()

// JWT 相关
const (
	JWTSecret = "your-secret-key" // 可以从环境变量读取更安全
)

// 业务相关
const (
	ContextUserID   = "userId"
	ContextTargetID = "targetId"
)

// 数据库相关
const (
	ID          = "_id"
	UserID      = "userId"
	Email       = "email"
	Phone       = "phone"
	Password    = "password"
	Status      = "status"
	Role        = "role"
	CreatedAt   = "createdAt"
	UpdatedAt   = "updatedAt"
	Birthday    = "birthday"
	Gender      = "gender"
	Avatar      = "avatar"
	Bio         = "bio"
	Address     = "address"
	Username    = "username"
	FirstName   = "firstName"
	LastName    = "lastName"
	LastLoginAt = "lastLoginAt"
)
