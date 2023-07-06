package data

import "github.com/google/uuid"

type DataContract struct {
	TransactionId uuid.UUID   `json:"transactionId"`
	Name          string      `json:"name"`
	DisplayName   string      `json:"displayName"`
	Description   string      `json:"description"`
	Version       string      `json:"version"`
	Data          interface{} `json:"data"`
}

func NewDataContract(data interface{}) *DataContract {
	return &DataContract{
		TransactionId: uuid.New(),
		Name:          "message_events",
		DisplayName:   "Message Events",
		Description:   "This data contract includes whatsapp message informations",
		Version:       "1.0",
		Data:          data,
	}
}
