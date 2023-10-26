package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Presenters interface {
	Response(
		ctx *gin.Context,
		arg any,
	)
}

type PresentersImpl struct {
	PresenterHTML          *PresenterHTML
	PresenterJSON          *PresenterJSON
	PresenterXML           *PresenterXML
	PresenterStandardError *PresenterStandardError
	PresenterString        *PresenterString
	PresenterRedirect      *PresenterRedirect
}

func (t *PresentersImpl) Response(
	ctx *gin.Context,
	arg any,
) {
	switch v := arg.(type) {
	case PresenterArgHTML:
		t.PresenterHTML.Response(ctx, &v)
		return
	case PresenterArgJSON:
		t.PresenterJSON.Response(ctx, &v)
		return
	case PresenterArgXML:
		t.PresenterXML.Response(ctx, &v)
		return
	case PresenterArgString:
		t.PresenterString.Response(ctx, &v)
		return
	case PresenterArgStandardError:
		t.PresenterStandardError.Response(ctx, &v)
		return
	case PresenterArgRedirect:
		t.PresenterRedirect.Response(ctx, &v)
		return
	}
	t.PresenterStandardError.Response(ctx, &PresenterArgStandardError{
		Err: fmt.Errorf("No presenter arg %+v", arg),
	})
}

func NewPresenters() Presenters {
	return &PresentersImpl{
		PresenterHTML:          &PresenterHTML{},
		PresenterJSON:          &PresenterJSON{},
		PresenterXML:           &PresenterXML{},
		PresenterStandardError: &PresenterStandardError{},
		PresenterString:        &PresenterString{},
		PresenterRedirect:      &PresenterRedirect{},
	}
}
