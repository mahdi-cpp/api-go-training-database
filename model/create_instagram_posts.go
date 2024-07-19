package model

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/mahdi-cpp/api-go-gallery/config"
//	"image"
//	_ "image/jpeg"
//	"io"
//	"log"
//	"os"
//	"path/filepath"
//	"sort"
//	"strings"
//)
//
//var username = "instagram"
//
////var directory = "/home/mahdi/.local/bin/" + username
//var directory = "/var/instagram/" + username
//
//var user User
//var posts []Post
//
//type Person struct {
//	Name string
//}
//
//func CreateProfile() {
//
//	var files []string
//
//	user.ID = 9 //temp for test
//
//	root := directory
//	err := filepath.Walk(root, func(directory string, info os.FileInfo, err error) error {
//		files = append(files, directory)
//		return nil
//	})
//	if err != nil {
//		panic(err)
//	}
//	for _, file := range files {
//		if strings.Contains(file, "profile_pic.jpg") {
//			user.Avatar = filepath.Base(file)
//			break
//		}
//	}
//	for _, file := range files {
//		if strings.Contains(file, username+"_") {
//			// Open a file
//			file, err := os.Open(file)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			// decompress buffer and write output to stdout
//			r, err := xz.NewReader(file)
//			if err != nil {
//				log.Fatalf("NewReader error %s", err)
//			}
//
//			json := new(strings.Builder)
//			_, err = io.Copy(json, r)
//
//			var str = json.String()
//			//fmt.Println(str)
//
//			var startIndex = strings.Index(str, "\"full_name\":")
//
//			var fullName = substr(str, startIndex, 13+40)
//			//fmt.Println(fullName)
//
//			var endIndex = strings.Index(fullName, ",")
//			//fmt.Println(endIndex)
//
//			fullName = substr(fullName, 13, endIndex-14)
//			fmt.Println("full_name:", fullName)
//
//			startIndex = strings.Index(str, "\"biography\":")
//			endIndex = strings.Index(str, "\",\"blocked_by_viewer\"")
//			fmt.Println(startIndex)
//			fmt.Println(endIndex)
//			var biography = substr(str, startIndex+12, endIndex-startIndex-11)
//			fmt.Println("biography:", biography)
//
//			user.Username = username
//			user.FullName = fullName
//			user.Biography = biography
//
//			config.DB.Create(&user) // pass pointer of data to Create
//			fmt.Println("user id:", user.ID)
//			break
//		}
//	}
//
//	for _, file := range files {
//		if strings.Contains(file, ".txt") {
//			buf, err := os.ReadFile(file)
//			if err != nil {
//				panic(err)
//			}
//			var name = filepath.Base(file)
//			name = strings.ReplaceAll(name, ".txt", "")
//
//			post := Post{}
//			post.UserID = user.ID
//			post.Caption = string(buf)
//
//			posts = append(posts, post)
//			config.DB.Create(&post)
//
//			var photos []string
//
//			for _, file := range files {
//				if strings.Contains(file, name) && strings.Contains(file, ".jpg") {
//					var photo = filepath.Base(file)
//					photo = strings.ReplaceAll(photo, ".jpg", "")
//					photos = append(photos, photo)
//				}
//			}
//
//			sort.Strings(photos)
//			fmt.Println(photos)
//
//			for _, photo := range photos {
//				var hasVideo = false
//				for _, file := range files {
//					if strings.Contains(file, photo+".mp4") {
//						hasVideo = true
//					}
//				}
//				width, height := getImageDimension(directory + "/" + photo + ".jpg")
//				fmt.Println("Width:", width, "Height:", height)
//
//				media := Media{}
//				media.PostRefer = post.ID
//				media.Photo = photo + ".jpg"
//				media.Width = width
//				media.Height = height
//				if hasVideo {
//					media.Video = photo + ".mp4"
//				}
//
//				config.DB.Create(&media)
//				config.DB.Create(&models.Tag{Name: "@mahdiabdolmaleki", MediaRefer: media.ID, X: 600, Y: 250})
//			}
//		}
//	}
//}
//
//func getImageDimension(imagePath string) (int, int) {
//	file, err := os.Open(imagePath)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "%v\n", err)
//	}
//
//	image, _, err := image.DecodeConfig(file)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
//	}
//	return image.Width, image.Height
//}
//
//func addProfileRoutes(rg *gin.RouterGroup) {
//	posts := rg.Group("/profile")
//	posts.GET("/create", func(context *gin.Context) {
//		CreateProfile()
//		//context.JSON(210, CreateProfile())
//	})
//}
//
////Original
//func substr(s string, pos, length int) string {
//	bytes := []byte(s)
//	l := pos + length
//	if l > len(bytes) {
//		l = len(bytes)
//	}
//	return string(bytes[pos:l])
//}
