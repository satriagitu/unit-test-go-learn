package math_test

import (
	"testing"
	"time"
	. "unit-test-go-learn"
)

func TestAdd(t *testing.T) {
	testTable := []struct {
		a, b            int
		expectedOutcome int
	}{
		{a: 1, b: 2, expectedOutcome: 3},
		{a: -1, b: 2, expectedOutcome: 1},
	}

	for _, test := range testTable {
		result := Add(test.a, test.b)
		if result != test.expectedOutcome {
			t.Logf("Got: %d But expect %d", result, 3)
			t.Fail()
		}
	}

}

func TestAddWithHierarchical(t *testing.T) {
	a := 10
	t.Run("a=positive", func(t *testing.T) {
		result := Add(a, 5)
		if result != 15 {
			t.Logf("Got: %d But expect %d", result, 15)
			t.Fail()
		}
	})

	t.Run("a=negative", func(t *testing.T) {
		result := Add(a, -5)
		if result != 5 {
			t.Logf("Got: %d But expect %d", result, 5)
			t.Fail()
		}
	})
}

func TestNeedsToBeSkip(t *testing.T) {
	t.Skip("this test will be skipped")
}

func TestCallToDb(t *testing.T) {
	if testing.Short() {
		t.Skip("skip because calling db is way to long")
	}
	<-time.After(3 * time.Second)
}
