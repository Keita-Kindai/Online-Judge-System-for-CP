/*
* 一番最初に作ったjudge.goです。ただ単にあっているかどうかを判定できるだけ
* 特に使うことがなさそうなので、archivedフォルダにぶち込みます。
* 
*/
package main 

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	cmd := exec.Command("g++", "main.cpp", "-o", "main")

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Error :\n", err)
	}
	
	file, err := os.Open("input1.txt")

	if err != nil {
    	log.Fatalf("File open error: %v\n", err)
	}
	defer file.Close()
	cmd = exec.Command("./main")
	cmd.Stdin = file

	out ,err = cmd.CombinedOutput()

		
	if err != nil {
		log.Fatalf("Error :\n", err)
	}

	ans, err := os.ReadFile("input1.ans")

	if err != nil {
		log.Fatalf("File open error: %v\n", err)
	}

	actual := strings.TrimSpace(string(out))
	expected := strings.TrimSpace(string(ans))

	if actual == expected {
		fmt.Println("Accecpted!!")
	} else {
		fmt.Println("Wrong Answer!")
	}

}
