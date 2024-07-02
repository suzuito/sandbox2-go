package rbac

import (
	"github.com/google/cel-go/cel"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity/pbrbac"
)

func init() {
	var err error
	celEnv, err = cel.NewEnv(
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
			"principalPhotoStudioMemberId",
			cel.StringType,
		),
		cel.Variable(
			"principalPhotoStudioId",
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
