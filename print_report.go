package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
 REPORT for %s
=============================
    `, baseURL)

	type page struct {
		url   string
		count int
	}

	var pagesVisited []page
	for url, count := range pages {
		pagesVisited = append(pagesVisited, page{url: url, count: count})
	}

	sort.Slice(pagesVisited, func(i, j int) bool {
		if pagesVisited[i].count == pagesVisited[j].count {
			return pagesVisited[i].url < pagesVisited[j].url
		}
		return pagesVisited[i].count > pagesVisited[j].count
	})

	for _, p := range pagesVisited {
		fmt.Printf("Found %d internal links to %s\n", p.count, p.url)
	}
}
