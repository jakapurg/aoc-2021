package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)
const SUM_RETENTION = 3
func sumNums(nums []string)int{
	var sum int=0
	var tmpVal int =0 
	for _,val:=range nums{
		tmpVal,_= strconv.Atoi(val)
		sum=sum+tmpVal
	}
	return sum
}
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
	prevSum:=sumNums(text[0:SUM_RETENTION])
	fmt.Println(prevSum)
	var tmpVal int =0 
	for i, _ := range text[1:] {
		tmpVal=sumNums(text[i+1:i+1+SUM_RETENTION])
		if(prevSum<tmpVal){
			increaseCount++
		}
		prevSum=tmpVal
	}
	fmt.Println(increaseCount)
}