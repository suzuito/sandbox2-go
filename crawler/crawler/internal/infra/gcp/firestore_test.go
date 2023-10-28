package gcp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/testhelper"
)

func Test_firestoreDocTimeSeriesData(t *testing.T) {
	fcli, err := testhelper.NewFirestoreClient(context.Background())
	if err != nil {
		t.Errorf("NewFirestoreClient is failed : %+v", err)
		t.Fail()
		return
	}
	repo := Repository{
		fcli:                    fcli,
		firestoreBaseCollection: "foo",
	}
	actual := repo.firestoreDocTimeSeriesData("hoge", "fuga")
	assert.Equal(t, "projects/dummy-prj/databases/(default)/documents/foo/TimeSeriesData/hoge/fuga", actual.Path)
}
