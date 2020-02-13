package tools

import (
	"encoding/json"
	"fmt"
	"github.com/namjagbrawa/go_chat/exception"
	"net/http"
	"strconv"
	"time"
)

func Ebbinghaus(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	startDate := query.Get("start_date")
	days := query.Get("days")

	if startDate == "" || days == "" {
		err := ""
		if startDate == "" {
			err = err + "param start_date is need \n"
		}
		if days == "" {
			err = err + "param days is need \n"
		}
		w.Write([]byte(err))

		return exception.UserError("参数不正确")
	}

	// todo 异常捕获
	// todo 参数校验返回

	fmt.Println("start_date = ", startDate, "days = ", days)

	resp := buildEbbinghausDate(startDate, days)

	w.Write(resp)

	return nil
}

func buildEbbinghausDate(startDate string, days string) []byte {
	LAYOUT := "2006-01-02"

	gap := [5]int{1, 2, 4, 7, 15}

	n, _ := strconv.Atoi(days)

	ebb := make([][]string, n)

	dd, _ := time.ParseDuration("24h")

	start, _ := time.Parse(LAYOUT, startDate)

	out, _ := time.Parse(LAYOUT, startDate)

	ebb[0] = []string{out.Format(LAYOUT)}
	for i := 1; i < n; i++ {
		out = out.Add(dd)
		ebbT := make([]string, 6)
		ebbT[0] = out.Format(LAYOUT)

		for j := 1; j <= len(gap); j++ {
			ddn, _ := time.ParseDuration("-" + strconv.Itoa(24*gap[j-1]) + "h")
			in := out.Add(ddn)
			if in.Unix() < start.Unix() {
				break
			}
			ebbT[j] = in.Format(LAYOUT)
		}
		ebb[i] = ebbT
	}

	for i := 0; i < len(ebb); i++ {
		fmt.Println(ebb[i])
	}

	ebbJson, _ := json.Marshal(ebb)

	fmt.Println(string(ebbJson))

	return ebbJson
}

/*func main() {
	buildEbbinghausDate("2020-02-12", "100")
}*/
