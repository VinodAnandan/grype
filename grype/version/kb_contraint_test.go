package version

import (
	"errors"
	"testing"
)

func TestVersionKbConstraint(t *testing.T) {
	tests := []testCase{
		{version: "", constraint: "", satisfied: true},
		{version: "", constraint: "foo", satisfied: false},
		{version: "1", constraint: "foo", satisfied: false},
		{version: "1", constraint: "1", satisfied: true},
		{version: "878787", constraint: "979797 || 101010 || 878787", satisfied: true},
		{version: "478787", constraint: "979797 || 101010 || 878787", satisfied: false},
	}

	for _, test := range tests {
		t.Run(test.name(), func(t *testing.T) {
			constraint, err := newKBConstraint(test.constraint)
			if !errors.Is(err, test.constErr) {
				t.Fatalf("unexpected constraint error: '%+v'!='%+v'", err, test.constErr)
			}

			test.assert(t, KBFormat, constraint)
		})
	}
}
