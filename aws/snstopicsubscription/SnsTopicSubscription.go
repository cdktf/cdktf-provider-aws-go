// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snstopicsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/snstopicsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/sns_topic_subscription aws_sns_topic_subscription}.
type SnsTopicSubscription interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfirmationTimeoutInMinutes() *float64
	SetConfirmationTimeoutInMinutes(val *float64)
	ConfirmationTimeoutInMinutesInput() *float64
	ConfirmationWasAuthenticated() cdktf.IResolvable
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
	DeliveryPolicy() *string
	SetDeliveryPolicy(val *string)
	DeliveryPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointAutoConfirms() interface{}
	SetEndpointAutoConfirms(val interface{})
	EndpointAutoConfirmsInput() interface{}
	EndpointInput() *string
	FilterPolicy() *string
	SetFilterPolicy(val *string)
	FilterPolicyInput() *string
	FilterPolicyScope() *string
	SetFilterPolicyScope(val *string)
	FilterPolicyScopeInput() *string
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
	OwnerId() *string
	PendingConfirmation() cdktf.IResolvable
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RawMessageDelivery() interface{}
	SetRawMessageDelivery(val interface{})
	RawMessageDeliveryInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RedrivePolicy() *string
	SetRedrivePolicy(val *string)
	RedrivePolicyInput() *string
	ReplayPolicy() *string
	SetReplayPolicy(val *string)
	ReplayPolicyInput() *string
	SubscriptionRoleArn() *string
	SetSubscriptionRoleArn(val *string)
	SubscriptionRoleArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TopicArn() *string
	SetTopicArn(val *string)
	TopicArnInput() *string
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
	ResetConfirmationTimeoutInMinutes()
	ResetDeliveryPolicy()
	ResetEndpointAutoConfirms()
	ResetFilterPolicy()
	ResetFilterPolicyScope()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRawMessageDelivery()
	ResetRedrivePolicy()
	ResetReplayPolicy()
	ResetSubscriptionRoleArn()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SnsTopicSubscription
type jsiiProxy_SnsTopicSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsTopicSubscription) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ConfirmationTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confirmationTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ConfirmationTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confirmationTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ConfirmationWasAuthenticated() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"confirmationWasAuthenticated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) DeliveryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) DeliveryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) EndpointAutoConfirms() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointAutoConfirms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) EndpointAutoConfirmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointAutoConfirmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) FilterPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) FilterPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) FilterPolicyScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPolicyScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) FilterPolicyScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterPolicyScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) PendingConfirmation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pendingConfirmation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RawMessageDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawMessageDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RawMessageDeliveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawMessageDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RedrivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redrivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) RedrivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redrivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ReplayPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replayPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) ReplayPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replayPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) SubscriptionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) SubscriptionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicSubscription) TopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArnInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/sns_topic_subscription aws_sns_topic_subscription} Resource.
func NewSnsTopicSubscription(scope constructs.Construct, id *string, config *SnsTopicSubscriptionConfig) SnsTopicSubscription {
	_init_.Initialize()

	if err := validateNewSnsTopicSubscriptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsTopicSubscription{}

	_jsii_.Create(
		"@cdktf/provider-aws.snsTopicSubscription.SnsTopicSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/sns_topic_subscription aws_sns_topic_subscription} Resource.
func NewSnsTopicSubscription_Override(s SnsTopicSubscription, scope constructs.Construct, id *string, config *SnsTopicSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.snsTopicSubscription.SnsTopicSubscription",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetConfirmationTimeoutInMinutes(val *float64) {
	if err := j.validateSetConfirmationTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confirmationTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetDeliveryPolicy(val *string) {
	if err := j.validateSetDeliveryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryPolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetEndpointAutoConfirms(val interface{}) {
	if err := j.validateSetEndpointAutoConfirmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointAutoConfirms",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetFilterPolicy(val *string) {
	if err := j.validateSetFilterPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetFilterPolicyScope(val *string) {
	if err := j.validateSetFilterPolicyScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPolicyScope",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetRawMessageDelivery(val interface{}) {
	if err := j.validateSetRawMessageDeliveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rawMessageDelivery",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetRedrivePolicy(val *string) {
	if err := j.validateSetRedrivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redrivePolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetReplayPolicy(val *string) {
	if err := j.validateSetReplayPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replayPolicy",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetSubscriptionRoleArn(val *string) {
	if err := j.validateSetSubscriptionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsTopicSubscription)SetTopicArn(val *string) {
	if err := j.validateSetTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicArn",
		val,
	)
}

// Generates CDKTF code for importing a SnsTopicSubscription resource upon running "cdktf plan <stack-name>".
func SnsTopicSubscription_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSnsTopicSubscription_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsTopicSubscription.SnsTopicSubscription",
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
func SnsTopicSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopicSubscription_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsTopicSubscription.SnsTopicSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsTopicSubscription_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopicSubscription_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsTopicSubscription.SnsTopicSubscription",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsTopicSubscription_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsTopicSubscription_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsTopicSubscription.SnsTopicSubscription",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsTopicSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.snsTopicSubscription.SnsTopicSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SnsTopicSubscription) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SnsTopicSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SnsTopicSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SnsTopicSubscription) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SnsTopicSubscription) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SnsTopicSubscription) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetConfirmationTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		s,
		"resetConfirmationTimeoutInMinutes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetDeliveryPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetEndpointAutoConfirms() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointAutoConfirms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetFilterPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetFilterPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetFilterPolicyScope() {
	_jsii_.InvokeVoid(
		s,
		"resetFilterPolicyScope",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetRawMessageDelivery() {
	_jsii_.InvokeVoid(
		s,
		"resetRawMessageDelivery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetRedrivePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetRedrivePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetReplayPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetReplayPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) ResetSubscriptionRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSubscriptionRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsTopicSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

