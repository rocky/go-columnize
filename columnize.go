// Format a string array into a single string with embedded newlines.
// On printing the string the columns are aligned.
//
// == Summary
// Display a list of strings as a compact set of columns.
//
//   For example, for a line width of 4 characters (arranged vertically):
//        ['1', '2,', '3', '4'] => '1  3\n2  4\n'
//   
//    or arranged horizontally:
//        ['1', '2,', '3', '4'] => '1  2\n3  4\n'
//        
// Each column is only as wide as necessary.  By default, columns are
// separated by two spaces. Options are avalable for setting
// * the display width
// * the column separator
// * the line prefix
// * whether to ignore terminal codes in text size calculation
// * whether to left justify text instead of right justify
//
// == License 
//
// Columnize is copyright (C) 2013 Rocky Bernstein
// <rocky@gnu.org>
//
// All rights reserved.  You can redistribute and/or modify it under
// the same terms as Ruby.
//
// Adapted from the routine of the same name in Python +cmd.py+.

package columnize

import (
	"fmt"
)
	
type Opts_t struct {
    Arrange_array    bool
    Arrange_vertical bool
    Array_prefix     string
    Array_suffix     string
    Colsep           string
    Displaywidth     int
    Lineprefix       string
    Ljustify         bool
    Term_adjust      bool
}

func Default_options(opts *Opts_t) {
	opts.Arrange_array    = false
	opts.Arrange_vertical = true
	opts.Array_prefix     = ""
	opts.Array_suffix     = ""
	opts.Colsep           = "  "
	opts.Displaywidth     = 80
	opts.Lineprefix       = ""
	opts.Ljustify         = true
	opts.Term_adjust      = false
}

// Return the length of String +cell+. If Boolean +term_adjust+ is true,
// ignore terminal sequences in +cell+.
func Cell_size(cell string, term_adjust bool) int {
	return len(cell)
}

func max(a, b int) int { 
	if a > b {return a } 
	return b
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


func Columnize(list [] string, opts Opts_t) string {
	if len(list) == 0 { 
		result :=
			fmt.Sprintf("%s%s", 
			opts.Array_prefix, opts.Array_suffix)
		return result
	}

    l := make([] string, len(list))
	for i:= 0; i<len(list); i++ {
		l[i] = fmt.Sprintf("%s", list[i])
	}

	if len(list) == 1 { 
		result := 
			fmt.Sprintf("%s%s%s", 
			opts.Array_prefix, l[0], opts.Array_suffix)
		return result
	}
	if opts.Displaywidth - len(opts.Lineprefix) < 4 {
		opts.Displaywidth = len(opts.Lineprefix)+ 4
	} else {
		opts.Displaywidth -= len(opts.Lineprefix)
	}
	var ncols, nrows int
	var colwidths [] int
	if opts.Arrange_vertical {
		array_index := func(num_rows, row, col int) int  {
	 		return num_rows*col + row 
	 	}
		// Try every row count from 1 upwards
		for nrows = 1; nrows < len(list); nrows++ {
			ncols = (len(list) + nrows-1) / nrows
			colwidths = make([] int, 0)
			totwidth := -len(opts.Colsep)
			
			for col := 0; col < ncols; col++ {
				// get max column width for this column
				colwidth := 0
				for row := 0; row < nrows; row++ {
					i := array_index(nrows, row, col)
					if i >= len(list) { break }
					colwidth = max(Cell_size(l[i], opts.Term_adjust),
					               colwidth)
					}
				colwidths = append(colwidths, colwidth)
				totwidth += colwidth + len(opts.Colsep)
				if totwidth > opts.Displaywidth {
					ncols = col
					break
				}
			}
			if totwidth <= opts.Displaywidth {
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
					x = l[i]
				}
				texts = append(texts, x)
			}
			// texts.pop while !texts.empty? and texts[-1] == ''
			if len(texts) > 0 {
				for col := 0; col < len(texts); col++ {
					if ncols != 1  {
						var fmt_str string
						if opts.Ljustify {
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
					line += fmt.Sprintf("%s%s", texts[i], opts.Colsep)
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
				totwidth = -len(opts.Colsep)
				var colwidth, row int
				for col := 0; col < ncols; col++ {
					// get max column width for this column
					for row = 1; row <= nrows; row++ {
						i = array_index(ncols, row, col)
						if i >= len(list) {	break }
						colwidth = max(colwidth, Cell_size(l[i], opts.Term_adjust))
					}
					colwidths = append(colwidths, colwidth)
					totwidth += colwidth + len(opts.Colsep)
					if totwidth > opts.Displaywidth { break };
				}
				if totwidth <= opts.Displaywidth {
					// Found the right nrows and ncols
					nrows  = row
					break
				} else { 
					if totwidth > opts.Displaywidth {
						// Need to reduce ncols
						break
					}
				}
			}
			if totwidth <= opts.Displaywidth && i >= rounded_size-1 {
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
		if len(opts.Array_prefix) == 0 {
            prefix = opts.Lineprefix 
        } else {
            prefix =  opts.Array_prefix
        }
		for row := 1; row <=nrows; row++ {
			texts := make([] string, 0)
			for col := 0;  col < ncols; col++ {
				var x string
				i = array_index(ncols, row, col)
				if i >= len(list) {
					break
				} else {
					x = l[i]
				}
				texts = append(texts, x)
			}
			for col := 0; col < len(texts); col++ {
				if ncols != 1  {
					var fmt_str string
					if opts.Ljustify {
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
				line += fmt.Sprintf("%s%s", texts[i], opts.Colsep)
			}
			if len(texts) > 0 {
				line += fmt.Sprintf("%s\n", texts[len(texts)-1])
			}
			s += line
			prefix = opts.Lineprefix
		}
		s += opts.Array_suffix
		return s
	}
	return "Not reached"
}
