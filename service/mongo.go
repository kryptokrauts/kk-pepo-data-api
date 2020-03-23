package service

import (
	"fmt"

	"../model"
)

// GetPepoVideos returns list of Pepo videos
func (s *Service) GetPepoVideos(limit int64) ([]model.PepoVideo, error) {
	// TODO get contributions & videos
	// transform to PepoVideo
	fmt.Println("get pepo videos ...")
	return nil, nil
}
