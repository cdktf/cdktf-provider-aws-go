package dbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/dbinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DbInstanceRestoreToPointInTimeOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DbInstanceRestoreToPointInTime
	SetInternalValue(val *DbInstanceRestoreToPointInTime)
	RestoreTime() *string
	SetRestoreTime(val *string)
	RestoreTimeInput() *string
	SourceDbInstanceAutomatedBackupsArn() *string
	SetSourceDbInstanceAutomatedBackupsArn(val *string)
	SourceDbInstanceAutomatedBackupsArnInput() *string
	SourceDbInstanceIdentifier() *string
	SetSourceDbInstanceIdentifier(val *string)
	SourceDbInstanceIdentifierInput() *string
	SourceDbiResourceId() *string
	SetSourceDbiResourceId(val *string)
	SourceDbiResourceIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
	UseLatestRestorableTimeInput() interface{}
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
	ResetRestoreTime()
	ResetSourceDbInstanceAutomatedBackupsArn()
	ResetSourceDbInstanceIdentifier()
	ResetSourceDbiResourceId()
	ResetUseLatestRestorableTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DbInstanceRestoreToPointInTimeOutputReference
type jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) InternalValue() *DbInstanceRestoreToPointInTime {
	var returns *DbInstanceRestoreToPointInTime
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) RestoreTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) RestoreTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) SourceDbInstanceAutomatedBackupsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) SourceDbInstanceAutomatedBackupsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) SourceDbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) SourceDbInstanceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) SourceDbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) SourceDbiResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) UseLatestRestorableTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTimeInput",
		&returns,
	)
	return returns
}


func NewDbInstanceRestoreToPointInTimeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DbInstanceRestoreToPointInTimeOutputReference {
	_init_.Initialize()

	if err := validateNewDbInstanceRestoreToPointInTimeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.dbInstance.DbInstanceRestoreToPointInTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDbInstanceRestoreToPointInTimeOutputReference_Override(d DbInstanceRestoreToPointInTimeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dbInstance.DbInstanceRestoreToPointInTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetInternalValue(val *DbInstanceRestoreToPointInTime) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetRestoreTime(val *string) {
	if err := j.validateSetRestoreTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreTime",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetSourceDbInstanceAutomatedBackupsArn(val *string) {
	if err := j.validateSetSourceDbInstanceAutomatedBackupsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetSourceDbInstanceIdentifier(val *string) {
	if err := j.validateSetSourceDbInstanceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetSourceDbiResourceId(val *string) {
	if err := j.validateSetSourceDbiResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbiResourceId",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ResetRestoreTime() {
	_jsii_.InvokeVoid(
		d,
		"resetRestoreTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ResetSourceDbInstanceAutomatedBackupsArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceDbInstanceAutomatedBackupsArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ResetSourceDbInstanceIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceDbInstanceIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ResetSourceDbiResourceId() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceDbiResourceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ResetUseLatestRestorableTime() {
	_jsii_.InvokeVoid(
		d,
		"resetUseLatestRestorableTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstanceRestoreToPointInTimeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

