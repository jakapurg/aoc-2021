package main
import (
	"fmt"
	"bufio"
	"os"
	//"strconv"
	//"strings"
)
//
func contains(s []rune, str rune) (bool,int){
	for i, v := range s {
		if v == str{
			return true, i
		}
	}

	return false, -1
}

func main(){
	file, _ := os.Open("reports.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text [] string
	oppeningTags := [4]rune{'(','[','{','<'}
	closingTags := [4]rune{')',']','}','>'}
	scores := [4]int{3,57,1197,25137}
	for scanner.Scan(){
		text = append(text,scanner.Text())
	}
	var expectedClosing [] rune
	globalScore:=0
	for _,val := range text {
		for _,val2 := range val{
			bl, in := contains(oppeningTags[:],val2)
			if(bl){
				expectedClosing = append(expectedClosing,closingTags[in])
			}else
			if(val2 == expectedClosing[len(expectedClosing)-1]){
				expectedClosing = expectedClosing[:len(expectedClosing)-1]
			}else
			{
				_,ind := contains(closingTags[:],val2)
				globalScore+=scores[ind]
				break;
			}
		}
	}
	fmt.Println(globalScore)

}