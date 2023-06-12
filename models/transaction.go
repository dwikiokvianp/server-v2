package models

type Transaction struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	UserId    uint64  `gorm:"not null" json:"user_id"`
	User      User    `gorm:"foreignKey:UserId"`
	VehicleId uint64  `gorm:"not null" json:"vehicle_id" `
	Vehicle   Vehicle `gorm:"foreignKey:VehicleId"`
	OilId     uint64  `gorm:"not null" json:"oil_id" `
	Oil       Oil     `gorm:"foreignKey:OilId"`
	QrCodeUrl string  `gorm:"not null" json:"qr_code_url" `
	CreatedAt int64   `gorm:"autoCreateTime" json:"created_at"`
}
