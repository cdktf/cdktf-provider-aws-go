// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrstudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrstudio/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/emr_studio aws_emr_studio}.
type EmrStudio interface {
	cdktf.TerraformResource
	Arn() *string
	AuthMode() *string
	SetAuthMode(val *string)
	AuthModeInput() *string
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
	DefaultS3Location() *string
	SetDefaultS3Location(val *string)
	DefaultS3LocationInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EncryptionKeyArn() *string
	SetEncryptionKeyArn(val *string)
	EncryptionKeyArnInput() *string
	EngineSecurityGroupId() *string
	SetEngineSecurityGroupId(val *string)
	EngineSecurityGroupIdInput() *string
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
	IdpAuthUrl() *string
	SetIdpAuthUrl(val *string)
	IdpAuthUrlInput() *string
	IdpRelayStateParameterName() *string
	SetIdpRelayStateParameterName(val *string)
	IdpRelayStateParameterNameInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	Url() *string
	UserRole() *string
	SetUserRole(val *string)
	UserRoleInput() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	WorkspaceSecurityGroupId() *string
	SetWorkspaceSecurityGroupId(val *string)
	WorkspaceSecurityGroupIdInput() *string
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
	ResetDescription()
	ResetEncryptionKeyArn()
	ResetId()
	ResetIdpAuthUrl()
	ResetIdpRelayStateParameterName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetTags()
	ResetTagsAll()
	ResetUserRole()
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

// The jsii proxy struct for EmrStudio
type jsiiProxy_EmrStudio struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrStudio) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) AuthMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) AuthModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) DefaultS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) DefaultS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) EncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) EngineSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) EngineSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdpAuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpAuthUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdpAuthUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpAuthUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdpRelayStateParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpRelayStateParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdpRelayStateParameterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpRelayStateParameterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) UserRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) UserRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) WorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) WorkspaceSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/emr_studio aws_emr_studio} Resource.
func NewEmrStudio(scope constructs.Construct, id *string, config *EmrStudioConfig) EmrStudio {
	_init_.Initialize()

	if err := validateNewEmrStudioParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrStudio{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrStudio.EmrStudio",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/emr_studio aws_emr_studio} Resource.
func NewEmrStudio_Override(e EmrStudio, scope constructs.Construct, id *string, config *EmrStudioConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrStudio.EmrStudio",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrStudio)SetAuthMode(val *string) {
	if err := j.validateSetAuthModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMode",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetDefaultS3Location(val *string) {
	if err := j.validateSetDefaultS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultS3Location",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetEncryptionKeyArn(val *string) {
	if err := j.validateSetEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetEngineSecurityGroupId(val *string) {
	if err := j.validateSetEngineSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetIdpAuthUrl(val *string) {
	if err := j.validateSetIdpAuthUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpAuthUrl",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetIdpRelayStateParameterName(val *string) {
	if err := j.validateSetIdpRelayStateParameterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpRelayStateParameterName",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetUserRole(val *string) {
	if err := j.validateSetUserRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRole",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_EmrStudio)SetWorkspaceSecurityGroupId(val *string) {
	if err := j.validateSetWorkspaceSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceSecurityGroupId",
		val,
	)
}

// Generates CDKTF code for importing a EmrStudio resource upon running "cdktf plan <stack-name>".
func EmrStudio_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEmrStudio_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrStudio.EmrStudio",
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
func EmrStudio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrStudio_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrStudio.EmrStudio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmrStudio_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrStudio_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrStudio.EmrStudio",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EmrStudio_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrStudio_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.emrStudio.EmrStudio",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrStudio_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.emrStudio.EmrStudio",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EmrStudio) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EmrStudio) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EmrStudio) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EmrStudio) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EmrStudio) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EmrStudio) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EmrStudio) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrStudio) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionKeyArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetIdpAuthUrl() {
	_jsii_.InvokeVoid(
		e,
		"resetIdpAuthUrl",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetIdpRelayStateParameterName() {
	_jsii_.InvokeVoid(
		e,
		"resetIdpRelayStateParameterName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetUserRole() {
	_jsii_.InvokeVoid(
		e,
		"resetUserRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrStudio) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

