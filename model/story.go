package model

import "gorm.io/gorm"

type StoryHighlight struct {
	gorm.Model
	Cover   string
	Name    string
	Stories []Story `gorm:"foreignKey:StoryHighlightRefer"`
}

type Story struct {
	gorm.Model
	StoryHighlightRefer uint
	Photo               string
	Video               string
	VideoDuration       int
	StoryMentions       []StoryMention   `gorm:"foreignKey:StoryRefer"`
	StoryQuestionBox    StoryQuestionBox `gorm:"foreignKey:StoryRefer"`
	//A belongs to association sets up a one-to-one connection with another models
	UserID uint
	User   User
}

type StoryMention struct {
	gorm.Model
	StoryRefer uint
	Username   string
	X          int     `gorm:"default:0"`
	Y          int     `gorm:"default:0"`
	Rotation   int     `gorm:"default:0"`
	Zoom       float32 `gorm:"default:0.0"`
	Color      int     `gorm:"default:0"`
}

type StoryQuestionBox struct {
	gorm.Model
	StoryRefer uint
	Question   string
	X          int     `gorm:"default:0"`
	Y          int     `gorm:"default:0"`
	Rotation   int     `gorm:"default:0"`
	Zoom       float32 `gorm:"default:0.0"`
	Color      int     `gorm:"default:0"`
}
