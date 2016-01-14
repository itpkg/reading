package cms

import (
	"fmt"

	"github.com/itpkg/reading/api/auth"
	"github.com/itpkg/reading/api/core"
)

type Attachment struct {
	core.Model
	UserID uint `sql:"not null"`
	User   auth.User

	Title       string `sql:"not null"`
	ContentType string `sql:"not null;type:varchar(16);index"`
	Body        []byte `sql:"not null"`
}

func (Attachment) TableName() string {
	return "attachments"
}

type Article struct {
	core.Model
	UserID uint `sql:"not null"`
	User   auth.User

	Aid     string `sql:"not null;type:varchar(36);unique"`
	Title   string `sql:"not null"`
	Summary string
	Lang    string `sql:"not null;type:char(5);index"`
	Body    string `sql:"not null;type:TEXT"`
	Type    string `sql:"not null;type:varchar(8);default:'markdown';index"`

	Tags []Tag `gorm:"many2many:cms_article_tags;"`
}

func (Article) TableName() string {
	return "cms_articles"
}

type Tag struct {
	core.Model
	Name     string    `sql:"not null;unique"`
	Articles []Article `gorm:"many2many:cms_article_tags;"`
}

func (Tag) TableName() string {
	return "cms_tags"
}

type Comment struct {
	core.Model
	UserID uint `sql:"not null"`
	User   auth.User

	Body string `sql:"not null;type:TEXT"`
	Type string `sql:"not null;type:varchar(8);default:'markdown';index"`
}

func (Comment) TableName() string {
	return "cms_comments"
}

type User struct {
	core.Model

	Uid  string `sql:"not null;index"`
	Type string `sql:"not null;type:varchar(8);index"`
}

func (p User) String() string {
	return fmt.Sprintf("%s@%s", p.Uid, p.Type)
}
func (User) TableName() string {
	return "video_users"
}

type Channel struct {
	core.Model

	Uid         string `sql:"not null;index" json:"uid"`
	Cid         string `sql:"not null;index" json:"cid"`
	Type        string `sql:"not null;type:varchar(8);index" json:"type"`
	Title       string `sql:"not null" json:"title"`
	Description string `sql:"not null;type:text" json:"description"`
}

func (Channel) TableName() string {
	return "video_channels"
}

type Playlist struct {
	core.Model
	Cid         string `sql:"not null;index" json:"cid"`
	Pid         string `sql:"not null;index" json:"pid"`
	Type        string `sql:"not null;type:varchar(8);index" json:"type"`
	Title       string `sql:"not null" json:"title"`
	Description string `sql:"not null;type:text" json:"description"`
}

func (Playlist) TableName() string {
	return "video_playlist"
}

type Video struct {
	core.Model

	Pid         string `sql:"not null;index" json:"pid"`
	Vid         string `sql:"not null;index" json:"vid"`
	Type        string `sql:"not null;type:varchar(8);index" json:"type"`
	Title       string `sql:"not null" json:"title"`
	Description string `sql:"not null;type:text" json:"description"`
}

func (Video) TableName() string {
	return "video_items"
}

type Book struct {
	core.Model
	Type      string `sql:"not null;index" json:"type"`
	Name      string `sql:"not null;unique" json:"name"`
	IndexHref string `json:"index_href"`
	IndexType string `json:"index_type"`
	CoverHref string `json:"cover_href"`
	CoverType string `json:"cover_type"`

	Title      string `sql:"not null;index" json:"title"`
	Author     string `sql:"not null;index" json:"author"`
	Language   string `json:"language"`
	Identifier string `json:"identifier"`
	Subject    string `json:"subject"`
	Publisher  string `json:"publisher"`
	Date       string `json:"date"`
}

func (Book) TableName() string {
	return "cms_books"
}

//-----------------------------------------------------------------------------
func (p *CmsEngine) Migrate() {
	p.Db.AutoMigrate(&Attachment{},
		&Article{}, &Tag{}, &Comment{},
		&User{}, &Channel{}, &Playlist{}, &Video{},
		&Book{},
	)
	p.Db.Model(&User{}).AddUniqueIndex("idx_video_users_uid_type", "uid", "type")
	p.Db.Model(&Channel{}).AddUniqueIndex("idx_video_channels_cid_type", "cid", "type")
	p.Db.Model(&Playlist{}).AddUniqueIndex("idx_video_playlist_cid_type", "pid", "type")
	p.Db.Model(&Video{}).AddUniqueIndex("idx_video_items_cid_type", "vid", "type")
}
