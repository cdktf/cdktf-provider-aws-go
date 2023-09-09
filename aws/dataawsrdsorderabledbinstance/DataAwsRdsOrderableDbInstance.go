// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsrdsorderabledbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/dataawsrdsorderabledbinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/rds_orderable_db_instance aws_rds_orderable_db_instance}.
type DataAwsRdsOrderableDbInstance interface {
	cdktf.TerraformDataSource
	AvailabilityZoneGroup() *string
	SetAvailabilityZoneGroup(val *string)
	AvailabilityZoneGroupInput() *string
	AvailabilityZones() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
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
	InstanceClass() *string
	SetInstanceClass(val *string)
	InstanceClassInput() *string
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxIopsPerDbInstance() *float64
	MaxIopsPerGib() *float64
	MaxStorageSize() *float64
	MinIopsPerDbInstance() *float64
	MinIopsPerGib() *float64
	MinStorageSize() *float64
	MultiAzCapable() cdktf.IResolvable
	// The tree node.
	Node() constructs.Node
	OutpostCapable() cdktf.IResolvable
	PreferredEngineVersions() *[]*string
	SetPreferredEngineVersions(val *[]*string)
	PreferredEngineVersionsInput() *[]*string
	PreferredInstanceClasses() *[]*string
	SetPreferredInstanceClasses(val *[]*string)
	PreferredInstanceClassesInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReadReplicaCapable() cdktf.IResolvable
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	SupportedEngineModes() *[]*string
	SupportedNetworkTypes() *[]*string
	SupportsEnhancedMonitoring() interface{}
	SetSupportsEnhancedMonitoring(val interface{})
	SupportsEnhancedMonitoringInput() interface{}
	SupportsGlobalDatabases() interface{}
	SetSupportsGlobalDatabases(val interface{})
	SupportsGlobalDatabasesInput() interface{}
	SupportsIamDatabaseAuthentication() interface{}
	SetSupportsIamDatabaseAuthentication(val interface{})
	SupportsIamDatabaseAuthenticationInput() interface{}
	SupportsIops() interface{}
	SetSupportsIops(val interface{})
	SupportsIopsInput() interface{}
	SupportsKerberosAuthentication() interface{}
	SetSupportsKerberosAuthentication(val interface{})
	SupportsKerberosAuthenticationInput() interface{}
	SupportsPerformanceInsights() interface{}
	SetSupportsPerformanceInsights(val interface{})
	SupportsPerformanceInsightsInput() interface{}
	SupportsStorageAutoscaling() interface{}
	SetSupportsStorageAutoscaling(val interface{})
	SupportsStorageAutoscalingInput() interface{}
	SupportsStorageEncryption() interface{}
	SetSupportsStorageEncryption(val interface{})
	SupportsStorageEncryptionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Vpc() interface{}
	SetVpc(val interface{})
	VpcInput() interface{}
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
	ResetAvailabilityZoneGroup()
	ResetEngineVersion()
	ResetId()
	ResetInstanceClass()
	ResetLicenseModel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreferredEngineVersions()
	ResetPreferredInstanceClasses()
	ResetStorageType()
	ResetSupportsEnhancedMonitoring()
	ResetSupportsGlobalDatabases()
	ResetSupportsIamDatabaseAuthentication()
	ResetSupportsIops()
	ResetSupportsKerberosAuthentication()
	ResetSupportsPerformanceInsights()
	ResetSupportsStorageAutoscaling()
	ResetSupportsStorageEncryption()
	ResetVpc()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsRdsOrderableDbInstance
type jsiiProxy_DataAwsRdsOrderableDbInstance struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) AvailabilityZoneGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) AvailabilityZoneGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) MaxIopsPerDbInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIopsPerDbInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) MaxIopsPerGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIopsPerGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) MaxStorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) MinIopsPerDbInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIopsPerDbInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) MinIopsPerGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIopsPerGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) MinStorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) MultiAzCapable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"multiAzCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) OutpostCapable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"outpostCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) PreferredEngineVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredEngineVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) PreferredEngineVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredEngineVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) PreferredInstanceClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredInstanceClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) PreferredInstanceClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredInstanceClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) ReadReplicaCapable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"readReplicaCapable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportedEngineModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedEngineModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportedNetworkTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedNetworkTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsEnhancedMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsEnhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsEnhancedMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsEnhancedMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsGlobalDatabases() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsGlobalDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsGlobalDatabasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsGlobalDatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsIamDatabaseAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsIamDatabaseAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsIops() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsIopsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsKerberosAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsKerberosAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsKerberosAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsKerberosAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsPerformanceInsights() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsPerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsPerformanceInsightsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsPerformanceInsightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsStorageAutoscaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsStorageAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsStorageAutoscalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsStorageAutoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsStorageEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsStorageEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) SupportsStorageEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportsStorageEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) Vpc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance) VpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/rds_orderable_db_instance aws_rds_orderable_db_instance} Data Source.
func NewDataAwsRdsOrderableDbInstance(scope constructs.Construct, id *string, config *DataAwsRdsOrderableDbInstanceConfig) DataAwsRdsOrderableDbInstance {
	_init_.Initialize()

	if err := validateNewDataAwsRdsOrderableDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRdsOrderableDbInstance{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRdsOrderableDbInstance.DataAwsRdsOrderableDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/rds_orderable_db_instance aws_rds_orderable_db_instance} Data Source.
func NewDataAwsRdsOrderableDbInstance_Override(d DataAwsRdsOrderableDbInstance, scope constructs.Construct, id *string, config *DataAwsRdsOrderableDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRdsOrderableDbInstance.DataAwsRdsOrderableDbInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetAvailabilityZoneGroup(val *string) {
	if err := j.validateSetAvailabilityZoneGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneGroup",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetPreferredEngineVersions(val *[]*string) {
	if err := j.validateSetPreferredEngineVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredEngineVersions",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetPreferredInstanceClasses(val *[]*string) {
	if err := j.validateSetPreferredInstanceClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredInstanceClasses",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetSupportsEnhancedMonitoring(val interface{}) {
	if err := j.validateSetSupportsEnhancedMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsEnhancedMonitoring",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetSupportsGlobalDatabases(val interface{}) {
	if err := j.validateSetSupportsGlobalDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsGlobalDatabases",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetSupportsIamDatabaseAuthentication(val interface{}) {
	if err := j.validateSetSupportsIamDatabaseAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetSupportsIops(val interface{}) {
	if err := j.validateSetSupportsIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsIops",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetSupportsKerberosAuthentication(val interface{}) {
	if err := j.validateSetSupportsKerberosAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsKerberosAuthentication",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetSupportsPerformanceInsights(val interface{}) {
	if err := j.validateSetSupportsPerformanceInsightsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsPerformanceInsights",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetSupportsStorageAutoscaling(val interface{}) {
	if err := j.validateSetSupportsStorageAutoscalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsStorageAutoscaling",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetSupportsStorageEncryption(val interface{}) {
	if err := j.validateSetSupportsStorageEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportsStorageEncryption",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsOrderableDbInstance)SetVpc(val interface{}) {
	if err := j.validateSetVpcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpc",
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
func DataAwsRdsOrderableDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRdsOrderableDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRdsOrderableDbInstance.DataAwsRdsOrderableDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRdsOrderableDbInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRdsOrderableDbInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRdsOrderableDbInstance.DataAwsRdsOrderableDbInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRdsOrderableDbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRdsOrderableDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRdsOrderableDbInstance.DataAwsRdsOrderableDbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsRdsOrderableDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsRdsOrderableDbInstance.DataAwsRdsOrderableDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetAvailabilityZoneGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailabilityZoneGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetInstanceClass() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		d,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetPreferredEngineVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredEngineVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetPreferredInstanceClasses() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredInstanceClasses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetStorageType() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetSupportsEnhancedMonitoring() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsEnhancedMonitoring",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetSupportsGlobalDatabases() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsGlobalDatabases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetSupportsIamDatabaseAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsIamDatabaseAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetSupportsIops() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsIops",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetSupportsKerberosAuthentication() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsKerberosAuthentication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetSupportsPerformanceInsights() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsPerformanceInsights",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetSupportsStorageAutoscaling() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsStorageAutoscaling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetSupportsStorageEncryption() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportsStorageEncryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ResetVpc() {
	_jsii_.InvokeVoid(
		d,
		"resetVpc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsOrderableDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

