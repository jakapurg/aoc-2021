package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)
func main(){
	file,_ := os.Open("reports.txt")
	defer file.Close()
	scanner:= bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text[] string
	for scanner.Scan(){
		text = append(text,scanner.Text())
	}
	var codes[] string
	leng:=0
	newCode:=""
	count:=0
	tmpCode:=0
	var dict[] string
	var spl[] string
	var correctDict [10] string
	for _,val := range text {
		spl = strings.Split(val,"|")
		
		dict = strings.Fields(spl[0])
		sort.Slice(dict, func(i, j int) bool {
			return len(dict[i]) < len(dict[j])
		})
		codes = strings.Fields(spl[1])
		for _,val3 := range dict{
			leng=len(val3)
			if(leng==2){
				correctDict[1]=val3
			}else
			if(leng==3){
				correctDict[7] = val3
			}else
			if(leng==4){
				correctDict[4] = val3
			}else
			if(leng==7){
				correctDict[8] = val3
			}else
			if(leng==5){
				if(strings.Contains(val3,string(correctDict[1][0])) && strings.Contains(val3,string(correctDict[1][1]))){
					correctDict[3] = val3
				}else
				{
					occ:=0
					for _,v := range correctDict[4]{
						if(strings.ContainsRune(val3,v)){
							occ++
						}
					}
					if(occ==3){
						correctDict[5] = val3
					}else {
						correctDict[2] = val3
					}
				}
			} else
				if(!(strings.Contains(val3,string(correctDict[1][0])) && strings.Contains(val3,string(correctDict[1][1])))){
					correctDict[6] = val3
				} else
				{
					occ:=0
					for _,v := range correctDict[4]{
						if(strings.ContainsRune(val3,v)){
							occ++
						}
					}
					if(occ==4){
						correctDict[9]=val3
					}else{
						correctDict[0]=val3
					}
				}
			
		}
		newCode=""
		for _,val2 := range codes {
			leng = len(val2)
			if(leng==2){
				newCode+="1"
			}else
			if(leng==3){
				newCode+="7"
			}else
			if(leng==4){
				newCode+="4"
			}else
			if(leng==7){
				newCode+="8"
			}else
			if(leng==5){
				is3:=true
				for _,val4 := range correctDict[3]{
					if(!strings.ContainsRune(val2,val4)){
						is3=false
					}
				}

				if(is3){
					newCode+="3"
				}else{
					is5:=true
					for _,val4 := range correctDict[5]{
						if(!strings.ContainsRune(val2,val4)){
							is5=false
						}
					}
					if(is5){
						newCode+="5"
					}else 
					{
						newCode+="2"
					}
				}

			} else
			{
				is6:=true
				for _,val4 := range correctDict[6]{
					if(!strings.ContainsRune(val2,val4)){
						is6=false
					}
				}

				if(is6){
					newCode+="6"
				}else{
					is9:=true
					for _,val4 := range correctDict[9]{
						if(!strings.ContainsRune(val2,val4)){
							is9=false
						}
					}
					if(is9){
						newCode+="9"
					}else 
					{
						newCode+="0"
					}
				}
			}
		}
		tmpCode,_ = strconv.Atoi(newCode)
		count+=tmpCode

	}
	fmt.Println(count)
}
