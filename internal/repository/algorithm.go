package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/tf63/code-api/internal/model"
)

// AlgorithmRepository
type AlgorithmRepository interface {
	ReadAlgorithms() (algorithms []model.Algorithm, err error) // Algorithmの取得
}

type algorithmRepository struct {
	Db  *sql.DB
	Rdb *redis.Client
}

// インターフェースを実装しているかチェック
var _ AlgorithmRepository = (*algorithmRepository)(nil)

// コンストラクタ
func NewAlgorithmRepository(db *sql.DB, rdb *redis.Client) AlgorithmRepository {
	return &algorithmRepository{db, rdb}
}

// Algorithmの取得
func (ar *algorithmRepository) ReadAlgorithms() ([]model.Algorithm, error) {

	// キャッシュを取得してみる
	ctx := context.Background()
	cachedList, err := ar.Rdb.Get(ctx, model.REDIS_ALGORITHM).Result()

	// キャッシュが存在したらキャッシュを返す
	if err == nil {
		log.Println("use cache of algorithm")

		// キャッシュを返す
		algorithms := []model.Algorithm{}
		if err := json.Unmarshal([]byte(cachedList), &algorithms); err != nil {
			return nil, err
		}

		return algorithms, nil
	}

	// Algorithmを取得するクエリ
	query := `SELECT algorithm_id, name FROM algorithm`

	// クエリを実行
	rows, err := ar.Db.Query(query)
	if err != nil {
		return nil, err
	}

	// レコードを割り当てる
	algorithms := []model.Algorithm{}
	for rows.Next() {
		algorithm := model.Algorithm{}
		if err := rows.Scan(&algorithm.AlgorithmId, &algorithm.Name); err != nil {
			return nil, err
		}
		algorithms = append(algorithms, algorithm)
	}

	// オブジェクトをjson形式にシリアライズ
	jsonAlgorithm, err := json.Marshal(algorithms)
	if err != nil {
		return nil, err
	}

	// シリアライズしたものをキャッシュとして保存する
	if err := ar.Rdb.Set(ctx, model.REDIS_ALGORITHM, jsonAlgorithm, 0).Err(); err != nil {
		return nil, err
	}

	return algorithms, nil
}
