// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksnodejsapplayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/opsworksnodejsapplayer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/opsworks_nodejs_app_layer aws_opsworks_nodejs_app_layer}.
type OpsworksNodejsAppLayer interface {
	cdktf.TerraformResource
	Arn() *string
	AutoAssignElasticIps() interface{}
	SetAutoAssignElasticIps(val interface{})
	AutoAssignElasticIpsInput() interface{}
	AutoAssignPublicIps() interface{}
	SetAutoAssignPublicIps(val interface{})
	AutoAssignPublicIpsInput() interface{}
	AutoHealing() interface{}
	SetAutoHealing(val interface{})
	AutoHealingInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudwatchConfiguration() OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksNodejsAppLayerCloudwatchConfiguration
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
	CustomConfigureRecipes() *[]*string
	SetCustomConfigureRecipes(val *[]*string)
	CustomConfigureRecipesInput() *[]*string
	CustomDeployRecipes() *[]*string
	SetCustomDeployRecipes(val *[]*string)
	CustomDeployRecipesInput() *[]*string
	CustomInstanceProfileArn() *string
	SetCustomInstanceProfileArn(val *string)
	CustomInstanceProfileArnInput() *string
	CustomJson() *string
	SetCustomJson(val *string)
	CustomJsonInput() *string
	CustomSecurityGroupIds() *[]*string
	SetCustomSecurityGroupIds(val *[]*string)
	CustomSecurityGroupIdsInput() *[]*string
	CustomSetupRecipes() *[]*string
	SetCustomSetupRecipes(val *[]*string)
	CustomSetupRecipesInput() *[]*string
	CustomShutdownRecipes() *[]*string
	SetCustomShutdownRecipes(val *[]*string)
	CustomShutdownRecipesInput() *[]*string
	CustomUndeployRecipes() *[]*string
	SetCustomUndeployRecipes(val *[]*string)
	CustomUndeployRecipesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DrainElbOnShutdown() interface{}
	SetDrainElbOnShutdown(val interface{})
	DrainElbOnShutdownInput() interface{}
	EbsVolume() OpsworksNodejsAppLayerEbsVolumeList
	EbsVolumeInput() interface{}
	ElasticLoadBalancer() *string
	SetElasticLoadBalancer(val *string)
	ElasticLoadBalancerInput() *string
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
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceShutdownTimeout() *float64
	SetInstanceShutdownTimeout(val *float64)
	InstanceShutdownTimeoutInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBasedAutoScaling() OpsworksNodejsAppLayerLoadBasedAutoScalingOutputReference
	LoadBasedAutoScalingInput() *OpsworksNodejsAppLayerLoadBasedAutoScaling
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodejsVersion() *string
	SetNodejsVersion(val *string)
	NodejsVersionInput() *string
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
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	SystemPackages() *[]*string
	SetSystemPackages(val *[]*string)
	SystemPackagesInput() *[]*string
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
	UseEbsOptimizedInstances() interface{}
	SetUseEbsOptimizedInstances(val interface{})
	UseEbsOptimizedInstancesInput() interface{}
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
	PutCloudwatchConfiguration(value *OpsworksNodejsAppLayerCloudwatchConfiguration)
	PutEbsVolume(value interface{})
	PutLoadBasedAutoScaling(value *OpsworksNodejsAppLayerLoadBasedAutoScaling)
	ResetAutoAssignElasticIps()
	ResetAutoAssignPublicIps()
	ResetAutoHealing()
	ResetCloudwatchConfiguration()
	ResetCustomConfigureRecipes()
	ResetCustomDeployRecipes()
	ResetCustomInstanceProfileArn()
	ResetCustomJson()
	ResetCustomSecurityGroupIds()
	ResetCustomSetupRecipes()
	ResetCustomShutdownRecipes()
	ResetCustomUndeployRecipes()
	ResetDrainElbOnShutdown()
	ResetEbsVolume()
	ResetElasticLoadBalancer()
	ResetId()
	ResetInstallUpdatesOnBoot()
	ResetInstanceShutdownTimeout()
	ResetLoadBasedAutoScaling()
	ResetName()
	ResetNodejsVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
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

// The jsii proxy struct for OpsworksNodejsAppLayer
type jsiiProxy_OpsworksNodejsAppLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CloudwatchConfiguration() OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksNodejsAppLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CloudwatchConfigurationInput() *OpsworksNodejsAppLayerCloudwatchConfiguration {
	var returns *OpsworksNodejsAppLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) EbsVolume() OpsworksNodejsAppLayerEbsVolumeList {
	var returns OpsworksNodejsAppLayerEbsVolumeList
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) EbsVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) LoadBasedAutoScaling() OpsworksNodejsAppLayerLoadBasedAutoScalingOutputReference {
	var returns OpsworksNodejsAppLayerLoadBasedAutoScalingOutputReference
	_jsii_.Get(
		j,
		"loadBasedAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) LoadBasedAutoScalingInput() *OpsworksNodejsAppLayerLoadBasedAutoScaling {
	var returns *OpsworksNodejsAppLayerLoadBasedAutoScaling
	_jsii_.Get(
		j,
		"loadBasedAutoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) NodejsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodejsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) NodejsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodejsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksNodejsAppLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/opsworks_nodejs_app_layer aws_opsworks_nodejs_app_layer} Resource.
func NewOpsworksNodejsAppLayer(scope constructs.Construct, id *string, config *OpsworksNodejsAppLayerConfig) OpsworksNodejsAppLayer {
	_init_.Initialize()

	if err := validateNewOpsworksNodejsAppLayerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksNodejsAppLayer{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksNodejsAppLayer.OpsworksNodejsAppLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/opsworks_nodejs_app_layer aws_opsworks_nodejs_app_layer} Resource.
func NewOpsworksNodejsAppLayer_Override(o OpsworksNodejsAppLayer, scope constructs.Construct, id *string, config *OpsworksNodejsAppLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksNodejsAppLayer.OpsworksNodejsAppLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetAutoAssignElasticIps(val interface{}) {
	if err := j.validateSetAutoAssignElasticIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetAutoAssignPublicIps(val interface{}) {
	if err := j.validateSetAutoAssignPublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetAutoHealing(val interface{}) {
	if err := j.validateSetAutoHealingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCustomConfigureRecipes(val *[]*string) {
	if err := j.validateSetCustomConfigureRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCustomDeployRecipes(val *[]*string) {
	if err := j.validateSetCustomDeployRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCustomInstanceProfileArn(val *string) {
	if err := j.validateSetCustomInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCustomJson(val *string) {
	if err := j.validateSetCustomJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCustomSecurityGroupIds(val *[]*string) {
	if err := j.validateSetCustomSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCustomSetupRecipes(val *[]*string) {
	if err := j.validateSetCustomSetupRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCustomShutdownRecipes(val *[]*string) {
	if err := j.validateSetCustomShutdownRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetCustomUndeployRecipes(val *[]*string) {
	if err := j.validateSetCustomUndeployRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetDrainElbOnShutdown(val interface{}) {
	if err := j.validateSetDrainElbOnShutdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetElasticLoadBalancer(val *string) {
	if err := j.validateSetElasticLoadBalancerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetInstallUpdatesOnBoot(val interface{}) {
	if err := j.validateSetInstallUpdatesOnBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetInstanceShutdownTimeout(val *float64) {
	if err := j.validateSetInstanceShutdownTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetNodejsVersion(val *string) {
	if err := j.validateSetNodejsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodejsVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetStackId(val *string) {
	if err := j.validateSetStackIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetSystemPackages(val *[]*string) {
	if err := j.validateSetSystemPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksNodejsAppLayer)SetUseEbsOptimizedInstances(val interface{}) {
	if err := j.validateSetUseEbsOptimizedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
		val,
	)
}

// Generates CDKTF code for importing a OpsworksNodejsAppLayer resource upon running "cdktf plan <stack-name>".
func OpsworksNodejsAppLayer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOpsworksNodejsAppLayer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksNodejsAppLayer.OpsworksNodejsAppLayer",
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
func OpsworksNodejsAppLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksNodejsAppLayer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksNodejsAppLayer.OpsworksNodejsAppLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksNodejsAppLayer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksNodejsAppLayer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksNodejsAppLayer.OpsworksNodejsAppLayer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksNodejsAppLayer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksNodejsAppLayer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksNodejsAppLayer.OpsworksNodejsAppLayer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksNodejsAppLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.opsworksNodejsAppLayer.OpsworksNodejsAppLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) PutCloudwatchConfiguration(value *OpsworksNodejsAppLayerCloudwatchConfiguration) {
	if err := o.validatePutCloudwatchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) PutEbsVolume(value interface{}) {
	if err := o.validatePutEbsVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEbsVolume",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) PutLoadBasedAutoScaling(value *OpsworksNodejsAppLayerLoadBasedAutoScaling) {
	if err := o.validatePutLoadBasedAutoScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLoadBasedAutoScaling",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetLoadBasedAutoScaling() {
	_jsii_.InvokeVoid(
		o,
		"resetLoadBasedAutoScaling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetNodejsVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetNodejsVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksNodejsAppLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

