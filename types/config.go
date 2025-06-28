package types

type DeleteConfig struct {
	ApiVersion string     `json:"apiVersion,omitEmpty"`
	Kind       string     `json:"kind,omitEmpty"`
	Selectors  []Selector `json:"selectors"`
}

type Selector struct {
	MatchExpression string            `json:"matchExpression,omitEmpty"`
	Tags            map[string]string `json:"tags"`
	LastAccess      *LastAccess       `json:"lastAccess"`
	S3ObjMetaData   map[string]any    `json:"objMetadata"`
	WithoutTagKeys  []string          `json:"withoutTagKeys"`
}

type LastAccess struct {
	TimeZone string `json:"timezone,omitEmpty"`
	FromDate string `json:"from,omitEmpty"`
	ToDate   string `json:"to,omitEmpty"`
}
