
# Unit Test - GO

Learn Unit Test
- how to write a test package file 
- table driven / test table
- test sort and skip
- test driven development



## Example

```golang
// Package Name
package math_test
package stack_test


// Table Drive or Test Table
testTable := []struct {
		a, b            int
		expectedOutcome int
	}{
		{a: 1, b: 2, expectedOutcome: 3},
		{a: -1, b: 2, expectedOutcome: 1},
	}


// Test Sort and Skip
func TestNeedsToBeSkip(t *testing.T) {
	t.Skip("this test will be skipped")
}

func TestCallToDb(t *testing.T) {
	if testing.Short() {
		t.Skip("skip because calling db is way to long")
	}
	<-time.After(3 * time.Second)
}


// Test Driven Development
// Requirements:
// - A stack is empty on construction
// - A stack has size 0 on construction

func TestNewSet(t *testing.T) {
	t.Run("A stack is empty on construction", func(t *testing.T) {
		s := stack.New()
		assert.True(t, s.IsEmpty())
	})

	t.Run("A stack has size 0 on construction", func(t *testing.T) {
		s := stack.New()
		assert.Equal(t, 0, s.Size())
	})
}
```


## How to Run

Run math_test

```bash
  go test / go test -v
```

Run stack_test

```bash
  cd stack
  go test / go test -v
```
