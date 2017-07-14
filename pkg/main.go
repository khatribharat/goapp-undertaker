package main

import (
	"github.com/casbin/casbin"
)

func main() {
	e := casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
}
