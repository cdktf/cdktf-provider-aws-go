// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtableexport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dynamodbtableexport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/dynamodb_table_export aws_dynamodb_table_export}.
type DynamodbTableExport interface {
	cdktf.TerraformResource
	Arn() *string
	BilledSizeInBytes() *float64
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndTime() *string
	ExportFormat() *string
	SetExportFormat(val *string)
	ExportFormatInput() *string
	ExportStatus() *string
	ExportTime() *string
	SetExportTime(val *string)
	ExportTimeInput() *string
	ExportType() *string
	SetExportType(val *string)
	ExportTypeInput() *string
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
	IncrementalExportSpecification() DynamodbTableExportIncrementalExportSpecificationOutputReference
	IncrementalExportSpecificationInput() *DynamodbTableExportIncrementalExportSpecification
	ItemCount() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManifestFilesS3Key() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3BucketOwner() *string
	SetS3BucketOwner(val *string)
	S3BucketOwnerInput() *string
	S3Prefix() *string
	SetS3Prefix(val *string)
	S3PrefixInput() *string
	S3SseAlgorithm() *string
	SetS3SseAlgorithm(val *string)
	S3SseAlgorithmInput() *string
	S3SseKmsKeyId() *string
	SetS3SseKmsKeyId(val *string)
	S3SseKmsKeyIdInput() *string
	StartTime() *string
	TableArn() *string
	SetTableArn(val *string)
	TableArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DynamodbTableExportTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutIncrementalExportSpecification(value *DynamodbTableExportIncrementalExportSpecification)
	PutTimeouts(value *DynamodbTableExportTimeouts)
	ResetExportFormat()
	ResetExportTime()
	ResetExportType()
	ResetId()
	ResetIncrementalExportSpecification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetS3BucketOwner()
	ResetS3Prefix()
	ResetS3SseAlgorithm()
	ResetS3SseKmsKeyId()
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

// The jsii proxy struct for DynamodbTableExport
type jsiiProxy_DynamodbTableExport struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DynamodbTableExport) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) BilledSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billedSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ExportFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ExportFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ExportStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ExportTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ExportTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ExportType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ExportTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) IncrementalExportSpecification() DynamodbTableExportIncrementalExportSpecificationOutputReference {
	var returns DynamodbTableExportIncrementalExportSpecificationOutputReference
	_jsii_.Get(
		j,
		"incrementalExportSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) IncrementalExportSpecificationInput() *DynamodbTableExportIncrementalExportSpecification {
	var returns *DynamodbTableExportIncrementalExportSpecification
	_jsii_.Get(
		j,
		"incrementalExportSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ItemCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"itemCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) ManifestFilesS3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestFilesS3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3BucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3BucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3SseAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SseAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3SseAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SseAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3SseKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SseKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) S3SseKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SseKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) TableArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) Timeouts() DynamodbTableExportTimeoutsOutputReference {
	var returns DynamodbTableExportTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableExport) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/dynamodb_table_export aws_dynamodb_table_export} Resource.
func NewDynamodbTableExport(scope constructs.Construct, id *string, config *DynamodbTableExportConfig) DynamodbTableExport {
	_init_.Initialize()

	if err := validateNewDynamodbTableExportParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbTableExport{}

	_jsii_.Create(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExport",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/dynamodb_table_export aws_dynamodb_table_export} Resource.
func NewDynamodbTableExport_Override(d DynamodbTableExport, scope constructs.Construct, id *string, config *DynamodbTableExportConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExport",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetExportFormat(val *string) {
	if err := j.validateSetExportFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportFormat",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetExportTime(val *string) {
	if err := j.validateSetExportTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportTime",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetExportType(val *string) {
	if err := j.validateSetExportTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportType",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetS3BucketOwner(val *string) {
	if err := j.validateSetS3BucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketOwner",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetS3Prefix(val *string) {
	if err := j.validateSetS3PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Prefix",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetS3SseAlgorithm(val *string) {
	if err := j.validateSetS3SseAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3SseAlgorithm",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetS3SseKmsKeyId(val *string) {
	if err := j.validateSetS3SseKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3SseKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableExport)SetTableArn(val *string) {
	if err := j.validateSetTableArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableArn",
		val,
	)
}

// Generates CDKTF code for importing a DynamodbTableExport resource upon running "cdktf plan <stack-name>".
func DynamodbTableExport_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDynamodbTableExport_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExport",
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
func DynamodbTableExport_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTableExport_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExport",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbTableExport_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTableExport_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExport",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbTableExport_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTableExport_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExport",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DynamodbTableExport_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dynamodbTableExport.DynamodbTableExport",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamodbTableExport) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DynamodbTableExport) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DynamodbTableExport) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbTableExport) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbTableExport) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbTableExport) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbTableExport) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbTableExport) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbTableExport) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbTableExport) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbTableExport) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbTableExport) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableExport) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DynamodbTableExport) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbTableExport) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbTableExport) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DynamodbTableExport) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbTableExport) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DynamodbTableExport) PutIncrementalExportSpecification(value *DynamodbTableExportIncrementalExportSpecification) {
	if err := d.validatePutIncrementalExportSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIncrementalExportSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableExport) PutTimeouts(value *DynamodbTableExportTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetExportFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetExportFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetExportTime() {
	_jsii_.InvokeVoid(
		d,
		"resetExportTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetExportType() {
	_jsii_.InvokeVoid(
		d,
		"resetExportType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetIncrementalExportSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetIncrementalExportSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetS3BucketOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetS3BucketOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetS3Prefix() {
	_jsii_.InvokeVoid(
		d,
		"resetS3Prefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetS3SseAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetS3SseAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetS3SseKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetS3SseKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableExport) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableExport) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableExport) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableExport) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableExport) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableExport) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

