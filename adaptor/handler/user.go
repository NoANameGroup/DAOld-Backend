package handler

import (
	"github.com/NoANameGroup/DAOld-Backend/adaptor"
	"github.com/NoANameGroup/DAOld-Backend/adaptor/dto/user"
	"github.com/NoANameGroup/DAOld-Backend/infra/consts/consts"
	"github.com/NoANameGroup/DAOld-Backend/provider"
	"github.com/gin-gonic/gin"
)

// Register .
// @router /api/users/register [POST]
func Register(c *gin.Context) {
	var err error
	var req user.RegisterReq
	var resp *user.RegisterResp

	if err = c.ShouldBindJSON(&req); err != nil {
		adaptor.PostProcess(c, &req, resp, err)
		return
	}

	tokenStr, _ := user.ExtractToken(c.Request.Header)
	c.Set(consts.ContextUserID, tokenStr)

	resp, err = provider.Get().UserService.Register(c, &req)
	adaptor.PostProcess(c, &req, resp, err)
}
