package main

import (
	"fmt"
	"time"
)

var ValDate string

const layout = "2006-01-02"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	date := "1999-12-31"
	t, _ := time.Parse(layout, date)
	ValDate = t.Format(layout)
	year, month, _ := t.Date()
	fmt.Println(t, ValDate, year, month.String(), int(month), time.November, t.Month())
	fmt.Println(Sunday)

	fmt.Println(CToF(12), FToC(0), float64(CToF(12)))

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
