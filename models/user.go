package models

type User struct {
	Id          int64  `gorm: "primaryKey;autoIncrement" json: "id"`
	NamaLengkap string `json: "nama_lengkap"`
	Username    string `json: "username"`
	Password    string `json: "password"`
}
