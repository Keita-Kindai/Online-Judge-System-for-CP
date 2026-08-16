/**
* 前回のACかWAを判定できるか、に加えてTLE判定もできるようにしたいファイルです。
* 基本的にはgo routineを使った並行処理、そしてchannelとtime、そしてselectを使って
* time.After(2)とchannelの戻りがどちらか早い方の処理が優先されるシステムになっています。
* まだ初期段階なので特にAC or WA or TLE以外を判定できる機能がないです。
*
*/
package main 

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"strings"
	"bytes"
)

func main() {


	cmd := exec.Command("g++", "main.cpp")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Error")
		fmt.Println(string(out))
		return
	}

	file, err := os.Open("input1.txt")

	if err != nil {
		log.Fatalf("Error")
	}

	runCmd := exec.Command("./a.out")

	
	runCmd.Stdin = file
	defer file.Close()

	var stderr bytes.Buffer
	var stdout bytes.Buffer
	runCmd.Stdout = &stdout
	runCmd.Stderr = &stderr

	err = runCmd.Start()

	if err != nil {
		return
	}

	done := make(chan error, 1)

	go func(){
		done <- runCmd.Wait()
	}()

	select {
	
		case <- done :
			output := stdout.String()

			actual := strings.TrimSpace(output)
			
			answer,read_err := os.ReadFile("input1.ans")
			
			if read_err != nil {
				return
			}

			expected := strings.TrimSpace(string(answer))

			if actual == expected {
				fmt.Println("Accepted")
			} else {
				fmt.Println(actual, expected)
				fmt.Println("Wrong Answer")
			}

		case <- time.After(time.Second * 2):
			runCmd.Process.Kill()
			fmt.Println("Time Limit Exceeded")
			return

	}

}
