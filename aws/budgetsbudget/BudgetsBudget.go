// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/budgetsbudget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/budgets_budget aws_budgets_budget}.
type BudgetsBudget interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Arn() *string
	AutoAdjustData() BudgetsBudgetAutoAdjustDataOutputReference
	AutoAdjustDataInput() *BudgetsBudgetAutoAdjustData
	BudgetType() *string
	SetBudgetType(val *string)
	BudgetTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostFilter() BudgetsBudgetCostFilterList
	CostFilterInput() interface{}
	CostTypes() BudgetsBudgetCostTypesOutputReference
	CostTypesInput() *BudgetsBudgetCostTypes
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LimitAmount() *string
	SetLimitAmount(val *string)
	LimitAmountInput() *string
	LimitUnit() *string
	SetLimitUnit(val *string)
	LimitUnitInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	// The tree node.
	Node() constructs.Node
	Notification() BudgetsBudgetNotificationList
	NotificationInput() interface{}
	PlannedLimit() BudgetsBudgetPlannedLimitList
	PlannedLimitInput() interface{}
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimePeriodEnd() *string
	SetTimePeriodEnd(val *string)
	TimePeriodEndInput() *string
	TimePeriodStart() *string
	SetTimePeriodStart(val *string)
	TimePeriodStartInput() *string
	TimeUnit() *string
	SetTimeUnit(val *string)
	TimeUnitInput() *string
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
	PutAutoAdjustData(value *BudgetsBudgetAutoAdjustData)
	PutCostFilter(value interface{})
	PutCostTypes(value *BudgetsBudgetCostTypes)
	PutNotification(value interface{})
	PutPlannedLimit(value interface{})
	ResetAccountId()
	ResetAutoAdjustData()
	ResetCostFilter()
	ResetCostTypes()
	ResetId()
	ResetLimitAmount()
	ResetLimitUnit()
	ResetName()
	ResetNamePrefix()
	ResetNotification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlannedLimit()
	ResetTimePeriodEnd()
	ResetTimePeriodStart()
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

// The jsii proxy struct for BudgetsBudget
type jsiiProxy_BudgetsBudget struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BudgetsBudget) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) AutoAdjustData() BudgetsBudgetAutoAdjustDataOutputReference {
	var returns BudgetsBudgetAutoAdjustDataOutputReference
	_jsii_.Get(
		j,
		"autoAdjustData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) AutoAdjustDataInput() *BudgetsBudgetAutoAdjustData {
	var returns *BudgetsBudgetAutoAdjustData
	_jsii_.Get(
		j,
		"autoAdjustDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) BudgetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) BudgetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostFilter() BudgetsBudgetCostFilterList {
	var returns BudgetsBudgetCostFilterList
	_jsii_.Get(
		j,
		"costFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"costFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostTypes() BudgetsBudgetCostTypesOutputReference {
	var returns BudgetsBudgetCostTypesOutputReference
	_jsii_.Get(
		j,
		"costTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostTypesInput() *BudgetsBudgetCostTypes {
	var returns *BudgetsBudgetCostTypes
	_jsii_.Get(
		j,
		"costTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) LimitAmount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitAmount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) LimitAmountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitAmountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) LimitUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) LimitUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Notification() BudgetsBudgetNotificationList {
	var returns BudgetsBudgetNotificationList
	_jsii_.Get(
		j,
		"notification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) NotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) PlannedLimit() BudgetsBudgetPlannedLimitList {
	var returns BudgetsBudgetPlannedLimitList
	_jsii_.Get(
		j,
		"plannedLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) PlannedLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"plannedLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimePeriodEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimePeriodEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimePeriodStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimePeriodStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnitInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/budgets_budget aws_budgets_budget} Resource.
func NewBudgetsBudget(scope constructs.Construct, id *string, config *BudgetsBudgetConfig) BudgetsBudget {
	_init_.Initialize()

	if err := validateNewBudgetsBudgetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetsBudget{}

	_jsii_.Create(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/budgets_budget aws_budgets_budget} Resource.
func NewBudgetsBudget_Override(b BudgetsBudget, scope constructs.Construct, id *string, config *BudgetsBudgetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudget",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetBudgetType(val *string) {
	if err := j.validateSetBudgetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetLimitAmount(val *string) {
	if err := j.validateSetLimitAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitAmount",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetLimitUnit(val *string) {
	if err := j.validateSetLimitUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitUnit",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetTimePeriodEnd(val *string) {
	if err := j.validateSetTimePeriodEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timePeriodEnd",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetTimePeriodStart(val *string) {
	if err := j.validateSetTimePeriodStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timePeriodStart",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget)SetTimeUnit(val *string) {
	if err := j.validateSetTimeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeUnit",
		val,
	)
}

// Generates CDKTF code for importing a BudgetsBudget resource upon running "cdktf plan <stack-name>".
func BudgetsBudget_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBudgetsBudget_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudget",
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
func BudgetsBudget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBudgetsBudget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BudgetsBudget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBudgetsBudget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BudgetsBudget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBudgetsBudget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BudgetsBudget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BudgetsBudget) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BudgetsBudget) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BudgetsBudget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BudgetsBudget) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BudgetsBudget) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BudgetsBudget) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BudgetsBudget) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BudgetsBudget) PutAutoAdjustData(value *BudgetsBudgetAutoAdjustData) {
	if err := b.validatePutAutoAdjustDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAutoAdjustData",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudget) PutCostFilter(value interface{}) {
	if err := b.validatePutCostFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCostFilter",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudget) PutCostTypes(value *BudgetsBudgetCostTypes) {
	if err := b.validatePutCostTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCostTypes",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudget) PutNotification(value interface{}) {
	if err := b.validatePutNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNotification",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudget) PutPlannedLimit(value interface{}) {
	if err := b.validatePutPlannedLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPlannedLimit",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetAccountId() {
	_jsii_.InvokeVoid(
		b,
		"resetAccountId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetAutoAdjustData() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoAdjustData",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetCostFilter() {
	_jsii_.InvokeVoid(
		b,
		"resetCostFilter",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetCostTypes() {
	_jsii_.InvokeVoid(
		b,
		"resetCostTypes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetLimitAmount() {
	_jsii_.InvokeVoid(
		b,
		"resetLimitAmount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetLimitUnit() {
	_jsii_.InvokeVoid(
		b,
		"resetLimitUnit",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		b,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetNotification() {
	_jsii_.InvokeVoid(
		b,
		"resetNotification",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetPlannedLimit() {
	_jsii_.InvokeVoid(
		b,
		"resetPlannedLimit",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetTimePeriodEnd() {
	_jsii_.InvokeVoid(
		b,
		"resetTimePeriodEnd",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetTimePeriodStart() {
	_jsii_.InvokeVoid(
		b,
		"resetTimePeriodStart",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

