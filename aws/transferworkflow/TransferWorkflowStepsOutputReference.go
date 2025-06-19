// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/transferworkflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferWorkflowStepsOutputReference interface {
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
	CopyStepDetails() TransferWorkflowStepsCopyStepDetailsOutputReference
	CopyStepDetailsInput() *TransferWorkflowStepsCopyStepDetails
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomStepDetails() TransferWorkflowStepsCustomStepDetailsOutputReference
	CustomStepDetailsInput() *TransferWorkflowStepsCustomStepDetails
	DecryptStepDetails() TransferWorkflowStepsDecryptStepDetailsOutputReference
	DecryptStepDetailsInput() *TransferWorkflowStepsDecryptStepDetails
	DeleteStepDetails() TransferWorkflowStepsDeleteStepDetailsOutputReference
	DeleteStepDetailsInput() *TransferWorkflowStepsDeleteStepDetails
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TagStepDetails() TransferWorkflowStepsTagStepDetailsOutputReference
	TagStepDetailsInput() *TransferWorkflowStepsTagStepDetails
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
	PutCopyStepDetails(value *TransferWorkflowStepsCopyStepDetails)
	PutCustomStepDetails(value *TransferWorkflowStepsCustomStepDetails)
	PutDecryptStepDetails(value *TransferWorkflowStepsDecryptStepDetails)
	PutDeleteStepDetails(value *TransferWorkflowStepsDeleteStepDetails)
	PutTagStepDetails(value *TransferWorkflowStepsTagStepDetails)
	ResetCopyStepDetails()
	ResetCustomStepDetails()
	ResetDecryptStepDetails()
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

// The jsii proxy struct for TransferWorkflowStepsOutputReference
type jsiiProxy_TransferWorkflowStepsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) CopyStepDetails() TransferWorkflowStepsCopyStepDetailsOutputReference {
	var returns TransferWorkflowStepsCopyStepDetailsOutputReference
	_jsii_.Get(
		j,
		"copyStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) CopyStepDetailsInput() *TransferWorkflowStepsCopyStepDetails {
	var returns *TransferWorkflowStepsCopyStepDetails
	_jsii_.Get(
		j,
		"copyStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) CustomStepDetails() TransferWorkflowStepsCustomStepDetailsOutputReference {
	var returns TransferWorkflowStepsCustomStepDetailsOutputReference
	_jsii_.Get(
		j,
		"customStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) CustomStepDetailsInput() *TransferWorkflowStepsCustomStepDetails {
	var returns *TransferWorkflowStepsCustomStepDetails
	_jsii_.Get(
		j,
		"customStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) DecryptStepDetails() TransferWorkflowStepsDecryptStepDetailsOutputReference {
	var returns TransferWorkflowStepsDecryptStepDetailsOutputReference
	_jsii_.Get(
		j,
		"decryptStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) DecryptStepDetailsInput() *TransferWorkflowStepsDecryptStepDetails {
	var returns *TransferWorkflowStepsDecryptStepDetails
	_jsii_.Get(
		j,
		"decryptStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) DeleteStepDetails() TransferWorkflowStepsDeleteStepDetailsOutputReference {
	var returns TransferWorkflowStepsDeleteStepDetailsOutputReference
	_jsii_.Get(
		j,
		"deleteStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) DeleteStepDetailsInput() *TransferWorkflowStepsDeleteStepDetails {
	var returns *TransferWorkflowStepsDeleteStepDetails
	_jsii_.Get(
		j,
		"deleteStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) TagStepDetails() TransferWorkflowStepsTagStepDetailsOutputReference {
	var returns TransferWorkflowStepsTagStepDetailsOutputReference
	_jsii_.Get(
		j,
		"tagStepDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) TagStepDetailsInput() *TransferWorkflowStepsTagStepDetails {
	var returns *TransferWorkflowStepsTagStepDetails
	_jsii_.Get(
		j,
		"tagStepDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewTransferWorkflowStepsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TransferWorkflowStepsOutputReference {
	_init_.Initialize()

	if err := validateNewTransferWorkflowStepsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransferWorkflowStepsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTransferWorkflowStepsOutputReference_Override(t TransferWorkflowStepsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) PutCopyStepDetails(value *TransferWorkflowStepsCopyStepDetails) {
	if err := t.validatePutCopyStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCopyStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) PutCustomStepDetails(value *TransferWorkflowStepsCustomStepDetails) {
	if err := t.validatePutCustomStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) PutDecryptStepDetails(value *TransferWorkflowStepsDecryptStepDetails) {
	if err := t.validatePutDecryptStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDecryptStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) PutDeleteStepDetails(value *TransferWorkflowStepsDeleteStepDetails) {
	if err := t.validatePutDeleteStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeleteStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) PutTagStepDetails(value *TransferWorkflowStepsTagStepDetails) {
	if err := t.validatePutTagStepDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagStepDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) ResetCopyStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) ResetCustomStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) ResetDecryptStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetDecryptStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) ResetDeleteStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) ResetTagStepDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetTagStepDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TransferWorkflowStepsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

