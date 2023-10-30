package v4

import "testing"

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test message"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test message"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, defination string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if defination != got {
		t.Errorf("got %q want %q", got, defination)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
