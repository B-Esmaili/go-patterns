package main

import "patterns/singletone/db"

func main() {
	for i := 0; i < 50; i++ {
		go db.GetDb()
	}
}
