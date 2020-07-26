package main

import (
	"Ariadne/Core"
	"context"
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"os"
)

func main(){
	parser := argparse.NewParser("Ariadne", "Guides you to root")

	target := parser.String("t", "target", &argparse.Options{Required: true, Help: "target ip address or hostname or website"})
	projectIndex := parser.String("i", "index", &argparse.Options{Required: true, Help: "new project index for elasticsearch"})
	httpExtensions := parser.String("x", "extensions", &argparse.Options{Required: false, Help: "http extensions for dirbusting",Default:"php,html,txt"})
	err := parser.Parse(os.Args)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Ariadne Starting On Target [%s] with ProjectIndex [%s] and httpExtensions [%s]\n",*target,*projectIndex,*httpExtensions)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ariadne := Core.NewAriadneTarget(*target,*projectIndex,ctx)
	ariadne.StartEnumerating()

}