package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/harshitrajsinha/van-man-go/models"
)

type QueryResponse struct{
	ID uuid.UUID `json:"-"`
	Displacement int64 `json:"displacement_in_cc"`
	NoOfCylinders int `json:"no_of_cylinders"`
	Material string `json:"material"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

type EngineStore struct{
	db *sql.DB
}

func NewEngineStore(db *sql.DB) *EngineStore{
	return &EngineStore{db: db}
}

func (e EngineStore) GetEngineById(ctx context.Context, id string)(interface{}, error){
	var queryData QueryResponse

	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		return QueryResponse{}, err
	}

	defer func(){
		if err != nil{
			if rbErr := tx.Rollback(); rbErr != nil{
				log.Println("Transaction rollback error: ", rbErr)
			}
		}else{
			if cmErr := tx.Commit(); cmErr != nil{
				log.Println("Commit rollback error: ", cmErr)
			}
		}
	}()
	err =tx.QueryRowContext(ctx, "SELECT * FROM engine WHERE id=$1", id).Scan(
		&queryData.ID, &queryData.Displacement, &queryData.NoOfCylinders, &queryData.Material, &queryData.CreatedAt, &queryData.UpdatedAt)
	if err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return QueryResponse{}, nil // return empty model
		}
		return QueryResponse{}, err // return empty model
	}	
	return queryData, err
}

func (e EngineStore) GetAllEngine(ctx context.Context)(interface{}, error){

	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		return QueryResponse{}, err
	}

	defer func(){
		if err != nil{
			if rbErr := tx.Rollback(); rbErr != nil{
				log.Println("Transaction rollback error: ", rbErr)
			}
		}else{
			if cmErr := tx.Commit(); cmErr != nil{
				log.Println("Commit rollback error: ", cmErr)
			}
		}
	}()
	rows, err := tx.QueryContext(ctx, "SELECT * FROM engine;")
	if err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return QueryResponse{}, nil // return empty model
		}
		return QueryResponse{}, err // return empty model
	}
	defer rows.Close()

	// slice to store all rows
	engineData := make([]interface{}, 0)

	// Get each row data into a slice
	for rows.Next() {
		var queryData QueryResponse
		if err := rows.Scan(
			&queryData.ID, &queryData.Displacement, &queryData.NoOfCylinders, &queryData.Material, &queryData.CreatedAt, &queryData.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		engineData = append(engineData, queryData)
	}

	return engineData, err
}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.Engine)(map[string]string, error){

	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		log.Println("Error while inserting data ", err)
		return nil, err
	}

	defer func(){
		if err != nil{
			if rbErr := tx.Rollback(); rbErr != nil{
				log.Println("Transaction rollback error: ", rbErr)
			}
		}else{
			if cmErr := tx.Commit(); cmErr != nil{
				log.Println("Commit rollback error: ", cmErr)
			}
		}
	}()

	var query string = "INSERT INTO engine (displacement_in_cc, no_of_cylinders, material) VALUES ($1, $2, $3)"
	result, err := tx.ExecContext(ctx, query, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.Material)

	if err != nil{
		log.Println("Error while inserting data ", err)
		return nil, err
	}	

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error while inserting data ", err)
		return nil, err
	}
	if rowsAffected > 0 {
		return map[string]string{"message": "Data inserted successfully!"}, nil
	} else {
		return map[string]string{"message": "No rows were inserted!"}, nil
	}
}

func (e EngineStore) UpdateEngine(ctx context.Context, engineID string, engineReq *models.Engine)(int64, error){

	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		log.Println("Error while updating data ", err)
		return -1, err
	}

	defer func(){
		if err != nil{
			if rbErr := tx.Rollback(); rbErr != nil{
				fmt.Printf("Transaction rollback error: %v\n", rbErr)
			}
		}else{
			if cmErr := tx.Commit(); cmErr != nil{
				fmt.Printf("Transaction commit error: %v\n", cmErr)
			}
		}
	}()

	var query strings.Builder
	var args []interface{}
	argCount := 1

	query.WriteString("UPDATE engine SET ")

	if engineReq.Displacement > 0 {
		query.WriteString(fmt.Sprintf("displacement_in_cc=$%d ", argCount))
		args = append(args, engineReq.Displacement)
		argCount++
	}
	if engineReq.NoOfCylinders > 0 {
		if argCount > 1{
			query.WriteString(", ")
		}
		query.WriteString(fmt.Sprintf("no_of_cylinders=$%d ", argCount))
		args = append(args, engineReq.NoOfCylinders)
		argCount++
	}
	if engineReq.Material != "" {
		if argCount > 1{
			query.WriteString(", ")
		}
		query.WriteString(fmt.Sprintf("material=$%d ", argCount))
		args = append(args, engineReq.Material)
		argCount++
	}
	query.WriteString(fmt.Sprintf("WHERE id=$%d ", argCount))
	args = append(args, engineID)

	result, err := tx.ExecContext(ctx, query.String(), args...)
	if err != nil{
		log.Println("Error while updating data ", err)
		return -1, err
	}

	rowAffected, err := result.RowsAffected()
	return rowAffected, nil
}

func (e EngineStore) DeleteEngine(ctx context.Context, id string)(int64, error){
	
	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		return -1, err
	}

	defer func(){
		if err != nil{
			if rbErr := tx.Rollback(); rbErr != nil{
				log.Println("Transaction rollback error: ", rbErr)
			}
		}else{
			if cmErr := tx.Commit(); cmErr != nil{
				log.Println("Commit rollback error: ", cmErr)
			}
		}
	}()

	var query string = "DELETE FROM engine WHERE id=$1"
	result, err := tx.ExecContext(ctx, query, id)
	if err != nil{
		return -1, err
	}
	rowAffected, err := result.RowsAffected()

	return rowAffected, nil
	
}