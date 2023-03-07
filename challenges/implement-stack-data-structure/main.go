package main

import "strconv"

const LIMIT = 4 // DONOT CHANGE IT!

type Stack struct {
	ix   int // first free position, so data[ix] == 0
	data [LIMIT]int
}

func (st *Stack) Push(n int) {
	if st.ix == LIMIT {
		return
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	if st.ix > 0 {
		st.ix--
		element := st.data[st.ix]
		st.data[st.ix] = 0
		return element
	}
	return -1
}

func (st Stack) String() string {
	str := ""
	for ix := 0; ix < st.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "]"
	}

	return str
}
