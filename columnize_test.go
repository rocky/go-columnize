package columnize

import "testing"

func check_columnize(expect string, data []string, opts Opts_t, t *testing.T) {
	got := ColumnizeS(data, opts)
	if expect != got  {
		t.Errorf("got:\n%s\nwant:\n%s\n", got, expect)
	}
}

func TestColumnize(t *testing.T) {
	bools := []bool{true, false}
	for i := 0; i < len(bools); i++ {
		got := CellSize("abc", bools[i])
		if 3 != got  {
			t.Errorf("Cell size mismatch; got %d, want %d", got, 3)
		}
        }

	opts := Default_options()

	data := []string{"1", "2", "3"}

	opts.ColSep = ", "
	opts.DisplayWidth = 10
	check_columnize("1, 2, 3\n", data, opts, t)

	data = []string{"1", "2", "3", "4"}
	opts.ColSep = "  "
	opts.DisplayWidth = 4
	check_columnize("1  3\n2  4\n", data, opts, t)
	opts.ArrangeVertical = false
	check_columnize("1  2\n3  4\n", data, opts, t)
}
