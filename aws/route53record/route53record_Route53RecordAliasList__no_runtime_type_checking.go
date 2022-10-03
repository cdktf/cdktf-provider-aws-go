//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package route53record

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53RecordAliasList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_Route53RecordAliasList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Route53RecordAliasList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Route53RecordAliasList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Route53RecordAliasList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Route53RecordAliasList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoute53RecordAliasListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

