// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudexadatainfrastructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/odbcloudexadatainfrastructure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_cloud_exadata_infrastructure aws_odb_cloud_exadata_infrastructure}.
type OdbCloudExadataInfrastructure interface {
	cdktf.TerraformResource
	ActivatedStorageCount() *float64
	AdditionalStorageCount() *float64
	Arn() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneId() *string
	SetAvailabilityZoneId(val *string)
	AvailabilityZoneIdInput() *string
	AvailabilityZoneInput() *string
	AvailableStorageSizeInGbs() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputeCount() *float64
	SetComputeCount(val *float64)
	ComputeCountInput() *float64
	ComputeModel() *string
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
	CpuCount() *float64
	CreatedAt() *string
	CustomerContactsToSendToOci() OdbCloudExadataInfrastructureCustomerContactsToSendToOciList
	CustomerContactsToSendToOciInput() interface{}
	DatabaseServerType() *string
	SetDatabaseServerType(val *string)
	DatabaseServerTypeInput() *string
	DataStorageSizeInTbs() *float64
	DbNodeStorageSizeInGbs() *float64
	DbServerVersion() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LastMaintenanceRunId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceWindow() OdbCloudExadataInfrastructureMaintenanceWindowList
	MaintenanceWindowInput() interface{}
	MaxCpuCount() *float64
	MaxDataStorageInTbs() *float64
	MaxDbNodeStorageSizeInGbs() *float64
	MaxMemoryInGbs() *float64
	MemorySizeInGbs() *float64
	MonthlyDbServerVersion() *string
	MonthlyStorageServerVersion() *string
	NextMaintenanceRunId() *string
	// The tree node.
	Node() constructs.Node
	Ocid() *string
	OciResourceAnchorName() *string
	OciUrl() *string
	PercentProgress() *float64
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
	Shape() *string
	SetShape(val *string)
	ShapeInput() *string
	Status() *string
	StatusReason() *string
	StorageCount() *float64
	SetStorageCount(val *float64)
	StorageCountInput() *float64
	StorageServerType() *string
	SetStorageServerType(val *string)
	StorageServerTypeInput() *string
	StorageServerVersion() *string
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
	Timeouts() OdbCloudExadataInfrastructureTimeoutsOutputReference
	TimeoutsInput() interface{}
	TotalStorageSizeInGbs() *float64
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
	PutCustomerContactsToSendToOci(value interface{})
	PutMaintenanceWindow(value interface{})
	PutTimeouts(value *OdbCloudExadataInfrastructureTimeouts)
	ResetAvailabilityZone()
	ResetComputeCount()
	ResetCustomerContactsToSendToOci()
	ResetDatabaseServerType()
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetStorageCount()
	ResetStorageServerType()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for OdbCloudExadataInfrastructure
type jsiiProxy_OdbCloudExadataInfrastructure struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) ActivatedStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activatedStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) AdditionalStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) AvailableStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) CpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) CustomerContactsToSendToOci() OdbCloudExadataInfrastructureCustomerContactsToSendToOciList {
	var returns OdbCloudExadataInfrastructureCustomerContactsToSendToOciList
	_jsii_.Get(
		j,
		"customerContactsToSendToOci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) CustomerContactsToSendToOciInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContactsToSendToOciInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) DatabaseServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseServerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) DatabaseServerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseServerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) DbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) DbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) LastMaintenanceRunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastMaintenanceRunId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MaintenanceWindow() OdbCloudExadataInfrastructureMaintenanceWindowList {
	var returns OdbCloudExadataInfrastructureMaintenanceWindowList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MaintenanceWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MaxCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MaxDataStorageInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataStorageInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MaxDbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MaxMemoryInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MonthlyDbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyDbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) MonthlyStorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyStorageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) NextMaintenanceRunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceRunId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) StorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) StorageCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) StorageServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) StorageServerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) StorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) Timeouts() OdbCloudExadataInfrastructureTimeoutsOutputReference {
	var returns OdbCloudExadataInfrastructureTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure) TotalStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalStorageSizeInGbs",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_cloud_exadata_infrastructure aws_odb_cloud_exadata_infrastructure} Resource.
func NewOdbCloudExadataInfrastructure(scope constructs.Construct, id *string, config *OdbCloudExadataInfrastructureConfig) OdbCloudExadataInfrastructure {
	_init_.Initialize()

	if err := validateNewOdbCloudExadataInfrastructureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OdbCloudExadataInfrastructure{}

	_jsii_.Create(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructure",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_cloud_exadata_infrastructure aws_odb_cloud_exadata_infrastructure} Resource.
func NewOdbCloudExadataInfrastructure_Override(o OdbCloudExadataInfrastructure, scope constructs.Construct, id *string, config *OdbCloudExadataInfrastructureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructure",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetDatabaseServerType(val *string) {
	if err := j.validateSetDatabaseServerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseServerType",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetStorageCount(val *float64) {
	if err := j.validateSetStorageCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCount",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetStorageServerType(val *string) {
	if err := j.validateSetStorageServerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageServerType",
		val,
	)
}

func (j *jsiiProxy_OdbCloudExadataInfrastructure)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a OdbCloudExadataInfrastructure resource upon running "cdktf plan <stack-name>".
func OdbCloudExadataInfrastructure_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOdbCloudExadataInfrastructure_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructure",
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
func OdbCloudExadataInfrastructure_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbCloudExadataInfrastructure_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructure",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OdbCloudExadataInfrastructure_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbCloudExadataInfrastructure_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructure",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OdbCloudExadataInfrastructure_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOdbCloudExadataInfrastructure_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructure",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OdbCloudExadataInfrastructure_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.odbCloudExadataInfrastructure.OdbCloudExadataInfrastructure",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OdbCloudExadataInfrastructure) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) PutCustomerContactsToSendToOci(value interface{}) {
	if err := o.validatePutCustomerContactsToSendToOciParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCustomerContactsToSendToOci",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) PutMaintenanceWindow(value interface{}) {
	if err := o.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) PutTimeouts(value *OdbCloudExadataInfrastructureTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		o,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetComputeCount() {
	_jsii_.InvokeVoid(
		o,
		"resetComputeCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetCustomerContactsToSendToOci() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomerContactsToSendToOci",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetDatabaseServerType() {
	_jsii_.InvokeVoid(
		o,
		"resetDatabaseServerType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		o,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetRegion() {
	_jsii_.InvokeVoid(
		o,
		"resetRegion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetStorageCount() {
	_jsii_.InvokeVoid(
		o,
		"resetStorageCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetStorageServerType() {
	_jsii_.InvokeVoid(
		o,
		"resetStorageServerType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OdbCloudExadataInfrastructure) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

