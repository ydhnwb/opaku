package entity

type Product struct {
	ID    uint64 `gorm:"primary_key:auto_increment" json:"-"`
	Name  string `gorm:"type:varchar(100)" json:"-"`
	Price uint64 `gorm:"type:bigint" json:"-"`
	Image string `gorm:"type:text" json:"-"`
	// UserID int64  `gorm:"not null" json:"-"`
	// User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
