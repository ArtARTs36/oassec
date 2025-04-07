package scope

type ObjectIDLocator string

const (
	ObjectIDLocatorInvalid = ""
	ObjectIDLocatorPath    = "path"
	ObjectIDLocatorQuery   = "query"
)

type ObjectID struct {
	In  ObjectIDLocator
	Key string
}

func (o ObjectIDLocator) Valid() bool {
	switch o {
	case ObjectIDLocatorPath, ObjectIDLocatorQuery:
		return true
	}
	return false
}
