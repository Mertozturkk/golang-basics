package v2

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a just test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a just test"
		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")

	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
