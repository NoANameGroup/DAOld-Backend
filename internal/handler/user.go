package handler

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/consts"
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/user"
	"github.com/NoANameGroup/DAOld-Backend/internal/errorx"
	"github.com/NoANameGroup/DAOld-Backend/internal/jwt"
	"github.com/NoANameGroup/DAOld-Backend/internal/provider"
	"github.com/NoANameGroup/DAOld-Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
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

// UpdateMyProfile .
// @router /api/users/me [PATCH]
func UpdateMyProfile(c *gin.Context) {
	var err error
	var req user.UpdateMyProfileReq
	var resp *user.UpdateMyProfileResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().UserService.UpdateMyProfile(c, &req)
	response.PostProcess(c, &req, resp, err)
}

// ChangePassword .
// @router /api/users/me/password [PATCH]
func ChangePassword(c *gin.Context) {
	var err error
	var req user.ChangePasswordReq
	var resp *user.ChangePasswordResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().UserService.ChangePassword(c, &req)
	response.PostProcess(c, &req, resp, err)
}

// DeleteAccount .
// @router /api/users/me [DELETE]
func DeleteAccount(c *gin.Context) {
	var err error
	var req user.DeleteAccountReq
	var resp *user.DeleteAccountResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().UserService.DeleteAccount(c, &req)
	response.PostProcess(c, &req, resp, err)
}

// Logout .
// @router /api/users/logout [POST]
func Logout(c *gin.Context) {
	var err error
	var resp *user.LogoutResp

	resp, err = provider.Get().UserService.Logout()
	response.PostProcess(c, nil, resp, err)
}

// UpdateUserRole .
// @router /api/users/:userId/role [PATCH]
func UpdateUserRole(c *gin.Context) {
	var err error
	var req user.UpdateUserRoleReq
	var resp *user.UpdateUserRoleResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	targetId, err := bson.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		response.PostProcess(c, &req, resp, errorx.New(10007, "invalid user id format"))
		return
	}

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	c.Set(consts.ContextTargetID, targetId)
	resp, err = provider.Get().UserService.UpdateUserRole(c, &req)
	response.PostProcess(c, &req, resp, err)
}
