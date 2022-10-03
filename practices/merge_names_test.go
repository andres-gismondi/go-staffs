package practices_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-staffs/practices"
)

func TestMergeNames_success(t *testing.T) {
	want := []string{"andres", "matias", "julian"}

	none := []string{"andres", "andres", "matias"}
	ntwo := []string{"andres", "julian"}

	merge := practices.MergeNames{}
	got := merge.UniqueNames(none, ntwo)

	assert.Equal(t, want, got)
}

func Uno(x, y int) int {
	if x == 0 {
		return y
	}
	a := Uno(x-1, x+y)
	return a
}
func TestUno(t *testing.T) {
	fmt.Println(Uno(5, 2))
}
