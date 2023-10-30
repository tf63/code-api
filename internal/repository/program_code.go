package repository

import (
	"database/sql"

	"github.com/tf63/code-api/internal/model"
)

// ProgramCodeRepository
type ProgramCodeRepository interface {
	ReadProgramCode(findProgramCode model.FindProgramCode) (programCodes []model.ProgramCode, err error) // ProgramCodeの取得
}

type programCodeRepository struct {
	Db *sql.DB
}

// インターフェースを実装しているかチェック
var _ ProgramCodeRepository = (*programCodeRepository)(nil)

// コンストラクタ
func NewProgramCodeRepository(db *sql.DB) ProgramCodeRepository {
	return &programCodeRepository{db}
}

// ProgramCodeの取得
func (pgcr *programCodeRepository) ReadProgramCode(findProgramCode model.FindProgramCode) ([]model.ProgramCode, error) {

	// ProgramCodeを取得するクエリ
	query := `SELECT program_code_id, content, nrow, created_at, tool_id 
				FROM program_code 
				WHERE tool_id = $1 AND nrow >= $2 AND nrow <= $3 
				ORDER BY program_code_id 
				OFFSET $4 LIMIT $5;
			`

	// プレースホルダへ渡す値
	args := []interface{}{
		findProgramCode.ToolId, findProgramCode.StartRow, findProgramCode.EndRow, findProgramCode.Offset, findProgramCode.Limit,
	}

	// クエリを実行
	rows, err := pgcr.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	// 実行結果を戻り値に変換
	programCodes := []model.ProgramCode{}
	for rows.Next() {
		programCode := model.ProgramCode{}
		if err := rows.Scan(&programCode.ProgramCodeId, &programCode.Content, &programCode.Nrow, &programCode.CreatedAt, &programCode.ToolId); err != nil {
			return nil, err
		}
		programCodes = append(programCodes, programCode)
	}

	return programCodes, nil
}
