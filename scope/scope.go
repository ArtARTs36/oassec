package scope

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Scope struct {
	Action     string
	ObjectType string
	ObjectID   ObjectID
}

func (o *Scope) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected scalar node, but got %q", n.Kind)
	}

	sc, err := ParseScope(n.Value)
	if err != nil {
		return fmt.Errorf("parse scope: %w", err)
	}

	*o = *sc

	return nil
}

func ParseScope(ability string) (*Scope, error) {
	parts := strings.SplitN(ability, ":", 3)
	if len(parts) < 2 {
		return nil, fmt.Errorf(
			"invalid value. expected action:object_type:query.field_name or action:object_type:path.field_name",
		)
	}

	scope := &Scope{
		Action:     parts[0],
		ObjectType: parts[1],
	}

	if len(parts) > 2 {
		oid, err := parseObjectID(parts[2])
		if err != nil {
			return nil, fmt.Errorf("parse object id: %w", err)
		}

		scope.ObjectID = *oid
	}

	return scope, nil
}

func parseObjectID(loc string) (*ObjectID, error) {
	parts := strings.SplitN(loc, ".", 2)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid value. expected query.field_name or path.field_name")
	}

	id := &ObjectID{}

	id.In = ObjectIDLocator(parts[0])
	if !id.In.Valid() {
		return nil, fmt.Errorf("invalid value. expected query.field_name or path.field_name, got %q", id.In)
	}

	id.Key = parts[1]

	return id, nil
}
