package models

import "time"

type LogEntry struct {
	NumberOrder     string    `json:"number_order,omitempty"`
	IdSession       string    `json:"id_session,omitempty"`
	Status          string    `json:"status"`
	Amount          int       `json:"amount"`
	AccountingDate  string    `json:"accounting_date"`
	TransactionDate time.Time `json:"transaction_date"`
	PaymentTypeCode string    `json:"payment_type_code"`

	CardNumber string `json:"card_number"`

	AuthorizationCode string  `json:"authorization_code"`
	UsuarioID         string  `json:"UsuarioID"`
	Usuario           Usuario `gorm:"foreignKey:UsuarioID"`
}
