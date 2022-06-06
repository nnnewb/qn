package utils

import (
	"fmt"

	"github.com/qiniu/go-sdk/v7/storage"
)

func CheckRegion(region string) (storage.RegionID, error) {
	switch region {
	case "":
		return "", nil
	case "huadong":
		return storage.RIDHuadong, nil
	case "huadong-zhejiang":
		return storage.RIDHuadongZheJiang, nil
	case "huabei":
		return storage.RIDHuabei, nil
	case "huanan":
		return storage.RIDHuanan, nil
	case "north-america":
		return storage.RIDNorthAmerica, nil
	case "singapore":
		return storage.RIDSingapore, nil
	case "fog-cn-east1":
		return storage.RIDFogCnEast1, nil
	default:
		return "", fmt.Errorf("unexpected region name %s", region)
	}
}
