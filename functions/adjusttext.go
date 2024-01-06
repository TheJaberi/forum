package forum

import "strings"

func AdjustText (text string) string{
	arrText := strings.Split(text, " ")
	finaltext := ""
	for i:=0;i<len(arrText);i++{
		if len(arrText[i]) > 25{
			for j:= 25;j<len(arrText[i]);{
				arrText[i] = arrText[i][:j] + " " + arrText[i][j:]
				j += 25
			}
			finaltext = finaltext + arrText[i]
		} else {
			finaltext = finaltext + " " + arrText[i]
		}
	} 
	return finaltext
}