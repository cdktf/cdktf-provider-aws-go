// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/eksnodegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/eks_node_group aws_eks_node_group}.
type EksNodeGroup interface {
	cdktf.TerraformResource
	AmiType() *string
	SetAmiType(val *string)
	AmiTypeInput() *string
	Arn() *string
	CapacityType() *string
	SetCapacityType(val *string)
	CapacityTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	DiskSize() *float64
	SetDiskSize(val *float64)
	DiskSizeInput() *float64
	ForceUpdateVersion() interface{}
	SetForceUpdateVersion(val interface{})
	ForceUpdateVersionInput() interface{}
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
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	InstanceTypesInput() *[]*string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LaunchTemplate() EksNodeGroupLaunchTemplateOutputReference
	LaunchTemplateInput() *EksNodeGroupLaunchTemplate
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NodeGroupName() *string
	SetNodeGroupName(val *string)
	NodeGroupNameInput() *string
	NodeGroupNamePrefix() *string
	SetNodeGroupNamePrefix(val *string)
	NodeGroupNamePrefixInput() *string
	NodeRoleArn() *string
	SetNodeRoleArn(val *string)
	NodeRoleArnInput() *string
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
	ReleaseVersion() *string
	SetReleaseVersion(val *string)
	ReleaseVersionInput() *string
	RemoteAccess() EksNodeGroupRemoteAccessOutputReference
	RemoteAccessInput() *EksNodeGroupRemoteAccess
	Resources() EksNodeGroupResourcesList
	ScalingConfig() EksNodeGroupScalingConfigOutputReference
	ScalingConfigInput() *EksNodeGroupScalingConfig
	Status() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	Taint() EksNodeGroupTaintList
	TaintInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EksNodeGroupTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateConfig() EksNodeGroupUpdateConfigOutputReference
	UpdateConfigInput() *EksNodeGroupUpdateConfig
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutLaunchTemplate(value *EksNodeGroupLaunchTemplate)
	PutRemoteAccess(value *EksNodeGroupRemoteAccess)
	PutScalingConfig(value *EksNodeGroupScalingConfig)
	PutTaint(value interface{})
	PutTimeouts(value *EksNodeGroupTimeouts)
	PutUpdateConfig(value *EksNodeGroupUpdateConfig)
	ResetAmiType()
	ResetCapacityType()
	ResetDiskSize()
	ResetForceUpdateVersion()
	ResetId()
	ResetInstanceTypes()
	ResetLabels()
	ResetLaunchTemplate()
	ResetNodeGroupName()
	ResetNodeGroupNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReleaseVersion()
	ResetRemoteAccess()
	ResetTags()
	ResetTagsAll()
	ResetTaint()
	ResetTimeouts()
	ResetUpdateConfig()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EksNodeGroup
type jsiiProxy_EksNodeGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EksNodeGroup) AmiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) AmiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) CapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) CapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) DiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ForceUpdateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ForceUpdateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) LaunchTemplate() EksNodeGroupLaunchTemplateOutputReference {
	var returns EksNodeGroupLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) LaunchTemplateInput() *EksNodeGroupLaunchTemplate {
	var returns *EksNodeGroupLaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) NodeGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) NodeGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) NodeGroupNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) NodeGroupNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) NodeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) NodeRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ReleaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) RemoteAccess() EksNodeGroupRemoteAccessOutputReference {
	var returns EksNodeGroupRemoteAccessOutputReference
	_jsii_.Get(
		j,
		"remoteAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) RemoteAccessInput() *EksNodeGroupRemoteAccess {
	var returns *EksNodeGroupRemoteAccess
	_jsii_.Get(
		j,
		"remoteAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Resources() EksNodeGroupResourcesList {
	var returns EksNodeGroupResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ScalingConfig() EksNodeGroupScalingConfigOutputReference {
	var returns EksNodeGroupScalingConfigOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) ScalingConfigInput() *EksNodeGroupScalingConfig {
	var returns *EksNodeGroupScalingConfig
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Taint() EksNodeGroupTaintList {
	var returns EksNodeGroupTaintList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) TaintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Timeouts() EksNodeGroupTimeoutsOutputReference {
	var returns EksNodeGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) UpdateConfig() EksNodeGroupUpdateConfigOutputReference {
	var returns EksNodeGroupUpdateConfigOutputReference
	_jsii_.Get(
		j,
		"updateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) UpdateConfigInput() *EksNodeGroupUpdateConfig {
	var returns *EksNodeGroupUpdateConfig
	_jsii_.Get(
		j,
		"updateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodeGroup) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/eks_node_group aws_eks_node_group} Resource.
func NewEksNodeGroup(scope constructs.Construct, id *string, config *EksNodeGroupConfig) EksNodeGroup {
	_init_.Initialize()

	if err := validateNewEksNodeGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksNodeGroup{}

	_jsii_.Create(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/eks_node_group aws_eks_node_group} Resource.
func NewEksNodeGroup_Override(e EksNodeGroup, scope constructs.Construct, id *string, config *EksNodeGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroup",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetAmiType(val *string) {
	if err := j.validateSetAmiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiType",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetCapacityType(val *string) {
	if err := j.validateSetCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityType",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetDiskSize(val *float64) {
	if err := j.validateSetDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetForceUpdateVersion(val interface{}) {
	if err := j.validateSetForceUpdateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateVersion",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetNodeGroupName(val *string) {
	if err := j.validateSetNodeGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupName",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetNodeGroupNamePrefix(val *string) {
	if err := j.validateSetNodeGroupNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupNamePrefix",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetNodeRoleArn(val *string) {
	if err := j.validateSetNodeRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeRoleArn",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetReleaseVersion(val *string) {
	if err := j.validateSetReleaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseVersion",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EksNodeGroup)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a EksNodeGroup resource upon running "cdktf plan <stack-name>".
func EksNodeGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEksNodeGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroup",
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
func EksNodeGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksNodeGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EksNodeGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksNodeGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EksNodeGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksNodeGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EksNodeGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.eksNodeGroup.EksNodeGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EksNodeGroup) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EksNodeGroup) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EksNodeGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EksNodeGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EksNodeGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EksNodeGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EksNodeGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EksNodeGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EksNodeGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EksNodeGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EksNodeGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EksNodeGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EksNodeGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EksNodeGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EksNodeGroup) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EksNodeGroup) PutLaunchTemplate(value *EksNodeGroupLaunchTemplate) {
	if err := e.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodeGroup) PutRemoteAccess(value *EksNodeGroupRemoteAccess) {
	if err := e.validatePutRemoteAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRemoteAccess",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodeGroup) PutScalingConfig(value *EksNodeGroupScalingConfig) {
	if err := e.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodeGroup) PutTaint(value interface{}) {
	if err := e.validatePutTaintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTaint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodeGroup) PutTimeouts(value *EksNodeGroupTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodeGroup) PutUpdateConfig(value *EksNodeGroupUpdateConfig) {
	if err := e.validatePutUpdateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUpdateConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetAmiType() {
	_jsii_.InvokeVoid(
		e,
		"resetAmiType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetCapacityType() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetDiskSize() {
	_jsii_.InvokeVoid(
		e,
		"resetDiskSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetForceUpdateVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetForceUpdateVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetInstanceTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetNodeGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetNodeGroupNamePrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeGroupNamePrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetReleaseVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetReleaseVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetRemoteAccess() {
	_jsii_.InvokeVoid(
		e,
		"resetRemoteAccess",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetTaint() {
	_jsii_.InvokeVoid(
		e,
		"resetTaint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetUpdateConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdateConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) ResetVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodeGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodeGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodeGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodeGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

