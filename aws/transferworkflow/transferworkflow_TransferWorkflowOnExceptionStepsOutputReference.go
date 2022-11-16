package transferworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/transferworkflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferWorkflowOnExceptionStepsOutputReference interface {
	cdktf.ComplexObject
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
	CopyStepDetails() TransferWorkflowOnExceptionStepsCopyStepDetailsOutputReference
	CopyStepDetailsInput() *TransferWorkflowOnExceptionStepsCopyStepDetails
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomStepDetails() TransferWorkflowOnExceptionStepsCustomStepDetailsOutputReference
	CustomStepDetailsInput() *TransferWorkflowOnExceptionStepsCustomStepDetails
	DeleteStepDetails() TransferWorkflowOnExceptionStepsDeleteStepDetailsOutputReference
	DeleteStepDetailsInput() *TransferWorkflowOnExceptionStepsDeleteStepDetails
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TagStepDetails() TransferWorkflowOnExceptionStepsTagStepDetailsOutputReference
	TagStepDetailsInput() *TransferWorkflowOnExceptionStepsTagStepDetails
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutCopyStepDetails(value *TransferWorkflowOnExceptionStepsCopyStepDetails)
	PutCustomStepDetails(value *TransferWorkflowOnExceptionStepsCustomStepDetails)
	PutDeleteStepDetails(value *TransferWorkflowOnExceptionStepsDeleteStepDetails)
	PutTagStepDetails(value *TransferWorkflowOnExceptionStepsTagStepDetails)
	ResetCopyStepDetails()
	ResetCustomStepDetails()
	ResetDeleteStepDetails()
	ResetTagStepDetails()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TransferWorkflowOnExceptionStepsOutputReference
type jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) CopyStepDetails() TransferWorkflowOnExceptionStepsCopyStepDetailsOutputReference {
	var returns TransferWorkflowOnExceptionStepsCopyStepDetailsOutputReference
	_jsii_.Get(
		j,
		"copyStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) CopyStepDetailsInput() *TransferWorkflowOnExceptionStepsCopyStepDetails {
	var returns *TransferWorkflowOnExceptionStepsCopyStepDetails
	_jsii_.Get(
		j,
		"copyStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) CustomStepDetails() TransferWorkflowOnExceptionStepsCustomStepDetailsOutputReference {
	var returns TransferWorkflowOnExceptionStepsCustomStepDetailsOutputReference
	_jsii_.Get(
		j,
		"customStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) CustomStepDetailsInput() *TransferWorkflowOnExceptionStepsCustomStepDetails {
	var returns *TransferWorkflowOnExceptionStepsCustomStepDetails
	_jsii_.Get(
		j,
		"customStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) DeleteStepDetails() TransferWorkflowOnExceptionStepsDeleteStepDetailsOutputReference {
	var returns TransferWorkflowOnExceptionStepsDeleteStepDetailsOutputReference
	_jsii_.Get(
		j,
		"deleteStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) DeleteStepDetailsInput() *TransferWorkflowOnExceptionStepsDeleteStepDetails {
	var returns *TransferWorkflowOnExceptionStepsDeleteStepDetails
	_jsii_.Get(
		j,
		"deleteStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) TagStepDetails() TransferWorkflowOnExceptionStepsTagStepDetailsOutputReference {
	var returns TransferWorkflowOnExceptionStepsTagStepDetailsOutputReference
	_jsii_.Get(
		j,
		"tagStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) TagStepDetailsInput() *TransferWorkflowOnExceptionStepsTagStepDetails {
	var returns *TransferWorkflowOnExceptionStepsTagStepDetails
	_jsii_.Get(
		j,
		"tagStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewTransferWorkflowOnExceptionStepsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TransferWorkflowOnExceptionStepsOutputReference {
	_init_.Initialize()

	if err := validateNewTransferWorkflowOnExceptionStepsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowOnExceptionStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTransferWorkflowOnExceptionStepsOutputReference_Override(t TransferWorkflowOnExceptionStepsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowOnExceptionStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) PutCopyStepDetails(value *TransferWorkflowOnExceptionStepsCopyStepDetails) {
	if err := t.validatePutCopyStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCopyStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) PutCustomStepDetails(value *TransferWorkflowOnExceptionStepsCustomStepDetails) {
	if err := t.validatePutCustomStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) PutDeleteStepDetails(value *TransferWorkflowOnExceptionStepsDeleteStepDetails) {
	if err := t.validatePutDeleteStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeleteStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) PutTagStepDetails(value *TransferWorkflowOnExceptionStepsTagStepDetails) {
	if err := t.validatePutTagStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) ResetCopyStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) ResetCustomStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) ResetDeleteStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) ResetTagStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetTagStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowOnExceptionStepsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

