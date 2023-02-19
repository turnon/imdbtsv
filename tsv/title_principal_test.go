package tsv

import "testing"

func TestTitlePrincipalUniqCategory(t *testing.T) {
	cates := TitlePrincipalUniqCategory("testdata/title.principals.tsv")
	t.Log(cates)
}
