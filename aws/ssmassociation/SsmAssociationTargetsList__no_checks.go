//go:build no_runtime_type_checking

package ssmassociation

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmAssociationTargetsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmAssociationTargetsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmAssociationTargetsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmAssociationTargetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmAssociationTargetsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmAssociationTargetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmAssociationTargetsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

