// Code generated by "stringer -type=tpchBench"; DO NOT EDIT.

package main

import "strconv"

const _tpchBench_name = "sql20"

var _tpchBench_index = [...]uint8{0, 5}

func (i tpchBench) String() string {
	if i < 0 || i >= tpchBench(len(_tpchBench_index)-1) {
		return "tpchBench(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _tpchBench_name[_tpchBench_index[i]:_tpchBench_index[i+1]]
}
