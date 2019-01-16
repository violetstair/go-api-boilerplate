package boilerplate_utiles

import (
	"github.com/eoscanada/eos-go"
	"github.com/violetstair/go-api-boilerplate/pkg/go-api-boilerplate/boilerplate-structure"
)

func GetServerVersion() (ei boilerplate_structure.EOSInfo, err error) {
	api := eos.New("https://eos-usw.owdin.network:9999")

	infoResp, err := api.GetInfo()
	if err != nil {
		return
	}

	ei.Version = infoResp.ServerVersion
	ei.VersionString = infoResp.ServerVersionString

	return
}
