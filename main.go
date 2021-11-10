package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Quel nombre de tirage souhaitez-vous ? ")
	fmt.Println("1 -         10 ")
	fmt.Println("2 -        100 ")
	fmt.Println("3 -      1 000 ")
	fmt.Println("4 -     10 000 ")
	fmt.Println("5 -    100 000 ")
	fmt.Println("6 -  1 000 000 ")
	fmt.Println("7 - 10 0000 00 ")

	fmt.Println("Quel est votre choix ?")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choixNbTirage, err := strconv.Atoi(strings.ToLower(in.Text()))
	if err != nil {
		println("Réponse non accepté !")
		fmt.Print("\n\nPress 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}

	var choix int
	switch choixNbTirage {
	case 1:
		choix = 10
	case 2:
		choix = 100
	case 3:
		choix = 1000
	case 4:
		choix = 10000
	case 5:
		choix = 100000
	case 6:
		choix = 1000000
	case 7:
		choix = 10000000
	default:
		println("Réponse non accepté !")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	var deux, trois, quatre, cinq, six, sept, huit, neuf, dix, onze, douze int
	nbTirage := float64(choix)

	for i := 0.0; i < nbTirage; i++ {
		d1 := rand.Intn(6) + 1.0
		d2 := rand.Intn(6) + 1.0
		result := d1 + d2
		switch result {
		case 2:
			deux++
		case 3:
			trois++
		case 4:
			quatre++
		case 5:
			cinq++
		case 6:
			six++
		case 7:
			sept++
		case 8:
			huit++
		case 9:
			neuf++
		case 10:
			dix++
		case 11:
			onze++
		case 12:
			douze++
		}
	}

	fmt.Println("Résultat\tChance\t\t% Chance\tRésultat lancés\t\t% Constaté\n")
	fmt.Printf("   2    \t 1/36 \t\t%.4f\t\t  %10d  \t\t   %.4f\n", 1.0/36*100, deux, (float64(deux)/nbTirage)*100)
	fmt.Printf("   3    \t 2/36 1/18 \t%.4f\t\t  %10d  \t\t   %.4f\n", 2.0/36*100, trois, (float64(trois)/nbTirage)*100)
	fmt.Printf("   4    \t 3/36 1/12\t%.4f\t\t  %10d  \t\t   %.4f\n", 3.0/36*100, quatre, (float64(quatre)/nbTirage)*100)
	fmt.Printf("   5    \t 4/36 1/9\t%.4f\t\t  %10d  \t\t   %.4f\n", 4.0/36*100, cinq, (float64(cinq)/nbTirage)*100)
	fmt.Printf("   6    \t 5/36 \t\t%.4f\t\t  %10d  \t\t   %.4f\n", 5.0/36*100, six, (float64(six)/nbTirage)*100)
	fmt.Printf("   7    \t 6/36 1/6\t%.4f\t\t  %10d  \t\t   %.4f\n", 6.0/36*100, sept, (float64(sept)/nbTirage)*100)
	fmt.Printf("   8    \t 5/36 \t\t%.4f\t\t  %10d  \t\t   %.4f\n", 5.0/36*100, huit, (float64(huit)/nbTirage)*100)
	fmt.Printf("   9    \t 4/36 1/9\t%.4f\t\t  %10d  \t\t   %.4f\n", 4.0/36*100, neuf, (float64(neuf)/nbTirage)*100)
	fmt.Printf("   10   \t 3/36 1/12\t%.4f\t\t  %10d  \t\t   %.4f\n", 3.0/36*100, dix, (float64(dix)/nbTirage)*100)
	fmt.Printf("   11   \t 2/36 1/18\t%.4f\t\t  %10d  \t\t   %.4f\n", 2.0/36*100, onze, (float64(onze)/nbTirage)*100)
	fmt.Printf("   12   \t 1/36 \t\t%.4f\t\t  %10d  \t\t   %.4f\n", 1.0/36*100, douze, (float64(douze)/nbTirage)*100)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	fmt.Print("\n\nPress 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
