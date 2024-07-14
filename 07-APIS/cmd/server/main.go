package main

import "github.com/torbenish/pos-golang/07-APIS/configs"

func main(){
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}