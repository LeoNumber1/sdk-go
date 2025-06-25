package tests

import (
	"sdk/api"
	"sdk/client"
	"sdk/models"
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

func TestCreateInstance(t *testing.T) {
	insClient := api.NewInstanceAPI(initClient())
	req := models.CreateInstanceRequest{
		MachineID:      "KMACPH01JJERBH8M9V4YWV7V5YVBQ3KP",
		ReqGPUAmount:   1,
		ImageType:      "official",
		Image:          "cuda12.1.1-1311",
		PrivateImage:   "",
		InstanceName:   "test",
		ExpandDataDisk: 0,
		SkuID:          "day",
		Duration:       1,
		AutoRenew:      "manual",
	}
	instance, err := insClient.CreateInstance(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(instance)
}

func TestStartInstance(t *testing.T) {
	insClient := api.NewInstanceAPI(initClient())
	instance, err := insClient.StartInstance("kpIns2abwtx3wwerbeyy1b16")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(instance)
}

func TestStopInstance(t *testing.T) {
	insClient := api.NewInstanceAPI(initClient())
	err := insClient.StopInstance("kpIns2abwtx3wwerbeyy1b16")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("stop instance success")
}

func TestDeleteInstance(t *testing.T) {
	insClient := api.NewInstanceAPI(initClient())
	err := insClient.DeleteInstance("kpIns2abwtx3wwerbeyy1b16")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("delete instance success")
}
