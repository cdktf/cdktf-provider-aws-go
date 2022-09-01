// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type ComprehendEntityRecognizerInputDataConfig struct {
	// entity_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#entity_types ComprehendEntityRecognizer#entity_types}
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// annotations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#annotations ComprehendEntityRecognizer#annotations}
	Annotations *ComprehendEntityRecognizerInputDataConfigAnnotations `field:"optional" json:"annotations" yaml:"annotations"`
	// augmented_manifests block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#augmented_manifests ComprehendEntityRecognizer#augmented_manifests}
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#data_format ComprehendEntityRecognizer#data_format}.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// documents block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#documents ComprehendEntityRecognizer#documents}
	Documents *ComprehendEntityRecognizerInputDataConfigDocuments `field:"optional" json:"documents" yaml:"documents"`
	// entity_list block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#entity_list ComprehendEntityRecognizer#entity_list}
	EntityList *ComprehendEntityRecognizerInputDataConfigEntityList `field:"optional" json:"entityList" yaml:"entityList"`
}

