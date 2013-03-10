/* 

Module to format a slice or array into a single string with embedded
newlines, On printing the string, the columns are aligned.

Summary

Return a string from an array with embedded newlines formatted
so that when printed the columns are aligned. 

	data := []int{1, 2, 3, 4}
	opts := columnize.DefaultOptions()
	opts.DisplayWidth = 10
	fmt.Println columnize.Columnize(data, opts)
	# prints "1  3\n2  4\n"


Options:
	ArrangeArray bool	format string as a go array
	ArrangeVertical bool	format entries top down and left to right
				Otherwise left to right and top to bottom
	ArrayPrefix string	Add this to the beginning of the string
	ArraySuffix string	Add this to the end of the string
	DisplayWidth int	the maximum line display width
	CellFmt			A format specify for formatting each item each array 
				item to a string
	ColSep string		Add this string between columns
	LinePrefix string	Add this prefix for each line
	LineSuffix string	Add this suffix for each line
	LAdjustify bool		whether to left justify text instead of right justify
	TermAdjust bool 	whether to ignore terminal codes in text size 
				calculation

Examples

	data := []int{1, 2, 3, 4}
	opts := columnize.DefaultOptions()
	fmt.Println columnize.Columnize(data, opts)

	opts.DisplayWidth = 8
	opts.CellFmt = "%02d"
	fmt.Println columnize.Columnize(data, opts)
    
	opts.ArrangeArray = true
	opts.CellFmt = ""
	fmt.Println columnize.Columnize(data, opts)

Author

	Rocky Bernstein	<rocky@gnu.org>

	Also available in Python (columnize), and Perl
	(Array::Columnize) and Ruby (columnize)

	Copyright 2013 Rocky Bernstein.

*/
package columnize
const VERSION string = "1.0"
