package stringSet

import (
	"bytes"
	"strings"
)

/*
 StringSet - an integer set
*/
type StringSet map[string]struct{}

/*
 returns a new StringSet
*/
func NewStringSet() StringSet {
	return make(StringSet)
}
/*
 adds a string to a StringSet
*/
func (s StringSet) Add(a string) {
	s[a] = struct{}{}
}

/*
 removes a string from a StringSet
*/
func (s StringSet) Remove(a string){
	delete(s, a)
}

/*
 returns true if s contains a, false otherwise
*/
func (s StringSet) Contains(a string) bool{
	_, ok := s[a]
	return ok
}

/*
 returns the union of s and r
*/
func (s StringSet) Union(r StringSet) StringSet {
	q := NewStringSet()
	for k,_ := range s {
		q[k] = struct{}{}
	}
	for k,_ := range r {
		q[k] = struct{}{}
	}
	return q
}

/*
 returns the intersection of s and r
*/
func (s StringSet) Intersect(r StringSet) StringSet {
	q:= NewStringSet()
	for k, _ := range s{
		if r.Contains(k){
			q[k] = struct{}{}
		}
	}
	return q
}

/*
 returns the set difference of s and r
*/
func (s StringSet) Diff(r StringSet) StringSet {
	q := NewStringSet()
	for k, _ := range s{
		if !r.Contains(k){
			q[k] = struct{}{}
		}
	}
	return q
}

/*
 returns the symmetric difference of s and r
*/
func (s StringSet) SymmDiff(r StringSet) StringSet {
	q := NewStringSet()
	for k, _ := range s{
		if !r.Contains(k){
			q[k] = struct{}{}
		}
	}
	for k, _ := range r{
		if !s.Contains(k){
			q[k] = struct{}{}
		}
	}
	return q
}

/*
  stringer function for a StringSet
*/
func (s StringSet) String() string {
	var buff bytes.Buffer
	buff.WriteString("{")
	for k, _ := range s{
		buff.WriteString(k + ", ")
	}
	buffString := buff.String()
	buffString = strings.TrimRight(buffString, ", ")
	return buffString + "}"
}