package registry_factories

import (
	procezor "github.com/mzdyhrave/procezorgo"
	constants "github.com/mzdyhrave/pzor-optimulago/internal/registry_constants"
)

type OptimulaConceptFactory struct {
	procezor.IConceptSpecFactory
}

func NewOptimulaConceptFactory() procezor.IConceptSpecFactory {
	return OptimulaConceptFactory{procezor.NewConceptSpecFactoryWithProviders(
		map[int32]procezor.IConceptSpecProvider {
			constants.CONCEPT_TIMESHEETS_PLAN.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_TIMESHEETS_WORK.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_TIMEACTUAL_WORK.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_PAYMENT_BASIS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_PAYMENT_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_PAYMENT_FIXED.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OPTIMUS_BASIS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OPTIMUS_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OPTIMUS_FIXED.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OPTIMUS_NETTO.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_REDUCED_BASIS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_REDUCED_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_REDUCED_FIXED.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_REDUCED_NETTO.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_AGRWORK_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_AGRTASK_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_ALLOWCE_MFULL.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_ALLOWCE_HFULL.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_ALLOWCE_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_ALLOWCE_DAILY.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_ALLDOWN_DAILY.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OFFWORK_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OFFTASK_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OFFSETS_HFULL.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OFFSETS_HOURS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OFFSETS_DAILY.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_OFFDOWN_DAILY.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_TARGETS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_TARNETT.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_AGRWORK.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_AGRTASK.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_ALLOWCE.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_ALLNETT.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_OFFWORK.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_OFFTASK.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_OFFSETS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_RESULTS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_SETTLEM_RESNETT.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_INCOMES_TAXFREE.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_INCOMES_TAXBASE.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_INCOMES_TAXWINS.Id(): NewTimeshtWorkingConProv(),
			constants.CONCEPT_INCOMES_SUMMARY.Id(): NewTimeshtWorkingConProv(),
		}),
	}
}

