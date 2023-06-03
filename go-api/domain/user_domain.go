package domain

type Client struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:nombre"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:contrasena"`
	Role     string `gorm:"column:rol"`
}
