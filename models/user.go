package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap string `gorm:"varchar(300);not null" json:"nama_lengkap`
	Username    string `gorm:"varchar(300);not null" json:"username"`
	Password    string `gorm:"varchar(300);not null" json:"password"`
}
