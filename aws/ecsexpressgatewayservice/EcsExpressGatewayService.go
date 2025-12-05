// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsexpressgatewayservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecsexpressgatewayservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_express_gateway_service aws_ecs_express_gateway_service}.
type EcsExpressGatewayService interface {
	cdktf.TerraformResource
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
	Cpu() *string
	SetCpu(val *string)
	CpuInput() *string
	CurrentDeployment() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	InfrastructureRoleArn() *string
	SetInfrastructureRoleArn(val *string)
	InfrastructureRoleArnInput() *string
	IngressPaths() EcsExpressGatewayServiceIngressPathsList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Memory() *string
	SetMemory(val *string)
	MemoryInput() *string
	NetworkConfiguration() EcsExpressGatewayServiceNetworkConfigurationList
	NetworkConfigurationInput() interface{}
	// The tree node.
	Node() constructs.Node
	PrimaryContainer() EcsExpressGatewayServicePrimaryContainerList
	PrimaryContainerInput() interface{}
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
	ScalingTarget() EcsExpressGatewayServiceScalingTargetList
	ScalingTargetInput() interface{}
	ServiceArn() *string
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	ServiceRevisionArn() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	TaskRoleArn() *string
	SetTaskRoleArn(val *string)
	TaskRoleArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EcsExpressGatewayServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
	WaitForSteadyState() interface{}
	SetWaitForSteadyState(val interface{})
	WaitForSteadyStateInput() interface{}
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
	PutNetworkConfiguration(value interface{})
	PutPrimaryContainer(value interface{})
	PutScalingTarget(value interface{})
	PutTimeouts(value *EcsExpressGatewayServiceTimeouts)
	ResetCluster()
	ResetCpu()
	ResetHealthCheckPath()
	ResetMemory()
	ResetNetworkConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryContainer()
	ResetRegion()
	ResetScalingTarget()
	ResetServiceName()
	ResetTags()
	ResetTaskRoleArn()
	ResetTimeouts()
	ResetWaitForSteadyState()
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

// The jsii proxy struct for EcsExpressGatewayService
type jsiiProxy_EcsExpressGatewayService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EcsExpressGatewayService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) CpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) CurrentDeployment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) InfrastructureRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) InfrastructureRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) IngressPaths() EcsExpressGatewayServiceIngressPathsList {
	var returns EcsExpressGatewayServiceIngressPathsList
	_jsii_.Get(
		j,
		"ingressPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) MemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) NetworkConfiguration() EcsExpressGatewayServiceNetworkConfigurationList {
	var returns EcsExpressGatewayServiceNetworkConfigurationList
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) PrimaryContainer() EcsExpressGatewayServicePrimaryContainerList {
	var returns EcsExpressGatewayServicePrimaryContainerList
	_jsii_.Get(
		j,
		"primaryContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) PrimaryContainerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ScalingTarget() EcsExpressGatewayServiceScalingTargetList {
	var returns EcsExpressGatewayServiceScalingTargetList
	_jsii_.Get(
		j,
		"scalingTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ScalingTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) ServiceRevisionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRevisionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) TaskRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) Timeouts() EcsExpressGatewayServiceTimeoutsOutputReference {
	var returns EcsExpressGatewayServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) WaitForSteadyState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSteadyState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayService) WaitForSteadyStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSteadyStateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_express_gateway_service aws_ecs_express_gateway_service} Resource.
func NewEcsExpressGatewayService(scope constructs.Construct, id *string, config *EcsExpressGatewayServiceConfig) EcsExpressGatewayService {
	_init_.Initialize()

	if err := validateNewEcsExpressGatewayServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsExpressGatewayService{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_express_gateway_service aws_ecs_express_gateway_service} Resource.
func NewEcsExpressGatewayService_Override(e EcsExpressGatewayService, scope constructs.Construct, id *string, config *EcsExpressGatewayServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayService",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetCpu(val *string) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetInfrastructureRoleArn(val *string) {
	if err := j.validateSetInfrastructureRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureRoleArn",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetMemory(val *string) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetServiceName(val *string) {
	if err := j.validateSetServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetTaskRoleArn(val *string) {
	if err := j.validateSetTaskRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayService)SetWaitForSteadyState(val interface{}) {
	if err := j.validateSetWaitForSteadyStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForSteadyState",
		val,
	)
}

// Generates CDKTF code for importing a EcsExpressGatewayService resource upon running "cdktf plan <stack-name>".
func EcsExpressGatewayService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEcsExpressGatewayService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayService",
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
func EcsExpressGatewayService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsExpressGatewayService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcsExpressGatewayService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsExpressGatewayService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcsExpressGatewayService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcsExpressGatewayService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EcsExpressGatewayService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EcsExpressGatewayService) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsExpressGatewayService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsExpressGatewayService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsExpressGatewayService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsExpressGatewayService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsExpressGatewayService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsExpressGatewayService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsExpressGatewayService) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsExpressGatewayService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsExpressGatewayService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsExpressGatewayService) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) PutNetworkConfiguration(value interface{}) {
	if err := e.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) PutPrimaryContainer(value interface{}) {
	if err := e.validatePutPrimaryContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPrimaryContainer",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) PutScalingTarget(value interface{}) {
	if err := e.validatePutScalingTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingTarget",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) PutTimeouts(value *EcsExpressGatewayServiceTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetCluster() {
	_jsii_.InvokeVoid(
		e,
		"resetCluster",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetCpu() {
	_jsii_.InvokeVoid(
		e,
		"resetCpu",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetMemory() {
	_jsii_.InvokeVoid(
		e,
		"resetMemory",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetPrimaryContainer() {
	_jsii_.InvokeVoid(
		e,
		"resetPrimaryContainer",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetScalingTarget() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingTarget",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetServiceName() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetTaskRoleArn() {
	_jsii_.InvokeVoid(
		e,
		"resetTaskRoleArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) ResetWaitForSteadyState() {
	_jsii_.InvokeVoid(
		e,
		"resetWaitForSteadyState",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

