package example

import (
	"testing"
)

func TestName(t *testing.T) {
	a := Example2{
		A: "111",
		B: &Example2_Nested{
			D: "1",
			E: 11,
		},
		//Status: Example2_EnumStatus(2),
		F: nil,
	}

	if err := a.Validate(); err != nil {
		t.Error(err)
	}
}
