package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Location  string  `json:"location"`
	Caption   string  `json:"caption"`
	HasVideo  bool    `json:"hasVideo"`
	Likes     uint    `json:"likes"`
	ChatId    uint    `json:"ChatId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	User      User    `gorm:"type:json;serializer:json" json:"user"`
	Medias    []Media `gorm:"type:json;serializer:json" json:"medias"`
}

type Media struct {
	Thumbnail     string `json:"thumbnail"`
	Photo         string `json:"photo"`
	Video         string `json:"video"`
	VideoDuration int    `json:"videoDuration"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	Tags          []Tag  `gorm:"type:json;serializer:json" json:"tags"`
}

type Tag struct {
	Username string  `json:"username"`
	Dx       float32 `json:"dx"`
	Dy       float32 `json:"dy"`
}
