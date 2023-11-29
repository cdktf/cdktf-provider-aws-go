// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/sagemakerworkforce/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_workforce aws_sagemaker_workforce}.
type SagemakerWorkforce interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CognitoConfig() SagemakerWorkforceCognitoConfigOutputReference
	CognitoConfigInput() *SagemakerWorkforceCognitoConfig
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
	// The tree node.
	Node() constructs.Node
	OidcConfig() SagemakerWorkforceOidcConfigOutputReference
	OidcConfigInput() *SagemakerWorkforceOidcConfig
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
	SourceIpConfig() SagemakerWorkforceSourceIpConfigOutputReference
	SourceIpConfigInput() *SagemakerWorkforceSourceIpConfig
	Subdomain() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkforceName() *string
	SetWorkforceName(val *string)
	WorkforceNameInput() *string
	WorkforceVpcConfig() SagemakerWorkforceWorkforceVpcConfigOutputReference
	WorkforceVpcConfigInput() *SagemakerWorkforceWorkforceVpcConfig
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCognitoConfig(value *SagemakerWorkforceCognitoConfig)
	PutOidcConfig(value *SagemakerWorkforceOidcConfig)
	PutSourceIpConfig(value *SagemakerWorkforceSourceIpConfig)
	PutWorkforceVpcConfig(value *SagemakerWorkforceWorkforceVpcConfig)
	ResetCognitoConfig()
	ResetId()
	ResetOidcConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSourceIpConfig()
	ResetWorkforceVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SagemakerWorkforce
type jsiiProxy_SagemakerWorkforce struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SagemakerWorkforce) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) CognitoConfig() SagemakerWorkforceCognitoConfigOutputReference {
	var returns SagemakerWorkforceCognitoConfigOutputReference
	_jsii_.Get(
		j,
		"cognitoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) CognitoConfigInput() *SagemakerWorkforceCognitoConfig {
	var returns *SagemakerWorkforceCognitoConfig
	_jsii_.Get(
		j,
		"cognitoConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) OidcConfig() SagemakerWorkforceOidcConfigOutputReference {
	var returns SagemakerWorkforceOidcConfigOutputReference
	_jsii_.Get(
		j,
		"oidcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) OidcConfigInput() *SagemakerWorkforceOidcConfig {
	var returns *SagemakerWorkforceOidcConfig
	_jsii_.Get(
		j,
		"oidcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) SourceIpConfig() SagemakerWorkforceSourceIpConfigOutputReference {
	var returns SagemakerWorkforceSourceIpConfigOutputReference
	_jsii_.Get(
		j,
		"sourceIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) SourceIpConfigInput() *SagemakerWorkforceSourceIpConfig {
	var returns *SagemakerWorkforceSourceIpConfig
	_jsii_.Get(
		j,
		"sourceIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) WorkforceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workforceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) WorkforceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workforceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) WorkforceVpcConfig() SagemakerWorkforceWorkforceVpcConfigOutputReference {
	var returns SagemakerWorkforceWorkforceVpcConfigOutputReference
	_jsii_.Get(
		j,
		"workforceVpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerWorkforce) WorkforceVpcConfigInput() *SagemakerWorkforceWorkforceVpcConfig {
	var returns *SagemakerWorkforceWorkforceVpcConfig
	_jsii_.Get(
		j,
		"workforceVpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_workforce aws_sagemaker_workforce} Resource.
func NewSagemakerWorkforce(scope constructs.Construct, id *string, config *SagemakerWorkforceConfig) SagemakerWorkforce {
	_init_.Initialize()

	if err := validateNewSagemakerWorkforceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerWorkforce{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerWorkforce.SagemakerWorkforce",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/sagemaker_workforce aws_sagemaker_workforce} Resource.
func NewSagemakerWorkforce_Override(s SagemakerWorkforce, scope constructs.Construct, id *string, config *SagemakerWorkforceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerWorkforce.SagemakerWorkforce",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerWorkforce)SetWorkforceName(val *string) {
	if err := j.validateSetWorkforceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workforceName",
		val,
	)
}

// Generates CDKTF code for importing a SagemakerWorkforce resource upon running "cdktf plan <stack-name>".
func SagemakerWorkforce_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerWorkforce_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sagemakerWorkforce.SagemakerWorkforce",
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
func SagemakerWorkforce_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerWorkforce_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sagemakerWorkforce.SagemakerWorkforce",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerWorkforce_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerWorkforce_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sagemakerWorkforce.SagemakerWorkforce",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerWorkforce_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerWorkforce_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.sagemakerWorkforce.SagemakerWorkforce",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerWorkforce_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.sagemakerWorkforce.SagemakerWorkforce",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerWorkforce) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerWorkforce) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerWorkforce) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerWorkforce) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerWorkforce) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerWorkforce) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerWorkforce) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerWorkforce) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerWorkforce) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerWorkforce) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerWorkforce) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) PutCognitoConfig(value *SagemakerWorkforceCognitoConfig) {
	if err := s.validatePutCognitoConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCognitoConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) PutOidcConfig(value *SagemakerWorkforceOidcConfig) {
	if err := s.validatePutOidcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOidcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) PutSourceIpConfig(value *SagemakerWorkforceSourceIpConfig) {
	if err := s.validatePutSourceIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSourceIpConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) PutWorkforceVpcConfig(value *SagemakerWorkforceWorkforceVpcConfig) {
	if err := s.validatePutWorkforceVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkforceVpcConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerWorkforce) ResetCognitoConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCognitoConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkforce) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkforce) ResetOidcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetOidcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkforce) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkforce) ResetSourceIpConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceIpConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkforce) ResetWorkforceVpcConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkforceVpcConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerWorkforce) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerWorkforce) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerWorkforce) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerWorkforce) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

