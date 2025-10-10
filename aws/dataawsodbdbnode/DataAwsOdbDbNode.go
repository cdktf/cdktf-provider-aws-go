// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsodbdbnode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsodbdbnode/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/odb_db_node aws_odb_db_node}.
type DataAwsOdbDbNode interface {
	cdktf.TerraformDataSource
	AdditionalDetails() *string
	Arn() *string
	BackupIpId() *string
	BackupVnic2Id() *string
	BackupVnicId() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudVmClusterId() *string
	SetCloudVmClusterId(val *string)
	CloudVmClusterIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCoreCount() *float64
	CreatedAt() *string
	DbServerId() *string
	DbStorageSizeInGbs() *float64
	DbSystemId() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FaultDomain() *string
	FloatingIpAddress() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostIpId() *string
	Hostname() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceType() *string
	MemorySizeInGbs() *float64
	// The tree node.
	Node() constructs.Node
	Ocid() *string
	OciResourceAnchorName() *string
	PrivateIpAddress() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SoftwareStorageSizeInGbs() *float64
	Status() *string
	StatusReason() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeMaintenanceWindowEnd() *string
	TimeMaintenanceWindowStart() *string
	TotalCpuCoreCount() *float64
	Vnic2Id() *string
	VnicId() *string
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsOdbDbNode
type jsiiProxy_DataAwsOdbDbNode struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsOdbDbNode) AdditionalDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) BackupIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) BackupVnic2Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVnic2Id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) BackupVnicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVnicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) CloudVmClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudVmClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) CloudVmClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudVmClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) DbServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) DbStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) FaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) FloatingIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floatingIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) HostIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) MaintenanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) SoftwareStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softwareStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) TimeMaintenanceWindowEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceWindowEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) TimeMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) TotalCpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalCpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) Vnic2Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnic2Id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbNode) VnicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnicId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/odb_db_node aws_odb_db_node} Data Source.
func NewDataAwsOdbDbNode(scope constructs.Construct, id *string, config *DataAwsOdbDbNodeConfig) DataAwsOdbDbNode {
	_init_.Initialize()

	if err := validateNewDataAwsOdbDbNodeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsOdbDbNode{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsOdbDbNode.DataAwsOdbDbNode",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/odb_db_node aws_odb_db_node} Data Source.
func NewDataAwsOdbDbNode_Override(d DataAwsOdbDbNode, scope constructs.Construct, id *string, config *DataAwsOdbDbNodeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsOdbDbNode.DataAwsOdbDbNode",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOdbDbNode)SetCloudVmClusterId(val *string) {
	if err := j.validateSetCloudVmClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudVmClusterId",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbNode)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbNode)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbNode)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbNode)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbNode)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbNode)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbNode)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsOdbDbNode resource upon running "cdktf plan <stack-name>".
func DataAwsOdbDbNode_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsOdbDbNode_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsOdbDbNode.DataAwsOdbDbNode",
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
func DataAwsOdbDbNode_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOdbDbNode_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsOdbDbNode.DataAwsOdbDbNode",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOdbDbNode_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOdbDbNode_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsOdbDbNode.DataAwsOdbDbNode",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsOdbDbNode_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsOdbDbNode_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsOdbDbNode.DataAwsOdbDbNode",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOdbDbNode_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsOdbDbNode.DataAwsOdbDbNode",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsOdbDbNode) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOdbDbNode) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOdbDbNode) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbNode) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

