package models

import (
	"time"
)

type Talker struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Talk struct {
		WatchedAt time.Time `json:"watched_at"`
		Rate      int8      `json:"rate"`
	} `json:"talk"`
}
