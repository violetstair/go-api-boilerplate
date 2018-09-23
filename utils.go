package main

import (
	"github.com/eoscanada/eos-go"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetServerVersion() (EOSInfo, error) {
	var ei EOSInfo
	api := eos.New("https://eos-usw.owdin.network:9999")

	infoResp, err := api.GetInfo()
	if err != nil {
		return ei, err
	}
	ei.Version = infoResp.ServerVersion
	ei.VersionString = infoResp.ServerVersionString
	return ei, nil

}
