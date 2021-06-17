package main

import (
	"io"
	"sort"
	"time"

	"github.com/turnon/smzdm/smzdm"
)

type resultSet struct {
	searches  []*smzdm.Search
	createdAt time.Time
}

func (rs *resultSet) collect(s *smzdm.Search) {
	if len(rs.searches) == 0 {
		rs.createdAt = time.Now()
	}
	rs.searches = append(rs.searches, s)
}

func (rs *resultSet) Len() int {
	return len(rs.searches)
}

func (rs *resultSet) Swap(i, j int) {
	rs.searches[i], rs.searches[j] = rs.searches[j], rs.searches[i]
}

func (rs *resultSet) Less(i, j int) bool {
	return rs.searches[i].Index < rs.searches[j].Index
}

func (rs *resultSet) sort() {
	sort.Sort(rs)
}

type output interface {
	collect(*smzdm.Search)
	sort()
	print(...io.Writer)
}
