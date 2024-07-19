package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-gallery/config"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseInit() {
	config.DatabaseInit()
	db = config.DB
}

type Person struct {
	ID        int
	Name      string
	Addresses []Address `gorm:"many2many:person_addresses;"`
	DeletedAt gorm.DeletedAt
}

type Address struct {
	ID   uint
	Name string
}

type PersonAddress struct {
	PersonID  int
	AddressID int
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func OverrideJoinTable() {

	DatabaseInit()

	db.Migrator().DropTable(&Person{}, &Address{}, &PersonAddress{})

	if err := db.SetupJoinTable(&Person{}, "Addresses", &PersonAddress{}); err != nil {
		fmt.Printf("Failed to setup join table for person, got error %v", err)
	}

	if err := db.AutoMigrate(&Person{}, &Address{}); err != nil {
		fmt.Printf("Failed to migrate, got %v", err)
	}

	address1 := Address{Name: "Mahdi address 1"}
	address2 := Address{Name: "Mahdi address 2"}
	address3 := Address{Name: "Mahdi address 3"}
	person := Person{Name: "Mahdi", Addresses: []Address{address1, address2, address3}}
	db.Create(&person)

	var addresses1 []Address
	if err := db.Model(&person).Association("Addresses").Find(&addresses1); err != nil || len(addresses1) != 2 {
		fmt.Printf("Failed to find address, got error %v, length: %v", err, len(addresses1))
	}

	if err := db.Model(&person).Association("Addresses").Delete(&person.Addresses[0]); err != nil {
		fmt.Printf("Failed to delete address, got error %v", err)
	}

	if len(person.Addresses) != 1 {
		fmt.Printf("Should have one address left")
	}

	if db.Find(&[]PersonAddress{}, "person_id = ?", person.ID).RowsAffected != 1 {
		fmt.Printf("Should found one address")
	}

	var addresses2 []Address
	if err := db.Model(&person).Association("Addresses").Find(&addresses2); err != nil || len(addresses2) != 1 {
		fmt.Printf("Failed to find address, got error %v, length: %v", err, len(addresses2))
	}

	if db.Model(&person).Association("Addresses").Count() != 1 {
		fmt.Printf("Should found one address")
	}

	var addresses3 []Address
	if err := db.Unscoped().Model(&person).Association("Addresses").Find(&addresses3); err != nil || len(addresses3) != 2 {
		fmt.Printf("Failed to find address, got error %v, length: %v", err, len(addresses3))
	}

	if db.Unscoped().Find(&[]PersonAddress{}, "person_id = ?", person.ID).RowsAffected != 2 {
		fmt.Printf("Should found soft deleted addresses with unscoped")
	}

	if db.Unscoped().Model(&person).Association("Addresses").Count() != 2 {
		fmt.Printf("Should found soft deleted addresses with unscoped")
	}

	db.Model(&person).Association("Addresses").Clear()

	if db.Model(&person).Association("Addresses").Count() != 0 {
		fmt.Printf("Should deleted all addresses")
	}

	if db.Unscoped().Model(&person).Association("Addresses").Count() != 2 {
		fmt.Printf("Should found soft deleted addresses with unscoped")
	}

	db.Unscoped().Model(&person).Association("Addresses").Clear()

	if db.Unscoped().Model(&person).Association("Addresses").Count() != 0 {
		fmt.Printf("address should be deleted when clear with unscoped")
	}

	//-----------------------------------------------
	address21 := Address{Name: "Ali address 1"}
	address22 := Address{Name: "Ali address 2"}
	person2 := Person{Name: "Ali", Addresses: []Address{address21, address22}}
	db.Create(&person2)

	//if err := db.Select(clause.Associations).Delete(&person2).Error; err != nil {
	//	fmt.Printf("failed to delete person, got error: %v", err)
	//}
	//
	//if count := db.Unscoped().Model(&person2).Association("Addresses").Count(); count != 2 {
	//	fmt.Printf("person's addresses expects 2, got %v", count)
	//}
	//
	//if count := db.Model(&person2).Association("Addresses").Count(); count != 0 {
	//	fmt.Printf("person's addresses expects 2, got %v", count)
	//}
}
