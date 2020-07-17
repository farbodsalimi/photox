package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"photox/pkg/util"
)

func TestMakePathByTakenDateTime(t *testing.T) {
	dt := time.Unix(1581975738, 0)
	val := util.MakePathByTakenDateTime("base_path", dt)
	assert.Equal(t, val, "base_path/2020-2-17")
}
