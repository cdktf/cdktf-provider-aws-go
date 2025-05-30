// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/ssmassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/ssm_association aws_ssm_association}.
type SsmAssociation interface {
	cdktf.TerraformResource
	ApplyOnlyAtCronInterval() interface{}
	SetApplyOnlyAtCronInterval(val interface{})
	ApplyOnlyAtCronIntervalInput() interface{}
	Arn() *string
	AssociationId() *string
	AssociationName() *string
	SetAssociationName(val *string)
	AssociationNameInput() *string
	AutomationTargetParameterName() *string
	SetAutomationTargetParameterName(val *string)
	AutomationTargetParameterNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComplianceSeverity() *string
	SetComplianceSeverity(val *string)
	ComplianceSeverityInput() *string
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
	DocumentVersion() *string
	SetDocumentVersion(val *string)
	DocumentVersionInput() *string
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
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxConcurrency() *string
	SetMaxConcurrency(val *string)
	MaxConcurrencyInput() *string
	MaxErrors() *string
	SetMaxErrors(val *string)
	MaxErrorsInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputLocation() SsmAssociationOutputLocationOutputReference
	OutputLocationInput() *SsmAssociationOutputLocation
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
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
	ScheduleExpression() *string
	SetScheduleExpression(val *string)
	ScheduleExpressionInput() *string
	SyncCompliance() *string
	SetSyncCompliance(val *string)
	SyncComplianceInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	Targets() SsmAssociationTargetsList
	TargetsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WaitForSuccessTimeoutSeconds() *float64
	SetWaitForSuccessTimeoutSeconds(val *float64)
	WaitForSuccessTimeoutSecondsInput() *float64
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
	PutOutputLocation(value *SsmAssociationOutputLocation)
	PutTargets(value interface{})
	ResetApplyOnlyAtCronInterval()
	ResetAssociationName()
	ResetAutomationTargetParameterName()
	ResetComplianceSeverity()
	ResetDocumentVersion()
	ResetId()
	ResetInstanceId()
	ResetMaxConcurrency()
	ResetMaxErrors()
	ResetOutputLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetScheduleExpression()
	ResetSyncCompliance()
	ResetTags()
	ResetTagsAll()
	ResetTargets()
	ResetWaitForSuccessTimeoutSeconds()
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

// The jsii proxy struct for SsmAssociation
type jsiiProxy_SsmAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmAssociation) ApplyOnlyAtCronInterval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyOnlyAtCronInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ApplyOnlyAtCronIntervalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyOnlyAtCronIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AssociationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AssociationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AutomationTargetParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationTargetParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) AutomationTargetParameterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationTargetParameterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ComplianceSeverity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ComplianceSeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceSeverityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) MaxConcurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) MaxConcurrencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) MaxErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) MaxErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) OutputLocation() SsmAssociationOutputLocationOutputReference {
	var returns SsmAssociationOutputLocationOutputReference
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) OutputLocationInput() *SsmAssociationOutputLocation {
	var returns *SsmAssociationOutputLocation
	_jsii_.Get(
		j,
		"outputLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) ScheduleExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) SyncCompliance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncCompliance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) SyncComplianceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncComplianceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) Targets() SsmAssociationTargetsList {
	var returns SsmAssociationTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) WaitForSuccessTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForSuccessTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmAssociation) WaitForSuccessTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForSuccessTimeoutSecondsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/ssm_association aws_ssm_association} Resource.
func NewSsmAssociation(scope constructs.Construct, id *string, config *SsmAssociationConfig) SsmAssociation {
	_init_.Initialize()

	if err := validateNewSsmAssociationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmAssociation{}

	_jsii_.Create(
		"@cdktf/provider-aws.ssmAssociation.SsmAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/ssm_association aws_ssm_association} Resource.
func NewSsmAssociation_Override(s SsmAssociation, scope constructs.Construct, id *string, config *SsmAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ssmAssociation.SsmAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmAssociation)SetApplyOnlyAtCronInterval(val interface{}) {
	if err := j.validateSetApplyOnlyAtCronIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyOnlyAtCronInterval",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetAssociationName(val *string) {
	if err := j.validateSetAssociationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associationName",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetAutomationTargetParameterName(val *string) {
	if err := j.validateSetAutomationTargetParameterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automationTargetParameterName",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetComplianceSeverity(val *string) {
	if err := j.validateSetComplianceSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceSeverity",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetDocumentVersion(val *string) {
	if err := j.validateSetDocumentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetMaxConcurrency(val *string) {
	if err := j.validateSetMaxConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetMaxErrors(val *string) {
	if err := j.validateSetMaxErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxErrors",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetScheduleExpression(val *string) {
	if err := j.validateSetScheduleExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleExpression",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetSyncCompliance(val *string) {
	if err := j.validateSetSyncComplianceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncCompliance",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_SsmAssociation)SetWaitForSuccessTimeoutSeconds(val *float64) {
	if err := j.validateSetWaitForSuccessTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForSuccessTimeoutSeconds",
		val,
	)
}

// Generates CDKTF code for importing a SsmAssociation resource upon running "cdktf plan <stack-name>".
func SsmAssociation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSsmAssociation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ssmAssociation.SsmAssociation",
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
func SsmAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSsmAssociation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ssmAssociation.SsmAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SsmAssociation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSsmAssociation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ssmAssociation.SsmAssociation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SsmAssociation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSsmAssociation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ssmAssociation.SsmAssociation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ssmAssociation.SsmAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SsmAssociation) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SsmAssociation) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SsmAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SsmAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SsmAssociation) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SsmAssociation) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SsmAssociation) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmAssociation) PutOutputLocation(value *SsmAssociationOutputLocation) {
	if err := s.validatePutOutputLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOutputLocation",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmAssociation) PutTargets(value interface{}) {
	if err := s.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTargets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmAssociation) ResetApplyOnlyAtCronInterval() {
	_jsii_.InvokeVoid(
		s,
		"resetApplyOnlyAtCronInterval",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetAssociationName() {
	_jsii_.InvokeVoid(
		s,
		"resetAssociationName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetAutomationTargetParameterName() {
	_jsii_.InvokeVoid(
		s,
		"resetAutomationTargetParameterName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetComplianceSeverity() {
	_jsii_.InvokeVoid(
		s,
		"resetComplianceSeverity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetInstanceId() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetMaxConcurrency() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetMaxErrors() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxErrors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetOutputLocation() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputLocation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetScheduleExpression() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduleExpression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetSyncCompliance() {
	_jsii_.InvokeVoid(
		s,
		"resetSyncCompliance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetTargets() {
	_jsii_.InvokeVoid(
		s,
		"resetTargets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) ResetWaitForSuccessTimeoutSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetWaitForSuccessTimeoutSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

