package database

import "time"

func Migrate() {
	db := ConnectToDatabase()
	db.AutoMigrate(&Product{}, &Option{}, &Image{}, &Description{}, &Variant{}, &User{}, &Order{}, &OrderDetail{}, &Bill{})
}

type Description struct {
	ID        int    `json:"id"`
	IDProduct int    `json:"id_product" gorm:"default:null"`
	Content   string `json:"content" gorm:"type:mediumtext"`
}

type Image struct {
	ID        int    `json:"id"`
	IDProduct int    `json:"id_product" gorm:"default:null"`
	Link      string `json:"link" gorm:"type:mediumtext"`
}

type Option struct {
	ID        int     `json:"id"`
	IDProduct int     `json:"id_product" gorm:"default:null"`
	Size      string  `json:"size" gorm:"type:tinytext"`
	Price     int     `json:"price" gorm:"type:int"`
	SalePrice int     `json:"sale_price" gorm:"type:int"`
	Quantity  int     `json:"quantity" gorm:"type:mediumint"`
	Variant   Variant `gorm:"foreignKey:IDOption;ASSOCIATION_FOREIGNKEY:ID"`
}

type Product struct {
	ID               int           `json:"id"`
	Name             string        `json:"name" gorm:"type:tinytext"`
	LinkDetail       string        `json:"link_detail" gorm:"type:mediumtext"`
	Technology       string        `json:"technology" gorm:"type:tinytext"`
	Resolution       string        `json:"resolution" gorm:"type:tinytext"`
	Type             string        `json:"type" gorm:"type:tinytext"`
	ES               bool          `json:"es"`
	ListDescriptions []Description `gorm:"foreignKey:IDProduct;ASSOCIATION_FOREIGNKEY:ID"`
	ListImages       []Image       `gorm:"foreignKey:IDProduct;ASSOCIATION_FOREIGNKEY:ID"`
	ListOptions      []Option      `gorm:"foreignKey:IDProduct;ASSOCIATION_FOREIGNKEY:ID"`
}

type Variant struct {
	ID               int           `json:"id"`
	IDProduct        int           `json:"id_product" gorm:"default:null"`
	IDOption         int           `json:"id_option" gorm:"default:null"`
	ProductName      string        `json:"product_name" gorm:"type:tinytext"`
	Size             string        `json:"size" gorm:"type:tinytext"`
	Price            int           `json:"price" gorm:"type:int"`
	SalePrice        int           `json:"sale_price" gorm:"type:int"`
	ListOrderDetails []OrderDetail `gorm:"foreignKey:IDVariant;ASSOCIATION_FOREIGNKEY:ID"`
}

type OrderDetail struct {
	ID          int    `json:"id"`
	IDVariant   int    `json:"id_variant" gorm:"default:null"`
	IDOrder     int    `json:"id_order"`
	Quantity    int    `json:"quantity" gorm:"type:mediumint"`
	ProductName string `json:"product_name" gorm:"type:tinytext"`
	TotalPrice  int    `json:"total_price" gorm:"type:int"`
}

type Order struct {
	ID               int           `json:"id"`
	IDUser           int           `json:"id_user" gorm:"default:null"`
	TotalPrice       int           `json:"total_price" gorm:"type:int"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	DeleteAt         time.Time     `json:"delete_at"`
	Username         string        `json:"username" gorm:"type:tinytext"`
	PhoneUser        string        `json:"phone_user" gorm:"type:tinytext"`
	AddressUser      string        `json:"address_user" gorm:"type:tinytext"`
	State            string        `json:"state" gorm:"type:tinytext"`
	Bill             Bill          `gorm:" foreignkey: IDOrder "`
	ListOrderDetails []OrderDetail `gorm:"foreignKey:IDOrder;ASSOCIATION_FOREIGNKEY:ID"`
}

type User struct {
	ID         int     `json:"id"`
	Name       string  `json:"name" gorm:"type:tinytext"`
	Username   string  `json:"username" gorm:"type:tinytext"`
	Password   string  `json:"password" gorm:"type:tinytext"`
	Email      string  `json:"email" gorm:"type:tinytext"`
	Address    string  `json:"address" gorm:"type:mediumtext"`
	Phone      string  `json:"phone" gorm:"type:tinytext"`
	IsAdmin    bool    `json:"isadmin"`
	ListOrders []Order `gorm:"foreignKey:IDUser;ASSOCIATION_FOREIGNKEY:ID"`
	ListBills  []Bill  `gorm:"foreignKey:IDUser;ASSOCIATION_FOREIGNKEY:ID"`
}

type Bill struct {
	ID          int       `json:"id"`
	IDOrder     int       `json:"id_order"`
	IDUser      int       `json:"id_user"`
	IdAdmin     int       `json:"id_admin"`
	Username    string    `json:"username" gorm:"type:tinytext"`
	AdminName   string    `json:"admin_name" gorm:"type:tinytext"`
	PhoneUser   string    `json:"phone_user" gorm:"type:tinytext"`
	TotalPrice  int       `json:"total_price" gorm:"type:int"`
	AddressUser string    `json:"address_user" gorm:"type:tinytext"`
	State       string    `json:"state" gorm:"type:tinytext"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
