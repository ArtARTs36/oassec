package specification

import (
	"fmt"
	"github.com/artarts36/go-swagger-auth/scope"
)

type Spec struct {
	Paths map[string]map[string]Operation `yaml:"paths"`
}

type Operation struct {
	Security []map[string][]string `yaml:"security"`
}

func (sp *Spec) Scopes() (map[string]*scope.Scope, error) {
	scopes := map[string]*scope.Scope{}

	for _, path := range sp.Paths {
		for _, operation := range path {
			for _, security := range operation.Security {
				for _, scopeStrings := range security {
					for _, scopeString := range scopeStrings {
						if _, ok := scopes[scopeString]; !ok {
							sc, err := scope.ParseScope(scopeString)
							if err != nil {
								return nil, fmt.Errorf("parse scope: %w", err)
							}

							scopes[scopeString] = sc
						}
					}
				}
			}
		}
	}

	return scopes, nil
}
