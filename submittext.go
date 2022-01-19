package main

import(
   "bufio"
    "fmt"
   "os"
   "strings"
)

func main(){

	getUserText()

}

func getUserText(){
	terminalReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a single word as a file name: ")
	filename, _ := terminalReader.ReadString('\n')
	filename = strings.TrimSpace(filename) 
	filename += ".txt"

	fmt.Println("The filename is: ", filename)
}