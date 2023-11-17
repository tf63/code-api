package repository

import (
	"database/sql"

	"github.com/tf63/code-api/pkg/domain/entity"
	"github.com/tf63/code-api/pkg/domain/repository"
)

type frameworkCodeRepository struct {
	db *sql.DB
}

// インターフェースを実装しているかチェック
var _ repository.FrameworkCodeRepository = (*frameworkCodeRepository)(nil)

// コンストラクタ
func NewFrameworkCodeRepository(db *sql.DB) repository.FrameworkCodeRepository {
	return &frameworkCodeRepository{db}
}

// FrameworkCodeの取得
func (pgcr *frameworkCodeRepository) ReadFrameworkCode(findFrameworkCode entity.FindFrameworkCode) ([]entity.FrameworkCode, error) {

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
	rows, err := pgcr.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	// 実行結果を戻り値に変換
	frameworkCodes := []entity.FrameworkCode{}
	for rows.Next() {
		frameworkCode := entity.FrameworkCode{}
		if err := rows.Scan(&frameworkCode.FrameworkCodeId, &frameworkCode.Content, &frameworkCode.Nrow, &frameworkCode.CreatedAt, &frameworkCode.ToolId); err != nil {
			return nil, err
		}
		frameworkCodes = append(frameworkCodes, frameworkCode)
	}

	return frameworkCodes, nil
}
