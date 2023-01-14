package domain_models

import "github.com/crclz/RightChain.cc/domain/utils"

type RecipeNode struct {
	Left          *RecipeNode `json:"left,omitempty"`
	Right         *RecipeNode `json:"right,omitempty"`
	Literal       string      `json:"literal,omitempty"`
	LiteralToHash string      `json:"literalToHash,omitempty"`

	outputCache string `json:"-"`
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
