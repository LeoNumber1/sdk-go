package tests

import (
	"sdk/api"
	"testing"
)

func TestMarketCommon(t *testing.T) {
	marketClient := api.NewMarketAPI(initClient())
	resp, err := marketClient.ListMarketCommon()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestListMarketMachines(t *testing.T) {
	marketClient := api.NewMarketAPI(initClient())
	resp, err := marketClient.ListMarketMachines("", "", "", 1, 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestListMarketImages(t *testing.T) {
	marketClient := api.NewMarketAPI(initClient())
	resp, err := marketClient.ListMarketImages()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
