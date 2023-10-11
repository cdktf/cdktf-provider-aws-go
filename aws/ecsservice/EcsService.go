// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/ecsservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/ecs_service aws_ecs_service}.
type EcsService interface {
	cdktf.TerraformResource
	Alarms() EcsServiceAlarmsOutputReference
	AlarmsInput() *EcsServiceAlarms
	CapacityProviderStrategy() EcsServiceCapacityProviderStrategyList
	CapacityProviderStrategyInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
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
	DeploymentCircuitBreaker() EcsServiceDeploymentCircuitBreakerOutputReference
	DeploymentCircuitBreakerInput() *EcsServiceDeploymentCircuitBreaker
	DeploymentController() EcsServiceDeploymentControllerOutputReference
	DeploymentControllerInput() *EcsServiceDeploymentController
	DeploymentMaximumPercent() *float64
	SetDeploymentMaximumPercent(val *float64)
	DeploymentMaximumPercentInput() *float64
	DeploymentMinimumHealthyPercent() *float64
	SetDeploymentMinimumHealthyPercent(val *float64)
	DeploymentMinimumHealthyPercentInput() *float64
	DesiredCount() *float64
	SetDesiredCount(val *float64)
	DesiredCountInput() *float64
	EnableEcsManagedTags() interface{}
	SetEnableEcsManagedTags(val interface{})
	EnableEcsManagedTagsInput() interface{}
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	EnableExecuteCommandInput() interface{}
	ForceNewDeployment() interface{}
	SetForceNewDeployment(val interface{})
	ForceNewDeploymentInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckGracePeriodSeconds() *float64
	SetHealthCheckGracePeriodSeconds(val *float64)
	HealthCheckGracePeriodSecondsInput() *float64
	IamRole() *string
	SetIamRole(val *string)
	IamRoleInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LaunchType() *string
	SetLaunchType(val *string)
	LaunchTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() EcsServiceLoadBalancerList
	LoadBalancerInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfiguration() EcsServiceNetworkConfigurationOutputReference
	NetworkConfigurationInput() *EcsServiceNetworkConfiguration
	// The tree node.
	Node() constructs.Node
	OrderedPlacementStrategy() EcsServiceOrderedPlacementStrategyList
	OrderedPlacementStrategyInput() interface{}
	PlacementConstraints() EcsServicePlacementConstraintsList
	PlacementConstraintsInput() interface{}
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	PlatformVersionInput() *string
	PropagateTags() *string
	SetPropagateTags(val *string)
	PropagateTagsInput() *string
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
	SchedulingStrategy() *string
	SetSchedulingStrategy(val *string)
	SchedulingStrategyInput() *string
	ServiceConnectConfiguration() EcsServiceServiceConnectConfigurationOutputReference
	ServiceConnectConfigurationInput() *EcsServiceServiceConnectConfiguration
	ServiceRegistries() EcsServiceServiceRegistriesOutputReference
	ServiceRegistriesInput() *EcsServiceServiceRegistries
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TaskDefinition() *string
	SetTaskDefinition(val *string)
	TaskDefinitionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EcsServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Triggers() *map[string]*string
	SetTriggers(val *map[string]*string)
	TriggersInput() *map[string]*string
	WaitForSteadyState() interface{}
	SetWaitForSteadyState(val interface{})
	WaitForSteadyStateInput() interface{}
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
	PutAlarms(value *EcsServiceAlarms)
	PutCapacityProviderStrategy(value interface{})
	PutDeploymentCircuitBreaker(value *EcsServiceDeploymentCircuitBreaker)
	PutDeploymentController(value *EcsServiceDeploymentController)
	PutLoadBalancer(value interface{})
	PutNetworkConfiguration(value *EcsServiceNetworkConfiguration)
	PutOrderedPlacementStrategy(value interface{})
	PutPlacementConstraints(value interface{})
	PutServiceConnectConfiguration(value *EcsServiceServiceConnectConfiguration)
	PutServiceRegistries(value *EcsServiceServiceRegistries)
	PutTimeouts(value *EcsServiceTimeouts)
	ResetAlarms()
	ResetCapacityProviderStrategy()
	ResetCluster()
	ResetDeploymentCircuitBreaker()
	ResetDeploymentController()
	ResetDeploymentMaximumPercent()
	ResetDeploymentMinimumHealthyPercent()
	ResetDesiredCount()
	ResetEnableEcsManagedTags()
	ResetEnableExecuteCommand()
	ResetForceNewDeployment()
	ResetHealthCheckGracePeriodSeconds()
	ResetIamRole()
	ResetId()
	ResetLaunchType()
	ResetLoadBalancer()
	ResetNetworkConfiguration()
	ResetOrderedPlacementStrategy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementConstraints()
	ResetPlatformVersion()
	ResetPropagateTags()
	ResetSchedulingStrategy()
	ResetServiceConnectConfiguration()
	ResetServiceRegistries()
	ResetTags()
	ResetTagsAll()
	ResetTaskDefinition()
	ResetTimeouts()
	ResetTriggers()
	ResetWaitForSteadyState()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EcsService
type jsiiProxy_EcsService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EcsService) Alarms() EcsServiceAlarmsOutputReference {
	var returns EcsServiceAlarmsOutputReference
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) AlarmsInput() *EcsServiceAlarms {
	var returns *EcsServiceAlarms
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) CapacityProviderStrategy() EcsServiceCapacityProviderStrategyList {
	var returns EcsServiceCapacityProviderStrategyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentCircuitBreaker() EcsServiceDeploymentCircuitBreakerOutputReference {
	var returns EcsServiceDeploymentCircuitBreakerOutputReference
	_jsii_.Get(
		j,
		"deploymentCircuitBreaker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentCircuitBreakerInput() *EcsServiceDeploymentCircuitBreaker {
	var returns *EcsServiceDeploymentCircuitBreaker
	_jsii_.Get(
		j,
		"deploymentCircuitBreakerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentController() EcsServiceDeploymentControllerOutputReference {
	var returns EcsServiceDeploymentControllerOutputReference
	_jsii_.Get(
		j,
		"deploymentController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentControllerInput() *EcsServiceDeploymentController {
	var returns *EcsServiceDeploymentController
	_jsii_.Get(
		j,
		"deploymentControllerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentMaximumPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMaximumPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentMaximumPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMaximumPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentMinimumHealthyPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMinimumHealthyPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DeploymentMinimumHealthyPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMinimumHealthyPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) DesiredCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ForceNewDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNewDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ForceNewDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNewDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) HealthCheckGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) HealthCheckGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) IamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) IamRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) LoadBalancer() EcsServiceLoadBalancerList {
	var returns EcsServiceLoadBalancerList
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) LoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) NetworkConfiguration() EcsServiceNetworkConfigurationOutputReference {
	var returns EcsServiceNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) NetworkConfigurationInput() *EcsServiceNetworkConfiguration {
	var returns *EcsServiceNetworkConfiguration
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) OrderedPlacementStrategy() EcsServiceOrderedPlacementStrategyList {
	var returns EcsServiceOrderedPlacementStrategyList
	_jsii_.Get(
		j,
		"orderedPlacementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) OrderedPlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedPlacementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlacementConstraints() EcsServicePlacementConstraintsList {
	var returns EcsServicePlacementConstraintsList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) SchedulingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) SchedulingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceConnectConfiguration() EcsServiceServiceConnectConfigurationOutputReference {
	var returns EcsServiceServiceConnectConfigurationOutputReference
	_jsii_.Get(
		j,
		"serviceConnectConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceConnectConfigurationInput() *EcsServiceServiceConnectConfiguration {
	var returns *EcsServiceServiceConnectConfiguration
	_jsii_.Get(
		j,
		"serviceConnectConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceRegistries() EcsServiceServiceRegistriesOutputReference {
	var returns EcsServiceServiceRegistriesOutputReference
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) ServiceRegistriesInput() *EcsServiceServiceRegistries {
	var returns *EcsServiceServiceRegistries
	_jsii_.Get(
		j,
		"serviceRegistriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TaskDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Timeouts() EcsServiceTimeoutsOutputReference {
	var returns EcsServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) WaitForSteadyState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSteadyState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsService) WaitForSteadyStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSteadyStateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/ecs_service aws_ecs_service} Resource.
func NewEcsService(scope constructs.Construct, id *string, config *EcsServiceConfig) EcsService {
	_init_.Initialize()

	if err := validateNewEcsServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsService{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsService.EcsService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/ecs_service aws_ecs_service} Resource.
func NewEcsService_Override(e EcsService, scope constructs.Construct, id *string, config *EcsServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsService.EcsService",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EcsService)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetDeploymentMaximumPercent(val *float64) {
	if err := j.validateSetDeploymentMaximumPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMaximumPercent",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetDeploymentMinimumHealthyPercent(val *float64) {
	if err := j.validateSetDeploymentMinimumHealthyPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMinimumHealthyPercent",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetDesiredCount(val *float64) {
	if err := j.validateSetDesiredCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCount",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetForceNewDeployment(val interface{}) {
	if err := j.validateSetForceNewDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceNewDeployment",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetHealthCheckGracePeriodSeconds(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetIamRole(val *string) {
	if err := j.validateSetIamRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRole",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetSchedulingStrategy(val *string) {
	if err := j.validateSetSchedulingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingStrategy",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetTaskDefinition(val *string) {
	if err := j.validateSetTaskDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_EcsService)SetWaitForSteadyState(val interface{}) {
	if err := j.validateSetWaitForSteadyStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForSteadyState",
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
func EcsService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsService.EcsService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcsService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsService.EcsService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcsService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsService.EcsService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EcsService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ecsService.EcsService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EcsService) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EcsService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsService) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsService) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EcsService) PutAlarms(value *EcsServiceAlarms) {
	if err := e.validatePutAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAlarms",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutCapacityProviderStrategy(value interface{}) {
	if err := e.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutDeploymentCircuitBreaker(value *EcsServiceDeploymentCircuitBreaker) {
	if err := e.validatePutDeploymentCircuitBreakerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeploymentCircuitBreaker",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutDeploymentController(value *EcsServiceDeploymentController) {
	if err := e.validatePutDeploymentControllerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeploymentController",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutLoadBalancer(value interface{}) {
	if err := e.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutNetworkConfiguration(value *EcsServiceNetworkConfiguration) {
	if err := e.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutOrderedPlacementStrategy(value interface{}) {
	if err := e.validatePutOrderedPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOrderedPlacementStrategy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutPlacementConstraints(value interface{}) {
	if err := e.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutServiceConnectConfiguration(value *EcsServiceServiceConnectConfiguration) {
	if err := e.validatePutServiceConnectConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putServiceConnectConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutServiceRegistries(value *EcsServiceServiceRegistries) {
	if err := e.validatePutServiceRegistriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putServiceRegistries",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) PutTimeouts(value *EcsServiceTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsService) ResetAlarms() {
	_jsii_.InvokeVoid(
		e,
		"resetAlarms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetCluster() {
	_jsii_.InvokeVoid(
		e,
		"resetCluster",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetDeploymentCircuitBreaker() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentCircuitBreaker",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetDeploymentController() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentController",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetDeploymentMaximumPercent() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentMaximumPercent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetDeploymentMinimumHealthyPercent() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentMinimumHealthyPercent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetDesiredCount() {
	_jsii_.InvokeVoid(
		e,
		"resetDesiredCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetForceNewDeployment() {
	_jsii_.InvokeVoid(
		e,
		"resetForceNewDeployment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetHealthCheckGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckGracePeriodSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetIamRole() {
	_jsii_.InvokeVoid(
		e,
		"resetIamRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetLaunchType() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		e,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetOrderedPlacementStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetOrderedPlacementStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		e,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetSchedulingStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetSchedulingStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetServiceConnectConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceConnectConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetServiceRegistries() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceRegistries",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetTaskDefinition() {
	_jsii_.InvokeVoid(
		e,
		"resetTaskDefinition",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetTriggers() {
	_jsii_.InvokeVoid(
		e,
		"resetTriggers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) ResetWaitForSteadyState() {
	_jsii_.InvokeVoid(
		e,
		"resetWaitForSteadyState",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

