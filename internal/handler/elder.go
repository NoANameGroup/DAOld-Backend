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

func DeleteMyElder(c *gin.Context) {
	var err error
	var resp *elder.CreateElderResp

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().ElderService.CreateElder(c)
	response.PostProcess(c, nil, resp, err)
}

func UpdateMyElder(c *gin.Context) {
	var err error
	var resp *elder.UpdateMyElderResp

	c.Set(consts.ContextUserID, jwt.ExtractUserIDFromContext(c))
	resp, err = provider.Get().ElderService.UpdateMyElder(c)
	response.PostProcess(c, nil, resp, err)
}
