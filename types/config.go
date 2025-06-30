package types

type DeleteConfig struct {
	ApiVersion string     `json:"apiVersion,omitempty"`
	Kind       string     `json:"kind,omitempty"`
	Selectors  []Selector `json:"selectors,omitempty"`
}

type Selector struct {
	MatchExpression string            `json:"matchExpression,omitempty"`
	Tags            map[string]string `json:"tags,omitempty"`
	LastAccess      *LastAccess       `json:"lastAccess,omitempty"`
	S3ObjMetaData   map[string]string `json:"objMetadata,omitempty"`
	WithoutTagKeys  []string          `json:"withoutTagKeys,omitempty"`
}

type LastAccess struct {
	TimeZone string `json:"timezone,omitempty"`
	FromDate string `json:"from,omitempty"`
	ToDate   string `json:"to,omitempty"`
}
