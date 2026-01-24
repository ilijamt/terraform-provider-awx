package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-viper/mapstructure/v2"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/models"
)

func UserId(ctx context.Context, client Client) (user *models.User, err error) {
	if client == nil {
		return user, fmt.Errorf("client is nil")
	}

	req, _ := client.NewRequest(ctx, http.MethodGet, "/api/v2/me", nil)
	var data map[string]any
	if data, err = client.Do(ctx, req); err == nil {
		if data, _, err = helpers.ExtractDataIfSearchResult(data); err == nil {
			user = new(models.User)
			err = mapstructure.Decode(data, user)
		}
	}

	return user, err
}
