package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"swiftDaily_myself/global"
	"swiftDaily_myself/model/response"
	"swiftDaily_myself/utils"
)

type GaodeService struct {
}

func (gaodeService *GaodeService) GetLocationByIP(ip string) (response.IPResponse, error) {
	data := response.IPResponse{}
	key := global.Config.Gaode.Key
	method := "GET"
	url := "https://restapi.amap.com/v3/ip"
	params := map[string]string{
		"key": key,
	}
	res, err := utils.HttpRequest(method, url, nil, params, nil)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return data, errors.New("request failed with status code" + strconv.Itoa(res.StatusCode))
	}
	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
