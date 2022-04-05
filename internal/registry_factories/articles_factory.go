package registry_factories

import (
	procezor "github.com/mzdyhrave/procezorgo"
)

type OptimulaArticleFactory struct {
	procezor.IArticleSpecFactory
}

func NewOptimulaArticleFactory() procezor.IArticleSpecFactory{
	providersConfig := []procezor.ProviderRecord {
	}
	return OptimulaArticleFactory{procezor.NewArticleSpecFactoryWithRecords(providersConfig)}
}

