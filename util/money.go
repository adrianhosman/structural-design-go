package util

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func IDR(money interface{}) string {

	var initVal float64
	switch val := money.(type) {
	case int:
		initVal = float64(money.(int))
	case int64:
		initVal = float64(money.(int64))
	case float64:
		initVal = money.(float64)
	case string:
		initVal, _ = strconv.ParseFloat(money.(string), 64)
	default:
		log.Println(val, reflect.TypeOf(money), "is unsupported data type, please input either int, int64, float64")
		return ""
	}

	var s string
	i := int(initVal)

	negative := false
	if i < 0 {
		negative = true
		i = i * -1
	}
	if i == 0 {
		s = "0"
	}
	for i > 0 {
		temp := strconv.Itoa(int(i) % 1000)
		i = i / 1000

		if i > 0 {
			zero := strings.Repeat("0", 3-len(temp))
			temp = zero + temp
			s = "." + temp + s
		} else {
			s = temp + s
		}

	}
	if negative == true {
		s = "-" + s
	}
	s = "Rp " + s

	//handle comma
	//-----------------------
	aDec := fmt.Sprintf("%.2f", initVal)
	splitted := strings.Split(aDec, ".")

	if splitted[1] != "00" {
		s = s + "," + splitted[1]
	}
	//-----------------------

	return s

}
