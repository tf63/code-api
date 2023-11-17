package repository

import (
	"database/sql"

	"github.com/tf63/code-api/pkg/domain/entity"
	"github.com/tf63/code-api/pkg/domain/repository"
)

type patternCodeRepository struct {
	db *sql.DB
}

// インターフェースを実装しているかチェック
var _ repository.PatternCodeRepository = (*patternCodeRepository)(nil)

// コンストラクタ
func NewPatternCodeRepository(db *sql.DB) repository.PatternCodeRepository {
	return &patternCodeRepository{db}
}

// PatternCodeの取得
func (ptcr *patternCodeRepository) ReadPatternCode(findPatternCode entity.FindPatternCode) ([]entity.PatternCode, error) {

	// PatternCodeを取得するクエリ
	query := `SELECT pattern_code_id, content, nrow, created_at, language_id, pattern_id
				FROM pattern_code 
				WHERE language_id = $1 AND pattern_id = $2 AND nrow >= $3 AND nrow <= $4 
				ORDER BY pattern_code_id 
				OFFSET $5 LIMIT $6;
			`

	// プレースホルダへ渡す値
	args := []interface{}{
		findPatternCode.LanguageId, findPatternCode.PatternId,
		findPatternCode.StartRow, findPatternCode.EndRow,
		findPatternCode.Offset, findPatternCode.Limit,
	}

	// クエリを実行
	rows, err := ptcr.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	// 実行結果を戻り値に変換
	patternCodes := []entity.PatternCode{}
	for rows.Next() {
		patternCode := entity.PatternCode{}
		if err := rows.Scan(&patternCode.PatternCodeId, &patternCode.Content, &patternCode.Nrow,
			&patternCode.CreatedAt, &patternCode.LanguageId, &patternCode.PatternId); err != nil {
			return nil, err
		}
		patternCodes = append(patternCodes, patternCode)
	}

	return patternCodes, nil
}
