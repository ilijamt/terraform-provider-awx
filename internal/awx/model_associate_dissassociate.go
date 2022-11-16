package awx

type associateDisassociateRequestModel struct {
	ID           int64 `json:"id"`
	Disassociate bool  `json:"disassociate,omitempty"`
}
