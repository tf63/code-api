package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/tf63/code-api/internal/model"
)

// PatternRepository
type PatternRepository interface {
	ReadPatterns() (patterns []model.Pattern, err error) // Patternの取得
}

type patternRepository struct {
	Db  *sql.DB
	Rdb *redis.Client
}

// インターフェースを実装しているかチェック
var _ PatternRepository = (*patternRepository)(nil)

// コンストラクタ
func NewPatternRepository(db *sql.DB, rdb *redis.Client) PatternRepository {
	return &patternRepository{db, rdb}
}

// Patternの取得
func (pr *patternRepository) ReadPatterns() ([]model.Pattern, error) {

	// キャッシュを取得してみる
	ctx := context.Background()
	cachedList, err := pr.Rdb.Get(ctx, model.REDIS_PATTERN).Result()

	// キャッシュが存在したらキャッシュを返す
	if err == nil {
		log.Println("use cache of pattern")

		// キャッシュを返す
		patterns := []model.Pattern{}
		if err := json.Unmarshal([]byte(cachedList), &patterns); err != nil {
			return nil, err
		}

		return patterns, nil
	}

	// Patternを取得するクエリ
	query := `SELECT pattern_id, name FROM pattern`

	// クエリを実行
	rows, err := pr.Db.Query(query)
	if err != nil {
		return nil, err
	}

	// レコードを割り当てる
	patterns := []model.Pattern{}
	for rows.Next() {
		pattern := model.Pattern{}
		if err := rows.Scan(&pattern.PatternId, &pattern.Name); err != nil {
			return nil, err
		}
		patterns = append(patterns, pattern)
	}

	// オブジェクトをjson形式にシリアライズ
	jsonPattern, err := json.Marshal(patterns)
	if err != nil {
		return nil, err
	}

	// シリアライズしたものをキャッシュとして保存する
	if err := pr.Rdb.Set(ctx, model.REDIS_PATTERN, jsonPattern, 0).Err(); err != nil {
		return nil, err
	}

	return patterns, nil
}
