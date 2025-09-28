package handler

import (
	"github.com/NoANameGroup/DAOld-Backend/consts"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/user"
	"github.com/NoANameGroup/DAOld-Backend/internal/jwt"
	"github.com/NoANameGroup/DAOld-Backend/internal/provider"
	"github.com/NoANameGroup/DAOld-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

// Register .
// @router /api/users/register [POST]
func Register(c *gin.Context) {
	var err error
	var req user.RegisterReq
	var resp *user.RegisterResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	resp, err = provider.Get().UserService.Register(c, &req)
	response.PostProcess(c, &req, resp, err)
}

// Login .
// @router /api/users/login [POST]
func Login(c *gin.Context) {
	var err error
	var req user.LoginReq
	var resp *user.LoginResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	resp, err = provider.Get().UserService.Login(c, &req)
	response.PostProcess(c, &req, resp, err)
}

// GetMyProfile .
// @router /api/users/me [GET]
func GetMyProfile(c *gin.Context) {
	var err error
	var resp *user.GetMyProfileResp

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().UserService.GetMyProfile(c)
	response.PostProcess(c, nil, resp, err)
}

func ChangePassword(c *gin.Context) {
	var err error

}
