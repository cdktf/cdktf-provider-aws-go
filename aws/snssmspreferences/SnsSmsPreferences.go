// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snssmspreferences

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/snssmspreferences/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/sns_sms_preferences aws_sns_sms_preferences}.
type SnsSmsPreferences interface {
	cdktf.TerraformResource
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
	DefaultSenderId() *string
	SetDefaultSenderId(val *string)
	DefaultSenderIdInput() *string
	DefaultSmsType() *string
	SetDefaultSmsType(val *string)
	DefaultSmsTypeInput() *string
	DeliveryStatusIamRoleArn() *string
	SetDeliveryStatusIamRoleArn(val *string)
	DeliveryStatusIamRoleArnInput() *string
	DeliveryStatusSuccessSamplingRate() *string
	SetDeliveryStatusSuccessSamplingRate(val *string)
	DeliveryStatusSuccessSamplingRateInput() *string
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
	MonthlySpendLimit() *float64
	SetMonthlySpendLimit(val *float64)
	MonthlySpendLimitInput() *float64
	// The tree node.
	Node() constructs.Node
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
	UsageReportS3Bucket() *string
	SetUsageReportS3Bucket(val *string)
	UsageReportS3BucketInput() *string
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
	ResetDefaultSenderId()
	ResetDefaultSmsType()
	ResetDeliveryStatusIamRoleArn()
	ResetDeliveryStatusSuccessSamplingRate()
	ResetId()
	ResetMonthlySpendLimit()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUsageReportS3Bucket()
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

// The jsii proxy struct for SnsSmsPreferences
type jsiiProxy_SnsSmsPreferences struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsSmsPreferences) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DefaultSenderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSenderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DefaultSenderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSenderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DefaultSmsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSmsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DefaultSmsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSmsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DeliveryStatusIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStatusIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DeliveryStatusIamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStatusIamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DeliveryStatusSuccessSamplingRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStatusSuccessSamplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DeliveryStatusSuccessSamplingRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStatusSuccessSamplingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) MonthlySpendLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlySpendLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) MonthlySpendLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthlySpendLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) UsageReportS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageReportS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsSmsPreferences) UsageReportS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageReportS3BucketInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/sns_sms_preferences aws_sns_sms_preferences} Resource.
func NewSnsSmsPreferences(scope constructs.Construct, id *string, config *SnsSmsPreferencesConfig) SnsSmsPreferences {
	_init_.Initialize()

	if err := validateNewSnsSmsPreferencesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsSmsPreferences{}

	_jsii_.Create(
		"@cdktf/provider-aws.snsSmsPreferences.SnsSmsPreferences",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/sns_sms_preferences aws_sns_sms_preferences} Resource.
func NewSnsSmsPreferences_Override(s SnsSmsPreferences, scope constructs.Construct, id *string, config *SnsSmsPreferencesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.snsSmsPreferences.SnsSmsPreferences",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetDefaultSenderId(val *string) {
	if err := j.validateSetDefaultSenderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSenderId",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetDefaultSmsType(val *string) {
	if err := j.validateSetDefaultSmsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSmsType",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetDeliveryStatusIamRoleArn(val *string) {
	if err := j.validateSetDeliveryStatusIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryStatusIamRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetDeliveryStatusSuccessSamplingRate(val *string) {
	if err := j.validateSetDeliveryStatusSuccessSamplingRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryStatusSuccessSamplingRate",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetMonthlySpendLimit(val *float64) {
	if err := j.validateSetMonthlySpendLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthlySpendLimit",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SnsSmsPreferences)SetUsageReportS3Bucket(val *string) {
	if err := j.validateSetUsageReportS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageReportS3Bucket",
		val,
	)
}

// Generates CDKTF code for importing a SnsSmsPreferences resource upon running "cdktf plan <stack-name>".
func SnsSmsPreferences_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSnsSmsPreferences_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsSmsPreferences.SnsSmsPreferences",
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
func SnsSmsPreferences_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsSmsPreferences_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsSmsPreferences.SnsSmsPreferences",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsSmsPreferences_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsSmsPreferences_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsSmsPreferences.SnsSmsPreferences",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsSmsPreferences_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsSmsPreferences_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsSmsPreferences.SnsSmsPreferences",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsSmsPreferences_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.snsSmsPreferences.SnsSmsPreferences",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SnsSmsPreferences) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SnsSmsPreferences) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SnsSmsPreferences) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SnsSmsPreferences) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SnsSmsPreferences) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SnsSmsPreferences) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SnsSmsPreferences) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SnsSmsPreferences) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SnsSmsPreferences) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SnsSmsPreferences) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SnsSmsPreferences) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SnsSmsPreferences) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsSmsPreferences) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SnsSmsPreferences) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SnsSmsPreferences) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SnsSmsPreferences) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SnsSmsPreferences) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SnsSmsPreferences) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetDefaultSenderId() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultSenderId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetDefaultSmsType() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultSmsType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetDeliveryStatusIamRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryStatusIamRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetDeliveryStatusSuccessSamplingRate() {
	_jsii_.InvokeVoid(
		s,
		"resetDeliveryStatusSuccessSamplingRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetMonthlySpendLimit() {
	_jsii_.InvokeVoid(
		s,
		"resetMonthlySpendLimit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) ResetUsageReportS3Bucket() {
	_jsii_.InvokeVoid(
		s,
		"resetUsageReportS3Bucket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsSmsPreferences) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsSmsPreferences) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsSmsPreferences) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsSmsPreferences) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsSmsPreferences) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsSmsPreferences) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

