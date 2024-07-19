package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Location string `json:"location"`
	Caption  string `json:"caption"`
	ChatId   uint   `json:"ChatId"`
	User     User   `gorm:"type:json;serializer:json" json:"user"`
	Map      Map    `gorm:"type:json;serializer:json" json:"map"`
	Photo    Photo  `gorm:"type:json;serializer:json" json:"photo"`
	Video    Video  `gorm:"type:json;serializer:json" json:"video"`
}
type Map struct {
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Photo struct {
	Photo     string `json:"photo"`
	Thumbnail string `json:"thumbnail"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Tags      []Tag  `gorm:"type:json;serializer:json" json:"tags"`
}
type Video struct {
	Thumbnail     string `json:"thumbnail"`
	Photo         string `json:"photo"`
	Video         string `json:"video"`
	VideoDuration int    `json:"videoDuration"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	Tags          []Tag  `gorm:"type:json;serializer:json" json:"tags"`
}
type Shape struct {
	Username string  `json:"username"`
	Dx       float32 `json:"dx"`
	Dy       float32 `json:"dy"`
	Width    float32 `json:"width"`
	Height   float32 `json:"height"`
}
