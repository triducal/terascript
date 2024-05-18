package compiler

import (
	"fmt"
	"terascript/src/ast"
)

// Returns the json-formatted go struct of the .sb3 file
func Compile(project ast.Project) Project {
	return compileProject(project)
}

func compileProject(project ast.Project) Project {
	var targets []Target
	for _, sprite := range project.Sprites {
		targets = append(targets, compileSprite(sprite))
	}

	var monitors = make([]Monitor, 0)

	var extensions []string
	extensions = append(extensions, "pen")

	return Project{
		Targets:    targets,
		Monitors:   monitors,
		Extensions: extensions,
		Meta: Meta{
			Semver: "3.0.0",
			VM:     "2.3.0",
			Agent:  "terascript",
		},
	}
}

func compileSprite(sprite ast.Sprite) Target {
	variables := make(map[string]variable, 0)
	for _, varName := range sprite.Vars {
		fmt.Println(varName)
		variables[varName] = variable{varName, "0"}
	}

	var lists = make(map[string]list, 0)
	var broadcasts = make(map[string]broadcast, 0)

	return Target{
		IsStage:    sprite.Name == "Stage",
		Name:       sprite.Name,
		Variables:  variables,
		Lists:      lists,
		Broadcasts: broadcasts,
	}
}
