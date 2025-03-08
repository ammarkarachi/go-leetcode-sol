package solution

import (
	"slices"
)

type DSU struct {
	parent map[string]string
	size   map[string]int
}

func NewDSU() *DSU {
	return &DSU{make(map[string]string), make(map[string]int)}
}

func (this *DSU) MakeSet(x string) {
	if _, exists := this.parent[x]; !exists {
		this.parent[x] = x
		this.size[x] = 0
	}
}

func (this *DSU) Find(x string) string {
	if _, exists := this.parent[x]; !exists {
		this.MakeSet(x)
	}

	if this.parent[x] != x {
		this.parent[x] = this.Find(this.parent[x])
	}

	return this.parent[x]
}

func (this *DSU) Union(x string, y string) {
	repX := this.Find(x)
	repY := this.Find(y)

	if repX == repY {
		return
	}

	if this.size[repX] > this.size[repY] {
		this.parent[repY] = repX
	} else if this.size[repX] < this.size[repY] {
		this.parent[repX] = repY
	} else {
		this.parent[repY] = repX
		this.size[repX]++
	}
}

func accountsMerge(accounts [][]string) [][]string {
	dsu := NewDSU()

	emailToName := make(map[string]string)

	for _, account := range accounts {
		name := account[0]

		for i := 1; i < len(account); i++ {
			email := account[i]
			emailToName[email] = name

			dsu.MakeSet(email)

			if i > 1 {
				dsu.Union(account[1], email)
			}

		}
	}

	emailGroups := make(map[string][]string)

	for email := range emailToName {
		root := dsu.Find(email)
		emailGroups[root] = append(emailGroups[root], email)
	}

	results := make([][]string, len(emailGroups))
	idx := 0
	for rootEmail, emails := range emailGroups {
		name := emailToName[rootEmail]
		allEmails := make([]string, len(emails)+1)
		allEmails[0] = name
		slices.Sort(emails)
		for i, email := range emails {
			allEmails[i+1] = email
		}
		results[idx] = allEmails
		idx++
	}

	return results

}
