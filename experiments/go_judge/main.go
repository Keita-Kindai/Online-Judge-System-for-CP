package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main () {
	cmd := exec.Command("ls", "-a")

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("実行エラー: %v\n", err)
	}

	fmt.Println("実行結果:\n", + string(out))
}
