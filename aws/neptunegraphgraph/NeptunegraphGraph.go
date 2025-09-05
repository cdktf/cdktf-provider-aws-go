// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package neptunegraphgraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/neptunegraphgraph/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/neptunegraph_graph aws_neptunegraph_graph}.
type NeptunegraphGraph interface {
	cdktf.TerraformResource
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
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GraphName() *string
	SetGraphName(val *string)
	GraphNameInput() *string
	GraphNamePrefix() *string
	SetGraphNamePrefix(val *string)
	GraphNamePrefixInput() *string
	Id() *string
	KmsKeyIdentifier() *string
	SetKmsKeyIdentifier(val *string)
	KmsKeyIdentifierInput() *string
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
	ProvisionedMemory() *float64
	SetProvisionedMemory(val *float64)
	ProvisionedMemoryInput() *float64
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicConnectivity() interface{}
	SetPublicConnectivity(val interface{})
	PublicConnectivityInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReplicaCount() *float64
	SetReplicaCount(val *float64)
	ReplicaCountInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NeptunegraphGraphTimeoutsOutputReference
	TimeoutsInput() interface{}
	VectorSearchConfiguration() NeptunegraphGraphVectorSearchConfigurationList
	VectorSearchConfigurationInput() interface{}
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
	PutTimeouts(value *NeptunegraphGraphTimeouts)
	PutVectorSearchConfiguration(value interface{})
	ResetDeletionProtection()
	ResetGraphName()
	ResetGraphNamePrefix()
	ResetKmsKeyIdentifier()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicConnectivity()
	ResetRegion()
	ResetReplicaCount()
	ResetTags()
	ResetTimeouts()
	ResetVectorSearchConfiguration()
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

// The jsii proxy struct for NeptunegraphGraph
type jsiiProxy_NeptunegraphGraph struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NeptunegraphGraph) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) GraphName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) GraphNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) GraphNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) GraphNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) KmsKeyIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) KmsKeyIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) ProvisionedMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) ProvisionedMemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) PublicConnectivity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicConnectivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) PublicConnectivityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicConnectivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) Timeouts() NeptunegraphGraphTimeoutsOutputReference {
	var returns NeptunegraphGraphTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) VectorSearchConfiguration() NeptunegraphGraphVectorSearchConfigurationList {
	var returns NeptunegraphGraphVectorSearchConfigurationList
	_jsii_.Get(
		j,
		"vectorSearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NeptunegraphGraph) VectorSearchConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vectorSearchConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/neptunegraph_graph aws_neptunegraph_graph} Resource.
func NewNeptunegraphGraph(scope constructs.Construct, id *string, config *NeptunegraphGraphConfig) NeptunegraphGraph {
	_init_.Initialize()

	if err := validateNewNeptunegraphGraphParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NeptunegraphGraph{}

	_jsii_.Create(
		"@cdktf/provider-aws.neptunegraphGraph.NeptunegraphGraph",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/neptunegraph_graph aws_neptunegraph_graph} Resource.
func NewNeptunegraphGraph_Override(n NeptunegraphGraph, scope constructs.Construct, id *string, config *NeptunegraphGraphConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.neptunegraphGraph.NeptunegraphGraph",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetGraphName(val *string) {
	if err := j.validateSetGraphNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphName",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetGraphNamePrefix(val *string) {
	if err := j.validateSetGraphNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphNamePrefix",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetKmsKeyIdentifier(val *string) {
	if err := j.validateSetKmsKeyIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyIdentifier",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetProvisionedMemory(val *float64) {
	if err := j.validateSetProvisionedMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedMemory",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetPublicConnectivity(val interface{}) {
	if err := j.validateSetPublicConnectivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicConnectivity",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_NeptunegraphGraph)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a NeptunegraphGraph resource upon running "cdktf plan <stack-name>".
func NeptunegraphGraph_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNeptunegraphGraph_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.neptunegraphGraph.NeptunegraphGraph",
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
func NeptunegraphGraph_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNeptunegraphGraph_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.neptunegraphGraph.NeptunegraphGraph",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NeptunegraphGraph_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNeptunegraphGraph_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.neptunegraphGraph.NeptunegraphGraph",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NeptunegraphGraph_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNeptunegraphGraph_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.neptunegraphGraph.NeptunegraphGraph",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NeptunegraphGraph_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.neptunegraphGraph.NeptunegraphGraph",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) PutTimeouts(value *NeptunegraphGraphTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) PutVectorSearchConfiguration(value interface{}) {
	if err := n.validatePutVectorSearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putVectorSearchConfiguration",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		n,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetGraphName() {
	_jsii_.InvokeVoid(
		n,
		"resetGraphName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetGraphNamePrefix() {
	_jsii_.InvokeVoid(
		n,
		"resetGraphNamePrefix",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetKmsKeyIdentifier() {
	_jsii_.InvokeVoid(
		n,
		"resetKmsKeyIdentifier",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetPublicConnectivity() {
	_jsii_.InvokeVoid(
		n,
		"resetPublicConnectivity",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetRegion() {
	_jsii_.InvokeVoid(
		n,
		"resetRegion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		n,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) ResetVectorSearchConfiguration() {
	_jsii_.InvokeVoid(
		n,
		"resetVectorSearchConfiguration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NeptunegraphGraph) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NeptunegraphGraph) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

