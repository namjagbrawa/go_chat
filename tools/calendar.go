package tools

import (
	"flag"
	"fmt"
)

var year int
var month int
var fs *flag.FlagSet
var showVersion, showUsage bool

func init() {
	//初始化解析参数
	fs = flag.NewFlagSet("calendar", flag.ExitOnError)
	fs.BoolVar(&showUsage, "help", false, "Show this message")
	fs.BoolVar(&showUsage, "h", false, "Show this message")
	fs.BoolVar(&showVersion, "version", false, "Print version information.")
	fs.BoolVar(&showVersion, "v", false, "Print version information.")

	fs.IntVar(&year, "year", 2019, "Input year.")
	fs.IntVar(&year, "y", 2019, "Input year.")
	fs.IntVar(&month, "month", 8, "Input month.")
	fs.IntVar(&month, "m", 8, "Input month.")
}

func perpetualCalendar(year int, month int) {
	var monthDays int = 0
	var totalDyas int = 0
	var isLeapYear bool = false
	fmt.Println("query date is:", year, month)

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		isLeapYear = true
		fmt.Println(year, "is leap year")
	} else {
		isLeapYear = false
		fmt.Println(year, "is not lear year")
	}

	//计算距离1900年的天数
	for i := 1900; i < year; i++ {
		if i%400 == 0 || (i%4 == 0 && i%100 != 0) {
			totalDyas += 366
		} else {
			totalDyas += 365
		}
	}

	//计算截至到当月1号前的总天数
	for j := 1; j <= month; j++ {
		switch j {
		case 1, 3, 5, 7, 8, 10, 12:
			monthDays = 31
			break
		case 2:
			if isLeapYear {
				monthDays = 29
			} else {
				monthDays = 28
			}
			break
		case 4, 6, 9, 11:
			monthDays = 30
			break
		default:
			fmt.Println("input month is error.")
		}

		if j != month {
			totalDyas += monthDays
		}
	}

	//计算当月1号是星期几
	var weekDay int = 0
	weekDay = 1 + totalDyas%7
	if weekDay == 7 {
		weekDay = 0
	}

	fmt.Println("weekDay is: ", weekDay)

	//展示日期
	fmt.Println("星期日\t" + "星期一\t" + "星期二\t" + "星期三\t" + "星期四\t" + "星期五\t" + "星期六\t")
	for k := 0; k < weekDay; k++ {
		fmt.Printf("\t")
	}

	for m := 1; m <= monthDays; m++ {
		fmt.Printf("%d\t", m)
		if (weekDay+m)%7 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

/*func main() {
	//开启main 函数的解析功能
	fs.Parse(os.Args[1:])
	if showUsage {
		fs.PrintDefaults()
		os.Exit(0)
	}
	if showVersion {
		fmt.Println("version:v1.0")
		os.Exit(0)
	}
	perpetualCalendar(year, month)
}*/
