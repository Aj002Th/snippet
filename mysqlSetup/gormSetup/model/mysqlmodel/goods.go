package mysqlmodel

type Goods struct {
	ID          string `gorm:"primaryKey;column:id"`
	Name        string `gorm:"column:name"`
	Price       int    `gorm:"column:price"`
	Description string `gorm:"column:description"`
}
