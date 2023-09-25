package repository

import (
	"database/sql"

	"github.com/tf63/code-api/internal/entity"
)

type TodoRepository interface {
	CreateTodo(input entity.NewTodo) (todoId *int, err error)
	ReadTodos() (todos []entity.Todo, err error)
	ReadTodo(todoId int) (todo *entity.Todo, err error)
}

type todoRepository struct {
	Db *sql.DB // repositoryはdbへの接続を担う
}

// インターフェースを実装しているかチェック
var _ TodoRepository = (*todoRepository)(nil)

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db}
}

// Todoの作成
func (tr *todoRepository) CreateTodo(input entity.NewTodo) (*int, error) {

	// 入力
	title := input.Title

	// トランザクションを開始
	tx, err := tr.Db.Begin()
	if err != nil {
		return nil, entity.STATUS_SERVICE_UNAVAILABLE
	}

	// エラーが生じた場合はロールバック
	defer tx.Rollback()

	// 実行するSQL (プレースホルダを使う)
	query := `
	INSERT INTO todos (title)
		VALUES ($1)
		RETURNING id
	`

	// 　SQL文の入力
	args := []interface{}{
		title,
	}

	// レコードの作成
	var todoId int
	if tr.Db.QueryRow(query, args...).Scan(&todoId); err != nil {
		return nil, entity.STATUS_SERVICE_UNAVAILABLE
	}

	// トランザクションのコミット
	err = tx.Commit()
	if err != nil {
		return nil, entity.STATUS_SERVICE_UNAVAILABLE
	}

	return &todoId, nil
}

// Todosの取得
func (tr *todoRepository) ReadTodos() ([]entity.Todo, error) {

	// 実行するSQL (プレースホルダを使う)
	query := `SELECT id, title, done FROM todos LIMIT 100`

	// レコードを割り当てる
	rows, err := tr.Db.Query(query)
	if err != nil {
		return nil, entity.STATUS_SERVICE_UNAVAILABLE
	}

	defer rows.Close()

	// returnするオブジェクトを作成
	todos := []entity.Todo{}
	for rows.Next() {
		todo := entity.Todo{}
		if rows.Scan(&todo.Id, &todo.Title, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// Todosの取得
func (tr *todoRepository) ReadTodo(todoId int) (*entity.Todo, error) {

	// 実行するSQL
	query := `SELECT id, title, done FROM todos WHERE id = $1`

	// レコードを割り当てる
	todo := entity.Todo{}
	if err := tr.Db.QueryRow(query, todoId).Scan(&todo.Id, &todo.Title, &todo.Done); err != nil {
		return nil, err
	}

	return &todo, nil
}
