package fsxopenzfsfilesystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/fsxopenzfsfilesystem/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference interface {
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
	CopyTagsToSnapshots() interface{}
	SetCopyTagsToSnapshots(val interface{})
	CopyTagsToSnapshotsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataCompressionType() *string
	SetDataCompressionType(val *string)
	DataCompressionTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *FsxOpenzfsFileSystemRootVolumeConfiguration
	SetInternalValue(val *FsxOpenzfsFileSystemRootVolumeConfiguration)
	NfsExports() FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference
	NfsExportsInput() *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	RecordSizeKib() *float64
	SetRecordSizeKib(val *float64)
	RecordSizeKibInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAndGroupQuotas() FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList
	UserAndGroupQuotasInput() interface{}
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
	PutNfsExports(value *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports)
	PutUserAndGroupQuotas(value interface{})
	ResetCopyTagsToSnapshots()
	ResetDataCompressionType()
	ResetNfsExports()
	ResetReadOnly()
	ResetRecordSizeKib()
	ResetUserAndGroupQuotas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference
type jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) CopyTagsToSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) CopyTagsToSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) InternalValue() *FsxOpenzfsFileSystemRootVolumeConfiguration {
	var returns *FsxOpenzfsFileSystemRootVolumeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) NfsExports() FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference {
	var returns FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference
	_jsii_.Get(
		j,
		"nfsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) NfsExportsInput() *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports {
	var returns *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports
	_jsii_.Get(
		j,
		"nfsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) RecordSizeKib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) RecordSizeKibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) UserAndGroupQuotas() FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList {
	var returns FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotasList
	_jsii_.Get(
		j,
		"userAndGroupQuotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) UserAndGroupQuotasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotasInput",
		&returns,
	)
	return returns
}


func NewFsxOpenzfsFileSystemRootVolumeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFsxOpenzfsFileSystemRootVolumeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOpenzfsFileSystem.FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFsxOpenzfsFileSystemRootVolumeConfigurationOutputReference_Override(f FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOpenzfsFileSystem.FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetCopyTagsToSnapshots(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshots",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetInternalValue(val *FsxOpenzfsFileSystemRootVolumeConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetRecordSizeKib(val *float64) {
	if err := j.validateSetRecordSizeKibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSizeKib",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) PutNfsExports(value *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports) {
	if err := f.validatePutNfsExportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putNfsExports",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) PutUserAndGroupQuotas(value interface{}) {
	if err := f.validatePutUserAndGroupQuotasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putUserAndGroupQuotas",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetCopyTagsToSnapshots() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToSnapshots",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		f,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetNfsExports() {
	_jsii_.InvokeVoid(
		f,
		"resetNfsExports",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		f,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetRecordSizeKib() {
	_jsii_.InvokeVoid(
		f,
		"resetRecordSizeKib",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetUserAndGroupQuotas() {
	_jsii_.InvokeVoid(
		f,
		"resetUserAndGroupQuotas",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

