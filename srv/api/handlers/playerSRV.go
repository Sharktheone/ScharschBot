package handlers

import (
	"Scharsch-Bot/srv/playersrv"
	"Scharsch-Bot/types"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PlayerSRVEventHandler(c *gin.Context) {
	var eventJson *types.EventJson
	if err := json.NewDecoder(c.Request.Body).Decode(&eventJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to decode json: %v", err),
		})
		return
	}
	err, status, pSrv := playersrv.DecodeV2(eventJson)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err,
		})
		return
	}
	pSrv.SwitchEvents()
}
