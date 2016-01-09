package assets

import (
	"time"
)

type Asset struct {
	ID        uint      `gorm:"primary_key"`
	Title     string    `sql:"not null"`
	Type      string    `sql:"not null;type:varchar(8);index"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp"`
}

type Channel struct {
	ID          uint      `gorm:"primary_key"`
	Cid         string    `sql:"not null;index"`
	Type        string    `sql:"not null;type:varchar(8);index"`
	Title       string    `sql:"not null"`
	Description string    `sql:"not null;type:text"`
	CreatedAt   time.Time `sql:"not null;default:current_timestamp"`
}

type Playlist struct {
	ID          uint      `gorm:"primary_key"`
	Cid         string    `sql:"not null;index"`
	Pid         string    `sql:"not null;index"`
	Type        string    `sql:"not null;type:varchar(8);index"`
	Title       string    `sql:"not null"`
	Description string    `sql:"not null;type:text"`
	CreatedAt   time.Time `sql:"not null;default:current_timestamp"`
}

type Video struct {
	ID          uint      `gorm:"primary_key"`
	Pid         string    `sql:"not null;index"`
	Vid         string    `sql:"not null;index"`
	Type        string    `sql:"not null;type:varchar(8);index"`
	Title       string    `sql:"not null"`
	Description string    `sql:"not null;type:text"`
	CreatedAt   time.Time `sql:"not null;default:current_timestamp"`
}

type Book struct {
	ID        uint      `gorm:"primary_key"`
	Url       string    `sql:"not null;index"`
	Title     string    `sql:"not null;index"`
	Author    string    `sql:"not null;index"`
	Type      string    `sql:"not null;type:varchar(8);index"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp"`
}

func (p *AssetsEngine) Migrate() {
	p.Db.AutoMigrate(&Asset{}, &Channel{}, &Playlist{}, &Video{}, &Book{})
}
