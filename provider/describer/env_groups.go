package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-render/pkg/sdk/models"
	"github.com/opengovern/og-describer-render/provider/model"
	"net/http"
	"net/url"
	"sync"
)

func ListEnvGroups(ctx context.Context, handler *RenderAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	renderChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(renderChan)
		defer close(errorChan)
		if err := processEnvGroups(ctx, handler, renderChan, &wg); err != nil {
			errorChan <- err // Send error to the error channel
		}
		wg.Wait()
	}()

	var values []models.Resource
	for {
		select {
		case value, ok := <-renderChan:
			if !ok {
				return values, nil
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err := <-errorChan:
			return nil, err
		}
	}
}

func GetEnvGroup(ctx context.Context, handler *RenderAPIHandler, resourceID string) (*models.Resource, error) {
	envGroup, err := processEnvGroup(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   envGroup.ID,
		Name: envGroup.Name,
		Description: JSONAllFieldsMarshaller{
			Value: envGroup,
		},
	}
	return &value, nil
}

func processEnvGroups(ctx context.Context, handler *RenderAPIHandler, renderChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var envGroups []model.EnvGroupDescription
	var envGroupResp []model.EnvGroupResponse
	var resp *http.Response
	baseURL := "https://api.render.com/v1/env-groups"
	cursor := ""

	for {
		params := url.Values{}
		params.Set("limit", limit)
		if cursor != "" {
			params.Set("cursor", cursor)
		}
		finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		requestFunc := func(req *http.Request) (*http.Response, error) {
			var e error
			resp, e = handler.Client.Do(req)
			if e != nil {
				return nil, fmt.Errorf("request execution failed: %w", e)
			}
			defer resp.Body.Close()

			if e = json.NewDecoder(resp.Body).Decode(&envGroupResp); e != nil {
				return nil, fmt.Errorf("failed to decode response: %w", e)
			}
			for i, envGroup := range envGroupResp {
				envGroups = append(envGroups, envGroup.EnvGroup)
				if i == len(envGroupResp)-1 {
					cursor = envGroup.Cursor
				}
			}
			return resp, nil
		}
		err = handler.DoRequest(ctx, req, requestFunc)
		if err != nil {
			return fmt.Errorf("error during request handling: %w", err)
		}

		if len(envGroupResp) < 100 {
			break
		}
	}
	for _, envGroup := range envGroups {
		wg.Add(1)
		go func(envGroup model.EnvGroupDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   envGroup.ID,
				Name: envGroup.Name,
				Description: JSONAllFieldsMarshaller{
					Value: envGroup,
				},
			}
			renderChan <- value
		}(envGroup)
	}
	return nil
}

func processEnvGroup(ctx context.Context, handler *RenderAPIHandler, resourceID string) (*model.EnvGroupDescription, error) {
	var envGroup model.EnvGroupDescription
	var resp *http.Response
	baseURL := "https://api.render.com/v1/env-groups/"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}

	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e = handler.Client.Do(req)
		if e != nil {
			return nil, fmt.Errorf("request execution failed: %w", e)
		}
		defer resp.Body.Close()

		if e = json.NewDecoder(resp.Body).Decode(&envGroup); e != nil {
			return nil, fmt.Errorf("failed to decode response: %w", e)
		}
		return resp, e
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, fmt.Errorf("error during request handling: %w", err)
	}
	return &envGroup, nil
}
