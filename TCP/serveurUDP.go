package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	/* "reflect" */
	"time"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    liste_mots := file_read("liste_francais.txt")
    index := rand.Intn(len(liste_mots))
    mot := liste_mots[index] 


	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleConnection(c,mot)
	}
}


func handleConnection(conn net.Conn,mot string) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
    input := string(buffer[:len(buffer)-1])

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("Client message:", input)

    if input != mot {
        log.Println("Faux")
        conn.Write([]byte("Faux"))
        handleConnection(conn,mot)
    } else {
        log.Println("Vrai")
        conn.Write([]byte("Vrai"))
        conn.Close()
    }

    conn.Write(buffer)
	
}

func file_read(file string)(liste []string) {
	
	readFile, err := os.Open(file)

	if err != nil {
			log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
			lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	return lines
}