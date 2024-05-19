package compiler

type Project struct {
	Targets    []Target  `json:"targets"`
	Monitors   []Monitor `json:"monitors"`
	Extensions []string  `json:"extensions"`
	Meta       Meta      `json:"meta"`
}

type Target struct {
	IsStage        bool                 `json:"isStage"`
	Name           string               `json:"name"`
	Variables      map[string][]any     `json:"variables"`
	Lists          map[string][]any     `json:"lists"`
	Broadcasts     map[string]Broadcast `json:"broadcasts"`
	Blocks         map[string]Block     `json:"blocks"`
	Comments       map[string]any       `json:"comments"`
	CurrentCostume int                  `json:"currentCostume"`
	Costumes       []Costume            `json:"costumes"`
	Sounds         []Sound              `json:"sounds"`
	Volume         int                  `json:"volume"`
	LayerOrder     int                  `json:"layerOrder"` // 0 if stage, 1 if sprite
}

type Monitor struct{}

type Meta struct {
	Semver string `json:"semver"`
	VM     string `json:"vm"`
	Agent  string `json:"agent"`
}

type Block struct {
}

type Costume struct {
	Name       string `json:"name"`
	DataFormat string `json:"dataFormat"`
	AssetId    string `json:"assetId"`
	Md5ext     string `json:"md5ext"`
}

type Sound struct {
}

type Broadcast struct{}
