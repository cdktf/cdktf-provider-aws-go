package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EbsSnapshotImportClientDataOutputReference interface {
	cdktf.ComplexObject
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *EbsSnapshotImportClientData
	SetInternalValue(val *EbsSnapshotImportClientData)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UploadEnd() *string
	SetUploadEnd(val *string)
	UploadEndInput() *string
	UploadSize() *float64
	SetUploadSize(val *float64)
	UploadSizeInput() *float64
	UploadStart() *string
	SetUploadStart(val *string)
	UploadStartInput() *string
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
	ResetComment()
	ResetUploadEnd()
	ResetUploadSize()
	ResetUploadStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EbsSnapshotImportClientDataOutputReference
type jsiiProxy_EbsSnapshotImportClientDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) InternalValue() *EbsSnapshotImportClientData {
	var returns *EbsSnapshotImportClientData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) UploadEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) UploadEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) UploadSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uploadSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) UploadSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uploadSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) UploadStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference) UploadStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadStartInput",
		&returns,
	)
	return returns
}


func NewEbsSnapshotImportClientDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EbsSnapshotImportClientDataOutputReference {
	_init_.Initialize()

	if err := validateNewEbsSnapshotImportClientDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EbsSnapshotImportClientDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.EbsSnapshotImportClientDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEbsSnapshotImportClientDataOutputReference_Override(e EbsSnapshotImportClientDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2.EbsSnapshotImportClientDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetInternalValue(val *EbsSnapshotImportClientData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetUploadEnd(val *string) {
	if err := j.validateSetUploadEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadEnd",
		val,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetUploadSize(val *float64) {
	if err := j.validateSetUploadSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadSize",
		val,
	)
}

func (j *jsiiProxy_EbsSnapshotImportClientDataOutputReference)SetUploadStart(val *string) {
	if err := j.validateSetUploadStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadStart",
		val,
	)
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		e,
		"resetComment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) ResetUploadEnd() {
	_jsii_.InvokeVoid(
		e,
		"resetUploadEnd",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) ResetUploadSize() {
	_jsii_.InvokeVoid(
		e,
		"resetUploadSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) ResetUploadStart() {
	_jsii_.InvokeVoid(
		e,
		"resetUploadStart",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EbsSnapshotImportClientDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

