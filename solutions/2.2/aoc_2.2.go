package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
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
	horizontalPos:=0
	depth:=0
	aim:=0
	tmpVal:=0
	for _,val := range text{
		tmpVal,_=strconv.Atoi(strings.Fields(val)[1])
		if(strings.Contains(val,"forward")){
			horizontalPos+=tmpVal
			depth+=aim*tmpVal
		}else
		if(strings.Contains(val,"down")){
			aim+=tmpVal
		}else
		if(strings.Contains(val,"up")){
			aim-=tmpVal
		}
	}
	fmt.Println(horizontalPos*depth)
}