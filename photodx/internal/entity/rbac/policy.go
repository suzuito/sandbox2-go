package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/google/cel-go/cel"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/pbrbac"
)

type Policy interface {
	EvalGinContext(
		permissions []*pbrbac.Permission,
		principalPhotoStudioMemberID string,
		principalPhotoStudioID string,
		ctx *gin.Context,
	) (bool, error)
}

var celEnv *cel.Env

type PolicyImpl struct {
	policyString string
	program      cel.Program
}

func (t *PolicyImpl) EvalGinContext(
	permissions []*pbrbac.Permission,
	principalPhotoStudioMemberID string,
	principalPhotoStudioID string,
	ctx *gin.Context,
) (bool, error) {
	pathParams := map[string]string{}
	for _, param := range ctx.Params {
		pathParams[param.Key] = param.Value
	}
	input := map[string]any{
		"permissions":                  permissions,
		"principalPhotoStudioMemberId": string(principalPhotoStudioMemberID),
		"principalPhotoStudioId":       string(principalPhotoStudioID),
		"pathParams":                   pathParams,
	}
	return t.Eval(input)
}

func (t *PolicyImpl) Eval(input any) (bool, error) {
	output, _, err := t.program.Eval(input)
	if err != nil {
		return false, terrors.Wrap(err)
	}
	if output.Type() != cel.BoolType {
		return false, terrors.Wrapf("result of cel.Eval is not boolean type")
	}
	raw, _ := output.Value().(bool)
	return raw, nil
}

func NewPolicy(
	policyString string,
) Policy {
	ast, issues := celEnv.Compile(policyString)
	if issues.Err() != nil {
		panic(terrors.Wrap(issues.Err()))
	}
	prog, err := celEnv.Program(ast)
	if err != nil {
		panic(terrors.Wrap(err))
	}
	return &PolicyImpl{
		policyString: policyString,
		program:      prog,
	}
}
