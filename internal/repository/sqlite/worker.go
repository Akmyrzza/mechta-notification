package sqlite

import (
	"time"

	"github.com/akmyrzza/mechta-notification/internal/models"
)

func (sql *Sqlite) GetAll() ([]string, error) {
	var workers []string

	statement := `
		SELECT * FROM info 
		WHERE day = $1 and month = $2 
	`

	day := time.Now().Day()
	month := int(time.Now().Month())

	rows, err := sql.Database.Query(statement, day, month)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var worker models.Worker
		rows.Scan(&worker.Id, &worker.Name, &worker.Day, &worker.Month)
		workers = append(workers, worker.Name)
	}

	return workers, nil
}
