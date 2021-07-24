package database

import "time"

func Migrate() {
	db := ConnectToDatabase()
	db.AutoMigrate(&Product{}, &Option{}, &Image{}, &Description{}, &Variant{}, &User{}, &Order{}, &OrderDetail{}, &Bill{})
}

type Description struct {
	ID         int    `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	ID_Product int    `json:"id_product" gorm:"default:null; type:mediumint"`
	Content    string `json:"content" gorm:"type:mediumtext"`
}

type Image struct {
	ID         int    `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	ID_Product int    `json:"id_product" gorm:"default:null; type:mediumint"`
	Link       string `json:"link" gorm:"type:mediumtext"`
}

type Option struct {
	ID           int       `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	ID_Product   int       `json:"id_product" gorm:"default:null; type:mediumint"`
	Size         string    `json:"size" gorm:"type:tinytext"`
	Price        string    `json:"price" gorm:"type:tinytext"`
	Sale_Price   string    `json:"sale_price" gorm:"type:tinytext"`
	Quantity     int       `json:"quantity" gorm:"type:mediumint"`
	List_Variant []Variant `gorm:"foreignKey:ID_Option;ASSOCIATION_FOREIGNKEY:ID"`
}

type Product struct {
	ID               int           `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	Name             string        `json:"name" gorm:"type:tinytext"`
	Link_Detail      string        `json:"link_detail" gorm:"type:mediumtext"`
	Technology       string        `json:"technology" gorm:"type:tinytext"`
	Resolution       string        `json:"resolution" gorm:"type:tinytext"`
	Type             string        `json:"type" gorm:"type:tinytext"`
	List_Description []Description `gorm:"foreignKey:ID_Product;ASSOCIATION_FOREIGNKEY:ID"`
	List_Image       []Image       `gorm:"foreignKey:ID_Product;ASSOCIATION_FOREIGNKEY:ID"`
	List_Option      []Option      `gorm:"foreignKey:ID_Product;ASSOCIATION_FOREIGNKEY:ID"`
}

type Variant struct {
	ID           int         `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	ID_Product   int         `json:"id_product" gorm:"default:null; type:mediumint"`
	ID_Option    int         `json:"id_option" gorm:"default:null; type:mediumint"`
	Name         string      `json:"name" gorm:"type:tinytext"`
	Size         string      `json:"size" gorm:"type:tinytext"`
	Price        string      `json:"price" gorm:"type:tinytext"`
	Sale_Price   string      `json:"sale_price" gorm:"type:tinytext"`
	Order_Detail OrderDetail `gorm:"foreignKey:ID_Variant;ASSOCIATION_FOREIGNKEY:ID"`
}

type OrderDetail struct {
	ID           int    `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	ID_Variant   int    `json:"id_variant" gorm:"default:null; type:mediumint"`
	ID_Order     int    `json:"id_order" gorm:"type:mediumint"`
	Quantity     int    `json:"quantity" gorm:"type:mediumint"`
	Product_Name string `json:"product_name" gorm:"type:tinytext"`
	Total_Price  int    `json:"total_price" gorm:"type:mediumint"`
}

type Order struct {
	ID                int           `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	ID_User           int           `json:"id_user" gorm:"default:null; type:mediumint"`
	Total_Price       int           `json:"total_price" gorm:"type:mediumint"`
	Created_At        time.Time     `json:"created_at"`
	Updated_At        time.Time     `json:"updated_at"`
	Delete_At         time.Time     `json:"delete_at"`
	User_Name         string        `json:"user_name" gorm:"type:tinytext"`
	Phone_User        int           `json:"phone_user" gorm:"type:tinyint"`
	Address_User      string        `json:"address_user" gorm:"type:tinytext"`
	State             string        `json:"state" gorm:"type:tinytext"`
	Bill              Bill          `gorm:" foreignkey: ID_Order "`
	List_OrderDetails []OrderDetail `gorm:"foreignKey:ID_Order;ASSOCIATION_FOREIGNKEY:ID"`
}

type User struct {
	ID         int     `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	Name       string  `json:"name" gorm:"type:tinytext"`
	Username   string  `json:"username" gorm:"type:tinytext"`
	Password   string  `json:"password" gorm:"type:tinytext"`
	Email      string  `json:"email" gorm:"type:tinytext"`
	Address    string  `json:"address" gorm:"type:mediumtext"`
	Telephone  string  `json:"telephone" gorm:"type:tinyint"`
	Role       string  `json:"role" gorm:"type:tinytext"`
	List_Order []Order `gorm:"foreignKey:ID_User;ASSOCIATION_FOREIGNKEY:ID"`
	List_Bill  []Bill  `gorm:"foreignKey:ID_User;ASSOCIATION_FOREIGNKEY:ID"`
}

type Bill struct {
	ID           int       `json:"id" gorm:"PRIMARY_KEY; type:mediumint"`
	ID_Order     int       `json:"id_order" gorm:"type:mediumint"`
	ID_User      int       `json:"id_user" gorm:"type:mediumint"`
	Id_Admin     int       `json:"id_admin" gorm:"type:mediumint"`
	User_Name    string    `json:"user_name" gorm:"type:tinytext"`
	Admin_Name   string    `json:"admin_name" gorm:"type:tinytext"`
	Phone_User   int       `json:"phone_user" gorm:"type:tinyint"`
	Total_Price  int       `json:"total_price" gorm:"type:mediumint"`
	Address_User string    `json:"address_user" gorm:"type:tinytext"`
	State        string    `json:"state" gorm:"type:tinytext"`
	Created_At   time.Time `json:"created_at"`
}
