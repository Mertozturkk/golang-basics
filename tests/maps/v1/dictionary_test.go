package v1

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a just test"}
	got := Search(dictionary, "test")
	want := "this is a just test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")

	}
}
