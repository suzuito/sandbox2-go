package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase/markdown2html"
	"go.uber.org/mock/gomock"
)

type PresentersDummy struct {
	RealArg any
}

func (t *PresentersDummy) Response(
	ctx *gin.Context,
	arg any,
) {
	t.RealArg = arg
}

func (t *PresentersDummy) AssertResponse(
	tt *testing.T,
	expectedArg any,
) {
	if assert.IsType(tt, expectedArg, t.RealArg) {
		switch expectedArgV := expectedArg.(type) {
		case PresenterArgHTML:
			realArgV := t.RealArg.(PresenterArgHTML)
			assert.Equal(tt, expectedArgV.Code, realArgV.Code)
			assert.Equal(tt, expectedArgV.Name, realArgV.Name)
			assert.EqualValues(tt, expectedArgV.Obj, realArgV.Obj)
			return
		case PresenterArgJSON:
			realArgV := t.RealArg.(PresenterArgJSON)
			assert.Equal(tt, expectedArgV.Code, realArgV.Code)
			assert.EqualValues(tt, expectedArgV.Obj, realArgV.Obj)
			return
		case PresenterArgXML:
			realArgV := t.RealArg.(PresenterArgXML)
			assert.Equal(tt, expectedArgV.Code, realArgV.Code)
			assert.EqualValues(tt, expectedArgV.Obj, realArgV.Obj)
			return
		case PresenterArgString:
			realArgV := t.RealArg.(PresenterArgString)
			assert.Equal(tt, expectedArgV.Code, realArgV.Code)
			assert.EqualValues(tt, expectedArgV.Body, realArgV.Body)
			return
		case PresenterArgStandardError:
			realArgV := t.RealArg.(PresenterArgStandardError)
			if expectedArgV.Err == nil {
				assert.Fail(tt, "Err is nil")
				return
			}
			assert.EqualValues(tt, expectedArgV.Err.Error(), realArgV.Err.Error())
			return
		case PresenterArgRedirect:
			realArgV := t.RealArg.(PresenterArgRedirect)
			assert.Equal(tt, expectedArgV.Code, realArgV.Code)
			assert.Equal(tt, expectedArgV.Location, realArgV.Location)
			return
		default:
			assert.Failf(tt, "Unknown expected arg type", "%+v", expectedArg)
		}
	}
}

func NewPresentersDummy() *PresentersDummy {
	return &PresentersDummy{}
}

type mocks struct {
	UC                      *usecase.MockUsecase
	RepositoryArticle       *usecase.MockRepositoryArticle
	RepositoryArticleSource *usecase.MockRepositoryArticleSource
	RepositoryArticleHTML   *usecase.MockRepositoryArticleHTML
	Markdown2HTML           *markdown2html.MockMarkdown2HTML
	PresentersDummy         *PresentersDummy
}

func newControllerImplForTest(ctrl *gomock.Controller, setting *WebSetting) (*ControllerImpl, *mocks) {
	m := &mocks{
		UC:                      usecase.NewMockUsecase(ctrl),
		RepositoryArticle:       usecase.NewMockRepositoryArticle(ctrl),
		RepositoryArticleSource: usecase.NewMockRepositoryArticleSource(ctrl),
		RepositoryArticleHTML:   usecase.NewMockRepositoryArticleHTML(ctrl),
		Markdown2HTML:           markdown2html.NewMockMarkdown2HTML(ctrl),
		PresentersDummy:         NewPresentersDummy(),
	}
	return &ControllerImpl{
		UC:                      m.UC,
		RepositoryArticle:       m.RepositoryArticle,
		RepositoryArticleSource: m.RepositoryArticleSource,
		RepositoryArticleHTML:   m.RepositoryArticleHTML,
		Markdown2HTML:           m.Markdown2HTML,
		Setting:                 setting,
		Presenters:              m.PresentersDummy,
	}, m
}

func newWebSetting() *WebSetting {
	return &WebSetting{
		DirPathTemplates: "templates",
		DirPathCSS:       "css",
		DirPathImages:    "images",
	}
}

type testCaseWebEndpoint struct {
	desc                 string
	inputRequestFunc     func() *http.Request
	setUp                func(m *mocks)
	beforeMiddleware     gin.HandlerFunc
	expectedPresenterArg any
}

func (t *testCaseWebEndpoint) String() string {
	r := t.inputRequestFunc()
	headers := []string{}
	for k, vv := range r.Header {
		for _, v := range vv {
			headers = append(headers, fmt.Sprintf("%s=%s", k, v))
		}
	}
	headersString := strings.Join(headers, ",")
	return fmt.Sprintf(
		"%s/path=%s/headers=%s/%s",
		r.Method,
		r.URL.Path,
		headersString,
		t.desc,
	)
}

func testWebEndpoint(
	setting *WebSetting,
	tC *testCaseWebEndpoint,
	setRouterFunc SetRouterFunc,
) func(*testing.T) {
	return func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		gin.DefaultWriter = io.Discard
		w := httptest.NewRecorder()
		router := gin.New()
		c, m := newControllerImplForTest(ctrl, setting)
		if tC.setUp != nil {
			tC.setUp(m)
		}
		if tC.beforeMiddleware != nil {
			router.Use(tC.beforeMiddleware)
		}
		setRouterFunc(router, c)
		router.ServeHTTP(w, tC.inputRequestFunc())
		m.PresentersDummy.AssertResponse(t, tC.expectedPresenterArg)
	}
}

func TestGetHealth(t *testing.T) {
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/health", nil)
				return r
			},
			expectedPresenterArg: PresenterArgJSON{
				Code: http.StatusOK,
				Obj:  map[string]any{"message": "ok"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(
			tC.String(),
			testWebEndpoint(newWebSetting(), &tC, SetRouter),
		)
	}
}

func TestGetTestLog(t *testing.T) {
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/test-log", nil)
				return r
			},
			expectedPresenterArg: PresenterArgJSON{
				Code: http.StatusOK,
				Obj:  map[string]any{"message": "ok"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(
			tC.String(),
			testWebEndpoint(newWebSetting(), &tC, SetRouter),
		)
	}
}

func TestMiddlewareXRobotsTag(t *testing.T) {
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/health", nil)
				return r
			},
			expectedPresenterArg: PresenterArgJSON{
				Code: http.StatusOK,
				Obj:  map[string]any{"message": "ok"},
			},
		},
	}
	setting := newWebSetting()
	setting.NoIndex = true
	for _, tC := range testCases {
		t.Run(
			tC.String(),
			testWebEndpoint(setting, &tC, SetRouter),
		)
	}
}
