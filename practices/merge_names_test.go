package practices_test

import (
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
