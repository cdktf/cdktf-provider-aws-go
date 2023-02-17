//go:build no_runtime_type_checking

package evidentlyfeature

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvidentlyFeatureVariationsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EvidentlyFeatureVariationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureVariationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureVariationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureVariationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EvidentlyFeatureVariationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEvidentlyFeatureVariationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

