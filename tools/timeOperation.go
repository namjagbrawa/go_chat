package tools

import (
	"fmt"
	"strings"
	"time"
)

func test() {
	// Add 时间相加
	now := time.Now()
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1)

	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)

	printSplit(50)

	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)

	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	printSplit(50)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)

	//获取时间戳

	timestamp := time.Now().Unix()

	fmt.Println(timestamp)

	//格式化为字符串,tm为Time类型

	tm := time.Unix(timestamp, 0)

	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))

	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串

	tm2, _ := time.Parse("01/02/2006", "02/08/2015")

	fmt.Println(tm2.Unix())
}

func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}

/*func main() {

	gap := [5]int{1, 2, 4, 7, 15}

	ebb := make([][]string, 100)

	dd, _ := time.ParseDuration("24h")

	start, _ := time.Parse("2006-01-02", "2020-02-20")

	out, _ := time.Parse("2006-01-02", "2020-02-20")

	ebb[0] = []string{out.Format("2006-01-02")}
	for i := 1; i < 100; i++ {
		out = out.Add(dd)
		ebbT := make([]string, 6)
		ebbT[0] = out.Format("2006-01-02")

		for j := 1; j <= len(gap); j++ {
			ddn, _ := time.ParseDuration("-" + strconv.Itoa(24*gap[j-1]) + "h")
			in := out.Add(ddn)
			if in.Unix() < start.Unix() {
				break
			}
			ebbT[j] = in.Format("2006-01-02")
		}
		ebb[i] = ebbT
	}

	for i := 0; i < len(ebb); i++ {
		fmt.Println(ebb[i])
	}

	ebbJson, _:= json.Marshal(ebb)

	fmt.Println(string(ebbJson))

}*/
