package handler

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/elder"
	"github.com/NoANameGroup/DAOld-Backend/internal/provider"
	"github.com/NoANameGroup/DAOld-Backend/internal/response"
	"github.com/NoANameGroup/DAOld-Backend/pkg/consts"
	"github.com/NoANameGroup/DAOld-Backend/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func CreateElder(c *gin.Context) {
	var err error
	var resp *elder.CreateElderResp

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().ElderService.CreateElder(c)
	response.PostProcess(c, nil, resp, err)
}

func GetMyElder(c *gin.Context) {
	var err error
	var resp *elder.GetMyElderResp

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().ElderService.GetMyElder(c)
	response.PostProcess(c, nil, resp, err)
}

func UpdateMyElder(c *gin.Context) {
	var err error
	var req elder.UpdateMyElderReq
	var resp *elder.UpdateMyElderResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, nil, err)
		return
	}

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().ElderService.UpdateMyElder(c, &req)
	response.PostProcess(c, &req, resp, err)
}

func DeleteMyElder(c *gin.Context) {
	var err error
	var resp *elder.DeleteMyElderResp

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().ElderService.DeleteMyElder(c)
	response.PostProcess(c, nil, resp, err)
}
