package secrets

import (
	"context"
	"io"
	"net/http"

	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/containers/podman/v2/pkg/domain/entities"
)

// List returns information about existing secrets in the form of a slice.
func List(ctx context.Context, options *ListOptions) ([]*entities.SecretInfoReport, error) {
	var (
		secrs []*entities.SecretInfoReport
	)
	conn, err := bindings.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	response, err := conn.DoRequest(nil, http.MethodGet, "/secrets/json", nil, nil)
	if err != nil {
		return secrs, err
	}
	return secrs, response.Process(&secrs)
}

// Inspect returns low-level information about a secret.
func Inspect(ctx context.Context, nameOrID string, options *InspectOptions) (*entities.SecretInfoReport, error) {
	var (
		inspect *entities.SecretInfoReport
	)
	conn, err := bindings.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	response, err := conn.DoRequest(nil, http.MethodGet, "/secrets/%s/json", nil, nil, nameOrID)
	if err != nil {
		return inspect, err
	}
	return inspect, response.Process(&inspect)
}

// Remove removes a secret from storage
func Remove(ctx context.Context, nameOrID string) error {
	conn, err := bindings.GetClient(ctx)
	if err != nil {
		return err
	}

	response, err := conn.DoRequest(nil, http.MethodDelete, "/secrets/%s", nil, nil, nameOrID)
	if err != nil {
		return err
	}
	return response.Process(nil)
}

// Create creates a secret given some data
func Create(ctx context.Context, reader io.Reader, options *CreateOptions) (*entities.SecretCreateReport, error) {
	var (
		create *entities.SecretCreateReport
	)
	conn, err := bindings.GetClient(ctx)
	if err != nil {
		return nil, err
	}

	params, err := options.ToParams()
	if err != nil {
		return nil, err
	}

	response, err := conn.DoRequest(reader, http.MethodPost, "/secrets/create", params, nil)
	if err != nil {
		return nil, err
	}
	return create, response.Process(&create)
}