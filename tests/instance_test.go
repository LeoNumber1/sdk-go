package tests

import (
	"sdk/api"
	"sdk/client"
	"testing"
)

func initClient() *client.Client {
	return client.NewClient("http://localhost:3000", "JWUNAEYW-BY-OWWTKKJY",
		"xxx")
}

func TestGetInstance(t *testing.T) {
	insClient := api.NewInstanceAPI(initClient())
	instance, err := insClient.GetInstance("kpIns2abwtx3wwerbeyy1b16")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(instance)
}

func TestListInstances(t *testing.T) {
	insClient := api.NewInstanceAPI(initClient())
	resp, err := insClient.ListInstances("", "", "", 1, 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
