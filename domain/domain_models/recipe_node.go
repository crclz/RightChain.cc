package domain_models

import "github.com/crclz/RightChain.cc/domain/utils"

type RecipeNode struct {
	Left          *RecipeNode `json:"left,omitempty"`
	Right         *RecipeNode `json:"right,omitempty"`
	Literal       string      `json:"literal,omitempty"`
	LiteralToHash string      `json:"literalToHash,omitempty"`

	outputCache string `json:"-"`

	Keep bool `json:"-"`
}

func (p *RecipeNode) IsValid() bool {
	if (p.Left == nil) != (p.Right == nil) {
		return false
	}

	var conditions = []bool{
		p.Literal != "",
		p.LiteralToHash != "",
		p.Left != nil,
	}

	var trueConditions = 0
	for _, cond := range conditions {
		if cond {
			trueConditions++
		}
	}

	return trueConditions == 1
}

func (p *RecipeNode) GetOutput() string {
	if p.outputCache == "" {
		p.outputCache = p.CalculateOutput()
	}

	return p.outputCache
}

func (p *RecipeNode) CalculateOutput() string {
	if p.Literal != "" {
		return p.Literal
	}

	if p.LiteralToHash != "" {
		return utils.GetSHA256(p.LiteralToHash)
	}

	return utils.GetSHA256(p.Left.GetOutput() + p.Right.GetOutput())
}

func (p *RecipeNode) ClearCache() {
	p.outputCache = ""

	if p.Left != nil {
		p.Left.ClearCache()
	}

	if p.Right != nil {
		p.Right.ClearCache()
	}
}

// return: target, targetParent
func (p *RecipeNode) FindNode(predicate func(*RecipeNode) bool) (*RecipeNode, *RecipeNode) {
	var target, targetParent *RecipeNode
	var path []*RecipeNode

	var dfs func(node *RecipeNode) bool

	dfs = func(node *RecipeNode) bool {
		if node == nil {
			return false
		}

		path = append(path, node)
		defer func() {
			path = path[:len(path)-1]
		}()

		if predicate(node) {
			target = node
			if len(path) >= 2 {
				targetParent = path[len(path)-2]
			}
			return true
		}

		for _, subNode := range []*RecipeNode{node.Left, node.Right} {
			var found = dfs(subNode)
			if found {
				return found
			}
		}

		return false
	}

	// act

	dfs(p)
	if len(path) != 0 {
		panic("logic error")
	}

	return target, targetParent
}
