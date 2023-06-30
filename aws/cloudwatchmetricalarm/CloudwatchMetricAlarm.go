package cloudwatchmetricalarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/cloudwatchmetricalarm/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/cloudwatch_metric_alarm aws_cloudwatch_metric_alarm}.
type CloudwatchMetricAlarm interface {
	cdktf.TerraformResource
	ActionsEnabled() interface{}
	SetActionsEnabled(val interface{})
	ActionsEnabledInput() interface{}
	AlarmActions() *[]*string
	SetAlarmActions(val *[]*string)
	AlarmActionsInput() *[]*string
	AlarmDescription() *string
	SetAlarmDescription(val *string)
	AlarmDescriptionInput() *string
	AlarmName() *string
	SetAlarmName(val *string)
	AlarmNameInput() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComparisonOperator() *string
	SetComparisonOperator(val *string)
	ComparisonOperatorInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatapointsToAlarm() *float64
	SetDatapointsToAlarm(val *float64)
	DatapointsToAlarmInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Dimensions() *map[string]*string
	SetDimensions(val *map[string]*string)
	DimensionsInput() *map[string]*string
	EvaluateLowSampleCountPercentiles() *string
	SetEvaluateLowSampleCountPercentiles(val *string)
	EvaluateLowSampleCountPercentilesInput() *string
	EvaluationPeriods() *float64
	SetEvaluationPeriods(val *float64)
	EvaluationPeriodsInput() *float64
	ExtendedStatistic() *string
	SetExtendedStatistic(val *string)
	ExtendedStatisticInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InsufficientDataActions() *[]*string
	SetInsufficientDataActions(val *[]*string)
	InsufficientDataActionsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	MetricQuery() CloudwatchMetricAlarmMetricQueryList
	MetricQueryInput() interface{}
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	OkActions() *[]*string
	SetOkActions(val *[]*string)
	OkActionsInput() *[]*string
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Statistic() *string
	SetStatistic(val *string)
	StatisticInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
	ThresholdMetricId() *string
	SetThresholdMetricId(val *string)
	ThresholdMetricIdInput() *string
	TreatMissingData() *string
	SetTreatMissingData(val *string)
	TreatMissingDataInput() *string
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutMetricQuery(value interface{})
	ResetActionsEnabled()
	ResetAlarmActions()
	ResetAlarmDescription()
	ResetDatapointsToAlarm()
	ResetDimensions()
	ResetEvaluateLowSampleCountPercentiles()
	ResetExtendedStatistic()
	ResetId()
	ResetInsufficientDataActions()
	ResetMetricName()
	ResetMetricQuery()
	ResetNamespace()
	ResetOkActions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriod()
	ResetStatistic()
	ResetTags()
	ResetTagsAll()
	ResetThreshold()
	ResetThresholdMetricId()
	ResetTreatMissingData()
	ResetUnit()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudwatchMetricAlarm
type jsiiProxy_CloudwatchMetricAlarm struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ActionsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ActionsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) AlarmActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) AlarmActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alarmActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) AlarmDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) AlarmDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) AlarmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ComparisonOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ComparisonOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) DatapointsToAlarm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) DatapointsToAlarmInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datapointsToAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Dimensions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) DimensionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) EvaluateLowSampleCountPercentiles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) EvaluateLowSampleCountPercentilesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluateLowSampleCountPercentilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) EvaluationPeriods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) EvaluationPeriodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ExtendedStatistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ExtendedStatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedStatisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) InsufficientDataActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) InsufficientDataActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"insufficientDataActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) MetricQuery() CloudwatchMetricAlarmMetricQueryList {
	var returns CloudwatchMetricAlarmMetricQueryList
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) MetricQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) OkActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) OkActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"okActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ThresholdMetricId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) ThresholdMetricIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdMetricIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) TreatMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) TreatMissingDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatMissingDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchMetricAlarm) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/cloudwatch_metric_alarm aws_cloudwatch_metric_alarm} Resource.
func NewCloudwatchMetricAlarm(scope constructs.Construct, id *string, config *CloudwatchMetricAlarmConfig) CloudwatchMetricAlarm {
	_init_.Initialize()

	if err := validateNewCloudwatchMetricAlarmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchMetricAlarm{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchMetricAlarm.CloudwatchMetricAlarm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/cloudwatch_metric_alarm aws_cloudwatch_metric_alarm} Resource.
func NewCloudwatchMetricAlarm_Override(c CloudwatchMetricAlarm, scope constructs.Construct, id *string, config *CloudwatchMetricAlarmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchMetricAlarm.CloudwatchMetricAlarm",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetActionsEnabled(val interface{}) {
	if err := j.validateSetActionsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionsEnabled",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetAlarmActions(val *[]*string) {
	if err := j.validateSetAlarmActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmActions",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetAlarmDescription(val *string) {
	if err := j.validateSetAlarmDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmDescription",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetAlarmName(val *string) {
	if err := j.validateSetAlarmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarmName",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetComparisonOperator(val *string) {
	if err := j.validateSetComparisonOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonOperator",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetDatapointsToAlarm(val *float64) {
	if err := j.validateSetDatapointsToAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datapointsToAlarm",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetDimensions(val *map[string]*string) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetEvaluateLowSampleCountPercentiles(val *string) {
	if err := j.validateSetEvaluateLowSampleCountPercentilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateLowSampleCountPercentiles",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetEvaluationPeriods(val *float64) {
	if err := j.validateSetEvaluationPeriodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationPeriods",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetExtendedStatistic(val *string) {
	if err := j.validateSetExtendedStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedStatistic",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetInsufficientDataActions(val *[]*string) {
	if err := j.validateSetInsufficientDataActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insufficientDataActions",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetOkActions(val *[]*string) {
	if err := j.validateSetOkActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"okActions",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetThresholdMetricId(val *string) {
	if err := j.validateSetThresholdMetricIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdMetricId",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetTreatMissingData(val *string) {
	if err := j.validateSetTreatMissingDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatMissingData",
		val,
	)
}

func (j *jsiiProxy_CloudwatchMetricAlarm)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CloudwatchMetricAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchMetricAlarm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudwatchMetricAlarm.CloudwatchMetricAlarm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudwatchMetricAlarm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchMetricAlarm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudwatchMetricAlarm.CloudwatchMetricAlarm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudwatchMetricAlarm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudwatchMetricAlarm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudwatchMetricAlarm.CloudwatchMetricAlarm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudwatchMetricAlarm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.cloudwatchMetricAlarm.CloudwatchMetricAlarm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) PutMetricQuery(value interface{}) {
	if err := c.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetActionsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetActionsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetAlarmActions() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetAlarmDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetAlarmDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetDatapointsToAlarm() {
	_jsii_.InvokeVoid(
		c,
		"resetDatapointsToAlarm",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetDimensions() {
	_jsii_.InvokeVoid(
		c,
		"resetDimensions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetEvaluateLowSampleCountPercentiles() {
	_jsii_.InvokeVoid(
		c,
		"resetEvaluateLowSampleCountPercentiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetExtendedStatistic() {
	_jsii_.InvokeVoid(
		c,
		"resetExtendedStatistic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetInsufficientDataActions() {
	_jsii_.InvokeVoid(
		c,
		"resetInsufficientDataActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetMetricName() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetOkActions() {
	_jsii_.InvokeVoid(
		c,
		"resetOkActions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetPeriod() {
	_jsii_.InvokeVoid(
		c,
		"resetPeriod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetStatistic() {
	_jsii_.InvokeVoid(
		c,
		"resetStatistic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetThresholdMetricId() {
	_jsii_.InvokeVoid(
		c,
		"resetThresholdMetricId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetTreatMissingData() {
	_jsii_.InvokeVoid(
		c,
		"resetTreatMissingData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ResetUnit() {
	_jsii_.InvokeVoid(
		c,
		"resetUnit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchMetricAlarm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchMetricAlarm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

