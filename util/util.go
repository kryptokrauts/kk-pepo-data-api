package util

import (
	"encoding/json"
	"strconv"

	"../model"
)

// Marshall struct
func Marshall(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}

// Transform transforms an updateEvent and a contributionEvent into model PepoVideo
func Transform(updateEvent model.EventPayload, contributionEvent *model.EventPayload) model.PepoVideo {
	var totalContributors int64 = 0
	var totalContributionAmount = "0.00"
	if contributionEvent != nil {
		totalContributors = contributionEvent.Data.Activity.Video.TotalContributors
		totalContributionAmount = contributionEvent.Data.Activity.Video.TotalContributionAmount
	}
	actor := updateEvent.Data.Users[strconv.FormatInt(updateEvent.Data.Activity.ActorID, 10)]
	creator := model.Creator{
		ID:                 actor.ID,
		Name:               actor.Name,
		ProfileImage:       actor.ProfileImage,
		TokenholderAddress: actor.TokenholderAddress,
		TwitterHandle:      actor.TwitterHandle,
		GithubHandle:       actor.GithubLogin,
	}
	pepoVideo := model.PepoVideo{
		ID:                      updateEvent.Data.Activity.Video.ID,
		LastModified:            updateEvent.CreatedAt,
		Creator:                 creator,
		URL:                     updateEvent.Data.Activity.Video.URL,
		VideoURL:                updateEvent.Data.Activity.Video.VideoURL,
		TotalContributors:       totalContributors,
		TotalContributionAmount: totalContributionAmount,
		Description:             updateEvent.Data.Activity.Video.Description,
		PosterImage:             updateEvent.Data.Activity.Video.PosterImage,
		Tags:                    updateEvent.Data.Activity.Video.Tags,
	}
	return pepoVideo
}
