package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/tf63/code-api/pkg/config"
	"github.com/tf63/code-api/pkg/domain/entity"
	"github.com/tf63/code-api/pkg/domain/repository"
)

type languageRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

// インターフェースを実装しているかチェック
var _ repository.LanguageRepository = (*languageRepository)(nil)

// コンストラクタ
func NewLanguageRepository(db *sql.DB, rdb *redis.Client) repository.LanguageRepository {
	return &languageRepository{db, rdb}
}

// Languageの取得
func (lr *languageRepository) ReadLanguages() ([]entity.Language, error) {

	// キャッシュを取得してみる
	ctx := context.Background()
	cachedList, err := lr.rdb.Get(ctx, config.REDIS_LANGUAGE).Result()

	// キャッシュが存在したらキャッシュを返す
	if err == nil {
		log.Println("use cache of language")

		// キャッシュを返す
		languages := []entity.Language{}
		if err := json.Unmarshal([]byte(cachedList), &languages); err != nil {
			return nil, err
		}

		return languages, nil
	}

	// Languageを取得するクエリ
	query := `SELECT language_id, name FROM language`

	// クエリを実行
	rows, err := lr.db.Query(query)
	if err != nil {
		return nil, err
	}

	// レコードを割り当てる
	languages := []entity.Language{}
	for rows.Next() {
		language := entity.Language{}
		if err := rows.Scan(&language.LanguageId, &language.Name); err != nil {
			return nil, err
		}
		languages = append(languages, language)
	}

	// オブジェクトをjson形式にシリアライズ
	jsonLanguage, err := json.Marshal(languages)
	if err != nil {
		return nil, err
	}

	// シリアライズしたものをキャッシュとして保存する
	if err := lr.rdb.Set(ctx, config.REDIS_LANGUAGE, jsonLanguage, 0).Err(); err != nil {
		return nil, err
	}

	return languages, nil
}
