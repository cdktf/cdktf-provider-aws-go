//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FisExperimentTemplateTargetList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FisExperimentTemplateTargetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FisExperimentTemplateTargetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FisExperimentTemplateTargetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FisExperimentTemplateTargetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FisExperimentTemplateTargetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFisExperimentTemplateTargetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

