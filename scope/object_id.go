package scope

type Location string

const (
	LocationInvalid = ""
	LocationPath    = "path"
	LocationQuery   = "query"
	LocationHeader  = "header"
	LocationCookie  = "cookie"
)

type ObjectID struct {
	In  Location
	Key string
}

func (o Location) Valid() bool {
	switch o {
	case LocationPath, LocationQuery, LocationHeader, LocationCookie:
		return true
	}
	return false
}
