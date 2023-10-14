package repo

import "hostel-service/internal/chat/domain"

var database = []domain.Payload{
	{
		Id:       "d6877823-ea96-425c-9b76-2c1dd9e42e48",
		ToId:     "1a82ced8-263c-46c6-bf0d-60237d933cec",
		SendTime: "5:00PM",
		Msg:      "1st message",
	},
	{
		Id:       "d6877823-ea96-425c-9b76-2c1dd9e42e48",
		ToId:     "1a82ced8-263c-46c6-bf0d-60237d933cec",
		SendTime: "6:00PM",
		Msg:      "2nd message",
	},
}

func FetchMessage() ([]domain.Payload, error) {
	// Data is sorted
	return database, nil
}

func InsertMessage(payload domain.Payload) error {
	database = append(database, payload)
	return nil
}
