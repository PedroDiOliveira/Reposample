package main

import (
	"errors"
	"fmt"
)

type Video struct {
	ID            int
	Duration      int
	Entertainment int
}

func escolherVideo(videos []Video, time int) (Video, error) {

	var VideoIdeal Video = Video{}
	for _, videoFeed := range videos {
		if videoFeed.Duration > time {
			time--
			continue
		}
		time--

		if videoFeed.Entertainment > VideoIdeal.Entertainment {
			VideoIdeal = videoFeed
		}
	}
	if VideoIdeal == (Video{}) {
		return Video{}, errors.New("erro.")
	}

	fmt.Println(VideoIdeal)

	return VideoIdeal, nil
}
