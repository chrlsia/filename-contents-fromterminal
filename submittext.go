package main

import(
   "bufio"
    "fmt"
   "os"
   "strings"
)

func main(){

	file,text :=getUserText()
	fmt.Println("The file name is: ", file)
	fmt.Println("The text to be saved is:",text)
	
}

func getUserText()(string,string){
	terminalReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a single word as a file name: ")
	filename, _ := terminalReader.ReadString('\n')
	filename = strings.TrimSpace(filename) 
	filename += ".txt"

	fmt.Print("Enter the text to be saved:")
	myText, _ := terminalReader.ReadString('\n')
	return filename, myText
}