// Format a string array into a single string with embedded newlines.
// On printing the string the columns are aligned.
//
// == Summary
// Format a simple (i.e. not nested) slice into aligned columns. 
// A string with embedded newline characters is returned.
//
// For example, for a line width of 4 characters (arranged vertically):
//        ['1', '2,', '3', '4'] => '1  3\n2  4\n'
//   
//    or arranged horizontally:
//        ['1', '2,', '3', '4'] => '1  2\n3  4\n'
//        
// Each column is only as wide as necessary.  By default, columns are
// separated by two spaces. Options are avalable for setting
//  * `DisplayWidth`:  the display width (`int`)
//  * `Colsep: the column separator (`string`)
//  * `Lineprefix`: the line prefix (`string`)
//  * `LJustify`: whether to left justify text instead of right justify (`bool`)
//  * `ArrayPrefix`: string to prefix the entire list with (`string`)
//  * `ArraySuffix` : string to suffix the entire list with (`string`)
//  * `LinePrefix`: string to add after each newline (`string`)
//  * `ArrangeArray`: whether to format as an array. This is really a combination of setting the `ArrayPrefix`, `ArraySuffix`, the `LinePrefix` and the `ColSep`
//   
//
// == License 
//
// Columnize is copyright (C) 2013 Rocky Bernstein
// <rocky@gnu.org>
//
// Copyright 2013 Rocky Bernstein.
//
// Adapted from the routine of the same name from Ruby.

package columnize

import "fmt"
import "reflect"

type Opts_t struct {
    ArrangeArray    bool
    ArrangeVertical bool
    ArrayPrefix     string
    ArraySuffix     string
    ColSep          string
    DisplayWidth    int
    LinePrefix      string
    LJustify        bool
    TermAdjust      bool
}

func Default_options() Opts_t {
	var opts Opts_t
	opts.ArrangeArray    = false
	opts.ArrangeVertical = true
	opts.ArrayPrefix     = ""
	opts.ArraySuffix     = ""
	opts.ColSep          = "  "
	opts.DisplayWidth    = 80
	opts.LinePrefix      = ""
	opts.LJustify        = true
	opts.TermAdjust      = false
	return opts
}

// Return the length of String +cell+. If Boolean +term_adjust+ is true,
// ignore terminal sequences in +cell+.
func CellSize(cell string, term_adjust bool) int {
	return len(cell)
}

func max(a, b int) int { 
	if a > b {return a } 
	return b
}

// The following two routines are from Carlos Castillo. Thanks Carlos! 
// http://play.golang.org/p/bxdcIj6ueH
// This only works (no panic) if x is value which
// has a length, and can be indexed (a slice/array)
func ToStringArray(x interface{}) []string {
	v := reflect.ValueOf(x)
	out := make([]string, v.Len())
	for i := range out {
		out[i] = fmt.Sprint(v.Index(i).Interface())
	}
	return out
}

// This version works as above for slices/arrays, and
// treats anything else as a single item
func ToStringArray2(x interface{}) []string {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return []string{fmt.Sprint(x)}
	}
	return ToStringArray(x)
}


// Return a list of strings with embedded newlines (\n) as a compact
// set of columns arranged horizontally or vertically.
//
// For example, for a line width of 4 characters (arranged vertically):
//     ['1', '2,', '3', '4'] => '1  3\n2  4\n'
	
// or arranged horizontally:
//     ['1', '2,', '3', '4'] => '1  2\n3  4\n'
//     
// Each column is only as wide as possible, no larger than
// +displaywidth'.  If +list+ is not an array, the empty string, '',
// is returned. By default, columns are separated by two spaces - one
// was not legible enough. Set +colsep+ to adjust the string separate
// columns. If +arrange_vertical+ is set false, consecutive items
// will go across, left to right, top to bottom.


func Columnize(list interface{}, opts Opts_t) string {
	l := ToStringArray2(list)
	return ColumnizeS(l, opts)
}

