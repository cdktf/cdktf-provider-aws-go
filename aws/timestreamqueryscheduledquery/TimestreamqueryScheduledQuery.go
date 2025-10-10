// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/timestreamqueryscheduledquery/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/timestreamquery_scheduled_query aws_timestreamquery_scheduled_query}.
type TimestreamqueryScheduledQuery interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreationTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ErrorReportConfiguration() TimestreamqueryScheduledQueryErrorReportConfigurationList
	ErrorReportConfigurationInput() interface{}
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LastRunSummary() TimestreamqueryScheduledQueryLastRunSummaryList
	LastRunSummaryInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NextInvocationTime() *string
	// The tree node.
	Node() constructs.Node
	NotificationConfiguration() TimestreamqueryScheduledQueryNotificationConfigurationList
	NotificationConfigurationInput() interface{}
	PreviousInvocationTime() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	// Experimental.
	RawOverrides() interface{}
	RecentlyFailedRuns() TimestreamqueryScheduledQueryRecentlyFailedRunsList
	RecentlyFailedRunsInput() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ScheduleConfiguration() TimestreamqueryScheduledQueryScheduleConfigurationList
	ScheduleConfigurationInput() interface{}
	State() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	TargetConfiguration() TimestreamqueryScheduledQueryTargetConfigurationList
	TargetConfigurationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() TimestreamqueryScheduledQueryTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutErrorReportConfiguration(value interface{})
	PutLastRunSummary(value interface{})
	PutNotificationConfiguration(value interface{})
	PutRecentlyFailedRuns(value interface{})
	PutScheduleConfiguration(value interface{})
	PutTargetConfiguration(value interface{})
	PutTimeouts(value *TimestreamqueryScheduledQueryTimeouts)
	ResetErrorReportConfiguration()
	ResetKmsKeyId()
	ResetLastRunSummary()
	ResetNotificationConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecentlyFailedRuns()
	ResetRegion()
	ResetScheduleConfiguration()
	ResetTags()
	ResetTargetConfiguration()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for TimestreamqueryScheduledQuery
type jsiiProxy_TimestreamqueryScheduledQuery struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) ErrorReportConfiguration() TimestreamqueryScheduledQueryErrorReportConfigurationList {
	var returns TimestreamqueryScheduledQueryErrorReportConfigurationList
	_jsii_.Get(
		j,
		"errorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) ErrorReportConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorReportConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) LastRunSummary() TimestreamqueryScheduledQueryLastRunSummaryList {
	var returns TimestreamqueryScheduledQueryLastRunSummaryList
	_jsii_.Get(
		j,
		"lastRunSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) LastRunSummaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastRunSummaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) NextInvocationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextInvocationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) NotificationConfiguration() TimestreamqueryScheduledQueryNotificationConfigurationList {
	var returns TimestreamqueryScheduledQueryNotificationConfigurationList
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) NotificationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) PreviousInvocationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"previousInvocationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) RecentlyFailedRuns() TimestreamqueryScheduledQueryRecentlyFailedRunsList {
	var returns TimestreamqueryScheduledQueryRecentlyFailedRunsList
	_jsii_.Get(
		j,
		"recentlyFailedRuns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) RecentlyFailedRunsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recentlyFailedRunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) ScheduleConfiguration() TimestreamqueryScheduledQueryScheduleConfigurationList {
	var returns TimestreamqueryScheduledQueryScheduleConfigurationList
	_jsii_.Get(
		j,
		"scheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) ScheduleConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) TargetConfiguration() TimestreamqueryScheduledQueryTargetConfigurationList {
	var returns TimestreamqueryScheduledQueryTargetConfigurationList
	_jsii_.Get(
		j,
		"targetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) TargetConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) Timeouts() TimestreamqueryScheduledQueryTimeoutsOutputReference {
	var returns TimestreamqueryScheduledQueryTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/timestreamquery_scheduled_query aws_timestreamquery_scheduled_query} Resource.
func NewTimestreamqueryScheduledQuery(scope constructs.Construct, id *string, config *TimestreamqueryScheduledQueryConfig) TimestreamqueryScheduledQuery {
	_init_.Initialize()

	if err := validateNewTimestreamqueryScheduledQueryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamqueryScheduledQuery{}

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQuery",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/timestreamquery_scheduled_query aws_timestreamquery_scheduled_query} Resource.
func NewTimestreamqueryScheduledQuery_Override(t TimestreamqueryScheduledQuery, scope constructs.Construct, id *string, config *TimestreamqueryScheduledQueryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQuery",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TimestreamqueryScheduledQuery)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a TimestreamqueryScheduledQuery resource upon running "cdktf plan <stack-name>".
func TimestreamqueryScheduledQuery_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateTimestreamqueryScheduledQuery_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQuery",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func TimestreamqueryScheduledQuery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamqueryScheduledQuery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQuery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TimestreamqueryScheduledQuery_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamqueryScheduledQuery_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQuery",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TimestreamqueryScheduledQuery_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTimestreamqueryScheduledQuery_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQuery",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TimestreamqueryScheduledQuery_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.timestreamqueryScheduledQuery.TimestreamqueryScheduledQuery",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) PutErrorReportConfiguration(value interface{}) {
	if err := t.validatePutErrorReportConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putErrorReportConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) PutLastRunSummary(value interface{}) {
	if err := t.validatePutLastRunSummaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLastRunSummary",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) PutNotificationConfiguration(value interface{}) {
	if err := t.validatePutNotificationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNotificationConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) PutRecentlyFailedRuns(value interface{}) {
	if err := t.validatePutRecentlyFailedRunsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecentlyFailedRuns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) PutScheduleConfiguration(value interface{}) {
	if err := t.validatePutScheduleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScheduleConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) PutTargetConfiguration(value interface{}) {
	if err := t.validatePutTargetConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) PutTimeouts(value *TimestreamqueryScheduledQueryTimeouts) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetErrorReportConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorReportConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetLastRunSummary() {
	_jsii_.InvokeVoid(
		t,
		"resetLastRunSummary",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetNotificationConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetNotificationConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetRecentlyFailedRuns() {
	_jsii_.InvokeVoid(
		t,
		"resetRecentlyFailedRuns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetScheduleConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetScheduleConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetTargetConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamqueryScheduledQuery) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

