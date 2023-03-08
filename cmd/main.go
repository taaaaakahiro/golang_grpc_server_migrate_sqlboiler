package main

import (
	"grpc_func_from_prcivate_repo/pkg/command"
)

//go:generate sqlboiler --wipe psql

func main() {
	command.Run()
}
