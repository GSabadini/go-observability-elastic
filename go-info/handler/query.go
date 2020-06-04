package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.elastic.co/apm"
)

func Query(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var vars = mux.Vars(r)
		name, ok := vars["name"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		count, err := updateRequestCount(db, r.Context(), name)
		if err != nil {
			log.Println(err)
		}

		var response = map[string]interface{}{
			"message":     "OK",
			"http_status": http.StatusOK,
			"count":       count,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

//updateRequestCount incrementa uma contagem para o nome no banco de dados, retornando a nova contagem.
func updateRequestCount(db *sql.DB, ctx context.Context, name string) (int, error) {
	span, ctx := apm.StartSpan(ctx, "updateRequestCount", "query")
	defer span.End()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	row := tx.QueryRowContext(ctx, "SELECT count FROM stats WHERE name=$1", name)
	var count int
	switch err := row.Scan(&count); err {
	case nil:
		count++
		if _, err := tx.ExecContext(ctx, "UPDATE stats SET count=$1 WHERE name=$2", count, name); err != nil {
			return -1, err
		}
	case sql.ErrNoRows:
		count = 1
		if _, err := tx.ExecContext(ctx, "INSERT INTO stats (name, count) VALUES ($1, $2)", name, count); err != nil {
			return -1, err
		}
	default:
		return -1, err
	}

	return count, tx.Commit()
}
