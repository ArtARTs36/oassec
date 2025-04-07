package specification

import (
	"fmt"

	"github.com/artarts36/oassec/scope"
)

type Spec struct {
	Paths map[string]map[string]Operation `yaml:"paths"`
}

type Operation struct {
	Parameters []Parameter
	Security   []map[string][]string `yaml:"security"`
}

type Parameter struct {
	Name string `yaml:"name"`
	In   string `yaml:"in"`
}

func (op *Operation) HasParameter(name string, in scope.ObjectIDLocator) bool {
	for _, parameter := range op.Parameters {
		if parameter.In == string(in) && parameter.Name == name {
			return true
		}
	}

	return false
}

func (sp *Spec) Scopes() (map[string]*scope.Scope, error) {
	scopes := map[string]*scope.Scope{}

	for pathName, path := range sp.Paths {
		for method, operation := range path {
			for _, security := range operation.Security {
				for _, scopeStrings := range security {
					for _, scopeString := range scopeStrings {
						if _, ok := scopes[scopeString]; !ok {
							sc, err := scope.ParseScope(scopeString)
							if err != nil {
								return nil, fmt.Errorf("parse scope: %w", err)
							}

							if !operation.HasParameter(sc.ObjectID.Key, sc.ObjectID.In) {
								return nil, fmt.Errorf(
									"operation %s %s must be have parameter with name %q in %s",
									method,
									pathName,
									sc.ObjectID.Key,
									sc.ObjectID.In,
								)
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
