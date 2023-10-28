package testhelper

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestCaseForFirestore struct {
	SetUp    func(ctx context.Context, fcli *firestore.Client) error
	TearDown func(ctx context.Context, fcli *firestore.Client) error
	Assert   func(ctx context.Context, fcli *firestore.Client, ass *FirestoreAssertion) error
}

func (th *TestCaseForFirestore) Run(
	ctx context.Context,
	desc string,
	t *testing.T,
	run func(t *testing.T, fcli *firestore.Client),
) bool {
	return t.Run(desc, func(t *testing.T) {
		fcli, err := NewFirestoreClient(ctx)
		if err != nil {
			t.Errorf("failed to NewFirestoreClient : %+v", err)
			t.Fail()
			return
		}
		defer func() {
			if th.TearDown != nil {
				th.TearDown(ctx, fcli)
			}
			fcli.Close()
		}()
		if th.SetUp != nil {
			if err := th.SetUp(ctx, fcli); err != nil {
				t.Errorf("failed to DeleteDocuments : %+v", err)
				t.Fail()
				return
			}
		}
		run(t, fcli)
		if th.Assert != nil {
			ass := FirestoreAssertion{Cli: fcli}
			th.Assert(ctx, fcli, &ass)
		}
	})
}

func NewFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8081")
	fcli, err := firestore.NewClient(ctx, "dummy-prj")
	if err != nil {
		return nil, err
	}
	return fcli, err
}

type SetDocumentsRef struct {
	Ref  *firestore.DocumentRef
	Data any
}

func SetDocuments(
	ctx context.Context,
	fcli *firestore.Client,
	data ...SetDocumentsRef,
) error {
	return fcli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		for _, d := range data {
			if err := tx.Set(d.Ref, d.Data); err != nil {
				return terrors.Wrap(err)
			}
		}
		return nil
	})
}

func DeleteDocuments(
	ctx context.Context,
	fcli *firestore.Client,
	data *firestore.CollectionRef,
) error {
	return fcli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		it := data.DocumentRefs(ctx)
		for {
			doc, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			if err := tx.Delete(doc); err != nil {
				return err
			}
		}
		return nil
	})
}

type FirestoreAssertion struct {
	Cli *firestore.Client
}

func (th *FirestoreAssertion) ExistsDoc(
	ctx context.Context,
	t *testing.T,
	ref *firestore.DocumentRef,
) {
	snp, err := ref.Get(ctx)
	if status.Code(err) == codes.NotFound {
		t.Errorf("'%s' does not exists", ref.Path)
		t.Fail()
		return
	}
	if err != nil {
		t.Errorf("failed to ref.Get : %+v", err)
		t.Fail()
		return
	}
	if !snp.Exists() {
		t.Errorf("'%s' does not exists", ref.Path)
		t.Fail()
	}
}

func (th *FirestoreAssertion) EqualsDoc(
	ctx context.Context,
	t *testing.T,
	ref *firestore.DocumentRef,
	expected map[string]interface{},
) {
	snp, err := ref.Get(ctx)
	if status.Code(err) == codes.NotFound {
		t.Errorf("'%s' does not exists", ref.Path)
		t.Fail()
		return
	}
	if err != nil {
		t.Errorf("failed to ref.Get : %+v", err)
		t.Fail()
		return
	}
	if !snp.Exists() {
		t.Errorf("'%s' does not exists", ref.Path)
		t.Fail()
	}
	actual := snp.Data()
	assert.Equal(t, expected, actual)
}
