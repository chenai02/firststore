package mysql

import "time"

type Tb_File struct {
	Id int `gorm:"column:id;primaryKey;autoIncrement;not null"`
	File_sha1 string `gorm:"type:char(40);not null;default:''"`
	File_name string `gorm:"type:varchar(256);not null;unique;default:''"`//unique 规范数据的唯一性
	File_size int64	`gorm:"type: bigint(20);default:'0'"`
	File_addr string `gorm:"type:varchar(1024);not null;default:''"`
	CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充  官网上有
	UpdatedAt time.Time // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	Status int `gorm:"type:int(11);not null;default:'0'"`
	Ext1 int `gorm:"type:int(11);not null;default:'0'"`
	Ext2 int
}
type Tb_User struct {
	Id int `gorm:"column:id;primaryKey;autoIncrement;not null"`
	User_name string `gorm:"type:varchar(64);not null;unique;default:''"`
	User_pwd string `gorm:"type:varchar(256);not null;default:''"`
	Email string `gorm:"type:varchar(64);not null;default:''"`
	Phone string `gorm:"varchar(128);default:''"`
	Email_validated int `gorm:"type:tinyint(1);default:'0'"`
	Phone_validated int`gorm:"type:tinyint(1);default:'0'"`
	CreatedAt time.Time `gorm:"type:datetime;"`
	UpdatedAt time.Time `gorm:"comment:'最后活跃时间戳';type:datetime"`
	Profile string `gorm:"type:text"`
	Status int`gorm:"type:int(11);not null;default:'0'"`
	//ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
}

type Tb_User_Token struct {
	Id int `gorm:"column:id;primaryKey;autoIncrement;not null"`
	User_name string `gorm:"type:varchar(64);not null;unique;default:''"`
	User_Token string `gorm:"type:varchar(40);not null;default:''"`
}


type UserFile struct {
	UserName    string
	FileHash    string
	FileName    string
	FileSize    int64
	UploadAt    time.Time
	LastUpdated time.Time
}

type Tb_User_File struct {
	Id int `gorm:"column:id;primaryKey;autoIncrement;not null"`
	User_name string `gorm:"type:varchar(64);not null;default:''"`
	File_sha1 string `gorm:"type:char(40);not null;default:''"`
	File_name string `gorm:"type:varchar(256);not null;unique;default:''"`//unique 规范数据的唯一性
	File_size int64	`gorm:"type: bigint(20);default:'0'"`
	CreatedAt time.Time `gorm:"type:datetime;"`
	UpdatedAt time.Time `gorm:"comment:'最后活跃时间戳';type:datetime"`
	Status int`gorm:"type:int(11);not null;default:'0'"`
	//ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
}
