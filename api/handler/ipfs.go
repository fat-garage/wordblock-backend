package handler

import (
	"github.com/fat-garage/wordblock-backend/api/request"
	"github.com/fat-garage/wordblock-backend/api/response"
	"github.com/fat-garage/wordblock-backend/internal/service/block"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Upload(c *gin.Context) {
	var param request.UploadBlock
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, errors.New("invalid param"))
		return
	}
	cid, err := block.Srv.AddBlock(c.Request.Context(), param.DID, param.Content)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, cid)
	return
}

func GetDIDCidList(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		response.Fail(c, errors.New("invalid param"))
		return
	}
	list, err := block.Srv.GetDIDBlockList(cid)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, list)
	return
}
