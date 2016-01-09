package assets

import (
	"time"
)

type Page struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `sql:"not null" json:"title"`
	Type      string    `sql:"not null;type:varchar(8);index" json:"type"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp" json:"created_at"`
}

type Channel struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Cid         string    `sql:"not null;index" json:"cid"`
	Type        string    `sql:"not null;type:varchar(8);index" json:"type"`
	Title       string    `sql:"not null" json:"title"`
	Description string    `sql:"not null;type:text" json:"description"`
	CreatedAt   time.Time `sql:"not null;default:current_timestamp" json:"created_at"`
}

type Playlist struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Cid         string    `sql:"not null;index" json:"cid"`
	Pid         string    `sql:"not null;index" json:"pid"`
	Type        string    `sql:"not null;type:varchar(8);index" json:"type"`
	Title       string    `sql:"not null" json:"title"`
	Description string    `sql:"not null;type:text" json:"description"`
	CreatedAt   time.Time `sql:"not null;default:current_timestamp" json:"created_at"`
}

type Video struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Pid         string    `sql:"not null;index" json:"pid"`
	Vid         string    `sql:"not null;index" json:"vid"`
	Type        string    `sql:"not null;type:varchar(8);index" json:"type"`
	Title       string    `sql:"not null" json:"title"`
	Description string    `sql:"not null;type:text" json:"description"`
	CreatedAt   time.Time `sql:"not null;default:current_timestamp" json:"created_at"`
}

type Book struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Url       string    `sql:"not null;index" json:"url"`
	Title     string    `sql:"not null;index" json:"title"`
	Author    string    `sql:"not null;index" json:"author"`
	Type      string    `sql:"not null;type:varchar(8);index" json:"type"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp" json:"created_at"`
}

func (p *AssetsEngine) Migrate() {
	p.Db.AutoMigrate(&Page{}, &Channel{}, &Playlist{}, &Video{}, &Book{})
}
