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

type EngineStore struct{
	db *sql.DB
}

type QueryResponse struct{
	ID uuid.UUID `json:"id"`
	Displacement int64 `json:"displacement_in_cc"`
	NoOfCylinders int `json:"no_of_cylinders"`
	Material string `json:"material"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewEngineStore(db *sql.DB) *EngineStore{
	return &EngineStore{db: db}
}

func (e EngineStore) EngineById(ctx context.Context, id string)(interface{}, error){
	var queryData QueryResponse

	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		return queryData, err
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
			return queryData, nil // return empty model
		}
		return queryData, err // return empty model
	}	
	return queryData, err
}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.Engine)(interface{}, error){
	var queryData QueryResponse

	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		return queryData, err
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
		return nil, err
	}	

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected > 0 {
		return map[string]string{"message": "Data inserted successfully!"}, nil
	} else {
		return map[string]string{"message": "No rows were inserted!"}, nil
	}
}

func (e EngineStore) EngineUpdate(ctx context.Context, id string, engineReq *models.Engine)(interface{}, error){
	engineID, err := uuid.Parse(id)
	if err != nil{
		return models.Engine{}, fmt.Errorf("invalid Engine ID: %w", err)
	}

	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		return models.Engine{}, err
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
		log.Println(err)
		return models.Engine{}, nil
	}

	// if err != nil{
	// 	return models.Engine{}, err
	// }
	// if rowAffected == 0{
	// 	return models.Engine{}, errors.New("No row updated")
	// }
	rowAffected, err := result.RowsAffected()
	if rowAffected > 0 {
		return map[string]string{"message": "Data updated successfully!"}, nil
	} else {
		return map[string]string{"message": "No rows were updated!"}, nil
	}
}

func (e EngineStore) EngineDelete(ctx context.Context, id string)(interface{}, error){
	
	var queryData QueryResponse

	// DB transaction
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil{
		return queryData, err
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
			return queryData, nil // return empty model
		}
		return queryData, err // return empty model
	}
	var query string = "DELETE FROM engine WHERE id=$1"
	result, err := tx.ExecContext(ctx, query, id)
	if err != nil{
		return models.Engine{}, nil
	}
	rowAffected, err := result.RowsAffected()
	// if err != nil{
	// 	return models.Engine{}, err
	// }
	// if rowAffected == 0{
	// 	return models.Engine{}, errors.New("No row deleted")
	// }
	if rowAffected > 0 {
		return map[string]string{"message": "Data deleted successfully!"}, nil
	} else {
		return map[string]string{"message": "No rows were deleted!"}, nil
	}
	
}