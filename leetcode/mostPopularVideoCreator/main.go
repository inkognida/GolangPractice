package main

import "fmt"

type Info struct {
	ID         string
	Views      int
	TotalViews int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	r := make([][]string, 0)
	maxViews := 0

	creatorsMap := make(map[string]*Info, 0)
	for i, v := range creators {
		if _, ok := creatorsMap[v]; ok {
			creatorsMap[v].TotalViews += views[i]
			if creatorsMap[v].TotalViews > maxViews {
				maxViews = creatorsMap[v].TotalViews
			}
			if views[i] > creatorsMap[v].Views {
				creatorsMap[v].Views = views[i]
				creatorsMap[v].ID = ids[i]
			} else if views[i] == creatorsMap[v].Views && ids[i] < creatorsMap[v].ID {
				creatorsMap[v].ID = ids[i]
			}
		} else {
			creatorsMap[v] = &Info{
				ID:         ids[i],
				TotalViews: views[i],
				Views:      views[i],
			}
			if creatorsMap[v].TotalViews > maxViews {
				maxViews = creatorsMap[v].TotalViews
			}
		}
	}

	localMin := creatorsMap[creators[0]]
	delete(creatorsMap, creators[0])
	for k, v := range creatorsMap {
		if v.TotalViews < maxViews {
			delete(creatorsMap, k)
		} else if v.TotalViews > maxViews {
			maxViews = v.TotalViews
		} else {
			continue
		}
	}

	for k, v := range creatorsMap {
		r = append(r, []string{k, v.ID})
	}

	if localMin.TotalViews == maxViews {
		r = append(r, []string{creators[0], localMin.ID})
	}

	return r
}

func main() {
	creators := []string{"msuuv", "hvmez", "smi", "eon", "uwcil"}
	ids := []string{"ag", "ycdt", "jy", "mfpj", "c"}
	views := []int{1, 3, 3, 2, 4}
	fmt.Println(mostPopularCreator(creators, ids, views), "a" < "b")
}
