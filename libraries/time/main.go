package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("now is ", t)
	fmt.Println("local now is ", t.Local())
	fmt.Println("utc now is ", t.UTC())
	fmt.Println("unix now is ", t.Unix())
	fmt.Println("t.Location is ", t.Location())
	fmt.Print("t.Zone is ")
	fmt.Println(t.Zone())
	fmt.Println("t.IsZero ? ", t.IsZero())
	fmt.Println("t.UTC() == t ? ", t.UTC() == t)
	fmt.Println("t.UTC().Equal(t) ? ", t.UTC().Equal(t))
	fmt.Print("t.Date is ")
	fmt.Println(t.Date())
	fmt.Print("t.Clock is ")
	fmt.Println(t.Clock())
	sec := time.Duration(1e9)
	fmt.Println("t.Add(sec) is ", t.Add(sec))
	fmt.Println("t.Sub(t.Add(-sec)) is ", t.Sub(t.Add(sec)))
	fmt.Println("t.Format(\"2006年1月2日 15:04:05\") is ", t.Format("2006年1月2日 15:04:05"))
	fmt.Println("t.Format(time.ANSIC) is ", t.Format(time.ANSIC))
	jsont, _ := t.MarshalJSON()
	fmt.Println("t.MarshalJSON is ", string(jsont))
	textt, _ := t.MarshalText()
	fmt.Println("t.MarshalText is ", string(textt))

	//d := time.Date(2017, time.January, 1, 12, 0, 0, 0, time.UTC)
	//fmt.Println(d.Local())
	//fmt.Println(d.Location())

}
