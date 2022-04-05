package registry

type OptimulaArticleConst int32

const ARTICLE_TIMESHEETS_PLAN = OptimulaArticleConst(1001)     // Full Timesheets Hours
const ARTICLE_TIMESHEETS_WORK = OptimulaArticleConst(1002)     // Work Timesheets Hours
const ARTICLE_TIMEACTUAL_WORK = OptimulaArticleConst(1003)     // Work Timeactual Hours
const ARTICLE_MSALARY_TARGETS = OptimulaArticleConst(2001)     // Base Salary
const ARTICLE_HSALARY_TARGETS = OptimulaArticleConst(2002)     // Base Salary
const ARTICLE_MAWARDS_TARGETS = OptimulaArticleConst(2003)     // Personal  Salary - Targets
const ARTICLE_HAWARDS_TARGETS = OptimulaArticleConst(2004)     // Personal  Salary - Targets
const ARTICLE_PREMIUM_TARGETS = OptimulaArticleConst(2005)     // Premium Bonus    - Targets
const ARTICLE_PREMADV_TARGETS = OptimulaArticleConst(2006)     // Premium Boss     - Targets
const ARTICLE_PREMEXT_TARGETS = OptimulaArticleConst(2007)     // Premium Personal - Targets
const ARTICLE_AGRWORK_TARGETS = OptimulaArticleConst(2010)     // Agreement Work Tariff - Targets
const ARTICLE_AGRTASK_TARGETS = OptimulaArticleConst(2011)     // Agreement Task Tariff - Targets
const ARTICLE_OFFWORK_TARGETS = OptimulaArticleConst(2012)     // Agreement Work Tariff - Targets Plus
const ARTICLE_OFFTASK_TARGETS = OptimulaArticleConst(2013)     // Agreement Task Tariff - Targets Plus
const ARTICLE_SETTLEM_TARGETS = OptimulaArticleConst(3001)     // Setlement - Targets
const ARTICLE_SETTLEM_TARNETT = OptimulaArticleConst(3002)     // Setlement - Targets
const ARTICLE_SETTLEM_AGRWORK = OptimulaArticleConst(3003)     // Setlement - Agreement Work
const ARTICLE_SETTLEM_AGRTASK = OptimulaArticleConst(3004)     // Setlement - Agreement Task
const ARTICLE_SETTLEM_ALLOWCE = OptimulaArticleConst(3005)     // Setlement - Allowance
const ARTICLE_SETTLEM_ALLNETT = OptimulaArticleConst(3006)     // Setlement - Allowance Netto
const ARTICLE_SETTLEM_OFFWORK = OptimulaArticleConst(3007)     // Setlement - Agreement Work Plus
const ARTICLE_SETTLEM_OFFTASK = OptimulaArticleConst(3008)     // Setlement - Agreement Task Plus
const ARTICLE_SETTLEM_OFFSETS = OptimulaArticleConst(3009)     // Setlement - Allowance Plus
const ARTICLE_PREMEXT_RESULTS = OptimulaArticleConst(4001)     // Premium Personal - Results
const ARTICLE_PREMADV_RESULTS = OptimulaArticleConst(4002)     // Premium Boss     - Results
const ARTICLE_PREMIUM_RESULTS = OptimulaArticleConst(4003)     // Premium Bonus    - Results
const ARTICLE_MAWARDS_RESULTS = OptimulaArticleConst(4004)     // Personal Award   - Results
const ARTICLE_HAWARDS_RESULTS = OptimulaArticleConst(4005)     // Personal Award   - Results
const ARTICLE_SETTLEM_RESULTS = OptimulaArticleConst(4011)     // Setlement - Results
const ARTICLE_SETTLEM_RESNETT = OptimulaArticleConst(4012)     // Setlement - Results
const ARTICLE_ALLOWCE_HOFFICE = OptimulaArticleConst(4021)     // HomeOffice Tariff
const ARTICLE_ALLOWCE_CLOTDAY = OptimulaArticleConst(4022)     // Clothing Daily Tarrif
const ARTICLE_ALLOWCE_CLOTHRS = OptimulaArticleConst(4023)     // Clothing Hours Tarrif
const ARTICLE_ALLOWCE_MEALDAY = OptimulaArticleConst(4024)     // Meal Contribution Tariff
const ARTICLE_OFFSETS_HOFFICE = OptimulaArticleConst(4031)     // HomeOffice Tariff
const ARTICLE_OFFSETS_CLOTDAY = OptimulaArticleConst(4032)     // Clothing Daily Tarrif
const ARTICLE_OFFSETS_CLOTHRS = OptimulaArticleConst(4033)     // Clothing Hours Tarrif
const ARTICLE_OFFSETS_MEALDAY = OptimulaArticleConst(4034)     // Meal Contribution Tariff
const ARTICLE_INCOMES_TAXFREE = OptimulaArticleConst(5001)     // Incomes Tax Free
const ARTICLE_INCOMES_TAXBASE = OptimulaArticleConst(5002)     // Incomes Tax
const ARTICLE_INCOMES_TAXWINS = OptimulaArticleConst(5003)     // Incomes Tax and Insurance
const ARTICLE_INCOMES_SUMMARY = OptimulaArticleConst(5004)     // Incomes Summary

