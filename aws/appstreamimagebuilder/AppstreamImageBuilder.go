package appstreamimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/appstreamimagebuilder/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/appstream_image_builder aws_appstream_image_builder}.
type AppstreamImageBuilder interface {
	cdktf.TerraformResource
	AccessEndpoint() AppstreamImageBuilderAccessEndpointList
	AccessEndpointInput() interface{}
	AppstreamAgentVersion() *string
	SetAppstreamAgentVersion(val *string)
	AppstreamAgentVersionInput() *string
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
	CreatedTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DomainJoinInfo() AppstreamImageBuilderDomainJoinInfoOutputReference
	DomainJoinInfoInput() *AppstreamImageBuilderDomainJoinInfo
	EnableDefaultInternetAccess() interface{}
	SetEnableDefaultInternetAccess(val interface{})
	EnableDefaultInternetAccessInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageArn() *string
	SetImageArn(val *string)
	ImageArnInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
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
	State() *string
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
	VpcConfig() AppstreamImageBuilderVpcConfigOutputReference
	VpcConfigInput() *AppstreamImageBuilderVpcConfig
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
	PutAccessEndpoint(value interface{})
	PutDomainJoinInfo(value *AppstreamImageBuilderDomainJoinInfo)
	PutVpcConfig(value *AppstreamImageBuilderVpcConfig)
	ResetAccessEndpoint()
	ResetAppstreamAgentVersion()
	ResetDescription()
	ResetDisplayName()
	ResetDomainJoinInfo()
	ResetEnableDefaultInternetAccess()
	ResetIamRoleArn()
	ResetId()
	ResetImageArn()
	ResetImageName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamImageBuilder
type jsiiProxy_AppstreamImageBuilder struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamImageBuilder) AccessEndpoint() AppstreamImageBuilderAccessEndpointList {
	var returns AppstreamImageBuilderAccessEndpointList
	_jsii_.Get(
		j,
		"accessEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AccessEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AppstreamAgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appstreamAgentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AppstreamAgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appstreamAgentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DomainJoinInfo() AppstreamImageBuilderDomainJoinInfoOutputReference {
	var returns AppstreamImageBuilderDomainJoinInfoOutputReference
	_jsii_.Get(
		j,
		"domainJoinInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DomainJoinInfoInput() *AppstreamImageBuilderDomainJoinInfo {
	var returns *AppstreamImageBuilderDomainJoinInfo
	_jsii_.Get(
		j,
		"domainJoinInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) EnableDefaultInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) EnableDefaultInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) VpcConfig() AppstreamImageBuilderVpcConfigOutputReference {
	var returns AppstreamImageBuilderVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) VpcConfigInput() *AppstreamImageBuilderVpcConfig {
	var returns *AppstreamImageBuilderVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/appstream_image_builder aws_appstream_image_builder} Resource.
func NewAppstreamImageBuilder(scope constructs.Construct, id *string, config *AppstreamImageBuilderConfig) AppstreamImageBuilder {
	_init_.Initialize()

	if err := validateNewAppstreamImageBuilderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppstreamImageBuilder{}

	_jsii_.Create(
		"@cdktf/provider-aws.appstreamImageBuilder.AppstreamImageBuilder",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/appstream_image_builder aws_appstream_image_builder} Resource.
func NewAppstreamImageBuilder_Override(a AppstreamImageBuilder, scope constructs.Construct, id *string, config *AppstreamImageBuilderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appstreamImageBuilder.AppstreamImageBuilder",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetAppstreamAgentVersion(val *string) {
	if err := j.validateSetAppstreamAgentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appstreamAgentVersion",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetEnableDefaultInternetAccess(val interface{}) {
	if err := j.validateSetEnableDefaultInternetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDefaultInternetAccess",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetIamRoleArn(val *string) {
	if err := j.validateSetIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetImageArn(val *string) {
	if err := j.validateSetImageArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
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
func AppstreamImageBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamImageBuilder_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appstreamImageBuilder.AppstreamImageBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppstreamImageBuilder_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamImageBuilder_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appstreamImageBuilder.AppstreamImageBuilder",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppstreamImageBuilder_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppstreamImageBuilder_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.appstreamImageBuilder.AppstreamImageBuilder",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamImageBuilder_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.appstreamImageBuilder.AppstreamImageBuilder",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) PutAccessEndpoint(value interface{}) {
	if err := a.validatePutAccessEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessEndpoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) PutDomainJoinInfo(value *AppstreamImageBuilderDomainJoinInfo) {
	if err := a.validatePutDomainJoinInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDomainJoinInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) PutVpcConfig(value *AppstreamImageBuilderVpcConfig) {
	if err := a.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetAccessEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetAppstreamAgentVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetAppstreamAgentVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDomainJoinInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainJoinInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetEnableDefaultInternetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDefaultInternetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetIamRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetImageArn() {
	_jsii_.InvokeVoid(
		a,
		"resetImageArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetImageName() {
	_jsii_.InvokeVoid(
		a,
		"resetImageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilder) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

