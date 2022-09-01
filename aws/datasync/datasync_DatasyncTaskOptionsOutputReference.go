package datasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/datasync/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncTaskOptionsOutputReference interface {
	cdktf.ComplexObject
	Atime() *string
	SetAtime(val *string)
	AtimeInput() *string
	BytesPerSecond() *float64
	SetBytesPerSecond(val *float64)
	BytesPerSecondInput() *float64
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
	Gid() *string
	SetGid(val *string)
	GidInput() *string
	InternalValue() *DatasyncTaskOptions
	SetInternalValue(val *DatasyncTaskOptions)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	Mtime() *string
	SetMtime(val *string)
	MtimeInput() *string
	OverwriteMode() *string
	SetOverwriteMode(val *string)
	OverwriteModeInput() *string
	PosixPermissions() *string
	SetPosixPermissions(val *string)
	PosixPermissionsInput() *string
	PreserveDeletedFiles() *string
	SetPreserveDeletedFiles(val *string)
	PreserveDeletedFilesInput() *string
	PreserveDevices() *string
	SetPreserveDevices(val *string)
	PreserveDevicesInput() *string
	TaskQueueing() *string
	SetTaskQueueing(val *string)
	TaskQueueingInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferMode() *string
	SetTransferMode(val *string)
	TransferModeInput() *string
	Uid() *string
	SetUid(val *string)
	UidInput() *string
	VerifyMode() *string
	SetVerifyMode(val *string)
	VerifyModeInput() *string
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
	ResetAtime()
	ResetBytesPerSecond()
	ResetGid()
	ResetLogLevel()
	ResetMtime()
	ResetOverwriteMode()
	ResetPosixPermissions()
	ResetPreserveDeletedFiles()
	ResetPreserveDevices()
	ResetTaskQueueing()
	ResetTransferMode()
	ResetUid()
	ResetVerifyMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatasyncTaskOptionsOutputReference
type jsiiProxy_DatasyncTaskOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) Atime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) AtimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) BytesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) BytesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bytesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) Gid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) GidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) InternalValue() *DatasyncTaskOptions {
	var returns *DatasyncTaskOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) Mtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) MtimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) OverwriteMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) OverwriteModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) PosixPermissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"posixPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) PosixPermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"posixPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) PreserveDeletedFiles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDeletedFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) PreserveDeletedFilesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDeletedFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) PreserveDevices() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) PreserveDevicesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) TaskQueueing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskQueueing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) TaskQueueingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskQueueingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) TransferMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) TransferModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) UidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) VerifyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference) VerifyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifyModeInput",
		&returns,
	)
	return returns
}


func NewDatasyncTaskOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatasyncTaskOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDatasyncTaskOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatasyncTaskOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.datasync.DatasyncTaskOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatasyncTaskOptionsOutputReference_Override(d DatasyncTaskOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.datasync.DatasyncTaskOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetAtime(val *string) {
	if err := j.validateSetAtimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atime",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetBytesPerSecond(val *float64) {
	if err := j.validateSetBytesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bytesPerSecond",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetGid(val *string) {
	if err := j.validateSetGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetInternalValue(val *DatasyncTaskOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetMtime(val *string) {
	if err := j.validateSetMtimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtime",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetOverwriteMode(val *string) {
	if err := j.validateSetOverwriteModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteMode",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetPosixPermissions(val *string) {
	if err := j.validateSetPosixPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"posixPermissions",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetPreserveDeletedFiles(val *string) {
	if err := j.validateSetPreserveDeletedFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveDeletedFiles",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetPreserveDevices(val *string) {
	if err := j.validateSetPreserveDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveDevices",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetTaskQueueing(val *string) {
	if err := j.validateSetTaskQueueingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskQueueing",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetTransferMode(val *string) {
	if err := j.validateSetTransferModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferMode",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetUid(val *string) {
	if err := j.validateSetUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

func (j *jsiiProxy_DatasyncTaskOptionsOutputReference)SetVerifyMode(val *string) {
	if err := j.validateSetVerifyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyMode",
		val,
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetAtime() {
	_jsii_.InvokeVoid(
		d,
		"resetAtime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetBytesPerSecond() {
	_jsii_.InvokeVoid(
		d,
		"resetBytesPerSecond",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetGid() {
	_jsii_.InvokeVoid(
		d,
		"resetGid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetMtime() {
	_jsii_.InvokeVoid(
		d,
		"resetMtime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetOverwriteMode() {
	_jsii_.InvokeVoid(
		d,
		"resetOverwriteMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetPosixPermissions() {
	_jsii_.InvokeVoid(
		d,
		"resetPosixPermissions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetPreserveDeletedFiles() {
	_jsii_.InvokeVoid(
		d,
		"resetPreserveDeletedFiles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetPreserveDevices() {
	_jsii_.InvokeVoid(
		d,
		"resetPreserveDevices",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetTaskQueueing() {
	_jsii_.InvokeVoid(
		d,
		"resetTaskQueueing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetTransferMode() {
	_jsii_.InvokeVoid(
		d,
		"resetTransferMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetUid() {
	_jsii_.InvokeVoid(
		d,
		"resetUid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ResetVerifyMode() {
	_jsii_.InvokeVoid(
		d,
		"resetVerifyMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatasyncTaskOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

