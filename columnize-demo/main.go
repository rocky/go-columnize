package main
import (
	"github.com/rocky/go-columnize"
	"fmt"
)

func main() {
	opts := columnize.DefaultOptions()

 	// line = 'require [1;29m"[0m[1;37mirb[0m[1;29m"[0m';
 	line := "testing"
 	fmt.Println(columnize.CellSize(line, true))
 	fmt.Println(columnize.CellSize(line, false))

	list := make([] string, 0)
 	fmt.Println(columnize.Columnize(list, opts))
	// for _, pair := [{4, 4}, {4, 7}, {100, 80}] {
	// 	width, num = pair
	// 	for i := 1;  i <= num; i++ {
	// 		for _, pair := [{false 'horizontal'},
        //                              {true, 'vertical'}] {
	// 			bool, dir := pair
	// 			opts[displaywidth] = width
	// 			opts[colsep] =  '  ',
	// 			opts[arrange_vertical] = bool
	// 			fmt.Println("Width: %d %s", width, dir)
	// 			fmt.Println(columnize.Columnize(data, opts))
	// 		}
	// 	}
	// }

 	// fmt.Println Columnize::columnize.Columnize(5)
 	// fmt.Println columnize.Columnize(["a", 2, "c"], :displaywidth =>10,
        //                                 :colsep => ', ')
	// fmt.Println columnize.Columnize(["oneitem"])
 	// fmt.Println columnize.Columnize(["one", "two", "three"])
	data := []string{
		"one",       "two",         "three",
		"for",       "five",        "six",
	 	"seven",     "eight",       "nine",
	 	"ten",       "eleven",      "twelve",
	 	"thirteen",  "fourteen",    "fifteen",
	 	"sixteen",   "seventeen",   "eightteen",
	 	"nineteen",  "twenty",      "twentyone",
	 	"twentytwo", "twentythree", "twentyfour",
	 	"twentyfive","twentysix",   "twentyseven"}

	opts.ArrangeVertical = true
	fmt.Println(columnize.ColumnizeS(data, opts))
	opts.DisplayWidth = 50
	opts.LJustify = false
	fmt.Println(columnize.ColumnizeS(data, opts))

	opts.ArrangeVertical = true
 	fmt.Println(columnize.ColumnizeS(data, opts))
	opts.DisplayWidth = 80
	opts.LJustify = true
	fmt.Println(columnize.ColumnizeS(data, opts))
	fmt.Println("----------------")

	a := []int{31, 4, 1, 59, 2, 6, 5, 3}
	// opts.ArrangeArray = true
	opts.ArrangeVertical = false
	opts.LJustify = false

	opts.DisplayWidth = 8
	fmt.Println(columnize.Columnize(a, opts))
	fmt.Println("----------------")

	opts.ArrangeArray = true
	fmt.Println(columnize.Columnize(a, opts))

}
