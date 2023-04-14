package codegurureviewerrepositoryassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/codegurureviewerrepositoryassociation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodegurureviewerRepositoryAssociationRepositoryOutputReference interface {
	cdktf.ComplexObject
	Bitbucket() CodegurureviewerRepositoryAssociationRepositoryBitbucketOutputReference
	BitbucketInput() *CodegurureviewerRepositoryAssociationRepositoryBitbucket
	Codecommit() CodegurureviewerRepositoryAssociationRepositoryCodecommitOutputReference
	CodecommitInput() *CodegurureviewerRepositoryAssociationRepositoryCodecommit
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
	GithubEnterpriseServer() CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServerOutputReference
	GithubEnterpriseServerInput() *CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServer
	InternalValue() *CodegurureviewerRepositoryAssociationRepository
	SetInternalValue(val *CodegurureviewerRepositoryAssociationRepository)
	S3Bucket() CodegurureviewerRepositoryAssociationRepositoryS3BucketOutputReference
	S3BucketInput() *CodegurureviewerRepositoryAssociationRepositoryS3Bucket
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
	PutBitbucket(value *CodegurureviewerRepositoryAssociationRepositoryBitbucket)
	PutCodecommit(value *CodegurureviewerRepositoryAssociationRepositoryCodecommit)
	PutGithubEnterpriseServer(value *CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServer)
	PutS3Bucket(value *CodegurureviewerRepositoryAssociationRepositoryS3Bucket)
	ResetBitbucket()
	ResetCodecommit()
	ResetGithubEnterpriseServer()
	ResetS3Bucket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodegurureviewerRepositoryAssociationRepositoryOutputReference
type jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) Bitbucket() CodegurureviewerRepositoryAssociationRepositoryBitbucketOutputReference {
	var returns CodegurureviewerRepositoryAssociationRepositoryBitbucketOutputReference
	_jsii_.Get(
		j,
		"bitbucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) BitbucketInput() *CodegurureviewerRepositoryAssociationRepositoryBitbucket {
	var returns *CodegurureviewerRepositoryAssociationRepositoryBitbucket
	_jsii_.Get(
		j,
		"bitbucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) Codecommit() CodegurureviewerRepositoryAssociationRepositoryCodecommitOutputReference {
	var returns CodegurureviewerRepositoryAssociationRepositoryCodecommitOutputReference
	_jsii_.Get(
		j,
		"codecommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) CodecommitInput() *CodegurureviewerRepositoryAssociationRepositoryCodecommit {
	var returns *CodegurureviewerRepositoryAssociationRepositoryCodecommit
	_jsii_.Get(
		j,
		"codecommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GithubEnterpriseServer() CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServerOutputReference {
	var returns CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServerOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GithubEnterpriseServerInput() *CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServer {
	var returns *CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServer
	_jsii_.Get(
		j,
		"githubEnterpriseServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) InternalValue() *CodegurureviewerRepositoryAssociationRepository {
	var returns *CodegurureviewerRepositoryAssociationRepository
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) S3Bucket() CodegurureviewerRepositoryAssociationRepositoryS3BucketOutputReference {
	var returns CodegurureviewerRepositoryAssociationRepositoryS3BucketOutputReference
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) S3BucketInput() *CodegurureviewerRepositoryAssociationRepositoryS3Bucket {
	var returns *CodegurureviewerRepositoryAssociationRepositoryS3Bucket
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCodegurureviewerRepositoryAssociationRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodegurureviewerRepositoryAssociationRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewCodegurureviewerRepositoryAssociationRepositoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codegurureviewerRepositoryAssociation.CodegurureviewerRepositoryAssociationRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodegurureviewerRepositoryAssociationRepositoryOutputReference_Override(c CodegurureviewerRepositoryAssociationRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codegurureviewerRepositoryAssociation.CodegurureviewerRepositoryAssociationRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference)SetInternalValue(val *CodegurureviewerRepositoryAssociationRepository) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) PutBitbucket(value *CodegurureviewerRepositoryAssociationRepositoryBitbucket) {
	if err := c.validatePutBitbucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBitbucket",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) PutCodecommit(value *CodegurureviewerRepositoryAssociationRepositoryCodecommit) {
	if err := c.validatePutCodecommitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCodecommit",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) PutGithubEnterpriseServer(value *CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServer) {
	if err := c.validatePutGithubEnterpriseServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGithubEnterpriseServer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) PutS3Bucket(value *CodegurureviewerRepositoryAssociationRepositoryS3Bucket) {
	if err := c.validatePutS3BucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3Bucket",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) ResetBitbucket() {
	_jsii_.InvokeVoid(
		c,
		"resetBitbucket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) ResetCodecommit() {
	_jsii_.InvokeVoid(
		c,
		"resetCodecommit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) ResetGithubEnterpriseServer() {
	_jsii_.InvokeVoid(
		c,
		"resetGithubEnterpriseServer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodegurureviewerRepositoryAssociationRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

