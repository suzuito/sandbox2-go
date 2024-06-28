package rbac

type Action string

const (
	// 定義済みのアクション
	// リソースを取得する
	ActionGet Action = "get"
	// リソースを一覧する
	ActionList Action = "list"
	// リソースを作成する
	ActionCreate Action = "create"
	// リソースを更新する(ただし削除は除く)
	ActionUpdate Action = "update"
	// リソースを削除する
	ActionDelete Action = "delete"
)
