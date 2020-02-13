package utils

import (
	"container/list"
)

/*func main() {

	// a := []int{1, 2, 4}

	b := list.New()

	b.PushBack("1")
	b.PushBack("2")
	b.PushBack("3")

	fmt.Print(toArray(b))
	fmt.Print(toArray(b))

}*/

func ToArrayString(l *list.List) []string {
	arr := make([]string, l.Len())
	count := 0
	for item := l.Front(); item != nil; item = item.Next() {
		arr[count] = item.Value.(string)
		count++
	}
	return arr
}
