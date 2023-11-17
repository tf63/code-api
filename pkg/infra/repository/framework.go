package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/tf63/code-api/pkg/domain/entity"
	"github.com/tf63/code-api/pkg/domain/repository"
)

type frameworkRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

// インターフェースを実装しているかチェック
var _ repository.FrameworkRepository = (*frameworkRepository)(nil)

// コンストラクタ
func NewFrameworkRepository(db *sql.DB, rdb *redis.Client) repository.FrameworkRepository {
	return &frameworkRepository{db, rdb}
}

// Frameworkの取得
func (fr *frameworkRepository) ReadFrameworks() ([]entity.Framework, error) {

	// キャッシュを取得してみる
	ctx := context.Background()
	cachedList, err := fr.rdb.Get(ctx, entity.REDIS_FRAMEWORK).Result()

	// キャッシュが存在したらキャッシュを返す
	if err == nil {
		log.Println("use cache of framework")

		// キャッシュを返す
		frameworks := []entity.Framework{}
		if err := json.Unmarshal([]byte(cachedList), &frameworks); err != nil {
			return nil, err
		}

		return frameworks, nil
	}

	// Frameworkを取得するクエリ
	query := `SELECT framework_id, name FROM framework`

	// クエリを実行
	rows, err := fr.db.Query(query)
	if err != nil {
		return nil, err
	}

	// レコードを割り当てる
	frameworks := []entity.Framework{}
	for rows.Next() {
		framework := entity.Framework{}
		if err := rows.Scan(&framework.FrameworkId, &framework.Name); err != nil {
			return nil, err
		}
		frameworks = append(frameworks, framework)
	}

	// オブジェクトをjson形式にシリアライズ
	jsonFramework, err := json.Marshal(frameworks)
	if err != nil {
		return nil, err
	}

	// シリアライズしたものをキャッシュとして保存する
	if err := fr.rdb.Set(ctx, entity.REDIS_FRAMEWORK, jsonFramework, 0).Err(); err != nil {
		return nil, err
	}

	return frameworks, nil
}
