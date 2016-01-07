package config

import (
	"fmt"
	"strconv"
)

type Database struct {
	Adapter  string            `toml:"adapter"`
	Host     string            `toml:"host"`
	Port     int               `toml:"port"`
	Name     string            `toml:"name"`
	User     string            `toml:"user"`
	Password string            `toml:"password"`
	Extra    map[string]string `toml:"extra"`
}

func (p *Database) Execute(sql string) (string, []string) {
	switch p.Adapter {
	case "postgres":
		return "psql", []string{"-U", p.User, "-h", p.Host, "-p", strconv.Itoa(p.Port), "-c", sql}
	default:
		return "echo", []string{fmt.Sprintf("Unsupported database adapter %s", p.Adapter)}
	}
}

func (p *Database) Console() (string, []string) {
	switch p.Adapter {
	case "postgres":
		return "psql", []string{"-U", p.User, "-h", p.Host, "-p", strconv.Itoa(p.Port), p.Name}
	default:
		return "echo", []string{fmt.Sprintf("Unsupported database adapter %s", p.Adapter)}
	}
}
