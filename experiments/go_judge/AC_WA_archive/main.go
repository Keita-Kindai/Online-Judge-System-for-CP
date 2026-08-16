/*
* これも同じくWAかACかだけを判定できるプログラムです。
* 多分復習用に作っただけ。archivedに移します。
*
*
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"os/exec"
)

func main () {

	cmd := exec.Command("./a.out")

	file, err := os.Open("./input1.txt")

	if err != nil {
		log.Fatalf("Error occured")
	}

	cmd.Stdin = file

	out, err_out := cmd.CombinedOutput()

	if err_out != nil {
		log.Fatalf("Error occured")
	}
	
	ans, err_ans := os.ReadFile("input1.ans")

	if err_ans != nil {
		log.Fatalf("Error")
	}
	
	actual := strings.TrimSpace(string(out))
	expected := strings.TrimSpace(string(ans))

	if actual == expected {
		fmt.Println("Accepted!!")
	} else {
		fmt.Println("Wrong Answer")
	}

}
