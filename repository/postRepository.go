package repository

import "github.com/mahdi-cpp/api-go-gallery/model"

func PostStart() {
	err := db.AutoMigrate(&model.Post{})
	if err != nil {
		return
	}

	for i := 0; i < 26; i++ {
		CreatUPosts()
	}

	//UpdatePost(4)
}

func UpdatePost(ID uint) {
	post := model.Post{}
	db.First(&post, ID)
	post.User.Avatar = "photos/new/ali.jpg"

	info := model.UserInfo{
		Website:   "www.helium.ir",
		Biography: "home cloud service",
		FullName:  "Sara Malek",
		Age:       42,
		City:      "Tehran",
		Height:    165,
		Weight:    56,
	}
	post.User.UserInfo = info

	db.Save(&post)
}

func GetPosts() []model.Post {
	var posts []model.Post
	db.Offset(0).Limit(050).Find(&posts)
	return posts
}

func CreatUPosts() {

	post1 := model.Post{
		Location: "Tehran",
		Caption:  "Mahdi",
		HasVideo: true,
		Likes:    32,
		ChatId:   456,
		Latitude: 32.4578,
		User: model.User{
			Name:       "Mahdi",
			Email:      "mahdi.cpp@gmail.com",
			Phone:      "09355512326",
			IsVerified: true,
			Avatar:     "photos/1234_0000_avatar.jpg",
			Count:      45.897,
			UserInfo: model.UserInfo{
				Website:   "www.mahdi.abdolmaleki.ir",
				Biography: "nothing",
				FullName:  "Mahdi abdolmaleki",
				Age:       30,
				City:      "Tehran",
				Height:    170,
				Weight:    70,
			},
			Persian: model.Persian{
				Type1:    "Verb",
				Type2:    "Noun",
				Persian1: []string{"mahdi", "ali", "reza"},
				Persian2: []string{"sara", "maryam", "leyla"},
			},
		},
		Medias: []model.Media{
			{Thumbnail: "photos/photo_125.jpg", Photo: "photos/thumbnail/photo_138.jpg", Video: "videos/video_45.mp4", VideoDuration: 135, Width: 1080, Height: 2217,
				Tags: []model.Tag{
					{Username: "@mahdi.45", Dx: 45, Dy: 100},
					{Username: "@sara.2008", Dx: 45, Dy: 100},
					{Username: "@maryam.2008", Dx: 345, Dy: 500},
				},
			},
			{Thumbnail: "photos/photo_226.jpg", Photo: "photos/thumbnail/photo_158.jpg", Video: "videos/video_45.mp4", VideoDuration: 135, Width: 1080, Height: 2217},
			{Thumbnail: "photos/photo_229.jpg", Photo: "photos/thumbnail/photo_178.jpg", Video: "videos/video_45.mp4", VideoDuration: 135, Width: 1080, Height: 2217},
			{Thumbnail: "photos/photo_328.jpg", Photo: "photos/thumbnail/photo_153.jpg", Video: "videos/video_45.mp4", VideoDuration: 135, Width: 1080, Height: 2217,
				Tags: []model.Tag{
					{Username: "@jac.45", Dx: 45, Dy: 100},
					{Username: "@messi.37", Dx: 45, Dy: 100},
					{Username: "@alvarez.55", Dx: 345, Dy: 500},
				},
			},
		},
	}

	post2 := model.Post{
		Location: "Hamedan",
		Caption:  "Mahyar Eskandari",
		HasVideo: true,
		Likes:    32,
		ChatId:   456,
		Latitude: 32.4578,
		User: model.User{
			Name:       "Mahyar",
			Email:      "mahyar.0935@gmail.com",
			Phone:      "09125640293",
			IsVerified: true,
			Avatar:     "photos/1234_0000_avatar.jpg",
			Count:      45.897,
			UserInfo: model.UserInfo{
				Website:   "www.mahdi.abdolmaleki.ir",
				Biography: "nothing",
				FullName:  "Mahdi abdolmaleki",
				Age:       30,
				City:      "Hamedan",
				Height:    170,
				Weight:    70,
			},
			Persian: model.Persian{
				Type1:    "Verb",
				Type2:    "Noun",
				Persian1: []string{"mahdi", "ali", "reza"},
				Persian2: []string{"sara", "maryam", "leyla"},
			},
		},
	}

	db.Create(&post1)
	db.Create(&post2)
}

//func QueryPosts() []model.Post {

//var posts []model.Post
//db.Find(&posts).Limit(10)
//return posts

//for _, u := range posts {
//var p = u.Persians
//fmt.Println(u.Name
//fmt.Println(p)

//persian1 := model.Persian{}
//err := PrettyJson(u.Persians["persian1"], &persian1)
//if err != nil {
//	fmt.Println(err)
//	c.JSON(400, err.Error())
//	return
//}

//}
