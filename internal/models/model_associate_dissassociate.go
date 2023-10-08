package models

type AssociateDisassociateRequestModel struct {
	ID           int64 `json:"id"`
	Disassociate bool  `json:"disassociate,omitempty"`
}
