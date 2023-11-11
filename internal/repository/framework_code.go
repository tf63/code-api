package repository

import (
	"database/sql"

	"github.com/tf63/code-api/internal/model"
)

// FrameworkCodeRepository
type FrameworkCodeRepository interface {
	ReadFrameworkCode(findFrameworkCode model.FindFrameworkCode) (frameworkCodes []model.FrameworkCode, err error) // FrameworkCodeの取得
}

type frameworkCodeRepository struct {
	Db *sql.DB
}

// インターフェースを実装しているかチェック
var _ FrameworkCodeRepository = (*frameworkCodeRepository)(nil)

// コンストラクタ
func NewFrameworkCodeRepository(db *sql.DB) FrameworkCodeRepository {
	return &frameworkCodeRepository{db}
}

// FrameworkCodeの取得
func (pgcr *frameworkCodeRepository) ReadFrameworkCode(findFrameworkCode model.FindFrameworkCode) ([]model.FrameworkCode, error) {

	// FrameworkCodeを取得するクエリ
	query := `SELECT framework_code_id, content, nrow, created_at, tool_id 
				FROM framework_code 
				WHERE tool_id = $1 AND nrow >= $2 AND nrow <= $3 
				ORDER BY framework_code_id 
				OFFSET $4 LIMIT $5;
			`

	// プレースホルダへ渡す値
	args := []interface{}{
		findFrameworkCode.ToolId, findFrameworkCode.StartRow, findFrameworkCode.EndRow, findFrameworkCode.Offset, findFrameworkCode.Limit,
	}

	// クエリを実行
	rows, err := pgcr.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	// 実行結果を戻り値に変換
	frameworkCodes := []model.FrameworkCode{}
	for rows.Next() {
		frameworkCode := model.FrameworkCode{}
		if err := rows.Scan(&frameworkCode.FrameworkCodeId, &frameworkCode.Content, &frameworkCode.Nrow, &frameworkCode.CreatedAt, &frameworkCode.ToolId); err != nil {
			return nil, err
		}
		frameworkCodes = append(frameworkCodes, frameworkCode)
	}

	return frameworkCodes, nil
}