func (tac OptimulaArticleConst) Id() int32 {
	return int32(tac)
}

func (tac OptimulaArticleConst) String() string {
	switch tac {
	case ARTICLE_TIMESHEETS_PLAN: return "ARTICLE_TIMESHEETS_PLAN"
	case ARTICLE_TIMESHEETS_WORK: return "ARTICLE_TIMESHEETS_WORK"
	case ARTICLE_TIMEACTUAL_WORK: return "ARTICLE_TIMEACTUAL_WORK"
	case ARTICLE_MSALARY_TARGETS: return "ARTICLE_MSALARY_TARGETS"
	case ARTICLE_HSALARY_TARGETS: return "ARTICLE_HSALARY_TARGETS"
	case ARTICLE_MAWARDS_TARGETS: return "ARTICLE_MAWARDS_TARGETS"
	case ARTICLE_HAWARDS_TARGETS: return "ARTICLE_HAWARDS_TARGETS"
	case ARTICLE_PREMIUM_TARGETS: return "ARTICLE_PREMIUM_TARGETS"
	case ARTICLE_PREMADV_TARGETS: return "ARTICLE_PREMADV_TARGETS"
	case ARTICLE_PREMEXT_TARGETS: return "ARTICLE_PREMEXT_TARGETS"
	case ARTICLE_AGRWORK_TARGETS: return "ARTICLE_AGRWORK_TARGETS"
	case ARTICLE_AGRTASK_TARGETS: return "ARTICLE_AGRTASK_TARGETS"
	case ARTICLE_OFFWORK_TARGETS: return "ARTICLE_OFFWORK_TARGETS"
	case ARTICLE_OFFTASK_TARGETS: return "ARTICLE_OFFTASK_TARGETS"
	case ARTICLE_SETTLEM_TARGETS: return "ARTICLE_SETTLEM_TARGETS"
	case ARTICLE_SETTLEM_TARNETT: return "ARTICLE_SETTLEM_TARNETT"
	case ARTICLE_SETTLEM_AGRWORK: return "ARTICLE_SETTLEM_AGRWORK"
	case ARTICLE_SETTLEM_AGRTASK: return "ARTICLE_SETTLEM_AGRTASK"
	case ARTICLE_SETTLEM_ALLOWCE: return "ARTICLE_SETTLEM_ALLOWCE"
	case ARTICLE_SETTLEM_ALLNETT: return "ARTICLE_SETTLEM_ALLNETT"
	case ARTICLE_SETTLEM_OFFWORK: return "ARTICLE_SETTLEM_OFFWORK"
	case ARTICLE_SETTLEM_OFFTASK: return "ARTICLE_SETTLEM_OFFTASK"
	case ARTICLE_SETTLEM_OFFSETS: return "ARTICLE_SETTLEM_OFFSETS"
	case ARTICLE_PREMEXT_RESULTS: return "ARTICLE_PREMEXT_RESULTS"
	case ARTICLE_PREMADV_RESULTS: return "ARTICLE_PREMADV_RESULTS"
	case ARTICLE_PREMIUM_RESULTS: return "ARTICLE_PREMIUM_RESULTS"
	case ARTICLE_MAWARDS_RESULTS: return "ARTICLE_MAWARDS_RESULTS"
	case ARTICLE_HAWARDS_RESULTS: return "ARTICLE_HAWARDS_RESULTS"
	case ARTICLE_SETTLEM_RESULTS: return "ARTICLE_SETTLEM_RESULTS"
	case ARTICLE_SETTLEM_RESNETT: return "ARTICLE_SETTLEM_RESNETT"
	case ARTICLE_ALLOWCE_HOFFICE: return "ARTICLE_ALLOWCE_HOFFICE"
	case ARTICLE_ALLOWCE_CLOTDAY: return "ARTICLE_ALLOWCE_CLOTDAY"
	case ARTICLE_ALLOWCE_CLOTHRS: return "ARTICLE_ALLOWCE_CLOTHRS"
	case ARTICLE_ALLOWCE_MEALDAY: return "ARTICLE_ALLOWCE_MEALDAY"
	case ARTICLE_OFFSETS_HOFFICE: return "ARTICLE_OFFSETS_HOFFICE"
	case ARTICLE_OFFSETS_CLOTDAY: return "ARTICLE_OFFSETS_CLOTDAY"
	case ARTICLE_OFFSETS_CLOTHRS: return "ARTICLE_OFFSETS_CLOTHRS"
	case ARTICLE_OFFSETS_MEALDAY: return "ARTICLE_OFFSETS_MEALDAY"
	case ARTICLE_INCOMES_TAXFREE: return "ARTICLE_INCOMES_TAXFREE"
	case ARTICLE_INCOMES_TAXBASE: return "ARTICLE_INCOMES_TAXBASE"
	case ARTICLE_INCOMES_TAXWINS: return "ARTICLE_INCOMES_TAXWINS"
	case ARTICLE_INCOMES_SUMMARY: return "ARTICLE_INCOMES_SUMMARY"
	}
	return "ARTICLE_UNKNOWN"
}

