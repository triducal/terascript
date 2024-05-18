package compiler

type Project struct {
	Targets    []Target  `json:"targets"`
	Monitors   []Monitor `json:"monitors"`
	Extensions []string  `json:"extensions"`
	Meta       Meta      `json:"meta"`
}

type Target struct {
	IsStage    bool                 `json:"isStage"`
	Name       string               `json:"name"`
	Variables  map[string]variable  `json:"variables"`
	Lists      map[string]list      `json:"lists"`
	Broadcasts map[string]broadcast `json:"broadcasts"`
}

type Monitor struct{}

type Meta struct {
	Semver string `json:"semver"`
	VM     string `json:"vm"`
	Agent  string `json:"agent"`
}

type variable struct {
	name  string
	value string
}

type list struct {
	name   string
	values []string
}

type broadcast struct{}
