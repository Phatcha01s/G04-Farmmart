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

		&Product{},
		&ProductType{},
		&Supplier{},
		&ProductStock{},
	)

	db = database

	passworduser1, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	passworduser2, err := bcrypt.GenerateFromPassword([]byte("123580"), 14)
	passwordstaff, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&User{}).Create(&User{
		Name:     "Phatcha",
		Email:    "Phatcha@gmail.com",
		Password: string(passworduser1),
	})
	db.Model(&User{}).Create(&User{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(passworduser2),
	})

	var Phatcha User
	var Name User
	db.Raw("SELECT * FROM users WHERE email = ?", "Phatcha@gmail.com").Scan(&Phatcha)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&Name)

	// --- Order Data
	order1ofPhatcha := Order{
		PreorderID: 4001,
		StatusID: 1,
		User: Phatcha,
	}
	db.Model(&Order{}).Create(&order1ofPhatcha)

	order2ofPhatcha := Order{
		PreorderID: 4002,
		StatusID: 1,
		User: Phatcha,
	}
	db.Model(&Order{}).Create(&order2ofPhatcha)

	order1ofName := Order{
		PreorderID: 4003,
		StatusID: 1,
		User: Name,
	}
	db.Model(&Order{}).Create(&order1ofName)

	order2ofName := Order{
		PreorderID: 4004,
		StatusID: 1,
		User: Name,
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

	// ข้อมูล user
	db.Model(&User{}).Create(&User{
		Name:     "Narudee Arunno",
		Email:    "narudee@gmail.com",
		Password: string(passworduser1),
	})
	db.Model(&User{}).Create(&User{
		Name:     "Nana Lanana",
		Email:    "nana@gmail.com",
		Password: string(passworduser1),
	})

	var narudee User
	var nana User
	db.Raw("SELECT * FROM users WHERE email = ?", "narudee@gmail.com").Scan(&narudee)
	db.Raw("SELECT * FROM users WHERE email = ?", "nana@gmail.com").Scan(&nana)

	// order data
	order1 := Order{
		User:       nana,
		PreorderID: 20,
		StatusID:   1,
		Ordertime:  time.Now(),
	}

	db.Model(&Order{}).Create(&order1)

	order2 := Order{
		User:       narudee,
		PreorderID: 21,
		StatusID:   1,
		Ordertime:  time.Now(),
	}
	db.Model(&Order{}).Create(&order2)

	// staff data

	db.Model(&Staff{}).Create(&Staff{
		Name:     "Suwanan",
		Email:    "suwanan@gmail.com",
		Password: string(passwordstaff),
	})

	db.Model(&Staff{}).Create(&Staff{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(passwordstaff),
	})

	//นุ่น

	// Staff Data
	db.Model(&Staff{}).Create(&Staff{
		Name:     "Suwanan",
		Email:    "suwanan@gmail.com",
		Password: string(passwordstaff),
	})

	db.Model(&Staff{}).Create(&Staff{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(passwordstaff),
	})

	var suwanan Staff
	var name Staff
	db.Raw("SELECT * FROM staffs WHERE email = ?", "suwanan@gmail.com").Scan(&suwanan)
	db.Raw("SELECT * FROM staffs WHERE email = ?", "name@example.com").Scan(&name)

	// ProductType Data
	drink := ProductType{
		Ptype: "เครื่องดื่ม",
	}
	db.Model(&ProductType{}).Create(&drink)

	vegetable := ProductType{
		Ptype: "ผัก",
	}
	db.Model(&ProductType{}).Create(&vegetable)

	fruit := ProductType{
		Ptype: "ผลไม้",
	}
	db.Model(&ProductType{}).Create(&fruit)

	meat := ProductType{
		Ptype: "เนื้อสัตว์",
	}
	db.Model(&ProductType{}).Create(&meat)

	another := ProductType{
		Ptype: "อื่นๆ",
	}
	db.Model(&ProductType{}).Create(&another)

	// --- Product Data
	mango := Product{
		Name:        "มะม่วงน้ำดอกไม้",
		ProductType: fruit,
	}
	db.Model(&Product{}).Create(&mango)

	egg := Product{
		Name:        "ไข่ไก่ No.0",
		ProductType: another,
	}
	db.Model(&Product{}).Create(&egg)

	egg_omega := Product{
		Name:        "ไข่ไก่ OMEGA-3",
		ProductType: another,
	}
	db.Model(&Product{}).Create(&egg_omega)

	sunflower_sprout := Product{
		Name:        "ต้นอ่อนทานตะวัน",
		ProductType: vegetable,
	}
	db.Model(&Product{}).Create(&sunflower_sprout)

	milk := Product{
		Name:        "นมสดพาสเจอร์ไรส์ รสจืด ขนาด 5 ลิตร",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&milk)

	milk_banana := Product{
		Name:        "นมรสกล้วย 150 ml",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&milk_banana)

	milk_lychee := Product{
		Name:        "นมรสลิ้นจี่ 150 ml",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&milk_lychee)

	milk_melon := Product{
		Name:        "นมรสเมล่อน 150 ml",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&milk_melon)

	pork := Product{
		Name:        "เนื้อหมู",
		ProductType: meat,
	}
	db.Model(&Product{}).Create(&pork)

	turkey := Product{
		Name:        "ไก่งวง 3-5 กิโลกรัม",
		ProductType: meat,
	}
	db.Model(&Product{}).Create(&turkey)

	durian := Product{
		Name:        "ทุเรียน",
		ProductType: fruit,
	}
	db.Model(&Product{}).Create(&durian)

	salad := Product{
		Name:        "ผักสลัด",
		ProductType: vegetable,
	}
	db.Model(&Product{}).Create(&salad)

	water350 := Product{
		Name:        "น้ำดื่ม ขนาด 350 ซีซี 12 ขวด",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&water350)

	water600 := Product{
		Name:        "น้ำดื่ม ขนาด 600 ซีซี 12 ขวด",
		ProductType: drink,
	}
	db.Model(&Product{}).Create(&water600)

	// --- Supplier Data
	sut := Supplier{
		Name: "มทส.",
	}
	db.Model(&Supplier{}).Create(&sut)

	student := Supplier{
		Name: "นักศึกษา",
	}
	db.Model(&Supplier{}).Create(&student)

	farmer := Supplier{
		Name: "เกษตรกร",
	}
	db.Model(&Supplier{}).Create(&farmer)

	// Stock 1
	db.Model(&ProductStock{}).Create(&ProductStock{

		Product:     milk,
		Supplier:    sut,
		Price:       10,
		Amount:      30,
		Staff:       suwanan,
		ProductTime: time.Now(),
	})
	// Stock 2
	db.Model(&ProductStock{}).Create(&ProductStock{
		Product:     mango,
		Supplier:    farmer,
		Price:       10,
		Amount:      22,
		Staff:       name,
		ProductTime: time.Now(),
	})
}