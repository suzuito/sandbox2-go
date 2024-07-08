package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/google/cel-go/cel"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/pbrbac"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity/rbac"
)

type AdminPrincipalAccessToken interface {
	GetPhotoStudioMemberID() PhotoStudioMemberID
	GetPhotoStudioID() PhotoStudioID
	GetRoles() []*rbac.Role
	GetPermissions() []*pbrbac.Permission
}

type AdminPrincipalAccessTokenImpl struct {
	PhotoStudioMemberID PhotoStudioMemberID
	PhotoStudioID       PhotoStudioID
	Roles               []*rbac.Role
}

func (t *AdminPrincipalAccessTokenImpl) GetPhotoStudioMemberID() PhotoStudioMemberID {
	return t.PhotoStudioMemberID
}

func (t *AdminPrincipalAccessTokenImpl) GetPhotoStudioID() PhotoStudioID {
	return t.PhotoStudioID
}

func (t *AdminPrincipalAccessTokenImpl) GetRoles() []*rbac.Role {
	return t.Roles
}

func (t *AdminPrincipalAccessTokenImpl) GetPermissions() []*pbrbac.Permission {
	permissions := []*pbrbac.Permission{}
	for _, role := range t.GetRoles() {
		permissions = append(permissions, role.Permissions...)
	}
	return permissions
}

type PolicyAdminPrincipalAccessToken interface {
	EvalGinContext(
		principal AdminPrincipalAccessToken,
		ctx *gin.Context,
	) (bool, error)
}

type PolicyAdminPrincipalAccessTokenImpl struct {
	policyString string
	program      cel.Program
}

func (t *PolicyAdminPrincipalAccessTokenImpl) EvalGinContext(
	principal AdminPrincipalAccessToken,
	ctx *gin.Context,
) (bool, error) {
	pathParams := map[string]string{}
	for _, param := range ctx.Params {
		pathParams[param.Key] = param.Value
	}
	input := map[string]any{
		"permissions":                       principal.GetPermissions(),
		"adminPrincipalPhotoStudioMemberId": string(principal.GetPhotoStudioMemberID()),
		"adminPrincipalPhotoStudioId":       string(principal.GetPhotoStudioID()),
		"pathParams":                        pathParams,
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

var celEnvAdminPrincipalAccessToken *cel.Env

func initCELEnvAdminPrincipalAccessToken() {
	var err error
	celEnvAdminPrincipalAccessToken, err = cel.NewEnv(
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
			"adminPrincipalPhotoStudioMemberId",
			cel.StringType,
		),
		cel.Variable(
			"adminPrincipalPhotoStudioId",
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

func NewPolicyAdminPrincipalAccessToken(
	policyString string,
) PolicyAdminPrincipalAccessToken {
	ast, issues := celEnvAdminPrincipalAccessToken.Compile(policyString)
	if issues.Err() != nil {
		panic(terrors.Wrap(issues.Err()))
	}
	prog, err := celEnvAdminPrincipalAccessToken.Program(ast)
	if err != nil {
		panic(terrors.Wrap(err))
	}
	return &PolicyAdminPrincipalAccessTokenImpl{
		policyString: policyString,
		program:      prog,
	}
}
