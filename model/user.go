package model

type User struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Avatar     string   `json:"avatar"`
	IsVerified bool     `json:"IsVerified"`
	Count      float32  `json:"count"`
	UserInfo   UserInfo `gorm:"type:json;serializer:json" json:"userInfo"`
	Persian    Persian  `gorm:"type:json;serializer:json" json:"Persian"`
}

type UserInfo struct {
	Website   string  `json:"website"`
	Biography string  `json:"biography"`
	FullName  string  `json:"fullName"`
	Age       int     `json:"age"`
	City      string  `json:"city"`
	Height    float32 `json:"height"`
	Weight    float32 `json:"weight"`
}

type Persian struct {
	Persian1 []string `json:"persian1"`
	Persian2 []string `json:"persian2"`
	Type1    string   `json:"type1"`
	Type2    string   `json:"type2"`
}
