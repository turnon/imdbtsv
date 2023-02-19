package tsv

import (
	"testing"
)

func TestRowCount(t *testing.T) {
	testRowCount(t, IterateNameBasic, "testdata/name.basics.tsv", func(row *NameBasicRow) error { return nil })
	testRowCount(t, IterateTitleBasic, "testdata/title.basics.tsv", func(row *TitleBasicRow) error { return nil })
	testRowCount(t, IterateTitlePrincipal, "testdata/title.principals.tsv", func(row *TitlePrincipalRow) error { return nil })
	testRowCount(t, IterateTitleRating, "testdata/title.ratings.tsv", func(row *TitleRatingRow) error { return nil })
}

func testRowCount[
	Row any,
	RowHandler func(row Row) error,
	IteratFn func(string, RowHandler) error,
](t *testing.T,
	iteratFn IteratFn,
	path string,
	rowHandler RowHandler) {

	expectedCount, err := wcCountLines(path)
	if err != nil {
		t.Fatal(err)
	}

	actualRowCount, wrappedRowHandler := countCall(rowHandler)

	iteratFn(path, wrappedRowHandler)

	if actualRowCount() != expectedCount {
		t.Errorf("%s is expected %d rows, but actually %d", path, expectedCount, actualRowCount())
		return
	}

	t.Logf("%s has %d rows", path, actualRowCount())
}

func countCall[row any](fn func(r row) error) (func() int, func(row) error) {
	var count int
	counter := func() int {
		return count
	}
	wrap := func(r row) error {
		count += 1
		return fn(r)
	}
	return counter, wrap
}