func ColumnizeS(list [] string, opts Opts_t) string {
	if len(list) == 0 { 
		result :=
			fmt.Sprintf("%s%s", 
			opts.ArrayPrefix, opts.ArraySuffix)
		return result
	}

	if len(list) == 1 { 
		result := 
			fmt.Sprintf("%s%s%s", 
			opts.ArrayPrefix, list[0], opts.ArraySuffix)
		return result
	}
	if opts.DisplayWidth - len(opts.LinePrefix) < 4 {
		opts.DisplayWidth = len(opts.LinePrefix)+ 4
	} else {
		opts.DisplayWidth -= len(opts.LinePrefix)
	}
	var ncols, nrows int
	var colwidths [] int
	if opts.ArrangeVertical {
		array_index := func(num_rows, row, col int) int  {
	 		return num_rows*col + row 
	 	}
		// Try every row count from 1 upwards
		for nrows = 1; nrows < len(list); nrows++ {
			ncols = (len(list) + nrows-1) / nrows
			colwidths = make([] int, 0)
			totwidth := -len(opts.ColSep)
			
			for col := 0; col < ncols; col++ {
				// get max column width for this column
				colwidth := 0
				for row := 0; row < nrows; row++ {
					i := array_index(nrows, row, col)
					if i >= len(list) { break }
					colwidth = max(CellSize(list[i], opts.TermAdjust),
					               colwidth)
					}
				colwidths = append(colwidths, colwidth)
				totwidth += colwidth + len(opts.ColSep)
				if totwidth > opts.DisplayWidth {
					ncols = col
					break
				}
			}
			if totwidth <= opts.DisplayWidth {
				break 
			}
		}
		if ncols < 1 {ncols = 1}
		if ncols == 1 { nrows = len(list) }
		// The smallest number of rows computed and the max widths for
		// each column has been obtained.  Now we just have to format
		// each of the rows.
		s := ""
		for row := 0; row < nrows; row++ {
			texts := make([] string, 0)
			for col := 0; col < ncols; col++ {
				var x string
				i := array_index(nrows, row, col)
				if i >= len(list) {
					x = ""
				} else {
					x = list[i]
				}
				texts = append(texts, x)
			}
			// texts.pop while !texts.empty? and texts[-1] == ''
			if len(texts) > 0 {
				for col := 0; col < len(texts); col++ {
					if ncols != 1  {
						var fmt_str string
						if opts.LJustify {
							fmt_str = fmt.Sprintf("%%%ds", -colwidths[col])
							texts[col] = fmt.Sprintf(fmt_str, texts[col])
						} else {
							fmt_str = fmt.Sprintf("%%%ds", colwidths[col])
							texts[col] = fmt.Sprintf(fmt_str, texts[col])
						}
					}
				}
				line := ""
				for i := 0; i <len(texts)-1; i++ {
					line += fmt.Sprintf("%s%s", texts[i], opts.ColSep)
				}
				if len(texts) > 0 {
					line += fmt.Sprintf("%s\n", texts[len(texts)-1])
				}
				s += line
			}
		}
		return s
	} else {
		array_index := func(ncols, row, col int) int {
			return ncols*(row-1) + col 
		}
		// Assign to make enlarge scope of loop variables.
		var totwidth, i, rounded_size int
		var ncols, nrows int
		// Try every column count from size downwards.
		for ncols = len(list); ncols >= 1; ncols-- {
			// Try every row count from 1 upwards
			min_rows := (len(list)+ncols-1) / ncols
			for nrows = min_rows; nrows <= (len(list)); nrows++ {
				rounded_size = nrows * ncols
				colwidths = [] int { }
				totwidth = -len(opts.ColSep)
				var colwidth, row int
				for col := 0; col < ncols; col++ {
					// get max column width for this column
					for row = 1; row <= nrows; row++ {
						i = array_index(ncols, row, col)
						if i >= len(list) { break }
						colwidth = max(colwidth, CellSize(list[i], opts.TermAdjust))
					}
					colwidths = append(colwidths, colwidth)
					totwidth += colwidth + len(opts.ColSep)
					if totwidth > opts.DisplayWidth { break };
				}
				if totwidth <= opts.DisplayWidth {
					// Found the right nrows and ncols
					nrows  = row
					break
				} else { 
					if totwidth > opts.DisplayWidth {
						// Need to reduce ncols
						break
					}
				}
			}
			if totwidth <= opts.DisplayWidth && i >= rounded_size-1 {
				break
			}
		}
		if ncols < 1 { ncols = 1 }
		if ncols == 1 { nrows = len(list) }
		// The smallest number of rows computed and the max widths for
		// each column has been obtained.  Now we just have to format
		// each of the rows.
		s := ""
		var prefix string
		if len(opts.ArrayPrefix) == 0 {
            prefix = opts.LinePrefix 
        } else {
            prefix =  opts.ArrayPrefix
        }
		for row := 1; row <=nrows; row++ {
			texts := make([] string, 0)
			for col := 0;  col < ncols; col++ {
				var x string
				i = array_index(ncols, row, col)
				if i >= len(list) {
					break
				} else {
					x = list[i]
				}
				texts = append(texts, x)
			}
			for col := 0; col < len(texts); col++ {
				if ncols != 1  {
					var fmt_str string
					if opts.LJustify {
						fmt_str = fmt.Sprintf("%%%ds", -colwidths[col])
						texts[col] = fmt.Sprintf(fmt_str, texts[col])
					} else {
						fmt_str = fmt.Sprintf("%%%ds", colwidths[col])
						texts[col] = fmt.Sprintf(fmt_str, texts[col])
					}
				}
			}
			line := prefix
			for i := 0; i <len(texts)-1; i++ {
				line += fmt.Sprintf("%s%s", texts[i], opts.ColSep)
			}
			if len(texts) > 0 {
				line += fmt.Sprintf("%s\n", texts[len(texts)-1])
			}
			s += line
			prefix = opts.LinePrefix
		}
		s += opts.ArraySuffix
		return s
	}
	return "Not reached"
}
