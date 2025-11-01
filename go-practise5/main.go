package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Job struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Company string `json:"company"`
	Salary  int    `json:"salary"`
}

type Response struct {
	Items       []Job `json:"items"`
	NextAfterID *int  `json:"next_after_id,omitempty"`
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/Jobs?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		company := r.URL.Query().Get("company")
		limitStr := r.URL.Query().Get("limit")
		afterIDStr := r.URL.Query().Get("after_id")

		limit := 10
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
				limit = l
			}
		}

		query := "SELECT id, title, company, salary FROM jobs"
		args := []interface{}{}
		where := []string{}

		if company != "" {
			where = append(where, "company = $1")
			args = append(args, company)
		}

		if afterIDStr != "" {
			if afterID, err := strconv.Atoi(afterIDStr); err == nil && afterID > 0 {
				where = append(where, "id < $"+strconv.Itoa(len(args)+1))
				args = append(args, afterID)
			}
		}

		if len(where) > 0 {
			query += " WHERE " + strings.Join(where, " AND ")
		}

		query += " ORDER BY created_at DESC, id DESC LIMIT $" + strconv.Itoa(len(args)+1)
		args = append(args, limit)

		rows, err := db.Query(query, args...)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var jobs []Job
		for rows.Next() {
			var job Job
			if err := rows.Scan(&job.ID, &job.Title, &job.Company, &job.Salary); err != nil {
				http.Error(w, "Error reading data", http.StatusInternalServerError)
				return
			}
			jobs = append(jobs, job)
		}

		response := Response{Items: jobs}
		if len(jobs) > 0 {
			lastID := jobs[len(jobs)-1].ID
			response.NextAfterID = &lastID
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Query-Time", time.Since(start).String())

		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
