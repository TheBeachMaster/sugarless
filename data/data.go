package data

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-faker/faker/v4"
	"github.com/redis/go-redis/v9"
	"golang.org/x/xerrors"
)

var EXTENSIONS_MARKETPLACE_KEY string = "marketplace"
var NATIVE_MODULE_MARKETPLACE_FIELD string = "modules"

func CreateData(redisClient *redis.Client) error {
	cacheData := make([]ExtensionsNode, 0)

	for range 10 {
		fData := ExtensionsNode{}
		if err := faker.FakeData(&fData); err != nil {
			log.Printf("unable to fake data %s", err.Error())
			break
		}

		cacheData = append(cacheData, fData)
	}

	_data := &ExtensionList{
		Extensions: cacheData,
	}

	log.Printf("data %s", fmt.Sprintf("%+v", _data))

	if err := redisClient.HSet(context.Background(), EXTENSIONS_MARKETPLACE_KEY, NATIVE_MODULE_MARKETPLACE_FIELD, _data).Err(); err != nil {
		log.Printf("unable to cache native modules %s", xerrors.Errorf("%w", err.Error()))
		return fmt.Errorf("unable to cache modules")
	}

	return nil
}

func FetchMarketplaceNativeModulesCache(redisClient *redis.Client) (*[]ExtensionsNode, error) {
	cachedExtensions, err := redisClient.HGetAll(context.Background(), EXTENSIONS_MARKETPLACE_KEY).Result()
	if errors.Is(redis.Nil, err) {
		return &[]ExtensionsNode{}, nil
	}
	if err != nil {
		log.Printf("unable to fetch native modules from cache %s", xerrors.Errorf("%w", err.Error()))
		return nil, fmt.Errorf("unable to fetch cached extensions")
	}

	// b_cachedExtension, err := json.Marshal(cachedExtensions["modules"])
	// if err != nil {
	// 	e.logger.ErrorContext(ctx, "FetchMarketplaceNativeModulesCache", "unable to serialize cache result", slog.Any("error", xerrors.New(err.Error())))
	// 	return nil, fmt.Errorf("unable to serialize cache response")
	// }

	m_cachedExtensions := cachedExtensions["modules"]

	redisData := &ExtensionList{}
	if err := redisData.UnmarshalBinary([]byte(m_cachedExtensions)); err != nil {
		log.Printf("unable to unmarshal %s into %+v due to %s", m_cachedExtensions, redisData, xerrors.New(err.Error()))
		return nil, fmt.Errorf("unable to parse cached extensions")
	}

	extData := make([]ExtensionsNode, 0)

	for _, v := range redisData.Extensions {
		_extData := v
		extData = append(extData, _extData)
	}

	return &extData, nil
}
