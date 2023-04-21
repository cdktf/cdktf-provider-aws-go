package ssmpatchbaseline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/ssmpatchbaseline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ssm_patch_baseline aws_ssm_patch_baseline}.
type SsmPatchBaseline interface {
	cdktf.TerraformResource
	ApprovalRule() SsmPatchBaselineApprovalRuleList
	ApprovalRuleInput() interface{}
	ApprovedPatches() *[]*string
	SetApprovedPatches(val *[]*string)
	ApprovedPatchesComplianceLevel() *string
	SetApprovedPatchesComplianceLevel(val *string)
	ApprovedPatchesComplianceLevelInput() *string
	ApprovedPatchesEnableNonSecurity() interface{}
	SetApprovedPatchesEnableNonSecurity(val interface{})
	ApprovedPatchesEnableNonSecurityInput() interface{}
	ApprovedPatchesInput() *[]*string
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
	GlobalFilter() SsmPatchBaselineGlobalFilterList
	GlobalFilterInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	OperatingSystemInput() *string
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
	RejectedPatches() *[]*string
	SetRejectedPatches(val *[]*string)
	RejectedPatchesAction() *string
	SetRejectedPatchesAction(val *string)
	RejectedPatchesActionInput() *string
	RejectedPatchesInput() *[]*string
	Source() SsmPatchBaselineSourceList
	SourceInput() interface{}
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
	PutApprovalRule(value interface{})
	PutGlobalFilter(value interface{})
	PutSource(value interface{})
	ResetApprovalRule()
	ResetApprovedPatches()
	ResetApprovedPatchesComplianceLevel()
	ResetApprovedPatchesEnableNonSecurity()
	ResetDescription()
	ResetGlobalFilter()
	ResetId()
	ResetOperatingSystem()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRejectedPatches()
	ResetRejectedPatchesAction()
	ResetSource()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SsmPatchBaseline
type jsiiProxy_SsmPatchBaseline struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovalRule() SsmPatchBaselineApprovalRuleList {
	var returns SsmPatchBaselineApprovalRuleList
	_jsii_.Get(
		j,
		"approvalRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovalRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approvedPatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesComplianceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvedPatchesComplianceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesComplianceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvedPatchesComplianceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesEnableNonSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvedPatchesEnableNonSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesEnableNonSecurityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvedPatchesEnableNonSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ApprovedPatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approvedPatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) GlobalFilter() SsmPatchBaselineGlobalFilterList {
	var returns SsmPatchBaselineGlobalFilterList
	_jsii_.Get(
		j,
		"globalFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) GlobalFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) OperatingSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RejectedPatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rejectedPatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RejectedPatchesAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rejectedPatchesAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RejectedPatchesActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rejectedPatchesActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) RejectedPatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rejectedPatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Source() SsmPatchBaselineSourceList {
	var returns SsmPatchBaselineSourceList
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaseline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ssm_patch_baseline aws_ssm_patch_baseline} Resource.
func NewSsmPatchBaseline(scope constructs.Construct, id *string, config *SsmPatchBaselineConfig) SsmPatchBaseline {
	_init_.Initialize()

	if err := validateNewSsmPatchBaselineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmPatchBaseline{}

	_jsii_.Create(
		"@cdktf/provider-aws.ssmPatchBaseline.SsmPatchBaseline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ssm_patch_baseline aws_ssm_patch_baseline} Resource.
func NewSsmPatchBaseline_Override(s SsmPatchBaseline, scope constructs.Construct, id *string, config *SsmPatchBaselineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ssmPatchBaseline.SsmPatchBaseline",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetApprovedPatches(val *[]*string) {
	if err := j.validateSetApprovedPatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvedPatches",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetApprovedPatchesComplianceLevel(val *string) {
	if err := j.validateSetApprovedPatchesComplianceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvedPatchesComplianceLevel",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetApprovedPatchesEnableNonSecurity(val interface{}) {
	if err := j.validateSetApprovedPatchesEnableNonSecurityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvedPatchesEnableNonSecurity",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetOperatingSystem(val *string) {
	if err := j.validateSetOperatingSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetRejectedPatches(val *[]*string) {
	if err := j.validateSetRejectedPatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rejectedPatches",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetRejectedPatchesAction(val *string) {
	if err := j.validateSetRejectedPatchesActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rejectedPatchesAction",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaseline)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
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
func SsmPatchBaseline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSsmPatchBaseline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ssmPatchBaseline.SsmPatchBaseline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SsmPatchBaseline_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSsmPatchBaseline_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ssmPatchBaseline.SsmPatchBaseline",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SsmPatchBaseline_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSsmPatchBaseline_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ssmPatchBaseline.SsmPatchBaseline",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SsmPatchBaseline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ssmPatchBaseline.SsmPatchBaseline",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SsmPatchBaseline) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SsmPatchBaseline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SsmPatchBaseline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmPatchBaseline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmPatchBaseline) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmPatchBaseline) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmPatchBaseline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmPatchBaseline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmPatchBaseline) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmPatchBaseline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmPatchBaseline) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmPatchBaseline) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SsmPatchBaseline) PutApprovalRule(value interface{}) {
	if err := s.validatePutApprovalRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putApprovalRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmPatchBaseline) PutGlobalFilter(value interface{}) {
	if err := s.validatePutGlobalFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGlobalFilter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmPatchBaseline) PutSource(value interface{}) {
	if err := s.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetApprovalRule() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovalRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetApprovedPatches() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovedPatches",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetApprovedPatchesComplianceLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovedPatchesComplianceLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetApprovedPatchesEnableNonSecurity() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovedPatchesEnableNonSecurity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetGlobalFilter() {
	_jsii_.InvokeVoid(
		s,
		"resetGlobalFilter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetOperatingSystem() {
	_jsii_.InvokeVoid(
		s,
		"resetOperatingSystem",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetRejectedPatches() {
	_jsii_.InvokeVoid(
		s,
		"resetRejectedPatches",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetRejectedPatchesAction() {
	_jsii_.InvokeVoid(
		s,
		"resetRejectedPatchesAction",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetSource() {
	_jsii_.InvokeVoid(
		s,
		"resetSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaseline) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmPatchBaseline) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmPatchBaseline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmPatchBaseline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

