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
	saveText(text, file)
}

func saveText(myText string, myFile string) {
	//fileHandler, _ := os.Create(myFile)
	fileHandler, _ := os.OpenFile(myFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)	
	defer fileHandler.Close()
	myWriterBelt := bufio.NewWriter(fileHandler)
	fmt.Fprintln(myWriterBelt, myText)
	myWriterBelt.Flush()
	fmt.Println("Data has been saved")
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