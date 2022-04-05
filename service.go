package pzor_optimulago

import (
	"fmt"
	factories "github.com/mzdyhrave/pzor-optimulago/internal/registry_factories"
	constants  "github.com/mzdyhrave/pzor-optimulago/internal/registry_constants"
	procezor "github.com/mzdyhrave/procezorgo"
)

const (
	OPTIMULA_VERSION_SCM int32 = 100
	OPTIMULA_VERSION_EPS int32 = 200
	OPTIMULA_VERSION_PUZZLE int32 = 300
	ARTICLE_DEFAULT_SEQUENS int16 = 0
)

type serviceBuilder struct {
}

func (t serviceBuilder) BuildArticleFactory(s *procezor.ProcezorService) bool {
	s.ArticleFactory = factories.NewOptimulaArticleFactory()
	if s.ArticleFactory == nil {
		return false
	}
	s.ArticleFactory.BuildFactory()
	return true
}

func (t serviceBuilder) BuildConceptFactory(s *procezor.ProcezorService) bool {
	s.ConceptFactory = factories.NewOptimulaConceptFactory()
	if s.ConceptFactory == nil {
		return false
	}
	s.ConceptFactory.BuildFactory()
	return true
}

type contractBuilder struct {
}

func (b contractBuilder) GetContractTerms(period procezor.IPeriod, targets procezor.ITermTargetList) procezor.IContractTermList {
	return procezor.IContractTermList{}
}

type positionBuilder struct {
}

func (b positionBuilder) GetPositionTerms(period procezor.IPeriod, contracts procezor.IContractTermList, targets procezor.ITermTargetList) procezor.IPositionTermList {
	return procezor.IPositionTermList{}
}

type OptimulaService struct {
	procezor.IProcezorService
}

func NewExampleServiceBuilder() procezor.IProcezorFactoryBuilder {
	return &serviceBuilder{}
}

func NewExampleContractBuilder() procezor.IProcezorContractBuilder {
	return &contractBuilder{}
}

func NewExamplePositionBuilder() procezor.IProcezorPositionBuilder {
	return &positionBuilder{}
}

func NewOptimulaService() procezor.IProcezorService{
	const (
		TestVersion      = OPTIMULA_VERSION_SCM
	)
	var (
		TestCalcsArticle = []procezor.ArticleCode{
			procezor.GetArticleCode(constants.ARTICLE_SETTLEM_TARGETS.Id()),
			procezor.GetArticleCode(constants.ARTICLE_SETTLEM_RESULTS.Id()),
			procezor.GetArticleCode(constants.ARTICLE_SETTLEM_ALLOWCE.Id()),
			procezor.GetArticleCode(constants.ARTICLE_SETTLEM_AGRWORK.Id()),
			procezor.GetArticleCode(constants.ARTICLE_INCOMES_TAXFREE.Id()),
			procezor.GetArticleCode(constants.ARTICLE_INCOMES_TAXBASE.Id()),
			procezor.GetArticleCode(constants.ARTICLE_INCOMES_TAXWINS.Id()),
			procezor.GetArticleCode(constants.ARTICLE_INCOMES_SUMMARY.Id()),
		}
	)

	service := OptimulaService{
		procezor.NewProcezorService(TestVersion, TestCalcsArticle,
			NewExampleContractBuilder(), NewExamplePositionBuilder(), NewExampleServiceBuilder()),
	}
	buildSuccess := service.BuildFactories()
	if buildSuccess == false {
		println(fmt.Sprintf("Version: %d, build factories failed", service.Version().Value()))
	}
	return &service
}

