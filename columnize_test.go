package columnize

import "testing"

// func assert_equal(a, b int, errmsg msg) {
//	t.Errorf(
// }

func TestColumnize(t *testing.T) {
	bools := []bool{true, false}
	for i := 0; i < len(bools); i++ {
		got := Cell_size("abc", bools[i])
		if 3 != got  {
			t.Errorf("Cell size mismatch; got %d, want %d", got, 3)
		}
        }
 	var opts Opts_t
	Default_options(&opts)

	data := []string{"1", "2", "3"}

	expect := "1, 2, 3\n"
	opts.Colsep = ", "
	opts.Displaywidth = 10
	got := Columnize(data, opts)
	if expect != got  {
		t.Errorf("got:\n%s\nwant:\n%s\n", got, expect)
	}

}
