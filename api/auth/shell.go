package auth

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/itpkg/reading/api/core"
	"log"
)

func (p *AuthEngine) Shell() []cli.Command {
	return []cli.Command{

		{
			Name:    "users",
			Aliases: []string{"us"},
			Usage:   "list all users",
			Flags:   []cli.Flag{core.ENV},
			Action: core.Action(func(env string) error {
				cfg, err := core.Load(env)
				if err != nil {
					return err
				}
				db, err := cfg.OpenDatabase()
				if err != nil {
					return err
				}
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
			Action: func(c *cli.Context) {
				u := c.String("user")
				r := c.String("name")
				d := c.Bool("deny")
				if u == "" || r == "" {
					log.Fatalln("need role's name and user's uid")
				}
				cfg, err := core.Load(c.String("environment"))
				if err != nil {
					log.Fatalln(err)
				}
				db, err := cfg.OpenDatabase()
				if err != nil {
					log.Fatalln(err)
				}
				dao := Dao{Db: db}
				role, err := dao.Role(r, "-", 0)
				if err != nil {
					log.Fatalln(err)
				}
				user, err := dao.GetUser(u)
				if err != nil {
					log.Fatalln(err)
				}
				if d {
					err = dao.Deny(role.ID, user.ID)
				} else {
					err = dao.Allow(role.ID, user.ID, 24*365*10*time.Hour)
				}
				if err != nil {
					log.Fatalln(err)
				}
				log.Println("Done...")
			},
		},
	}
}
