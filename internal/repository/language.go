package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/tf63/code-api/internal/model"
)

// LanguageRepository
type LanguageRepository interface {
	ReadLanguages() (languages []model.Language, err error) // Languageの取得
}

type languageRepository struct {
	Db  *sql.DB
	Rdb *redis.Client
}

// インターフェースを実装しているかチェック
var _ LanguageRepository = (*languageRepository)(nil)

// コンストラクタ
func NewLanguageRepository(db *sql.DB, rdb *redis.Client) LanguageRepository {
	return &languageRepository{db, rdb}
}

// Languageの取得
func (lr *languageRepository) ReadLanguages() ([]model.Language, error) {

	// キャッシュを取得してみる
	ctx := context.Background()
	cachedList, err := lr.Rdb.Get(ctx, model.REDIS_LANGUAGE).Result()

	// キャッシュが存在したらキャッシュを返す
	if err == nil {
		log.Println("use cache of language")

		// キャッシュを返す
		languages := []model.Language{}
		if err := json.Unmarshal([]byte(cachedList), &languages); err != nil {
			return nil, err
		}

		return languages, nil
	}

	// Languageを取得するクエリ
	query := `SELECT language_id, name FROM language`

	// クエリを実行
	rows, err := lr.Db.Query(query)
	if err != nil {
		return nil, err
	}

	// レコードを割り当てる
	languages := []model.Language{}
	for rows.Next() {
		language := model.Language{}
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
	if err := lr.Rdb.Set(ctx, model.REDIS_LANGUAGE, jsonLanguage, 0).Err(); err != nil {
		return nil, err
	}

	return languages, nil
}
