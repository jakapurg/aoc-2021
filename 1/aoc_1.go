package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
func main(){
	file, _ := os.Open("reports.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text [] string
	for scanner.Scan(){
		text = append(text,scanner.Text())
	}
	increaseCount:=0
	prevVal,_:=strconv.Atoi(text[0])
	var tmpVal int =0 
	for _, val := range text[1:] {
		tmpVal,_=strconv.Atoi(val)
		if(prevVal<tmpVal){
			increaseCount++
		}
		prevVal=tmpVal
	}
	fmt.Println(increaseCount)
}