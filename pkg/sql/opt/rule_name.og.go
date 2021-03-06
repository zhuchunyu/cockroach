// Code generated by optgen; DO NOT EDIT.

package opt

const (
	startAutoRule RuleName = iota + NumManualRuleNames

	// ------------------------------------------------------------
	// Normalize Rule Names
	// ------------------------------------------------------------
	EliminateEmptyAnd
	EliminateEmptyOr
	EliminateSingletonAndOr
	SimplifyAnd
	SimplifyOr
	SimplifyFilters
	FoldNullAndOr
	NegateComparison
	EliminateNot
	NegateAnd
	NegateOr
	CommuteVarInequality
	CommuteConstInequality
	NormalizeCmpPlusConst
	NormalizeCmpMinusConst
	NormalizeCmpConstMinus
	NormalizeTupleEquality
	FoldNullComparisonLeft
	FoldNullComparisonRight
	EliminateDistinct
	EnsureJoinFiltersAnd
	EnsureJoinFilters
	PushFilterIntoJoinLeft
	PushFilterIntoJoinRight
	PushLimitIntoProject
	PushOffsetIntoProject
	FoldPlusZero
	FoldZeroPlus
	FoldMinusZero
	FoldMultOne
	FoldOneMult
	FoldDivOne
	InvertMinus
	EliminateUnaryMinus
	EliminateProject
	EliminateProjectProject
	FilterUnusedProjectCols
	FilterUnusedScanCols
	FilterUnusedSelectCols
	FilterUnusedLimitCols
	FilterUnusedOffsetCols
	FilterUnusedJoinLeftCols
	FilterUnusedJoinRightCols
	FilterUnusedAggCols
	FilterUnusedGroupByCols
	FilterUnusedValueCols
	CommuteVar
	CommuteConst
	EliminateCoalesce
	SimplifyCoalesce
	EliminateCast
	FoldNullCast
	FoldNullUnary
	FoldNullBinaryLeft
	FoldNullBinaryRight
	FoldNullInNonEmpty
	FoldNullInEmpty
	FoldNullNotInEmpty
	NormalizeInConst
	FoldInNull
	EnsureSelectFiltersAnd
	EnsureSelectFilters
	EliminateSelect
	MergeSelects
	PushSelectIntoProject
	PushSelectIntoJoinLeft
	PushSelectIntoJoinRight
	MergeSelectInnerJoin
	PushSelectIntoGroupBy

	// ------------------------------------------------------------
	// Explore Rule Names
	// ------------------------------------------------------------
	PushLimitIntoScan
	GenerateIndexScans
	ConstrainScan

	// NumRuleNames tracks the total count of rule names.
	NumRuleNames
)
