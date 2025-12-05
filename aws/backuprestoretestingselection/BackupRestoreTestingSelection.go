// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuprestoretestingselection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/backuprestoretestingselection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/backup_restore_testing_selection aws_backup_restore_testing_selection}.
type BackupRestoreTestingSelection interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProtectedResourceArns() *[]*string
	SetProtectedResourceArns(val *[]*string)
	ProtectedResourceArnsInput() *[]*string
	ProtectedResourceConditions() BackupRestoreTestingSelectionProtectedResourceConditionsList
	ProtectedResourceConditionsInput() interface{}
	ProtectedResourceType() *string
	SetProtectedResourceType(val *string)
	ProtectedResourceTypeInput() *string
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
	RestoreMetadataOverrides() *map[string]*string
	SetRestoreMetadataOverrides(val *map[string]*string)
	RestoreMetadataOverridesInput() *map[string]*string
	RestoreTestingPlanName() *string
	SetRestoreTestingPlanName(val *string)
	RestoreTestingPlanNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidationWindowHours() *float64
	SetValidationWindowHours(val *float64)
	ValidationWindowHoursInput() *float64
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
	PutProtectedResourceConditions(value interface{})
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtectedResourceArns()
	ResetProtectedResourceConditions()
	ResetRegion()
	ResetRestoreMetadataOverrides()
	ResetValidationWindowHours()
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

// The jsii proxy struct for BackupRestoreTestingSelection
type jsiiProxy_BackupRestoreTestingSelection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BackupRestoreTestingSelection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ProtectedResourceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protectedResourceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ProtectedResourceArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protectedResourceArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ProtectedResourceConditions() BackupRestoreTestingSelectionProtectedResourceConditionsList {
	var returns BackupRestoreTestingSelectionProtectedResourceConditionsList
	_jsii_.Get(
		j,
		"protectedResourceConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ProtectedResourceConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedResourceConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ProtectedResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ProtectedResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) RestoreMetadataOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"restoreMetadataOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) RestoreMetadataOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"restoreMetadataOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) RestoreTestingPlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTestingPlanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) RestoreTestingPlanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTestingPlanNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ValidationWindowHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validationWindowHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupRestoreTestingSelection) ValidationWindowHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"validationWindowHoursInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/backup_restore_testing_selection aws_backup_restore_testing_selection} Resource.
func NewBackupRestoreTestingSelection(scope constructs.Construct, id *string, config *BackupRestoreTestingSelectionConfig) BackupRestoreTestingSelection {
	_init_.Initialize()

	if err := validateNewBackupRestoreTestingSelectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupRestoreTestingSelection{}

	_jsii_.Create(
		"@cdktf/provider-aws.backupRestoreTestingSelection.BackupRestoreTestingSelection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/backup_restore_testing_selection aws_backup_restore_testing_selection} Resource.
func NewBackupRestoreTestingSelection_Override(b BackupRestoreTestingSelection, scope constructs.Construct, id *string, config *BackupRestoreTestingSelectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.backupRestoreTestingSelection.BackupRestoreTestingSelection",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetIamRoleArn(val *string) {
	if err := j.validateSetIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetProtectedResourceArns(val *[]*string) {
	if err := j.validateSetProtectedResourceArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectedResourceArns",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetProtectedResourceType(val *string) {
	if err := j.validateSetProtectedResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectedResourceType",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetRestoreMetadataOverrides(val *map[string]*string) {
	if err := j.validateSetRestoreMetadataOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreMetadataOverrides",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetRestoreTestingPlanName(val *string) {
	if err := j.validateSetRestoreTestingPlanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreTestingPlanName",
		val,
	)
}

func (j *jsiiProxy_BackupRestoreTestingSelection)SetValidationWindowHours(val *float64) {
	if err := j.validateSetValidationWindowHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationWindowHours",
		val,
	)
}

// Generates CDKTF code for importing a BackupRestoreTestingSelection resource upon running "cdktf plan <stack-name>".
func BackupRestoreTestingSelection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBackupRestoreTestingSelection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.backupRestoreTestingSelection.BackupRestoreTestingSelection",
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
func BackupRestoreTestingSelection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupRestoreTestingSelection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.backupRestoreTestingSelection.BackupRestoreTestingSelection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BackupRestoreTestingSelection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupRestoreTestingSelection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.backupRestoreTestingSelection.BackupRestoreTestingSelection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BackupRestoreTestingSelection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupRestoreTestingSelection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.backupRestoreTestingSelection.BackupRestoreTestingSelection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BackupRestoreTestingSelection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.backupRestoreTestingSelection.BackupRestoreTestingSelection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) PutProtectedResourceConditions(value interface{}) {
	if err := b.validatePutProtectedResourceConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putProtectedResourceConditions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ResetProtectedResourceArns() {
	_jsii_.InvokeVoid(
		b,
		"resetProtectedResourceArns",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ResetProtectedResourceConditions() {
	_jsii_.InvokeVoid(
		b,
		"resetProtectedResourceConditions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ResetRegion() {
	_jsii_.InvokeVoid(
		b,
		"resetRegion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ResetRestoreMetadataOverrides() {
	_jsii_.InvokeVoid(
		b,
		"resetRestoreMetadataOverrides",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ResetValidationWindowHours() {
	_jsii_.InvokeVoid(
		b,
		"resetValidationWindowHours",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupRestoreTestingSelection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupRestoreTestingSelection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

