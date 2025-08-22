// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2classificationjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/macie2classificationjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/macie2_classification_job aws_macie2_classification_job}.
type Macie2ClassificationJob interface {
	cdktf.TerraformResource
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
	CreatedAt() *string
	CustomDataIdentifierIds() *[]*string
	SetCustomDataIdentifierIds(val *[]*string)
	CustomDataIdentifierIdsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	InitialRun() interface{}
	SetInitialRun(val interface{})
	InitialRunInput() interface{}
	JobArn() *string
	JobId() *string
	JobStatus() *string
	SetJobStatus(val *string)
	JobStatusInput() *string
	JobType() *string
	SetJobType(val *string)
	JobTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
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
	S3JobDefinition() Macie2ClassificationJobS3JobDefinitionOutputReference
	S3JobDefinitionInput() *Macie2ClassificationJobS3JobDefinition
	SamplingPercentage() *float64
	SetSamplingPercentage(val *float64)
	SamplingPercentageInput() *float64
	ScheduleFrequency() Macie2ClassificationJobScheduleFrequencyOutputReference
	ScheduleFrequencyInput() *Macie2ClassificationJobScheduleFrequency
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
	Timeouts() Macie2ClassificationJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserPausedDetails() Macie2ClassificationJobUserPausedDetailsList
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
	PutS3JobDefinition(value *Macie2ClassificationJobS3JobDefinition)
	PutScheduleFrequency(value *Macie2ClassificationJobScheduleFrequency)
	PutTimeouts(value *Macie2ClassificationJobTimeouts)
	ResetCustomDataIdentifierIds()
	ResetDescription()
	ResetId()
	ResetInitialRun()
	ResetJobStatus()
	ResetName()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSamplingPercentage()
	ResetScheduleFrequency()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for Macie2ClassificationJob
type jsiiProxy_Macie2ClassificationJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Macie2ClassificationJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) CustomDataIdentifierIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDataIdentifierIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) CustomDataIdentifierIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customDataIdentifierIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) InitialRun() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) InitialRunInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) JobTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) S3JobDefinition() Macie2ClassificationJobS3JobDefinitionOutputReference {
	var returns Macie2ClassificationJobS3JobDefinitionOutputReference
	_jsii_.Get(
		j,
		"s3JobDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) S3JobDefinitionInput() *Macie2ClassificationJobS3JobDefinition {
	var returns *Macie2ClassificationJobS3JobDefinition
	_jsii_.Get(
		j,
		"s3JobDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) SamplingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) SamplingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) ScheduleFrequency() Macie2ClassificationJobScheduleFrequencyOutputReference {
	var returns Macie2ClassificationJobScheduleFrequencyOutputReference
	_jsii_.Get(
		j,
		"scheduleFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) ScheduleFrequencyInput() *Macie2ClassificationJobScheduleFrequency {
	var returns *Macie2ClassificationJobScheduleFrequency
	_jsii_.Get(
		j,
		"scheduleFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) Timeouts() Macie2ClassificationJobTimeoutsOutputReference {
	var returns Macie2ClassificationJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2ClassificationJob) UserPausedDetails() Macie2ClassificationJobUserPausedDetailsList {
	var returns Macie2ClassificationJobUserPausedDetailsList
	_jsii_.Get(
		j,
		"userPausedDetails",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/macie2_classification_job aws_macie2_classification_job} Resource.
func NewMacie2ClassificationJob(scope constructs.Construct, id *string, config *Macie2ClassificationJobConfig) Macie2ClassificationJob {
	_init_.Initialize()

	if err := validateNewMacie2ClassificationJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Macie2ClassificationJob{}

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/macie2_classification_job aws_macie2_classification_job} Resource.
func NewMacie2ClassificationJob_Override(m Macie2ClassificationJob, scope constructs.Construct, id *string, config *Macie2ClassificationJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJob",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetCustomDataIdentifierIds(val *[]*string) {
	if err := j.validateSetCustomDataIdentifierIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDataIdentifierIds",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetInitialRun(val interface{}) {
	if err := j.validateSetInitialRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialRun",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetJobStatus(val *string) {
	if err := j.validateSetJobStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobStatus",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetJobType(val *string) {
	if err := j.validateSetJobTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobType",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetSamplingPercentage(val *float64) {
	if err := j.validateSetSamplingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingPercentage",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Macie2ClassificationJob)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a Macie2ClassificationJob resource upon running "cdktf plan <stack-name>".
func Macie2ClassificationJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMacie2ClassificationJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJob",
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
func Macie2ClassificationJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMacie2ClassificationJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Macie2ClassificationJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMacie2ClassificationJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Macie2ClassificationJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMacie2ClassificationJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Macie2ClassificationJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.macie2ClassificationJob.Macie2ClassificationJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) PutS3JobDefinition(value *Macie2ClassificationJobS3JobDefinition) {
	if err := m.validatePutS3JobDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putS3JobDefinition",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) PutScheduleFrequency(value *Macie2ClassificationJobScheduleFrequency) {
	if err := m.validatePutScheduleFrequencyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScheduleFrequency",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) PutTimeouts(value *Macie2ClassificationJobTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetCustomDataIdentifierIds() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomDataIdentifierIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetInitialRun() {
	_jsii_.InvokeVoid(
		m,
		"resetInitialRun",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetJobStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetJobStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetSamplingPercentage() {
	_jsii_.InvokeVoid(
		m,
		"resetSamplingPercentage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetScheduleFrequency() {
	_jsii_.InvokeVoid(
		m,
		"resetScheduleFrequency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2ClassificationJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2ClassificationJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

