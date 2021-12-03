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
	leng:=len(text[0])
	most_common:=""
	least_common:=""
	oneCount:=0
	zeroCount:=0
	commonTxt:=text
	leastTxt:=text
	for x:=0;x<leng;x++{
		oneCount=0
		zeroCount=0

		if( len(commonTxt) > 1){
			for _,val := range commonTxt{
				if(val[x]=='1'){
					oneCount++
				}else{
					zeroCount++
				}
			}
			if(oneCount>=zeroCount){
				most_common="1"
			}else{
				most_common="0"
			}
			var tmpCommon[] string
			fmt.Println(most_common)
			for _,val:= range commonTxt{
				if(string(val[x])==most_common){
					tmpCommon = append(tmpCommon,val)
				}
			}
			commonTxt=tmpCommon
		}
		if( len(leastTxt) > 1){
			oneCount=0
			zeroCount=0
			for _,val := range leastTxt{
				if(val[x]=='1'){
					oneCount++
				}else{
					zeroCount++
				}
			}
			if(oneCount>=zeroCount){
				least_common="0"
			}else{
				least_common="1"
			}
			

			var tmpLeast[] string
			for _,val:= range leastTxt{
				if(string(val[x])==least_common){
					tmpLeast=append(tmpLeast,val)
				}
			}
			leastTxt = tmpLeast
		}


	}
	oxRating,_:=strconv.ParseInt(commonTxt[0],2,64)
	coRating,_:=strconv.ParseInt(leastTxt[0],2,64)
	fmt.Println(oxRating*coRating)

}