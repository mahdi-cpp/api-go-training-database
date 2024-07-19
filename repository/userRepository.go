package repository

import "github.com/mahdi-cpp/api-go-gallery/model"

func UserStart() {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
	//CreatUsers()
	QueryUsers()
}

func CreatUsers() {
	//
	//user1 := model.User{
	//	Name:       "Mahdi",
	//	Email:      "mahdi.cpp@gmail.com",
	//	Phone:      "09355512326",
	//	IsVerified: true,
	//	Avatar:     "photos/1234_0000_avatar.jpg",
	//	UserInfo: map[string]interface{}{
	//		"website":   "www.mahdi.abdolmaleki.ir",
	//		"biography": "nothing",
	//		"fullName":  "Mahdi abdolmaleki",
	//		"age":       30,
	//		"city":      "Tehran",
	//		"height":    170,
	//		"weight":    70,
	//	},
	//	Persians: map[string]interface{}{
	//		"type1":    "Verb",
	//		"type2":    "Noun",
	//		"persian1": []string{"mahdi", "ali", "reza"},
	//		"persian2": []string{"sara", "maryam", "leyla"},
	//	},
	//}
	//
	//user2 := model.User{
	//	Name:       "Ali",
	//	Email:      "ali.456@gmail.com",
	//	Phone:      "0912563658",
	//	IsVerified: true,
	//	Avatar:     "photos/4334_0000_avatar.jpg",
	//	UserInfo: map[string]interface{}{
	//		"website":   "www.ali.nasiri.ir",
	//		"biography": "nothing",
	//		"fullName":  "Ali Nasiri",
	//		"age":       33,
	//		"city":      "Fardis",
	//		"height":    181,
	//		"weight":    86,
	//	},
	//	Persians: map[string]interface{}{
	//		"type1":    "Verb",
	//		"type2":    "Noun",
	//		"persian1": []string{"cat", "dog", "pets"},
	//		"persian2": []string{"beautiful", "nice", "cute"},
	//	},
	//}
	//
	//db.Create(&user1)
	//db.Create(&user2)
}

func GetUsers() []model.User {
	var users []model.User
	db.Where("persians->>'type1' = ?", "Verb").Find(&users)
	return users
}

func QueryUsers() []model.User {

	var users []model.User
	db.Where("persians->>'type1' = ?", "Verb").Find(&users)
	return users

	//for _, u := range users {
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

	// Convert the map to JSON
	//jsonData, err := json.Marshal(u)
	//if err != nil {
	//	//fmt.Println("Error marshaling JSON:", err)
	//} else {
	//	persian := model.Persian{}
	//	err2 := json.Unmarshal(jsonData, &persian)
	//	if err2 == nil {
	//		//fmt.Println(persian.Type2)
	//	}
	//}

	//a = append(a, string(jsonData))
	//}
}

//
//func QueryUsers2() {
//	var users []Post
//
//	db.Where("info->>'width' = ?", 11).Find(&users)
//	db.Where("info->>'width' = ?", 11).Find(&users)
//
//	fmt.Println("query result: ")
//	for _, u := range users {
//		var city = u.UserInfo["city"]
//		fmt.Println(city)
//	}
//
//	fmt.Println("-----------------------")
//}

//type UserWithJSON struct {
//	gorm.Model
//	Name       string
//	Attributes datatypes.JSON
//}
//
//func JsonQuery() {
//
//	config.db.Create(&UserWithJSON{
//		Name:       "json-1",
//		Attributes: datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
//	}
//
//	// Check JSON has keys
//	datatypes.JSONQuery("attributes").HasKey(value, keys...)
//
//	db.Find(&user, datatypes.JSONQuery("attributes").HasKey("role"))
//	db.Find(&user, datatypes.JSONQuery("attributes").HasKey("orgs", "orga"))
//	// MySQL
//
//}
