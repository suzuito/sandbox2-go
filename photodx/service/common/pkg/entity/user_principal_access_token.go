package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/google/cel-go/cel"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
)

type UserPrincipalAccessToken interface {
	GetUserID() UserID
	IsGuest() bool
	GetPermissions() []*pbrbac.Permission
}

type PolicyUserPrincipalAccessToken interface {
	EvalGinContext(
		principal UserPrincipalAccessToken,
		ctx *gin.Context,
	) (bool, error)
}

type PolicyUserPrincipalAccessTokenImpl struct {
	policyString string
	program      cel.Program
}

func (t *PolicyUserPrincipalAccessTokenImpl) EvalGinContext(
	principal UserPrincipalAccessToken,
	ctx *gin.Context,
) (bool, error) {
	pathParams := map[string]string{}
	for _, param := range ctx.Params {
		pathParams[param.Key] = param.Value
	}
	input := map[string]any{
		"permissions":         principal.GetPermissions(),
		"userPrincipalUserId": string(principal.GetUserID()),
		"pathParams":          pathParams,
	}

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

var celEnvUserPrincipalAccessToken *cel.Env

func initCELEnvUserPrincipalAccessToken() {
	var err error
	celEnvUserPrincipalAccessToken, err = cel.NewEnv(
		cel.Types(
			&pbrbac.Permission{},
		),
		cel.Variable(
			"permissions",
			cel.ListType(
				cel.ObjectType("pbrbac.Permission"),
			),
		),
		cel.Variable(
			"userPrincipalUserId",
			cel.StringType,
		),
		cel.Variable(
			"pathParams",
			cel.MapType(
				cel.StringType,
				cel.StringType,
			),
		),
	)
	if err != nil {
		panic(terrors.Wrap(err))
	}
}

func NewPolicyUserPrincipalAccessToken(
	policyString string,
) PolicyUserPrincipalAccessToken {
	ast, issues := celEnvUserPrincipalAccessToken.Compile(policyString)
	if issues.Err() != nil {
		panic(terrors.Wrap(issues.Err()))
	}
	prog, err := celEnvUserPrincipalAccessToken.Program(ast)
	if err != nil {
		panic(terrors.Wrap(err))
	}
	return &PolicyUserPrincipalAccessTokenImpl{
		policyString: policyString,
		program:      prog,
	}
}
