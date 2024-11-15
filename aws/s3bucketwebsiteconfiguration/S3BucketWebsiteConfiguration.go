// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketwebsiteconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucketwebsiteconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/s3_bucket_website_configuration aws_s3_bucket_website_configuration}.
type S3BucketWebsiteConfiguration interface {
	cdktf.TerraformResource
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ErrorDocument() S3BucketWebsiteConfigurationErrorDocumentOutputReference
	ErrorDocumentInput() *S3BucketWebsiteConfigurationErrorDocument
	ExpectedBucketOwner() *string
	SetExpectedBucketOwner(val *string)
	ExpectedBucketOwnerInput() *string
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
	IndexDocument() S3BucketWebsiteConfigurationIndexDocumentOutputReference
	IndexDocumentInput() *S3BucketWebsiteConfigurationIndexDocument
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	RedirectAllRequestsTo() S3BucketWebsiteConfigurationRedirectAllRequestsToOutputReference
	RedirectAllRequestsToInput() *S3BucketWebsiteConfigurationRedirectAllRequestsTo
	RoutingRule() S3BucketWebsiteConfigurationRoutingRuleList
	RoutingRuleInput() interface{}
	RoutingRules() *string
	SetRoutingRules(val *string)
	RoutingRulesInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WebsiteDomain() *string
	WebsiteEndpoint() *string
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
	PutErrorDocument(value *S3BucketWebsiteConfigurationErrorDocument)
	PutIndexDocument(value *S3BucketWebsiteConfigurationIndexDocument)
	PutRedirectAllRequestsTo(value *S3BucketWebsiteConfigurationRedirectAllRequestsTo)
	PutRoutingRule(value interface{})
	ResetErrorDocument()
	ResetExpectedBucketOwner()
	ResetId()
	ResetIndexDocument()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRedirectAllRequestsTo()
	ResetRoutingRule()
	ResetRoutingRules()
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

// The jsii proxy struct for S3BucketWebsiteConfiguration
type jsiiProxy_S3BucketWebsiteConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) ErrorDocument() S3BucketWebsiteConfigurationErrorDocumentOutputReference {
	var returns S3BucketWebsiteConfigurationErrorDocumentOutputReference
	_jsii_.Get(
		j,
		"errorDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) ErrorDocumentInput() *S3BucketWebsiteConfigurationErrorDocument {
	var returns *S3BucketWebsiteConfigurationErrorDocument
	_jsii_.Get(
		j,
		"errorDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) ExpectedBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) ExpectedBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) IndexDocument() S3BucketWebsiteConfigurationIndexDocumentOutputReference {
	var returns S3BucketWebsiteConfigurationIndexDocumentOutputReference
	_jsii_.Get(
		j,
		"indexDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) IndexDocumentInput() *S3BucketWebsiteConfigurationIndexDocument {
	var returns *S3BucketWebsiteConfigurationIndexDocument
	_jsii_.Get(
		j,
		"indexDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) RedirectAllRequestsTo() S3BucketWebsiteConfigurationRedirectAllRequestsToOutputReference {
	var returns S3BucketWebsiteConfigurationRedirectAllRequestsToOutputReference
	_jsii_.Get(
		j,
		"redirectAllRequestsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) RedirectAllRequestsToInput() *S3BucketWebsiteConfigurationRedirectAllRequestsTo {
	var returns *S3BucketWebsiteConfigurationRedirectAllRequestsTo
	_jsii_.Get(
		j,
		"redirectAllRequestsToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) RoutingRule() S3BucketWebsiteConfigurationRoutingRuleList {
	var returns S3BucketWebsiteConfigurationRoutingRuleList
	_jsii_.Get(
		j,
		"routingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) RoutingRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) RoutingRules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) RoutingRulesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) WebsiteDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration) WebsiteEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteEndpoint",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/s3_bucket_website_configuration aws_s3_bucket_website_configuration} Resource.
func NewS3BucketWebsiteConfiguration(scope constructs.Construct, id *string, config *S3BucketWebsiteConfigurationConfig) S3BucketWebsiteConfiguration {
	_init_.Initialize()

	if err := validateNewS3BucketWebsiteConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketWebsiteConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-aws.s3BucketWebsiteConfiguration.S3BucketWebsiteConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/s3_bucket_website_configuration aws_s3_bucket_website_configuration} Resource.
func NewS3BucketWebsiteConfiguration_Override(s S3BucketWebsiteConfiguration, scope constructs.Construct, id *string, config *S3BucketWebsiteConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.s3BucketWebsiteConfiguration.S3BucketWebsiteConfiguration",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetExpectedBucketOwner(val *string) {
	if err := j.validateSetExpectedBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBucketOwner",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfiguration)SetRoutingRules(val *string) {
	if err := j.validateSetRoutingRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingRules",
		val,
	)
}

// Generates CDKTF code for importing a S3BucketWebsiteConfiguration resource upon running "cdktf plan <stack-name>".
func S3BucketWebsiteConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateS3BucketWebsiteConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3BucketWebsiteConfiguration.S3BucketWebsiteConfiguration",
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
func S3BucketWebsiteConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3BucketWebsiteConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3BucketWebsiteConfiguration.S3BucketWebsiteConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3BucketWebsiteConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3BucketWebsiteConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3BucketWebsiteConfiguration.S3BucketWebsiteConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func S3BucketWebsiteConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3BucketWebsiteConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.s3BucketWebsiteConfiguration.S3BucketWebsiteConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func S3BucketWebsiteConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.s3BucketWebsiteConfiguration.S3BucketWebsiteConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketWebsiteConfiguration) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) PutErrorDocument(value *S3BucketWebsiteConfigurationErrorDocument) {
	if err := s.validatePutErrorDocumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putErrorDocument",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) PutIndexDocument(value *S3BucketWebsiteConfigurationIndexDocument) {
	if err := s.validatePutIndexDocumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIndexDocument",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) PutRedirectAllRequestsTo(value *S3BucketWebsiteConfigurationRedirectAllRequestsTo) {
	if err := s.validatePutRedirectAllRequestsToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRedirectAllRequestsTo",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) PutRoutingRule(value interface{}) {
	if err := s.validatePutRoutingRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRoutingRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ResetErrorDocument() {
	_jsii_.InvokeVoid(
		s,
		"resetErrorDocument",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ResetExpectedBucketOwner() {
	_jsii_.InvokeVoid(
		s,
		"resetExpectedBucketOwner",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ResetIndexDocument() {
	_jsii_.InvokeVoid(
		s,
		"resetIndexDocument",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ResetRedirectAllRequestsTo() {
	_jsii_.InvokeVoid(
		s,
		"resetRedirectAllRequestsTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ResetRoutingRule() {
	_jsii_.InvokeVoid(
		s,
		"resetRoutingRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ResetRoutingRules() {
	_jsii_.InvokeVoid(
		s,
		"resetRoutingRules",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

