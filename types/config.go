package types

type DeleteConfig struct {
	ApiVersion string     `json:"apiVersion,omitempty"`
	Kind       string     `json:"kind,omitempty"`
	Selectors  []Selector `json:"selectors"`
}

type Selector struct {
	MatchExpression string            `json:"matchExpression,omitempty"`
	Tags            map[string]string `json:"tags"`
	LastAccess      *LastAccess       `json:"lastAccess"`
	S3ObjMetaData   map[string]any    `json:"objMetadata"`
	WithoutTagKeys  []string          `json:"withoutTagKeys"`
}

type LastAccess struct {
	TimeZone string `json:"timezone,omitempty"`
	FromDate string `json:"from,omitempty"`
	ToDate   string `json:"to,omitempty"`
}
