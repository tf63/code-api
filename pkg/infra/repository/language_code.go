package repository

import (
	"database/sql"

	"github.com/tf63/code-api/pkg/domain/entity"
	"github.com/tf63/code-api/pkg/domain/repository"
)

type languageCodeRepository struct {
	Db *sql.DB
}

// インターフェースを実装しているかチェック
var _ repository.LanguageCodeRepository = (*languageCodeRepository)(nil)

// コンストラクタ
func NewLanguageCodeRepository(db *sql.DB) repository.LanguageCodeRepository {
	return &languageCodeRepository{db}
}

// LanguageCodeの取得
func (pgcr *languageCodeRepository) ReadLanguageCode(findLanguageCode entity.FindLanguageCode) ([]entity.LanguageCode, error) {

	// LanguageCodeを取得するクエリ
	query := `SELECT language_code_id, content, nrow, created_at, language_id 
				FROM language_code 
				WHERE language_id = $1 AND nrow >= $2 AND nrow <= $3 
				ORDER BY language_code_id 
				OFFSET $4 LIMIT $5;
			`

	// プレースホルダへ渡す値
	args := []interface{}{
		findLanguageCode.LanguageId, findLanguageCode.StartRow, findLanguageCode.EndRow, findLanguageCode.Offset, findLanguageCode.Limit,
	}

	// クエリを実行
	rows, err := pgcr.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	// 実行結果を戻り値に変換
	languageCodes := []entity.LanguageCode{}
	for rows.Next() {
		languageCode := entity.LanguageCode{}
		if err := rows.Scan(&languageCode.LanguageCodeId, &languageCode.Content, &languageCode.Nrow, &languageCode.CreatedAt, &languageCode.LanguageId); err != nil {
			return nil, err
		}
		languageCodes = append(languageCodes, languageCode)
	}

	return languageCodes, nil
}
