package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Order{},
		&User{},

		&PaymentMethod{},
		&DeliveryType{},
		&Payment{},

		&Return{},
		&Staff{},
	)

	db = database

	password1, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123580"), 14)

	db.Model(&User{}).Create(&User{
		Name:     "Phatcha",
		Email:    "Phatcha@gmail.com",
		Password: string(password1),
	})
	db.Model(&User{}).Create(&User{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(password2),
	})

	var Phatcha User
	var name User
	db.Raw("SELECT * FROM users WHERE email = ?", "Phatcha@gmail.com").Scan(&Phatcha)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)

	// --- Order Data
	order1ofPhatcha := Order{
		PreorderID: 4001,
		Statusorder: "Confirm",
		Owner: Phatcha,
	}
	db.Model(&Order{}).Create(&order1ofPhatcha)

	order2ofPhatcha := Order{
		PreorderID: 4002,
		Statusorder: "Confirm",
		Owner: Phatcha,
	}
	db.Model(&Order{}).Create(&order2ofPhatcha)

	order1ofName := Order{
		PreorderID: 4003,
		Statusorder: "Confirm",
		Owner: name,
	}
	db.Model(&Order{}).Create(&order1ofName)

	order2ofName := Order{
		PreorderID: 4004,
		Statusorder: "Confirm",
		Owner: name,
	}
	db.Model(&Order{}).Create(&order2ofName)

	// DeliveryType Data
	type1 := DeliveryType{
		Type: "รับสินค้าที่ร้าน",
	}
	db.Model(&DeliveryType{}).Create(&type1)

	type2 := DeliveryType{
		Type: "จัดส่งถึงบ้าน",
	}
	db.Model(&DeliveryType{}).Create(&type2)


	// PaymentMethod Data
	Method1 := PaymentMethod{
		Method: "Bank",
	}
	db.Model(&PaymentMethod{}).Create(&Method1)

	Method2 := PaymentMethod{
		Method: "Promtpay",
	}
	db.Model(&PaymentMethod{}).Create(&Method2)

	Method3 := PaymentMethod{
		Method: "จ่ายสดหน้าร้าน",
	}
	db.Model(&PaymentMethod{}).Create(&Method3)

	Method4 := PaymentMethod{
		Method: "เก็บเงินปลายทาง",
	}
	db.Model(&PaymentMethod{}).Create(&Method4)

	//จอย

	passwordReturn, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	// ข้อมูล user
	db.Model(&User{}).Create(&User{
		Name:     "Narudee Arunno",
		Email:    "narudee@gmail.com",
		Password: string(passwordReturn),
	})
	db.Model(&User{}).Create(&User{
		Name:     "Nana Lanana",
		Email:    "nana@gmail.com",
		Password: string(passwordReturn),
	})

	var narudee User
	var nana User
	db.Raw("SELECT * FROM users WHERE email = ?", "narudee@gmail.com").Scan(&narudee)
	db.Raw("SELECT * FROM users WHERE email = ?", "nana@gmail.com").Scan(&nana)

	// order data
	order1 := Order{
		Owner:       nana,
		PreorderID: 20,
		StatusID:   1,
		Ordertime:  time.Now(),
	}

	db.Model(&Order{}).Create(&order1)

	order2 := Order{
		Owner:       narudee,
		PreorderID: 21,
		StatusID:   1,
		Ordertime:  time.Now(),
	}
	db.Model(&Order{}).Create(&order2)

	// staff data

	db.Model(&Staff{}).Create(&Staff{
		Name:     "Suwanan",
		Email:    "suwanan@gmail.com",
		Password: string(passwordReturn),
	})

	db.Model(&Staff{}).Create(&Staff{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(passwordReturn),
	})

}