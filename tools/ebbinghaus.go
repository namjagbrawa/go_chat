package tools

import (
	"container/list"
	"encoding/json"
	"github.com/namjagbrawa/go_chat/exception"
	"github.com/namjagbrawa/go_chat/log"
	"github.com/namjagbrawa/go_chat/utils"
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

	log.Logger.Println("start_date = ", startDate, "days = ", days)

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
		ebbT := list.New()
		ebbT.PushBack(out.Format(LAYOUT))

		for j := 1; j <= len(gap); j++ {
			ddn, _ := time.ParseDuration("-" + strconv.Itoa(24*gap[j-1]) + "h")
			in := out.Add(ddn)
			if in.Unix() < start.Unix() {
				break
			}
			ebbT.PushBack(in.Format(LAYOUT))
		}
		ebb[i] = utils.ToArrayString(ebbT)
		ebbT.Init()
	}

	for i := 0; i < len(ebb); i++ {
		log.Logger.Println(ebb[i])
	}

	ebbJson, _ := json.Marshal(ebb)

	log.Logger.Println(string(ebbJson))

	return ebbJson
}

/*func main() {
	buildEbbinghausDate("2020-02-12", "100")
}*/
