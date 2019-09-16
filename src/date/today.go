package date

import (
    "strconv"
    "time"
    
    "cloud.google.com/go/civil"
)

func GetToday() string {
	n := time.Now()
	return strconv.Itoa(n.Year()) + "-" + strconv.Itoa(int(n.Month())) + "-" + strconv.Itoa(n.Day())
}

func GetTodayCivil() string {
    n := time.Now()
    return civil.DateOf(n).String()
}
