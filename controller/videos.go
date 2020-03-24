package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"../httputil"
)

// GetPepoVideos godoc
// @Summary Get pepo videos
// @Description get videos of the Pepo community kryptokrauts sorted by creation date
// @Tags videos
// @Accept  json
// @Produce  json
// @Param limit query int false "maximum amount of videos to return"
// @Success 200 {array} model.PepoVideo
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos [get]
func (c *Controller) GetPepoVideos(ctx *gin.Context) {
	limit := ctx.Request.URL.Query().Get("limit")
	var limitNo int64 = 0
	if limit != "" {
		n, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
		limitNo = n
	}
	pepoVideos, err := c.Service.GetPepoVideos(limitNo)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, pepoVideos)
}
