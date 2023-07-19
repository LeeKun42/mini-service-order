package main

import "order/app/cmd"

func main() {
	cmd.Main.AddCommand(&cmd.GormDto)
	cmd.Main.Execute()
}
