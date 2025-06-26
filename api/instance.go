package api

import (
	"encoding/json"
	"github.com/LeoNumber1/sdk-go/client"
	"github.com/LeoNumber1/sdk-go/models"
	"net/http"
	"net/url"
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
		root:   "/open/api/v1/instances",
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
	uri := u.String()
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

func (i *InstanceAPI) CreateInstance(req models.CreateInstanceRequest) (string, error) {
	uri := i.root
	method := http.MethodPost
	body, _ := json.Marshal(req)
	payload := strings.NewReader(string(body))
	headers := generateHeader(i.client, method, uri, nil, []byte(body))
	var instanceID string
	_, err := httpRequest(i.client.HttpClient, i.client.BaseURL+uri, method, headers, payload, &instanceID)
	if err != nil {
		return "", err
	}
	return instanceID, nil
}

func (i *InstanceAPI) StartInstance(id string) error {
	uri := i.root + "/" + id + "/start"
	method := http.MethodPut
	body := ""
	payload := strings.NewReader(body)
	headers := generateHeader(i.client, method, uri, nil, []byte(body))
	_, err := httpRequest(i.client.HttpClient, i.client.BaseURL+uri, method, headers, payload, nil)
	if err != nil {
		return err
	}
	return nil
}

func (i *InstanceAPI) StopInstance(id string) error {
	uri := i.root + "/" + id + "/stop"
	method := http.MethodPut
	body := ""
	payload := strings.NewReader(body)
	headers := generateHeader(i.client, method, uri, nil, []byte(body))
	_, err := httpRequest(i.client.HttpClient, i.client.BaseURL+uri, method, headers, payload, nil)
	if err != nil {
		return err
	}
	return nil
}

func (i *InstanceAPI) DeleteInstance(id string) error {
	uri := i.root + "/" + id
	method := http.MethodDelete
	body := ""
	payload := strings.NewReader(body)
	headers := generateHeader(i.client, method, uri, nil, []byte(body))
	_, err := httpRequest(i.client.HttpClient, i.client.BaseURL+uri, method, headers, payload, nil)
	if err != nil {
		return err
	}
	return nil
}
