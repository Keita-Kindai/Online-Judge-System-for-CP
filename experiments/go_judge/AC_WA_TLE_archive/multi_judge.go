/**
* TLEの判定に加えて複数のテストケースの判定もできるようにしたプログラムです。
* 工夫としてはcompile()とtestcase(i)でそれぞれの処理を関数化しています
* inputファイルとoutputファイルの形式は今のところinput1.txt input1.ans、という形にしています。
*
*/
package main

import (
    "fmt"   
	"os"
	"time"
	"bytes"
	"os/exec"
	"log"
	"strings"
)

func compile() bool {

	cmd := exec.Command("g++", "main.cpp") 
	
	var err error

	err = cmd.Run()

	if err != nil {
		return false
	}

	return true

}

func testcase(case_number int) int {
	fmt.Printf("Testcase#%d ", case_number)	
	
	cmd := exec.Command("prlimit", "--as=67108864", "./a.out")

	var Stdout bytes.Buffer
	var Stderr bytes.Buffer

	cmd.Stdout = &Stdout
	cmd.Stderr = &Stderr
	
	input_file_name := fmt.Sprintf("input%d.txt", case_number)
	input_file, err := os.Open(input_file_name)

	if err != nil {return -1}

	cmd.Stdin = input_file

	defer input_file.Close()

	// エラーが起こったかどうかを返すchannelを作成する
	done := make(chan error, 1)

	fatal_error := cmd.Start()

	if fatal_error != nil {
		return -1
	}

	go func() {
		done <- cmd.Wait()	
	}()
	
	select {
		case done_error := <- done: 

			if done_error != nil {
				fmt.Println(done_error)
				fmt.Println("stderr: ", Stderr.String())
				return -1
			}
			
			output_file_name := fmt.Sprintf("input%d.ans", case_number)
			output_file, err := os.ReadFile(output_file_name)
		
			if err != nil {
				return -1
			}
			
			actual := strings.TrimSpace(Stdout.String())
			expected := strings.TrimSpace(string(output_file))

			if actual == expected {
				fmt.Println("AC")
				return 0
			} else {
				fmt.Println("WA")
				return 1
			}

		case <- time.After(time.Second * 2):
			fmt.Println("TLE")
			cmd.Process.Kill()
			return 2
	}

}

func main() {

	if compile() == false {
		log.Fatalf("Compile Error")
	}	
	
	var err bool = false
	var success int = 0

	for i:=1; i<=3; i++ {
		result := testcase(i)
		if result == -1 {
			err = true
		}
		success = max(success, result)
	}

	if err == true {
		fmt.Println("Internal Error Occured")	
		return
	}

	if success == 0 {
		fmt.Println("ACCEPTED")
	} else if success == 1 {
		fmt.Println("Wrong Answer")
	} else if success == 2 {
		fmt.Println("Time Limit Exceeded")	
	}
}	
