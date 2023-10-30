package repository

import (
	"database/sql"

	"github.com/tf63/code-api/internal/model"
)

// AlgorithmCodeRepository
type AlgorithmCodeRepository interface {
	ReadAlgorithmCode(findAlgorithmCode model.FindAlgorithmCode) (algorithmCodes []model.AlgorithmCode, err error) // AlgorithmCodeの取得
}

type algorithmCodeRepository struct {
	Db *sql.DB
}

// インターフェースを実装しているかチェック
var _ AlgorithmCodeRepository = (*algorithmCodeRepository)(nil)

// コンストラクタ
func NewAlgorithmCodeRepository(db *sql.DB) AlgorithmCodeRepository {
	return &algorithmCodeRepository{db}
}

// AlgorithmCodeの取得
func (arcr *algorithmCodeRepository) ReadAlgorithmCode(findAlgorithmCode model.FindAlgorithmCode) ([]model.AlgorithmCode, error) {

	// AlgorithmCodeを取得するクエリ
	query := `SELECT algorithm_code_id, content, nrow, created_at, language_id, algorithm_id
				FROM algorithm_code 
				WHERE language_id = $1 AND algorithm_id = $2 AND nrow >= $3 AND nrow <= $4 
				ORDER BY algorithm_code_id 
				OFFSET $5 LIMIT $6;
			`

	// プレースホルダへ渡す値
	args := []interface{}{
		findAlgorithmCode.LanguageId, findAlgorithmCode.AlgorithmId,
		findAlgorithmCode.StartRow, findAlgorithmCode.EndRow,
		findAlgorithmCode.Offset, findAlgorithmCode.Limit,
	}

	// クエリを実行
	rows, err := arcr.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	// 実行結果を戻り値に変換
	algorithmCodes := []model.AlgorithmCode{}
	for rows.Next() {
		algorithmCode := model.AlgorithmCode{}
		if err := rows.Scan(&algorithmCode.AlgorithmCodeId, &algorithmCode.Content, &algorithmCode.Nrow,
			&algorithmCode.CreatedAt, &algorithmCode.LanguageId, &algorithmCode.AlgorithmId); err != nil {
			return nil, err
		}
		algorithmCodes = append(algorithmCodes, algorithmCode)
	}

	return algorithmCodes, nil
}
