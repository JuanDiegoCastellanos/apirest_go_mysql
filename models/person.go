package models

type Person struct {

	// `gorm:"primary_key;auto_increment"`
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"lastname"`
	Doc          string `json:"doc"`
	Doc_Type     string `json:"doc_type"`
	Bith_Date    string `json:"birth_date"`
	Phone_Number string `json:"phone_number"`
	Url_Image    string `json:"url_image"`
	User_Id      int64  `json:"user_id"`
}
