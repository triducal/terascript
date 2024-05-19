package compiler

import (
	"encoding/json"
	"fmt"
	"terascript/src/ast"
	"terascript/src/util"

	"github.com/sanity-io/litter"
)

// Outputs an .sb3 archive after compiling the ast
func Compile(project ast.Project) {
	compileProject(project)
}

func compileProject(project ast.Project) {
	var zipArchive, _ = util.CreateArchive("terascript.sb3")
	var targets []Target
	for _, sprite := range project.Sprites {
		target := compileSprite(sprite)
		for i, costume := range sprite.Costumes {
			fileInfo := util.AddFile(zipArchive, costume)
			target.Costumes = append(target.Costumes, Costume{
				Name:       fmt.Sprint(i),
				DataFormat: fileInfo.DataFormat,
				AssetId:    fileInfo.Checksum,
				Md5ext:     fileInfo.Filename,
			})
		}
		targets = append(targets, *target)
	}

	var monitors = make([]Monitor, 0)

	var extensions []string
	extensions = append(extensions, "pen")

	compiledProject := &Project{
		Targets:    targets,
		Monitors:   monitors,
		Extensions: extensions,
		Meta: Meta{
			Semver: "3.0.0",
			VM:     "2.3.0",
			Agent:  "terascript",
		},
	}
	compiledJson, _ := json.Marshal(compiledProject)
	litter.Dump(string(compiledJson))
	jsonWriter, _ := zipArchive.Create("project.json")
	jsonWriter.Write(compiledJson)
	zipArchive.Close()
}

func compileSprite(sprite ast.Sprite) *Target {
	variables := make(map[string][]any)
	for _, varName := range sprite.Vars {
		variables[varName] = []any{varName, 0}
	}

	var lists = make(map[string][]any)
	var broadcasts = make(map[string]Broadcast)

	var layerOrder int
	if sprite.Name == "Stage" {
		layerOrder = 0
	} else {
		layerOrder = 1
	}

	return &Target{
		IsStage:        sprite.Name == "Stage",
		Name:           sprite.Name,
		Variables:      variables,
		Lists:          lists,
		Broadcasts:     broadcasts,
		Blocks:         make(map[string]Block),
		Comments:       make(map[string]any),
		CurrentCostume: 0,
		// takes care of costume in compileProject
		Sounds:     make([]Sound, 0),
		Volume:     0,
		LayerOrder: layerOrder,
	}
}
