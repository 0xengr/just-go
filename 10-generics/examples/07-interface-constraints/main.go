package main

import "fmt"

// NOTE: An interface can be a constraint
// * Any interface can restrict a type parameter, not just any and comparable.
// * Describer says a type can describe itself. Labeled stores a value of any
// * type that implements Describer and puts a label next to it.

type Describer interface {
	Describe() string
}

type Labeled[T Describer] struct {
	Label string
	Value T
}

func (l Labeled[T]) Show() string {
	return l.Label + " -> " + l.Value.Describe()
}

// NOTE: Unwrap returns T, not Describer. A plain interface field would erase
// * the concrete type; the type parameter keeps it, so the caller can still use
// * the fields and methods of the original type.
func (l Labeled[T]) Unwrap() T {
	return l.Value
}

type Book struct {
	Title string
	Pages int
}

func (b Book) Describe() string {
	return fmt.Sprintf("%s (%d pages)", b.Title, b.Pages)
}

type Movie struct {
	Title string
	Year  int
}

func (m Movie) Describe() string {
	return fmt.Sprintf("%s (%d)", m.Title, m.Year)
}

func main() {
	book := Labeled[Book]{
		Label: "book",
		Value: Book{Title: "Learning Go", Pages: 450},
	}
	fmt.Println(book.Show())
	fmt.Println(book.Unwrap().Pages)

	movie := Labeled[Movie]{
		Label: "movie",
		Value: Movie{Title: "Arrival", Year: 2016},
	}
	fmt.Println(movie.Show())

	// BUG: a type without Describe does not satisfy the constraint:
	// type Plain struct{ Name string }
	// fmt.Println(Labeled[Plain]{Label: "plain", Value: Plain{Name: "x"}})
	// Plain does not satisfy Describer (missing method Describe)
}
