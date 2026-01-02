package models

type MasterPengadaan struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func (MasterPengadaan) TableName() string {
	return "master_type"
}
