package main
import (
	"fmt"
	"bufio"
	"os"
	"sort"
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
func reverseArray(arr []rune) []rune{
	for i, j := 0, len(arr)-1; i<j; i, j = i+1, j-1 {
	   arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
 }
func main(){
	file, _ := os.Open("reports.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text [] string
	oppeningTags := [4]rune{'(','[','{','<'}
	closingTags := [4]rune{')',']','}','>'}
	for scanner.Scan(){
		text = append(text,scanner.Text())
	}
	var expectedClosing [] rune
	var scores [] int
	correct:=false
	for _,val := range text {
		score :=0
		correct=true
		expectedClosing = nil
		for _,val2 := range val{
			bl, in := contains(oppeningTags[:],val2)
			if(bl){
				expectedClosing = append(expectedClosing,closingTags[in])
			}else
			if(val2 == expectedClosing[len(expectedClosing)-1]){
				expectedClosing = expectedClosing[:len(expectedClosing)-1]
			}else
			{
				correct=false
				break;
			}
		}
		if(correct){
			for _,val3 := range reverseArray(expectedClosing){
				_,indS := contains(closingTags[:],val3)
				score=score*5+indS+1
			}
			scores = append(scores,score)
		}
	}
	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])


}