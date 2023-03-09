package models

type InputChannel struct {
	YoutubeId string `json:"channelYoutubeId" binding:"required"`
	IsForeign bool   `json:"channelIsForeign"`
}
