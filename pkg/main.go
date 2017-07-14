package main

import (
	"github.com/casbin/casbin"
	_ "github.com/abbot/go-http-auth"
)

func main() {
	e := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
}
