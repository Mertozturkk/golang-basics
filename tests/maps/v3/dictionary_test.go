package v3

import "testing"

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	defination := "this is a test message"
	dictionary.Add(word, defination)

	assertDefination(t, dictionary, word, defination)
}

func assertDefination(t testing.TB, dictionary Dictionary, word, defination string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if defination != got {
		t.Errorf("got %q want %q", got, defination)
	}
}
