package snsplatformapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/snsplatformapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/sns_platform_application aws_sns_platform_application}.
type SnsPlatformApplication interface {
	cdktf.TerraformResource
	ApplePlatformBundleId() *string
	SetApplePlatformBundleId(val *string)
	ApplePlatformBundleIdInput() *string
	ApplePlatformTeamId() *string
	SetApplePlatformTeamId(val *string)
	ApplePlatformTeamIdInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventDeliveryFailureTopicArn() *string
	SetEventDeliveryFailureTopicArn(val *string)
	EventDeliveryFailureTopicArnInput() *string
	EventEndpointCreatedTopicArn() *string
	SetEventEndpointCreatedTopicArn(val *string)
	EventEndpointCreatedTopicArnInput() *string
	EventEndpointDeletedTopicArn() *string
	SetEventEndpointDeletedTopicArn(val *string)
	EventEndpointDeletedTopicArnInput() *string
	EventEndpointUpdatedTopicArn() *string
	SetEventEndpointUpdatedTopicArn(val *string)
	EventEndpointUpdatedTopicArnInput() *string
	FailureFeedbackRoleArn() *string
	SetFailureFeedbackRoleArn(val *string)
	FailureFeedbackRoleArnInput() *string
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
	Platform() *string
	SetPlatform(val *string)
	PlatformCredential() *string
	SetPlatformCredential(val *string)
	PlatformCredentialInput() *string
	PlatformInput() *string
	PlatformPrincipal() *string
	SetPlatformPrincipal(val *string)
	PlatformPrincipalInput() *string
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
	SuccessFeedbackRoleArn() *string
	SetSuccessFeedbackRoleArn(val *string)
	SuccessFeedbackRoleArnInput() *string
	SuccessFeedbackSampleRate() *string
	SetSuccessFeedbackSampleRate(val *string)
	SuccessFeedbackSampleRateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetApplePlatformBundleId()
	ResetApplePlatformTeamId()
	ResetEventDeliveryFailureTopicArn()
	ResetEventEndpointCreatedTopicArn()
	ResetEventEndpointDeletedTopicArn()
	ResetEventEndpointUpdatedTopicArn()
	ResetFailureFeedbackRoleArn()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatformPrincipal()
	ResetSuccessFeedbackRoleArn()
	ResetSuccessFeedbackSampleRate()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SnsPlatformApplication
type jsiiProxy_SnsPlatformApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SnsPlatformApplication) ApplePlatformBundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePlatformBundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) ApplePlatformBundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePlatformBundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) ApplePlatformTeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePlatformTeamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) ApplePlatformTeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePlatformTeamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventDeliveryFailureTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliveryFailureTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventDeliveryFailureTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliveryFailureTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointCreatedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointCreatedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointCreatedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointCreatedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointDeletedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointDeletedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointDeletedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointDeletedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointUpdatedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointUpdatedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) EventEndpointUpdatedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointUpdatedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) FailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) FailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformCredential() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformCredentialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) PlatformPrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformPrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) SuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) SuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) SuccessFeedbackSampleRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) SuccessFeedbackSampleRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPlatformApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/sns_platform_application aws_sns_platform_application} Resource.
func NewSnsPlatformApplication(scope constructs.Construct, id *string, config *SnsPlatformApplicationConfig) SnsPlatformApplication {
	_init_.Initialize()

	if err := validateNewSnsPlatformApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsPlatformApplication{}

	_jsii_.Create(
		"@cdktf/provider-aws.snsPlatformApplication.SnsPlatformApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/sns_platform_application aws_sns_platform_application} Resource.
func NewSnsPlatformApplication_Override(s SnsPlatformApplication, scope constructs.Construct, id *string, config *SnsPlatformApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.snsPlatformApplication.SnsPlatformApplication",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetApplePlatformBundleId(val *string) {
	if err := j.validateSetApplePlatformBundleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applePlatformBundleId",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetApplePlatformTeamId(val *string) {
	if err := j.validateSetApplePlatformTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applePlatformTeamId",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetEventDeliveryFailureTopicArn(val *string) {
	if err := j.validateSetEventDeliveryFailureTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventDeliveryFailureTopicArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetEventEndpointCreatedTopicArn(val *string) {
	if err := j.validateSetEventEndpointCreatedTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventEndpointCreatedTopicArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetEventEndpointDeletedTopicArn(val *string) {
	if err := j.validateSetEventEndpointDeletedTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventEndpointDeletedTopicArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetEventEndpointUpdatedTopicArn(val *string) {
	if err := j.validateSetEventEndpointUpdatedTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventEndpointUpdatedTopicArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetPlatform(val *string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetPlatformCredential(val *string) {
	if err := j.validateSetPlatformCredentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformCredential",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetPlatformPrincipal(val *string) {
	if err := j.validateSetPlatformPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformPrincipal",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_SnsPlatformApplication)SetSuccessFeedbackSampleRate(val *string) {
	if err := j.validateSetSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successFeedbackSampleRate",
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
func SnsPlatformApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsPlatformApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsPlatformApplication.SnsPlatformApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsPlatformApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsPlatformApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsPlatformApplication.SnsPlatformApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SnsPlatformApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSnsPlatformApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.snsPlatformApplication.SnsPlatformApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SnsPlatformApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.snsPlatformApplication.SnsPlatformApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SnsPlatformApplication) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SnsPlatformApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SnsPlatformApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SnsPlatformApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SnsPlatformApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SnsPlatformApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SnsPlatformApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SnsPlatformApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SnsPlatformApplication) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SnsPlatformApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SnsPlatformApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SnsPlatformApplication) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetApplePlatformBundleId() {
	_jsii_.InvokeVoid(
		s,
		"resetApplePlatformBundleId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetApplePlatformTeamId() {
	_jsii_.InvokeVoid(
		s,
		"resetApplePlatformTeamId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetEventDeliveryFailureTopicArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEventDeliveryFailureTopicArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetEventEndpointCreatedTopicArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEventEndpointCreatedTopicArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetEventEndpointDeletedTopicArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEventEndpointDeletedTopicArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetEventEndpointUpdatedTopicArn() {
	_jsii_.InvokeVoid(
		s,
		"resetEventEndpointUpdatedTopicArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetPlatformPrincipal() {
	_jsii_.InvokeVoid(
		s,
		"resetPlatformPrincipal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) ResetSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SnsPlatformApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsPlatformApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsPlatformApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsPlatformApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

