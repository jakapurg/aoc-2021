package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	//"strings"
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
	len:=len(text[0])
	most_common:=""
	least_common:=""
	oneCount:=0
	zeroCount:=0
	for x:=0;x<len;x++{
		oneCount=0
		zeroCount=0
		for _,val := range text{
			if(val[x]=='1'){
				oneCount++
			}else
			{
				zeroCount++
			}
		}
		if(oneCount>zeroCount){
			most_common+="1"
			least_common+="0"
		}else{
			most_common+="0"
			least_common+="1"
		}

	}
	
	gammaRate,_:=strconv.ParseInt(most_common,2,64)
	epsilonRate,_:=strconv.ParseInt(least_common,2,64)
	fmt.Println(gammaRate*epsilonRate)

}