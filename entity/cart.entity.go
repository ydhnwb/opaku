package entity

type Cart struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"-"`
	ProductID uint64  `gorm:"not null" json:"-"`
	Product   Product `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	UserID    uint64  `gorm:"not null" json:"-"`
	User      User    `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
