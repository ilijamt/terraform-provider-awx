package framework

// ProviderBase holds the shared fields for all generated resources and data sources.
type ProviderBase struct {
	Client   Requester
	Endpoint string
	TypeName string
}

func (b *ProviderBase) configureClient(providerData any) {
	if providerData == nil {
		return
	}
	b.Client = providerData.(Requester)
}
