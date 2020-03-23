package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"../httputil"
)

// GetVideos godoc
// @Summary Get videos
// @Description get videos of the kryptokrauts community on Pepo sorted by creation date
// @Tags videos
// @Accept  json
// @Produce  json
// @Param limit query int false "amount of videos to return maximal"
// @Success 200 {array} model.PepoVideo
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /videos [get]
func (c *Controller) GetVideos(ctx *gin.Context) {
	limit := ctx.Request.URL.Query().Get("limit")
	n, err := strconv.ParseInt(limit, 10, 64)
	if err == nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
	}
	pepoVideos, err := c.Service.GetPepoVideos(n)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, pepoVideos)
}
