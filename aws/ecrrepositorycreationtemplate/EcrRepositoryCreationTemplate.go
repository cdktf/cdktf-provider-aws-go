// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrrepositorycreationtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecrrepositorycreationtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/ecr_repository_creation_template aws_ecr_repository_creation_template}.
type EcrRepositoryCreationTemplate interface {
	cdktf.TerraformResource
	AppliedFor() *[]*string
	SetAppliedFor(val *[]*string)
	AppliedForInput() *[]*string
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
	CustomRoleArn() *string
	SetCustomRoleArn(val *string)
	CustomRoleArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EncryptionConfiguration() EcrRepositoryCreationTemplateEncryptionConfigurationList
	EncryptionConfigurationInput() interface{}
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
	ImageTagMutability() *string
	SetImageTagMutability(val *string)
	ImageTagMutabilityInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LifecyclePolicy() *string
	SetLifecyclePolicy(val *string)
	LifecyclePolicyInput() *string
	// The tree node.
	Node() constructs.Node
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
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
	RegistryId() *string
	RepositoryPolicy() *string
	SetRepositoryPolicy(val *string)
	RepositoryPolicyInput() *string
	ResourceTags() *map[string]*string
	SetResourceTags(val *map[string]*string)
	ResourceTagsInput() *map[string]*string
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
	PutEncryptionConfiguration(value interface{})
	ResetCustomRoleArn()
	ResetDescription()
	ResetEncryptionConfiguration()
	ResetId()
	ResetImageTagMutability()
	ResetLifecyclePolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRepositoryPolicy()
	ResetResourceTags()
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

// The jsii proxy struct for EcrRepositoryCreationTemplate
type jsiiProxy_EcrRepositoryCreationTemplate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) AppliedFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appliedFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) AppliedForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appliedForInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) CustomRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) CustomRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) EncryptionConfiguration() EcrRepositoryCreationTemplateEncryptionConfigurationList {
	var returns EcrRepositoryCreationTemplateEncryptionConfigurationList
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) EncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) ImageTagMutability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTagMutability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) ImageTagMutabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTagMutabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) LifecyclePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecyclePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) LifecyclePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecyclePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) RegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) RepositoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) RepositoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) ResourceTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) ResourceTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/ecr_repository_creation_template aws_ecr_repository_creation_template} Resource.
func NewEcrRepositoryCreationTemplate(scope constructs.Construct, id *string, config *EcrRepositoryCreationTemplateConfig) EcrRepositoryCreationTemplate {
	_init_.Initialize()

	if err := validateNewEcrRepositoryCreationTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrRepositoryCreationTemplate{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/ecr_repository_creation_template aws_ecr_repository_creation_template} Resource.
func NewEcrRepositoryCreationTemplate_Override(e EcrRepositoryCreationTemplate, scope constructs.Construct, id *string, config *EcrRepositoryCreationTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplate",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetAppliedFor(val *[]*string) {
	if err := j.validateSetAppliedForParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appliedFor",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetCustomRoleArn(val *string) {
	if err := j.validateSetCustomRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customRoleArn",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetImageTagMutability(val *string) {
	if err := j.validateSetImageTagMutabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageTagMutability",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetLifecyclePolicy(val *string) {
	if err := j.validateSetLifecyclePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecyclePolicy",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetRepositoryPolicy(val *string) {
	if err := j.validateSetRepositoryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryPolicy",
		val,
	)
}

func (j *jsiiProxy_EcrRepositoryCreationTemplate)SetResourceTags(val *map[string]*string) {
	if err := j.validateSetResourceTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTags",
		val,
	)
}

// Generates CDKTF code for importing a EcrRepositoryCreationTemplate resource upon running "cdktf plan <stack-name>".
func EcrRepositoryCreationTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEcrRepositoryCreationTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplate",
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
func EcrRepositoryCreationTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcrRepositoryCreationTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcrRepositoryCreationTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcrRepositoryCreationTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcrRepositoryCreationTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcrRepositoryCreationTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EcrRepositoryCreationTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ecrRepositoryCreationTemplate.EcrRepositoryCreationTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcrRepositoryCreationTemplate) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) PutEncryptionConfiguration(value interface{}) {
	if err := e.validatePutEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetCustomRoleArn() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomRoleArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetImageTagMutability() {
	_jsii_.InvokeVoid(
		e,
		"resetImageTagMutability",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetLifecyclePolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetLifecyclePolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetRepositoryPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRepositoryPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ResetResourceTags() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrRepositoryCreationTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

