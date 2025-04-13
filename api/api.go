package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kratos69/url-shortner/db/sqlc"
	"github.com/kratos69/url-shortner/util"
)

type request struct {
	URL string `json:"url" binding:"required,url"`
}

func (server *Server) handleShorten(ctx *gin.Context) {
	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errResponse(err),
		})
		return
	}

	code := util.GenerateShortCode(6)

	arg := db.CreateURLParams{
		Code:        code,
		OriginalUrl: req.URL,
	}

	_, err := server.store.CreateURL(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": errResponse(err),
		})
		return
	}

	shortURL := fmt.Sprintf("http://localhost:8080/kratos69/%s", code)

	ctx.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
}

func (server *Server) handleRedirect(ctx *gin.Context) {
	codeReq := ctx.Param("code")

	exists, err := server.store.CheckCodeExists(ctx, codeReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	if exists != 1 {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	urlInDB, err := server.store.GetURLByCode(ctx, codeReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Redirect(http.StatusFound, urlInDB.OriginalUrl)
}
