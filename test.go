package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
			panic(err)
	}
	defer func() {
			_ = keyboard.Close()
	}()

	for isFileOk := checkFile(); !isFileOk; {
			log.Println("File requires manual update.")
			log.Println("Hit [RETURN] when resolved.")

			char, key, err := keyboard.GetKey()
			if err != nil {
					panic(err)
			}
			fmt.Println(char, key, err)

			if key == keyboard.KeyEsc {
					return
			}
	}
}

func checkFile() bool {
	return false
}