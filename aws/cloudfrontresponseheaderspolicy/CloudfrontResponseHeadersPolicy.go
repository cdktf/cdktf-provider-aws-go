// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cloudfrontresponseheaderspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cloudfront_response_headers_policy aws_cloudfront_response_headers_policy}.
type CloudfrontResponseHeadersPolicy interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsConfig() CloudfrontResponseHeadersPolicyCorsConfigOutputReference
	CorsConfigInput() *CloudfrontResponseHeadersPolicyCorsConfig
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomHeadersConfig() CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference
	CustomHeadersConfigInput() *CloudfrontResponseHeadersPolicyCustomHeadersConfig
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Etag() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	RemoveHeadersConfig() CloudfrontResponseHeadersPolicyRemoveHeadersConfigOutputReference
	RemoveHeadersConfigInput() *CloudfrontResponseHeadersPolicyRemoveHeadersConfig
	SecurityHeadersConfig() CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference
	SecurityHeadersConfigInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfig
	ServerTimingHeadersConfig() CloudfrontResponseHeadersPolicyServerTimingHeadersConfigOutputReference
	ServerTimingHeadersConfigInput() *CloudfrontResponseHeadersPolicyServerTimingHeadersConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutCorsConfig(value *CloudfrontResponseHeadersPolicyCorsConfig)
	PutCustomHeadersConfig(value *CloudfrontResponseHeadersPolicyCustomHeadersConfig)
	PutRemoveHeadersConfig(value *CloudfrontResponseHeadersPolicyRemoveHeadersConfig)
	PutSecurityHeadersConfig(value *CloudfrontResponseHeadersPolicySecurityHeadersConfig)
	PutServerTimingHeadersConfig(value *CloudfrontResponseHeadersPolicyServerTimingHeadersConfig)
	ResetComment()
	ResetCorsConfig()
	ResetCustomHeadersConfig()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRemoveHeadersConfig()
	ResetSecurityHeadersConfig()
	ResetServerTimingHeadersConfig()
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

// The jsii proxy struct for CloudfrontResponseHeadersPolicy
type jsiiProxy_CloudfrontResponseHeadersPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CorsConfig() CloudfrontResponseHeadersPolicyCorsConfigOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigOutputReference
	_jsii_.Get(
		j,
		"corsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CorsConfigInput() *CloudfrontResponseHeadersPolicyCorsConfig {
	var returns *CloudfrontResponseHeadersPolicyCorsConfig
	_jsii_.Get(
		j,
		"corsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CustomHeadersConfig() CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference {
	var returns CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference
	_jsii_.Get(
		j,
		"customHeadersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CustomHeadersConfigInput() *CloudfrontResponseHeadersPolicyCustomHeadersConfig {
	var returns *CloudfrontResponseHeadersPolicyCustomHeadersConfig
	_jsii_.Get(
		j,
		"customHeadersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) RemoveHeadersConfig() CloudfrontResponseHeadersPolicyRemoveHeadersConfigOutputReference {
	var returns CloudfrontResponseHeadersPolicyRemoveHeadersConfigOutputReference
	_jsii_.Get(
		j,
		"removeHeadersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) RemoveHeadersConfigInput() *CloudfrontResponseHeadersPolicyRemoveHeadersConfig {
	var returns *CloudfrontResponseHeadersPolicyRemoveHeadersConfig
	_jsii_.Get(
		j,
		"removeHeadersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SecurityHeadersConfig() CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference
	_jsii_.Get(
		j,
		"securityHeadersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SecurityHeadersConfigInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfig {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfig
	_jsii_.Get(
		j,
		"securityHeadersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) ServerTimingHeadersConfig() CloudfrontResponseHeadersPolicyServerTimingHeadersConfigOutputReference {
	var returns CloudfrontResponseHeadersPolicyServerTimingHeadersConfigOutputReference
	_jsii_.Get(
		j,
		"serverTimingHeadersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) ServerTimingHeadersConfigInput() *CloudfrontResponseHeadersPolicyServerTimingHeadersConfig {
	var returns *CloudfrontResponseHeadersPolicyServerTimingHeadersConfig
	_jsii_.Get(
		j,
		"serverTimingHeadersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cloudfront_response_headers_policy aws_cloudfront_response_headers_policy} Resource.
func NewCloudfrontResponseHeadersPolicy(scope constructs.Construct, id *string, config *CloudfrontResponseHeadersPolicyConfig) CloudfrontResponseHeadersPolicy {
	_init_.Initialize()

	if err := validateNewCloudfrontResponseHeadersPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontResponseHeadersPolicy{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cloudfront_response_headers_policy aws_cloudfront_response_headers_policy} Resource.
func NewCloudfrontResponseHeadersPolicy_Override(c CloudfrontResponseHeadersPolicy, scope constructs.Construct, id *string, config *CloudfrontResponseHeadersPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicy",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a CloudfrontResponseHeadersPolicy resource upon running "cdktf plan <stack-name>".
func CloudfrontResponseHeadersPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudfrontResponseHeadersPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicy",
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
func CloudfrontResponseHeadersPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontResponseHeadersPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfrontResponseHeadersPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontResponseHeadersPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfrontResponseHeadersPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfrontResponseHeadersPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontResponseHeadersPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) PutCorsConfig(value *CloudfrontResponseHeadersPolicyCorsConfig) {
	if err := c.validatePutCorsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCorsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) PutCustomHeadersConfig(value *CloudfrontResponseHeadersPolicyCustomHeadersConfig) {
	if err := c.validatePutCustomHeadersConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCustomHeadersConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) PutRemoveHeadersConfig(value *CloudfrontResponseHeadersPolicyRemoveHeadersConfig) {
	if err := c.validatePutRemoveHeadersConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRemoveHeadersConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) PutSecurityHeadersConfig(value *CloudfrontResponseHeadersPolicySecurityHeadersConfig) {
	if err := c.validatePutSecurityHeadersConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecurityHeadersConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) PutServerTimingHeadersConfig(value *CloudfrontResponseHeadersPolicyServerTimingHeadersConfig) {
	if err := c.validatePutServerTimingHeadersConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServerTimingHeadersConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetCorsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCorsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetCustomHeadersConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomHeadersConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetRemoveHeadersConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoveHeadersConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetSecurityHeadersConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityHeadersConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetServerTimingHeadersConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServerTimingHeadersConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

