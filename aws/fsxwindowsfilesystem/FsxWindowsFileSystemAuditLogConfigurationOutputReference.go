package fsxwindowsfilesystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/fsxwindowsfilesystem/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxWindowsFileSystemAuditLogConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuditLogDestination() *string
	SetAuditLogDestination(val *string)
	AuditLogDestinationInput() *string
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
	FileAccessAuditLogLevel() *string
	SetFileAccessAuditLogLevel(val *string)
	FileAccessAuditLogLevelInput() *string
	FileShareAccessAuditLogLevel() *string
	SetFileShareAccessAuditLogLevel(val *string)
	FileShareAccessAuditLogLevelInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *FsxWindowsFileSystemAuditLogConfiguration
	SetInternalValue(val *FsxWindowsFileSystemAuditLogConfiguration)
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
	ResetAuditLogDestination()
	ResetFileAccessAuditLogLevel()
	ResetFileShareAccessAuditLogLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxWindowsFileSystemAuditLogConfigurationOutputReference
type jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) AuditLogDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) AuditLogDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) FileAccessAuditLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAccessAuditLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) FileAccessAuditLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAccessAuditLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) FileShareAccessAuditLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareAccessAuditLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) FileShareAccessAuditLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareAccessAuditLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) InternalValue() *FsxWindowsFileSystemAuditLogConfiguration {
	var returns *FsxWindowsFileSystemAuditLogConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFsxWindowsFileSystemAuditLogConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FsxWindowsFileSystemAuditLogConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFsxWindowsFileSystemAuditLogConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxWindowsFileSystem.FsxWindowsFileSystemAuditLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxWindowsFileSystemAuditLogConfigurationOutputReference_Override(f FsxWindowsFileSystemAuditLogConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxWindowsFileSystem.FsxWindowsFileSystemAuditLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference)SetAuditLogDestination(val *string) {
	if err := j.validateSetAuditLogDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogDestination",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference)SetFileAccessAuditLogLevel(val *string) {
	if err := j.validateSetFileAccessAuditLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileAccessAuditLogLevel",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference)SetFileShareAccessAuditLogLevel(val *string) {
	if err := j.validateSetFileShareAccessAuditLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileShareAccessAuditLogLevel",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference)SetInternalValue(val *FsxWindowsFileSystemAuditLogConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ResetAuditLogDestination() {
	_jsii_.InvokeVoid(
		f,
		"resetAuditLogDestination",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ResetFileAccessAuditLogLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetFileAccessAuditLogLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ResetFileShareAccessAuditLogLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetFileShareAccessAuditLogLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

