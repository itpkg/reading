package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/config"
	"github.com/itpkg/reading/api/core"
	"github.com/jinzhu/gorm"
)

func (p *AuthEngine) Shell() []cli.Command {
	return []cli.Command{

		{
			Name:    "users",
			Aliases: []string{"us"},
			Usage:   "list all users",
			Flags:   []cli.Flag{core.ENV},
			Action: config.DatabaseAction(func(db *gorm.DB, _ *cli.Context) error {
				dao := Dao{Db: db}
				fmt.Println("UID USER")
				for _, u := range dao.ListUser() {
					fmt.Printf("%s %s<%s>\n", u.Uid, u.Name, u.Email)
				}
				return nil
			}),
		},
		{
			Name:    "role",
			Aliases: []string{"r"},
			Usage:   "apple role to user",

			Flags: []cli.Flag{
				core.ENV,
				cli.StringFlag{
					Name:  "user, u",
					Value: "",
					Usage: "uid of user",
				},
				cli.StringFlag{
					Name:  "name, n",
					Value: "",
					Usage: "name of role",
				},
				cli.BoolFlag{
					Name:  "deny, d",
					Usage: "remove role from user",
				},
			},
			Action: config.DatabaseAction(func(db *gorm.DB, ctx *cli.Context) error {
				u := ctx.String("user")
				r := ctx.String("name")
				d := ctx.Bool("deny")
				if u == "" || r == "" {
					return errors.New("need role's name and user's uid")
				}

				dao := Dao{Db: db}
				role, err := dao.Role(r, "-", 0)
				if err != nil {
					return err
				}
				user, err := dao.GetUser(u)
				if err != nil {
					return err
				}
				if d {
					err = dao.Deny(role.ID, user.ID)
				} else {
					err = dao.Allow(role.ID, user.ID, 24*365*10*time.Hour)
				}
				return err
			}),
		},
	}
}
