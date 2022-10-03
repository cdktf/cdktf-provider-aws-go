//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2managedprefixlist

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2ManagedPrefixListEntryList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2ManagedPrefixListEntryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2ManagedPrefixListEntryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2ManagedPrefixListEntryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2ManagedPrefixListEntryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2ManagedPrefixListEntryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2ManagedPrefixListEntryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

