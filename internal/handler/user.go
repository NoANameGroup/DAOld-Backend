package handler

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/user"
	"github.com/NoANameGroup/DAOld-Backend/internal/provider"
	"github.com/NoANameGroup/DAOld-Backend/internal/response"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/errorx"
	"github.com/NoANameGroup/DAOld-Backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// CreateUser .
// @router /api/users/register [POST]
func CreateUser(c *gin.Context) {
	var err error
	var req user.CreateUserReq
	var resp *user.CreateUserResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	resp, err = provider.Get().UserService.CreateUser(c, &req)
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

// UpdateMyPassword .
// @router /api/users/me/password [PATCH]
func UpdateMyPassword(c *gin.Context) {
	var err error
	var req user.UpdateMyPasswordReq
	var resp *user.UpdateMyPasswordResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().UserService.UpdateMyPassword(c, &req)
	response.PostProcess(c, &req, resp, err)
}

// DeleteMyAccount .
// @router /api/users/me [DELETE]
func DeleteMyAccount(c *gin.Context) {
	var err error
	var req user.DeleteMyAccountReq
	var resp *user.DeleteMyAccountResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().UserService.DeleteMyAccount(c, &req)
	response.PostProcess(c, &req, resp, err)
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

//// UpdateMyEmail .
//// @router /api/users/me/email [PATCH]
//func UpdateMyEmail(c *gin.Context) {
//	var err error
//	var req user.UpdateMyEmailReq
//	var resp *user.UpdateMyEmailResp
//
//	if err = c.ShouldBindJSON(&req); err != nil {
//		response.PostProcess(c, &req, resp, err)
//		return
//	}
//
//	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
//	resp, err = provider.Get().UserService.UpdateMyEmail(c, &req)
//	response.PostProcess(c, &req, resp, err)
//}

//// UpdateMyPhone .
//// @router /api/users/me/phone [PATCH]
//func UpdateMyPhone(c *gin.Context) {
//	var err error
//	var req user.UpdateMyPhoneReq
//	var resp *user.UpdateMyPhoneResp
//
//	if err = c.ShouldBindJSON(&req); err != nil {
//		response.PostProcess(c, &req, resp, err)
//		return
//	}
//
//	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
//	resp, err = provider.Get().UserService.UpdateMyPhone(c, &req)
//	response.PostProcess(c, &req, resp, err)
//}

// UpdateMyUsername .
// @router /api/users/me/username [PATCH]
func UpdateMyUsername(c *gin.Context) {
	var err error
	var req user.UpdateMyUsernameReq
	var resp *user.UpdateMyUsernameResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().UserService.UpdateMyUsername(c, &req)
	response.PostProcess(c, &req, resp, err)
}
