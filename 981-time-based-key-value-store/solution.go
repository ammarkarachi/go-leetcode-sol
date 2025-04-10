package solution

type Value struct {
	Value     string
	Timestamp int
}
type TimeMap struct {
	keys map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{
		keys: make(map[string][]Value),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.keys[key] = append(this.keys[key], Value{
		Value:     value,
		Timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {

	items := this.keys[key]

	l := 0
	r := len(items)
	for l < r {
		mid := (l + r) / 2
		if items[mid].Timestamp <= timestamp {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if r == 0 {
		return ""
	}

	return items[l].Value

}
