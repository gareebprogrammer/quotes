package main

import (
	"os"
)

//EnvVars struct to load env vars
type EnvVars struct {
	addr string
	port string
}

//ParseEnvVars ...Retuns a new EnvVars struct with all th env values
func ParseEnvVars() EnvVars {
	vars := EnvVars{}
	port := os.Getenv("PORT")
	addr := os.Getenv("ADDR")
	vars.addr = addr
	vars.port = port
	return vars
}

// JoinAddr ...takes address and port and join them
func JoinAddr(addr string, port string) string {
	return (addr + port)
}
