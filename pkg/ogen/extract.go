package ogen

import (
	"fmt"

	"github.com/ogen-go/ogen/middleware"

	"github.com/artarts36/oassec/scope"
)

func ExtractObjectID(sc scope.Scope, params middleware.Parameters) (any, error) {
	extractor, err := objectIDExtractor(sc.ObjectID.In, params)
	if err != nil {
		return nil, err
	}

	value, ok := extractor(sc.ObjectID.Key)
	if !ok {
		return nil, fmt.Errorf("key %q not found", sc.ObjectID.Key)
	}

	return value, nil
}

func objectIDExtractor(loc scope.Location, params middleware.Parameters) (func(key string) (any, bool), error) {
	switch loc {
	case scope.LocationPath:
		return params.Path, nil
	case scope.LocationQuery:
		return params.Query, nil
	case scope.LocationHeader:
		return params.Header, nil
	case scope.LocationCookie:
		return params.Cookie, nil
	default:
		return nil, fmt.Errorf("unexpected location %v", loc)
	}
}
