package main

import (
	"golang_grpc_proto/pkg/command"
)

//go:generate sqlboiler --wipe psql

func main() {
	command.Run()
}
