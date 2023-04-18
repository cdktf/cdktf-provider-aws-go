package ssmpatchbaseline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/ssmpatchbaseline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmPatchBaselineApprovalRuleOutputReference interface {
	cdktf.ComplexObject
	ApproveAfterDays() *float64
	SetApproveAfterDays(val *float64)
	ApproveAfterDaysInput() *float64
	ApproveUntilDate() *string
	SetApproveUntilDate(val *string)
	ApproveUntilDateInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComplianceLevel() *string
	SetComplianceLevel(val *string)
	ComplianceLevelInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableNonSecurity() interface{}
	SetEnableNonSecurity(val interface{})
	EnableNonSecurityInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PatchFilter() SsmPatchBaselineApprovalRulePatchFilterList
	PatchFilterInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutPatchFilter(value interface{})
	ResetApproveAfterDays()
	ResetApproveUntilDate()
	ResetComplianceLevel()
	ResetEnableNonSecurity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SsmPatchBaselineApprovalRuleOutputReference
type jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ApproveAfterDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approveAfterDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ApproveAfterDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approveAfterDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ApproveUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ApproveUntilDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveUntilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ComplianceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ComplianceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) EnableNonSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNonSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) EnableNonSecurityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNonSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) PatchFilter() SsmPatchBaselineApprovalRulePatchFilterList {
	var returns SsmPatchBaselineApprovalRulePatchFilterList
	_jsii_.Get(
		j,
		"patchFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) PatchFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"patchFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSsmPatchBaselineApprovalRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SsmPatchBaselineApprovalRuleOutputReference {
	_init_.Initialize()

	if err := validateNewSsmPatchBaselineApprovalRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ssmPatchBaseline.SsmPatchBaselineApprovalRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSsmPatchBaselineApprovalRuleOutputReference_Override(s SsmPatchBaselineApprovalRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ssmPatchBaseline.SsmPatchBaselineApprovalRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetApproveAfterDays(val *float64) {
	if err := j.validateSetApproveAfterDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approveAfterDays",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetApproveUntilDate(val *string) {
	if err := j.validateSetApproveUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approveUntilDate",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetComplianceLevel(val *string) {
	if err := j.validateSetComplianceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceLevel",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetEnableNonSecurity(val interface{}) {
	if err := j.validateSetEnableNonSecurityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNonSecurity",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) PutPatchFilter(value interface{}) {
	if err := s.validatePutPatchFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPatchFilter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ResetApproveAfterDays() {
	_jsii_.InvokeVoid(
		s,
		"resetApproveAfterDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ResetApproveUntilDate() {
	_jsii_.InvokeVoid(
		s,
		"resetApproveUntilDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ResetComplianceLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetComplianceLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ResetEnableNonSecurity() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableNonSecurity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

