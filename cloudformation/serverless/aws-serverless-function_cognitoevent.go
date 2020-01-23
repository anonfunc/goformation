package serverless

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Function_CognitoEvent AWS CloudFormation Resource (AWS::Serverless::Function.CognitoEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#cognito
type Function_CognitoEvent struct {

	// Trigger AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#cognito
	Trigger string `json:"Trigger,omitempty"`

	// UserPool AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#cognito
	UserPool string `json:"UserPool,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_CognitoEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.CognitoEvent"
}
