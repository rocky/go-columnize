package main
import (
	"columnize"
	"fmt"
)

func main() {
 	var opts columnize.Opts_t
	columnize.Default_options(&opts)
	
 	// line = 'require [1;29m"[0m[1;37mirb[0m[1;29m"[0m';
 	line := "testing"
 	fmt.Println(columnize.Cell_size(line, true))
 	fmt.Println(columnize.Cell_size(line, false))

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

	opts.Arrange_vertical = true
	fmt.Println(columnize.Columnize(data, opts))
	opts.Displaywidth = 50
	opts.Ljustify = false
	fmt.Println(columnize.Columnize(data, opts))

	opts.Arrange_vertical = true
 	fmt.Println(columnize.Columnize(data, opts))
	opts.Displaywidth = 80
	opts.Ljustify = true
	fmt.Println(columnize.Columnize(data, opts))

}
