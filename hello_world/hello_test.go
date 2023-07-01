package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty space string defaults to 'world'", func(t *testing.T) {
		got := Hello("  ", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Alberto", "es")
		want := "Hola, Alberto"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Elia", "fr")
		want := "Bonjour, Elia"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in portuguese", func(t *testing.T) {
		got := Hello("Tiago", "pt")
		want := "Olá, Tiago"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
