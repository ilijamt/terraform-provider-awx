package resource_test

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

type dummyResource struct {
	diags    diag.Diagnostics
	err      error
	data     map[string]any
	getIdId  string
	getIdErr error
}

func (r *dummyResource) GetId() (string, error) { return r.getIdId, r.getIdErr }

func (r *dummyResource) RequestBody() ([]byte, error) { return json.Marshal(r.data) }

func (r *dummyResource) UpdateWithApiData(callee resource.Callee, source resource.Source, data map[string]any) (_ diag.Diagnostics, _ error) {
	r.data = data
	return r.diags, r.err
}

var _ resource.RequestBody = (*dummyResource)(nil)
var _ resource.Updater = (*dummyResource)(nil)
var _ resource.Id = (*dummyResource)(nil)
