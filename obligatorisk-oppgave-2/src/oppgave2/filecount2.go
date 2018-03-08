package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"strings"
	"sort"

)

func main() {
	/**
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter filename: ")
	fileName, _ := reader.ReadString('\n')
	*/

	fil, err := os.Open("C:/Users/magnu/GolandProjects/TeamLenny/obligatorisk-oppgave-2/src/oppgave2/text.txt") //Henter fil via os.Open.
	errorCheck(err)
	fil2, err := ioutil.ReadFile("C:/Users/magnu/GolandProjects/TeamLenny/obligatorisk-oppgave-2/src/oppgave2/text.txt") //Henter fil via ioutil.ReadFile.
	errorCheck(err)

	fileScanner := bufio.NewScanner(fil)

	/**
	En løkke som benytter fileScanner for å gå igjennom hver linje, og adderer 1 i en variabel.
	 */
	lineCounter := 0
	for fileScanner.Scan() {
		lineCounter++
	}

	str := string(fil2) //Konverterer filen til en string.

	/**
	Alfabet; adderer tilsvarende store og små bokstaver sammen.
	 */
	amountRuneA := strings.Count(str, "A") + strings.Count(str, "a")
	amountRuneB := strings.Count(str, "B") + strings.Count(str, "b")
	amountRuneC := strings.Count(str, "C") + strings.Count(str, "c")
	amountRuneD := strings.Count(str, "D") + strings.Count(str, "d")
	amountRuneE := strings.Count(str, "E") + strings.Count(str, "e")
	amountRuneF := strings.Count(str, "F") + strings.Count(str, "f")
	amountRuneG := strings.Count(str, "G") + strings.Count(str, "g")
	amountRuneH := strings.Count(str, "H") + strings.Count(str, "h")
	amountRuneI := strings.Count(str, "I") + strings.Count(str, "i")
	amountRuneJ := strings.Count(str, "J") + strings.Count(str, "j")
	amountRuneK := strings.Count(str, "K") + strings.Count(str, "k")
	amountRuneL := strings.Count(str, "L") + strings.Count(str, "l")
	amountRuneM := strings.Count(str, "M") + strings.Count(str, "m")
	amountRuneN := strings.Count(str, "N") + strings.Count(str, "n")
	amountRuneO := strings.Count(str, "O") + strings.Count(str, "o")
	amountRuneP := strings.Count(str, "P") + strings.Count(str, "p")
	amountRuneQ := strings.Count(str, "Q") + strings.Count(str, "q")
	amountRuneR := strings.Count(str, "R") + strings.Count(str, "r")
	amountRuneS := strings.Count(str, "S") + strings.Count(str, "s")
	amountRuneT := strings.Count(str, "T") + strings.Count(str, "t")
	amountRuneU := strings.Count(str, "U") + strings.Count(str, "u")
	amountRuneV := strings.Count(str, "V") + strings.Count(str, "v")
	amountRuneW := strings.Count(str, "W") + strings.Count(str, "w")
	amountRuneX := strings.Count(str, "X") + strings.Count(str, "x")
	amountRuneY := strings.Count(str, "Y") + strings.Count(str, "y")
	amountRuneZ := strings.Count(str, "Z") + strings.Count(str, "z")

	amountRunes := []int {
		amountRuneA, amountRuneB, amountRuneC, amountRuneD, amountRuneE,
		amountRuneF, amountRuneG, amountRuneH, amountRuneI, amountRuneJ,
		amountRuneK, amountRuneL, amountRuneM, amountRuneN, amountRuneO,
		amountRuneP, amountRuneQ, amountRuneR, amountRuneS, amountRuneT,
		amountRuneU, amountRuneV, amountRuneW, amountRuneX, amountRuneY, amountRuneZ,

	}

	fmt.Println(amountRunes[0])

	letters := []struct {
		Letter string
		Count  int
	}{
		{"A", amountRuneA},
		{"B", amountRuneB},
		{"C", amountRuneC},
		{"D", amountRuneD},
		{"E", amountRuneE},
		{"F", amountRuneF},
		{"G", amountRuneG},
		{"H", amountRuneH},
		{"I", amountRuneI},
		{"J", amountRuneJ},
		{"K", amountRuneK},
		{"L", amountRuneL},
		{"M", amountRuneM},
		{"N", amountRuneN},
		{"O", amountRuneO},
		{"P", amountRuneP},
		{"Q", amountRuneQ},
		{"R", amountRuneR},
		{"S", amountRuneS},
		{"T", amountRuneT},
		{"U", amountRuneU},
		{"V", amountRuneV},
		{"W", amountRuneW},
		{"X", amountRuneX},
		{"Y", amountRuneY},
		{"Z", amountRuneZ},
	}
	sort.Slice(letters, func(i, j int) bool { return letters[i].Letter < letters[j].Letter })
	fmt.Println("By Letter:", letters)

	sort.Slice(letters, func(i, j int) bool { return letters[i].Count < letters[j].Count })
	fmt.Println("By Count:", letters)
}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}