package opsworksecsclusterlayer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/opsworksecsclusterlayer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_ecs_cluster_layer aws_opsworks_ecs_cluster_layer}.
type OpsworksEcsClusterLayer interface {
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
	CloudwatchConfiguration() OpsworksEcsClusterLayerCloudwatchConfigurationOutputReference
	CloudwatchConfigurationInput() *OpsworksEcsClusterLayerCloudwatchConfiguration
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
	EbsVolume() OpsworksEcsClusterLayerEbsVolumeList
	EbsVolumeInput() interface{}
	EcsClusterArn() *string
	SetEcsClusterArn(val *string)
	EcsClusterArnInput() *string
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
	LoadBasedAutoScaling() OpsworksEcsClusterLayerLoadBasedAutoScalingOutputReference
	LoadBasedAutoScalingInput() *OpsworksEcsClusterLayerLoadBasedAutoScaling
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
	PutCloudwatchConfiguration(value *OpsworksEcsClusterLayerCloudwatchConfiguration)
	PutEbsVolume(value interface{})
	PutLoadBasedAutoScaling(value *OpsworksEcsClusterLayerLoadBasedAutoScaling)
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSystemPackages()
	ResetTags()
	ResetTagsAll()
	ResetUseEbsOptimizedInstances()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OpsworksEcsClusterLayer
type jsiiProxy_OpsworksEcsClusterLayer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) AutoAssignElasticIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) AutoAssignElasticIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignElasticIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) AutoAssignPublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) AutoAssignPublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssignPublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) AutoHealing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) AutoHealingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHealingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CloudwatchConfiguration() OpsworksEcsClusterLayerCloudwatchConfigurationOutputReference {
	var returns OpsworksEcsClusterLayerCloudwatchConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CloudwatchConfigurationInput() *OpsworksEcsClusterLayerCloudwatchConfiguration {
	var returns *OpsworksEcsClusterLayerCloudwatchConfiguration
	_jsii_.Get(
		j,
		"cloudwatchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomConfigureRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomConfigureRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customConfigureRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomDeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomDeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomSetupRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomSetupRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customSetupRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomShutdownRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomShutdownRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customShutdownRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomUndeployRecipes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) CustomUndeployRecipesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customUndeployRecipesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) DrainElbOnShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) DrainElbOnShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drainElbOnShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) EbsVolume() OpsworksEcsClusterLayerEbsVolumeList {
	var returns OpsworksEcsClusterLayerEbsVolumeList
	_jsii_.Get(
		j,
		"ebsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) EbsVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) EcsClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) EcsClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) ElasticLoadBalancer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) ElasticLoadBalancerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticLoadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) InstanceShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) InstanceShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) LoadBasedAutoScaling() OpsworksEcsClusterLayerLoadBasedAutoScalingOutputReference {
	var returns OpsworksEcsClusterLayerLoadBasedAutoScalingOutputReference
	_jsii_.Get(
		j,
		"loadBasedAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) LoadBasedAutoScalingInput() *OpsworksEcsClusterLayerLoadBasedAutoScaling {
	var returns *OpsworksEcsClusterLayerLoadBasedAutoScaling
	_jsii_.Get(
		j,
		"loadBasedAutoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) SystemPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) SystemPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) UseEbsOptimizedInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksEcsClusterLayer) UseEbsOptimizedInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useEbsOptimizedInstancesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_ecs_cluster_layer aws_opsworks_ecs_cluster_layer} Resource.
func NewOpsworksEcsClusterLayer(scope constructs.Construct, id *string, config *OpsworksEcsClusterLayerConfig) OpsworksEcsClusterLayer {
	_init_.Initialize()

	if err := validateNewOpsworksEcsClusterLayerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksEcsClusterLayer{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksEcsClusterLayer.OpsworksEcsClusterLayer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_ecs_cluster_layer aws_opsworks_ecs_cluster_layer} Resource.
func NewOpsworksEcsClusterLayer_Override(o OpsworksEcsClusterLayer, scope constructs.Construct, id *string, config *OpsworksEcsClusterLayerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksEcsClusterLayer.OpsworksEcsClusterLayer",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetAutoAssignElasticIps(val interface{}) {
	if err := j.validateSetAutoAssignElasticIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAssignElasticIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetAutoAssignPublicIps(val interface{}) {
	if err := j.validateSetAutoAssignPublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAssignPublicIps",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetAutoHealing(val interface{}) {
	if err := j.validateSetAutoHealingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHealing",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCustomConfigureRecipes(val *[]*string) {
	if err := j.validateSetCustomConfigureRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customConfigureRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCustomDeployRecipes(val *[]*string) {
	if err := j.validateSetCustomDeployRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCustomInstanceProfileArn(val *string) {
	if err := j.validateSetCustomInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCustomJson(val *string) {
	if err := j.validateSetCustomJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customJson",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCustomSecurityGroupIds(val *[]*string) {
	if err := j.validateSetCustomSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCustomSetupRecipes(val *[]*string) {
	if err := j.validateSetCustomSetupRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSetupRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCustomShutdownRecipes(val *[]*string) {
	if err := j.validateSetCustomShutdownRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customShutdownRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetCustomUndeployRecipes(val *[]*string) {
	if err := j.validateSetCustomUndeployRecipesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customUndeployRecipes",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetDrainElbOnShutdown(val interface{}) {
	if err := j.validateSetDrainElbOnShutdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drainElbOnShutdown",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetEcsClusterArn(val *string) {
	if err := j.validateSetEcsClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecsClusterArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetElasticLoadBalancer(val *string) {
	if err := j.validateSetElasticLoadBalancerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticLoadBalancer",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetInstallUpdatesOnBoot(val interface{}) {
	if err := j.validateSetInstallUpdatesOnBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetInstanceShutdownTimeout(val *float64) {
	if err := j.validateSetInstanceShutdownTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetStackId(val *string) {
	if err := j.validateSetStackIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetSystemPackages(val *[]*string) {
	if err := j.validateSetSystemPackagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemPackages",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OpsworksEcsClusterLayer)SetUseEbsOptimizedInstances(val interface{}) {
	if err := j.validateSetUseEbsOptimizedInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useEbsOptimizedInstances",
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
func OpsworksEcsClusterLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksEcsClusterLayer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksEcsClusterLayer.OpsworksEcsClusterLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksEcsClusterLayer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksEcsClusterLayer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksEcsClusterLayer.OpsworksEcsClusterLayer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksEcsClusterLayer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksEcsClusterLayer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksEcsClusterLayer.OpsworksEcsClusterLayer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksEcsClusterLayer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.opsworksEcsClusterLayer.OpsworksEcsClusterLayer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OpsworksEcsClusterLayer) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) PutCloudwatchConfiguration(value *OpsworksEcsClusterLayerCloudwatchConfiguration) {
	if err := o.validatePutCloudwatchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCloudwatchConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) PutEbsVolume(value interface{}) {
	if err := o.validatePutEbsVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEbsVolume",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) PutLoadBasedAutoScaling(value *OpsworksEcsClusterLayerLoadBasedAutoScaling) {
	if err := o.validatePutLoadBasedAutoScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLoadBasedAutoScaling",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetAutoAssignElasticIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignElasticIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetAutoAssignPublicIps() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoAssignPublicIps",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetAutoHealing() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoHealing",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCloudwatchConfiguration() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudwatchConfiguration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCustomConfigureRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomConfigureRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCustomDeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomDeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCustomInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCustomJson() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCustomSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCustomSetupRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomSetupRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCustomShutdownRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomShutdownRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetCustomUndeployRecipes() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomUndeployRecipes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetDrainElbOnShutdown() {
	_jsii_.InvokeVoid(
		o,
		"resetDrainElbOnShutdown",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetEbsVolume() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsVolume",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetElasticLoadBalancer() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticLoadBalancer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetInstanceShutdownTimeout() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceShutdownTimeout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetLoadBasedAutoScaling() {
	_jsii_.InvokeVoid(
		o,
		"resetLoadBasedAutoScaling",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetSystemPackages() {
	_jsii_.InvokeVoid(
		o,
		"resetSystemPackages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ResetUseEbsOptimizedInstances() {
	_jsii_.InvokeVoid(
		o,
		"resetUseEbsOptimizedInstances",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksEcsClusterLayer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

