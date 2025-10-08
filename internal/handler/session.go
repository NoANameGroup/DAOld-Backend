package handler

import (
	"github.com/NoANameGroup/DAOld-Backend/internal/dto/session"
	"github.com/NoANameGroup/DAOld-Backend/internal/provider"
	"github.com/NoANameGroup/DAOld-Backend/internal/response"
	"github.com/gin-gonic/gin"
)

// CreateSession .
// @router /api/sessions [POST]
func CreateSession(c *gin.Context) {
	var err error
	var req session.CreateSessionReq
	var resp *session.CreateSessionResp

	if err = c.ShouldBindJSON(&req); err != nil {
		response.PostProcess(c, &req, resp, err)
		return
	}

	resp, err = provider.Get().SessionService.CreateSession(c, &req)
	response.PostProcess(c, &req, resp, err)
}

// DeleteSession .
// @router /api/sessions [DELETE]
func DeleteSession(c *gin.Context) {
	var err error
	var resp *session.DeleteSessionResp

	resp, err = provider.Get().SessionService.DeleteSession()
	response.PostProcess(c, nil, resp, err)
}
