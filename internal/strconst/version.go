package strconst

const (
	ProjectVersion                 = "v0.1.0"
	LatestGoVersionFallback        = "1.25.5"
	RecommendedGofumptVersion      = "v0.9.2"
	RecommendedGoimportsVersion    = "v0.40.0"
	RecommendedGolangciLintVersion = "v2.7.2"
)

const ProjectVersionTemplateFormat = `Version: {{.Name}} {{.Version}} (%s)
Runtime: %s (%s/%s)
Organization: Thought2Code`
