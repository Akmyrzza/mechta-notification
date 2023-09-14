package sqlite

import "github.com/akmyrzza/mechta-notification/internal/models"

func (sql *Sqlite) GetAllChannels() ([]models.Channel, error) {
	var channels []models.Channel

	statement := `
		SELECT * FROM channels  
	`
	rows, err := sql.Database.Query(statement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var channel models.Channel
		rows.Scan(&channel.Id, &channel.TelegramId)
		channels = append(channels, channel)
	}

	return channels, nil
}
