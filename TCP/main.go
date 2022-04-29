package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
		"math/rand"
		"time"
)

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

func main() {
	rand.Seed(time.Now().UnixNano())

	liste_mots := file_read("liste_francais.txt")
	index := rand.Intn(len(liste_mots)-1)

	var mot string = liste_mots[index]

	fmt.Println("Trouver le mot :",string(mot[0]))


	b := true
	var msgf string
	var nblettre_trouve int = 1
	var tab_lettre_trouve []string = []string{string(mot[0])}
	reader := bufio.NewReader(os.Stdin)

	for b {
		fmt.Print("Type here : ")
		msg, _ := reader.ReadBytes('\n')
		msg = msg[:len(msg)-1]
		msgf = string(msg)
		
		if msgf == mot {
			fmt.Println("Gagné")
			break
		} else {
			fmt.Println(string(mot))
			for i := 0; i < len(msgf) ; i++ {
				for j := 0 ; j < len(mot) ; j++ {
					if msgf[i] == mot[j] {
						fmt.Println("+1")
						for k:= 0 ; k<len(tab_lettre_trouve) ; k++ {
							if string(msgf[i]) == tab_lettre_trouve[k] {
								break
							} else {
								tab_lettre_trouve = append(tab_lettre_trouve, string(msgf[i]))
							}
						}
		
					}
				}
			}
			fmt.Println("Vous avez trouvé ",nblettre_trouve," lettres")
			fmt.Println("Ce sont les lettres suivantes : ", tab_lettre_trouve)
		}
	}
}

