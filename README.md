A Go module to format a simple (i.e. not nested) slice into aligned
columns. A string with embedded newline characters is returned.

For example:

```go
import (
  "fmt"
   "code.google.com/p/go-columnize"
)

opts := columnize.DefaultOptions()
data := []string{"1", "2", "3", "4"}
opts.DisplayWidth = 6
fmt.Println(columnize.Columnize(data, opts)
```

gives output:
```
1  3
2  4
```

while

```go
opts.Displaywidth = 6
opts.ColSep = ', '
opts.ArrangeVertical = false
fmt.Println(columnize.Columnize(data, opts)
```

gives output:

```
1, 2
3, 4
```

and

```go
data := []string{"1", "2", "3", "4", "5"}
opts.DisplayWidth = 8
print columnize.Columnize(data, opts)
```

gives output:

```
1  3  5
2  4
```

By default entries are left justified, Columns are separated by two spaces.

Each column is only as wide as necessary. Set `opts.ColSep` to adjust the string separate columns. Set `opts.DisplayWidth` to set the line width.

Normally, consecutive items go down from the top to bottom from the left-most column to the right-most. If `opts.ArrangeVertical` is set false, consecutive items will go across, left to right, top to bottom.

A list of all options:

* `DisplayWidth`:  the display width (`int`)
* `ColSep`: the column separator (`string`)
* `ArrayPrefix`: string to prefix the entire list with (`string`)
* `ArraySuffix` : string to suffix the entire list with (`string`)
* `LinePrefix`: string to add after each newline (`string`)
* `LJustify`: whether to left justify text instead of right justify (`bool`)
* `ArrangeArray`: whether to format as an array. This is really a combination of setting the `ArrayPrefix`, `ArraySuffix`, the `LinePrefix` and the `ColSep`


This package (essentially one function) is port of my [Ruby gem](https://rubygems.org/gems/columnize) which in turn is a port of my Python module [pycolumnize](http://code.google.com/p/pycolumnize).
