// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudgetaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/budgetsbudgetaction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/budgets_budget_action aws_budgets_budget_action}.
type BudgetsBudgetAction interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	ActionId() *string
	ActionThreshold() BudgetsBudgetActionActionThresholdOutputReference
	ActionThresholdInput() *BudgetsBudgetActionActionThreshold
	ActionType() *string
	SetActionType(val *string)
	ActionTypeInput() *string
	ApprovalModel() *string
	SetApprovalModel(val *string)
	ApprovalModelInput() *string
	Arn() *string
	BudgetName() *string
	SetBudgetName(val *string)
	BudgetNameInput() *string
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
	Definition() BudgetsBudgetActionDefinitionOutputReference
	DefinitionInput() *BudgetsBudgetActionDefinition
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotificationType() *string
	SetNotificationType(val *string)
	NotificationTypeInput() *string
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
	Status() *string
	Subscriber() BudgetsBudgetActionSubscriberList
	SubscriberInput() interface{}
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
	Timeouts() BudgetsBudgetActionTimeoutsOutputReference
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
	PutActionThreshold(value *BudgetsBudgetActionActionThreshold)
	PutDefinition(value *BudgetsBudgetActionDefinition)
	PutSubscriber(value interface{})
	PutTimeouts(value *BudgetsBudgetActionTimeouts)
	ResetAccountId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for BudgetsBudgetAction
type jsiiProxy_BudgetsBudgetAction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BudgetsBudgetAction) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionThreshold() BudgetsBudgetActionActionThresholdOutputReference {
	var returns BudgetsBudgetActionActionThresholdOutputReference
	_jsii_.Get(
		j,
		"actionThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionThresholdInput() *BudgetsBudgetActionActionThreshold {
	var returns *BudgetsBudgetActionActionThreshold
	_jsii_.Get(
		j,
		"actionThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ApprovalModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ApprovalModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) BudgetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) BudgetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Definition() BudgetsBudgetActionDefinitionOutputReference {
	var returns BudgetsBudgetActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) DefinitionInput() *BudgetsBudgetActionDefinition {
	var returns *BudgetsBudgetActionDefinition
	_jsii_.Get(
		j,
		"definitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) NotificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) NotificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Subscriber() BudgetsBudgetActionSubscriberList {
	var returns BudgetsBudgetActionSubscriberList
	_jsii_.Get(
		j,
		"subscriber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) SubscriberInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subscriberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Timeouts() BudgetsBudgetActionTimeoutsOutputReference {
	var returns BudgetsBudgetActionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/budgets_budget_action aws_budgets_budget_action} Resource.
func NewBudgetsBudgetAction(scope constructs.Construct, id *string, config *BudgetsBudgetActionConfig) BudgetsBudgetAction {
	_init_.Initialize()

	if err := validateNewBudgetsBudgetActionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetsBudgetAction{}

	_jsii_.Create(
		"@cdktf/provider-aws.budgetsBudgetAction.BudgetsBudgetAction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/budgets_budget_action aws_budgets_budget_action} Resource.
func NewBudgetsBudgetAction_Override(b BudgetsBudgetAction, scope constructs.Construct, id *string, config *BudgetsBudgetActionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.budgetsBudgetAction.BudgetsBudgetAction",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetActionType(val *string) {
	if err := j.validateSetActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetApprovalModel(val *string) {
	if err := j.validateSetApprovalModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalModel",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetBudgetName(val *string) {
	if err := j.validateSetBudgetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"budgetName",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetNotificationType(val *string) {
	if err := j.validateSetNotificationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a BudgetsBudgetAction resource upon running "cdktf plan <stack-name>".
func BudgetsBudgetAction_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBudgetsBudgetAction_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.budgetsBudgetAction.BudgetsBudgetAction",
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
func BudgetsBudgetAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBudgetsBudgetAction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.budgetsBudgetAction.BudgetsBudgetAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BudgetsBudgetAction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBudgetsBudgetAction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.budgetsBudgetAction.BudgetsBudgetAction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BudgetsBudgetAction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBudgetsBudgetAction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.budgetsBudgetAction.BudgetsBudgetAction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BudgetsBudgetAction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.budgetsBudgetAction.BudgetsBudgetAction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BudgetsBudgetAction) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BudgetsBudgetAction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BudgetsBudgetAction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BudgetsBudgetAction) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BudgetsBudgetAction) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BudgetsBudgetAction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BudgetsBudgetAction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BudgetsBudgetAction) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BudgetsBudgetAction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BudgetsBudgetAction) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAction) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BudgetsBudgetAction) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) PutActionThreshold(value *BudgetsBudgetActionActionThreshold) {
	if err := b.validatePutActionThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putActionThreshold",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) PutDefinition(value *BudgetsBudgetActionDefinition) {
	if err := b.validatePutDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDefinition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) PutSubscriber(value interface{}) {
	if err := b.validatePutSubscriberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSubscriber",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) PutTimeouts(value *BudgetsBudgetActionTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) ResetAccountId() {
	_jsii_.InvokeVoid(
		b,
		"resetAccountId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) ResetTagsAll() {
	_jsii_.InvokeVoid(
		b,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAction) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAction) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

