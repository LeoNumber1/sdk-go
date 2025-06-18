package api

import (
	"net/http"
	"net/url"
	"sdk/client"
	"sdk/models"
	"strconv"
	"strings"
)

type InstanceAPI struct {
	client *client.Client
	root   string
}

func NewInstanceAPI(client *client.Client) *InstanceAPI {
	return &InstanceAPI{
		client: client,
		root:   "/api/open/v1/instances",
	}
}

func (i *InstanceAPI) GetInstance(id string) (*models.InstanceInfo, error) {
	uri := i.root + "/" + id
	method := http.MethodGet
	body := ""
	payload := strings.NewReader(body)
	headers := generateHeader(i.client, method, uri, nil, []byte(body))
	resp, err := httpRequest(i.client.HttpClient, i.client.BaseURL+uri, method, headers, payload, &models.InstanceInfo{})
	if err != nil {
		return nil, err
	}
	return resp.Data.(*models.InstanceInfo), nil
}

func (i *InstanceAPI) ListInstances(instanceID, status, chargeType string,
	page, size int) (*models.ListInstancesResponse, error) {
	uri := i.root
	u, err := url.Parse(i.root)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if instanceID != "" {
		q.Add("instance_id", instanceID)
	}
	if status != "" {
		q.Add("status", status)
	}
	if chargeType != "" {
		q.Add("charge_type", chargeType)
	}
	if page > 0 {
		q.Add("page", strconv.Itoa(page))
	}
	if size > 0 {
		q.Add("size", strconv.Itoa(size))
	}
	u.RawQuery = q.Encode()
	uri = u.String()
	method := http.MethodGet
	body := ""
	payload := strings.NewReader(body)
	headers := generateHeader(i.client, method, uri, nil, []byte(body))
	respSchema := &models.ListInstancesResponse{}
	resp, err := httpRequest(i.client.HttpClient, i.client.BaseURL+uri, method, headers, payload, respSchema)
	if err != nil {
		return nil, err
	}
	return resp.Data.(*models.ListInstancesResponse), nil
}
