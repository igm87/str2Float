package str2Float

import  (
	"strings"
	"strconv"
)

func Str2F(s string)(float64,error){
	s = strings.TrimSpace(s)
	converted,err := strconv.ParseFloat(s,64)
	if err != nil{
		return 0.0,err
	}
	return converted,nil
}
