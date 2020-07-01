package main

import "fmt"

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
//  */

// func insert(intervals []Interval, newInterval []int]) res [][]int {
// 	res := make([]Interval, 0)
// 	if len(intervals) == 0 {
// 		res = append(res, newInterval)
// 		return res
// 	}
// 	curIndex := 0
// 	for curIndex < len(intervals) && intervals[curIndex].End < newInterval.Start {
// 		res = append(res, intervals[curIndex])
// 		curIndex++
// 	}

// 	for curIndex < len(intervals) && intervals[curIndex].Start <= newInterval.End {
// 		newInterval = Interval{Start: min(newInterval.Start, intervals[curIndex].Start), End: max(newInterval.End, intervals[curIndex].End)}
// 		curIndex++
// 	}
// 	res = append(res, newInterval)

// 	for curIndex < len(intervals) {
// 		res = append(res, intervals[curIndex])
// 		curIndex++
// 	}
// 	return res
// }

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	intervals = append(intervals, newInterval)
	for i := len(intervals) - 2; i >= 0; i-- {
		if intervals[i+1][0] > intervals[i][1] {
			//后一个interva的start值大于前一个interval的end，遍历结束
			break
		} else if intervals[i+1][1] < intervals[i][0] {
			//后一个interval的end值小于前一个interval的start值，交换之
			intervals[i], intervals[i+1] = intervals[i+1], intervals[i]
			continue
		} else {
			//后一个interval与前一个有交集，则合并这两个interval存储于前一个interval中，丢弃后一个
			intervals[i][0] = min(intervals[i][0], intervals[i+1][0])
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1])
			intervals = append(intervals[:i+1], intervals[i+2:]...)
		}
	}
	return intervals
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{3, 5}
	fmt.Println(insert(intervals, newInterval))
}
