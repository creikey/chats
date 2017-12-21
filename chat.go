package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	errHandle "github.com/creikey/chats/errorHandles"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("(s) for server, (c) for client: ")
	text, _, err := reader.ReadRune()
	errHandle.CheckError(err)
	if text == 's' {
		runServer()
	} else if text == 'c' {
		runClient()
	} else {
		log.Fatalln(errors.New("THAT'S NOT AN OPTION. I'M QUITTING"))
	}
}

func runServer() {

}

func runClient() {

}
