package tree

import "testing"

func TestTrie(t *testing.T) {
	tt := Constructor()
	tt.Insert("apple")
	t.Logf("apple: %v", tt.Search("apple"))
	t.Logf("app: %v", tt.Search("app"))
	t.Logf("app: %v", tt.StartsWith("app"))
	tt.Insert("app")
	t.Logf("app: %v", tt.Search("app"))
}
