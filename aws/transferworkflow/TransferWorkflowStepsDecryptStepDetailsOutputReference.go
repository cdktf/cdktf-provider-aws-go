package transferworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/transferworkflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferWorkflowStepsDecryptStepDetailsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationFileLocation() TransferWorkflowStepsDecryptStepDetailsDestinationFileLocationOutputReference
	DestinationFileLocationInput() *TransferWorkflowStepsDecryptStepDetailsDestinationFileLocation
	// Experimental.
	Fqn() *string
	InternalValue() *TransferWorkflowStepsDecryptStepDetails
	SetInternalValue(val *TransferWorkflowStepsDecryptStepDetails)
	Name() *string
	SetName(val *string)
	NameInput() *string
	OverwriteExisting() *string
	SetOverwriteExisting(val *string)
	OverwriteExistingInput() *string
	SourceFileLocation() *string
	SetSourceFileLocation(val *string)
	SourceFileLocationInput() *string
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
	PutDestinationFileLocation(value *TransferWorkflowStepsDecryptStepDetailsDestinationFileLocation)
	ResetDestinationFileLocation()
	ResetName()
	ResetOverwriteExisting()
	ResetSourceFileLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TransferWorkflowStepsDecryptStepDetailsOutputReference
type jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) DestinationFileLocation() TransferWorkflowStepsDecryptStepDetailsDestinationFileLocationOutputReference {
	var returns TransferWorkflowStepsDecryptStepDetailsDestinationFileLocationOutputReference
	_jsii_.Get(
		j,
		"destinationFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) DestinationFileLocationInput() *TransferWorkflowStepsDecryptStepDetailsDestinationFileLocation {
	var returns *TransferWorkflowStepsDecryptStepDetailsDestinationFileLocation
	_jsii_.Get(
		j,
		"destinationFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) InternalValue() *TransferWorkflowStepsDecryptStepDetails {
	var returns *TransferWorkflowStepsDecryptStepDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) OverwriteExisting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteExisting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) OverwriteExistingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteExistingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) SourceFileLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) SourceFileLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewTransferWorkflowStepsDecryptStepDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TransferWorkflowStepsDecryptStepDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewTransferWorkflowStepsDecryptStepDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowStepsDecryptStepDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTransferWorkflowStepsDecryptStepDetailsOutputReference_Override(t TransferWorkflowStepsDecryptStepDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.transferWorkflow.TransferWorkflowStepsDecryptStepDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetInternalValue(val *TransferWorkflowStepsDecryptStepDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetOverwriteExisting(val *string) {
	if err := j.validateSetOverwriteExistingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteExisting",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetSourceFileLocation(val *string) {
	if err := j.validateSetSourceFileLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFileLocation",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) PutDestinationFileLocation(value *TransferWorkflowStepsDecryptStepDetailsDestinationFileLocation) {
	if err := t.validatePutDestinationFileLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationFileLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) ResetDestinationFileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationFileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) ResetOverwriteExisting() {
	_jsii_.InvokeVoid(
		t,
		"resetOverwriteExisting",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) ResetSourceFileLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceFileLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TransferWorkflowStepsDecryptStepDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

