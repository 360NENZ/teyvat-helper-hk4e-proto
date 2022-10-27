package main

import (
	"log"

	"github.com/teyvat-helper/hk4e-proto/pb"
)

func main() {
	log.Println("Loaded", len(pb.ProtoCommandNewFuncMap), "proto command(s).")
}
