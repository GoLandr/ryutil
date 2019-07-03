package ryutil

import (
	"fmt"
	"strings"
	"sort"
)

func HasId(j []string, id string) bool {
	if len(j) > 0 {
		for _, l := range j {
			if l == id {
				return true
			}
		}
		return false
	} else {
		return false
	}
}

/*
func Min(values ...int) int {
	min := values[0]
	for _, val := range values {
		if val <= min {
			min = val
		}
	}
	return min
}
*/
func MinIn(values []int) int {
	min := values[0]
	for _, val := range values {
		if val <= min {
			min = val
		}
	}
	return min
}

//Normal Function
func add(x int, y int) int { //add(x, y int)
	return x + y
}

//function with multiple return values
func multipleReturnValues(x int) (int, int) {
	return x + 1, x - 1
}

//funtion with named return values
func namedReturnType(x int) (add, mul int) {
	add = x + 1
	mul = x * x
	return // no need to specify the
}

//Vardiac function
func sum(Name string, nums ...int) {
	fmt.Println("name = ", Name)
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func makeEvenGenerator(i int) func(i int) int {
	return func(i int) (ret int) {
		ret = i
		i += 2
		return
	}
}

func test_main() {
	fmt.Println("hello World")
	fmt.Println(add(42, 13))
	add, sub := multipleReturnValues(5)
	fmt.Println("add = ", add, " sub = ", sub)
	fmt.Println(namedReturnType(5))
	sum("AAA", 1, 2, 3, 4, 5)

	nums := []int{1, 2, 3, 4}
	sum("BBB", nums...)

	nextEven := makeEvenGenerator(2)
	fmt.Println(nextEven(2)) // 0
	fmt.Println(nextEven(2)) // 2
	fmt.Println(nextEven(2)) // 4
}

/**===============
     INT ARRAY
 **=============**/
type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// length of the array
func (p *IntArray) Length() int {
	return len(*p)
}

// length of the array
func (p *IntArray) Size() int {
	return len(*p)
}

// add to the existing array
func (p *IntArray) Push(value ...int) {
	tmp := *p
	for _, v := range value {
		tmp = append(tmp, v)
	}
	*p = tmp
}

// remove last value of the array
func (p *IntArray) Pop() int {
	tmp := *p
	last := tmp[len(tmp)-1 : len(tmp)]
	tmp = tmp[0 : len(tmp)-1]

	*p = tmp
	return last[0]
}

// get first value of the array
func (p IntArray) First() int {
	return p[0]
}

// get last value of the array
func (p IntArray) Last() int {
	last := p[len(p)-1 : len(p)]
	return last[0]
}

// remove first value of the array
func (p *IntArray) Shift() {
	tmp := *p
	if len(tmp) < 1 {
		return
	}
	*p = tmp[1:len(tmp)]
}

// add value the first index to the existing array
func (p *IntArray) Unshift(arry IntArray) {
	tmp := *p
	arry.Concat(tmp)
	*p = arry
}

// add new array to the existing array
func (p *IntArray) Concat(arry IntArray) {
	tmp := *p
	tmp = append(tmp, arry...)
	*p = tmp
}

// initialize of the array
func (p *IntArray) Clear() {
	var tmp IntArray
	*p = tmp
}

// remove del_target from array
func (p *IntArray) Delete(del_value int) {
	tmp := *p
	var new_array IntArray
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != del_value {
			new_array = append(new_array, tmp[i])
		}
	}
	*p = new_array
}

// remove del_index of index from array
func (p *IntArray) Delete_at(del_index int) {
	tmp := *p
	var new_array IntArray
	for i := 0; i < len(tmp); i++ {
		if i != del_index {
			new_array = append(new_array, tmp[i])
		}
	}
	*p = new_array
}

// get index of value from array
func (p IntArray) Index(n int) int {
	for i, v := range p {
		if v == n {
			return i
		}
	}
	return -1
}

// get a new reverse array
func (p *IntArray) Reverse() {
	tmp := *p
	var new_array IntArray
	length := len(tmp)

	for i := length - 1; i >= 0; i-- {
		new_array = append(new_array, tmp[i])
	}
	*p = new_array
}

// get a new string from array
func (p *IntArray) Join(sep string) string {
	tmp := *p
	var str_ary []string

	for _, v := range tmp {
		str_ary = append(str_ary, fmt.Sprint(v))
	}
	str := strings.Join(str_ary, sep)
	return str
}

// get a new unique array from array
func (p *IntArray) Uniq() IntArray {
	tmp := *p
	var new_array IntArray

	for _, v := range tmp {
		if new_array.Index(v) == -1 {
			new_array = append(new_array, v)
		}
	}
	return new_array
}

// get a new sort array from array
func (p *IntArray) Sort() {
	tmp := *p
	// var tmp2 IntArray
	sort.Sort(tmp)
	*p = tmp
}

func StrArrContains(p []string, needle string) bool {
	for _, v := range p {
		if v == needle {
			return true
		}
	}
	return false
}

func Inter(arrs ...[]int) []int {
	res := []int{}
	x := arrs[0][0]
	i := 1
	for {
		off := sort.SearchInts(arrs[i], x)
		if off == len(arrs[i]) {
			// we emptied one slice, we're done.
			break
		}
		if arrs[i][off] == x {
			i++
			if i == len(arrs) {
				// x was in all the slices
				res = append(res, x)
				x++ // search for the next possible x.
				i = 0
			}
		} else {
			x = arrs[i][off]
			i = 0 // This can be done a bit more optimally.
		}
	}
	return res
}

// 截取字符串，start 开始下标，length 截取长度
/*
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
*/