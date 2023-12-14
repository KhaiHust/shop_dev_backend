package main

import (
	"go.uber.org/fx"
	"shop_dev/bootstrap"
)

func main() {
	fx.New(
		bootstrap.All(),
	).Run()
}
