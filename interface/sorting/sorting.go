package sorting

import "fmt"

// Sorter interface
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// IntArray sorting
type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// StringArray sorting
type StringArray []string

func (p StringArray) Len() int           { return len(p) }
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// DayArray sorting
type DayArray struct {
	data []*day
}

type day struct {
	num       int
	shortName string
	longName  string
}

func (d DayArray) Len() int           { return len(d.data) }
func (d DayArray) Less(i, j int) bool { return d.data[i].num < d.data[j].num }
func (d DayArray) Swap(i, j int)      { d.data[i], d.data[j] = d.data[j], d.data[i] }

func SortInts(a []int)       { Sort(IntArray(a)) }
func SortStrings(a []string) { Sort(StringArray(a)) }
func SortDays(a DayArray)    { Sort(a) }

func IntsAreSorted(a []int) bool       { return IsSorted(IntArray(a)) }
func StringsAreSorted(a []string) bool { return IsSorted(StringArray(a)) }

func Sorting() {
	fmt.Println("sorting int array")
	ia := []int{1, 12, 5, 30, 76, 3, 2}
	SortInts(ia)
	fmt.Println(ia)

	fmt.Println("sorting string array")
	sa := []string{"dart", "delphi", "rust", "javascript", "python", "ruby"}
	SortStrings(sa)
	fmt.Println(sa)

	fmt.Println("sorting day array")
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednessday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	Sunday := day{0, "SUN", "Sunday"}
	data := []*day{&Friday, &Sunday, &Thursday, &Saturday, &Monday, &Wednesday, &Tuesday}
	da := DayArray{data: data}
	Sort(&da)
	if !IsSorted(&da) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}
