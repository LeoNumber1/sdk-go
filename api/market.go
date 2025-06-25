package api

import (
	"net/http"
	"net/url"
	"sdk/client"
	"sdk/models"
	"strconv"
	"strings"
)

type MarketAPI struct {
	client *client.Client
	root   string
}

func NewMarketAPI(client *client.Client) *MarketAPI {
	return &MarketAPI{
		client: client,
		root:   "/open/api/v1/market",
	}
}

func (m *MarketAPI) ListMarketCommon() (*models.MarketResponse, error) {
	uri := m.root + "/common"
	method := http.MethodGet
	body := ""
	payload := strings.NewReader(body)
	headers := generateHeader(m.client, method, uri, nil, []byte(body))
	resp, err := httpRequest(m.client.HttpClient, m.client.BaseURL+uri, method, headers, payload, &models.MarketResponse{})
	if err != nil {
		return nil, err
	}
	return resp.Data.(*models.MarketResponse), nil
}

func (m *MarketAPI) ListMarketMachines(regionID, skuID, gpuType string, page, size int) (*models.ListMachinesResponse, error) {
	uri := m.root + "/machines"
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if regionID != "" {
		q.Add("region_id", regionID)
	}
	if skuID != "" {
		q.Add("sku_id", skuID)
	}
	if gpuType != "" {
		q.Add("gpu_type", gpuType)
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
	headers := generateHeader(m.client, method, uri, nil, []byte(body))
	resp, err := httpRequest(m.client.HttpClient, m.client.BaseURL+uri, method, headers, payload, &models.ListMachinesResponse{})
	if err != nil {
		return nil, err
	}
	return resp.Data.(*models.ListMachinesResponse), nil
}

func (m *MarketAPI) ListMarketImages() (*[]models.ImageItem, error) {
	uri := m.root + "/images"
	method := http.MethodGet
	body := ""
	payload := strings.NewReader(body)
	headers := generateHeader(m.client, method, uri, nil, []byte(body))
	resp, err := httpRequest(m.client.HttpClient, m.client.BaseURL+uri, method, headers, payload, &[]models.ImageItem{})
	if err != nil {
		return nil, err
	}
	return resp.Data.(*[]models.ImageItem), nil
}
