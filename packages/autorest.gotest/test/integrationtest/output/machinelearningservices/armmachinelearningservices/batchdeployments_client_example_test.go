//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/list.json
func ExampleBatchDeploymentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewBatchDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("test-rg", "my-aml-workspace", "testEndpointName", &armmachinelearningservices.BatchDeploymentsClientListOptions{OrderBy: to.Ptr("string"),
		Top:  to.Ptr[int32](1),
		Skip: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.BatchDeploymentTrackedResourceArmPaginatedResult = armmachinelearningservices.BatchDeploymentTrackedResourceArmPaginatedResult{
		// 	Value: []*armmachinelearningservices.BatchDeploymentData{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armmachinelearningservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("string"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmachinelearningservices.ManagedServiceIdentity{
		// 				Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 				UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
		// 					"string": &armmachinelearningservices.UserAssignedIdentity{
		// 						ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 						PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					},
		// 				},
		// 			},
		// 			Kind: to.Ptr("string"),
		// 			Properties: &armmachinelearningservices.BatchDeploymentDetails{
		// 				Description: to.Ptr("string"),
		// 				CodeConfiguration: &armmachinelearningservices.CodeConfiguration{
		// 					CodeID: to.Ptr("string"),
		// 					ScoringScript: to.Ptr("string"),
		// 				},
		// 				EnvironmentID: to.Ptr("string"),
		// 				EnvironmentVariables: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				Properties: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				Compute: to.Ptr("string"),
		// 				ErrorThreshold: to.Ptr[int32](1),
		// 				LoggingLevel: to.Ptr(armmachinelearningservices.BatchLoggingLevelInfo),
		// 				MaxConcurrencyPerInstance: to.Ptr[int32](1),
		// 				MiniBatchSize: to.Ptr[int64](1),
		// 				Model: &armmachinelearningservices.IDAssetReference{
		// 					ReferenceType: to.Ptr(armmachinelearningservices.ReferenceTypeID),
		// 					AssetID: to.Ptr("string"),
		// 				},
		// 				OutputAction: to.Ptr(armmachinelearningservices.BatchOutputActionSummaryOnly),
		// 				OutputFileName: to.Ptr("string"),
		// 				ProvisioningState: to.Ptr(armmachinelearningservices.DeploymentProvisioningStateSucceeded),
		// 				Resources: &armmachinelearningservices.ResourceConfiguration{
		// 					InstanceCount: to.Ptr[int32](1),
		// 					InstanceType: to.Ptr("string"),
		// 					Properties: map[string]any{
		// 						"string": map[string]any{
		// 							"a3c13e2e-a213-4cac-9f5a-b49966906ad6": nil,
		// 						},
		// 					},
		// 				},
		// 				RetrySettings: &armmachinelearningservices.BatchRetrySettings{
		// 					MaxRetries: to.Ptr[int32](1),
		// 					Timeout: to.Ptr("PT5M"),
		// 				},
		// 			},
		// 			SKU: &armmachinelearningservices.SKU{
		// 				Name: to.Ptr("string"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Family: to.Ptr("string"),
		// 				Size: to.Ptr("string"),
		// 				Tier: to.Ptr(armmachinelearningservices.SKUTierFree),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/delete.json
func ExampleBatchDeploymentsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewBatchDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/get.json
func ExampleBatchDeploymentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewBatchDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BatchDeploymentData = armmachinelearningservices.BatchDeploymentData{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearningservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearningservices.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
	// 			"string": &armmachinelearningservices.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearningservices.BatchDeploymentDetails{
	// 		Description: to.Ptr("string"),
	// 		CodeConfiguration: &armmachinelearningservices.CodeConfiguration{
	// 			CodeID: to.Ptr("string"),
	// 			ScoringScript: to.Ptr("string"),
	// 		},
	// 		EnvironmentID: to.Ptr("string"),
	// 		EnvironmentVariables: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Compute: to.Ptr("string"),
	// 		ErrorThreshold: to.Ptr[int32](1),
	// 		LoggingLevel: to.Ptr(armmachinelearningservices.BatchLoggingLevelInfo),
	// 		MaxConcurrencyPerInstance: to.Ptr[int32](1),
	// 		MiniBatchSize: to.Ptr[int64](1),
	// 		Model: &armmachinelearningservices.IDAssetReference{
	// 			ReferenceType: to.Ptr(armmachinelearningservices.ReferenceTypeID),
	// 			AssetID: to.Ptr("string"),
	// 		},
	// 		OutputAction: to.Ptr(armmachinelearningservices.BatchOutputActionSummaryOnly),
	// 		OutputFileName: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armmachinelearningservices.DeploymentProvisioningStateSucceeded),
	// 		Resources: &armmachinelearningservices.ResourceConfiguration{
	// 			InstanceCount: to.Ptr[int32](1),
	// 			InstanceType: to.Ptr("string"),
	// 			Properties: map[string]any{
	// 				"string": map[string]any{
	// 					"843c2bb4-e5f1-4267-98c8-ba22a99dbb00": nil,
	// 				},
	// 			},
	// 		},
	// 		RetrySettings: &armmachinelearningservices.BatchRetrySettings{
	// 			MaxRetries: to.Ptr[int32](1),
	// 			Timeout: to.Ptr("PT5M"),
	// 		},
	// 	},
	// 	SKU: &armmachinelearningservices.SKU{
	// 		Name: to.Ptr("string"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("string"),
	// 		Size: to.Ptr("string"),
	// 		Tier: to.Ptr(armmachinelearningservices.SKUTierFree),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/update.json
func ExampleBatchDeploymentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewBatchDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", armmachinelearningservices.PartialBatchDeploymentPartialTrackedResource{
		Identity: &armmachinelearningservices.PartialManagedServiceIdentity{
			Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
			UserAssignedIdentities: map[string]any{
				"string": map[string]any{},
			},
		},
		Kind:     to.Ptr("string"),
		Location: to.Ptr("string"),
		Properties: &armmachinelearningservices.PartialBatchDeployment{
			Description: to.Ptr("string"),
			CodeConfiguration: &armmachinelearningservices.PartialCodeConfiguration{
				CodeID:        to.Ptr("string"),
				ScoringScript: to.Ptr("string"),
			},
			Compute:       to.Ptr("string"),
			EnvironmentID: to.Ptr("string"),
			EnvironmentVariables: map[string]*string{
				"string": to.Ptr("string"),
			},
			ErrorThreshold:            to.Ptr[int32](1),
			LoggingLevel:              to.Ptr(armmachinelearningservices.BatchLoggingLevelInfo),
			MaxConcurrencyPerInstance: to.Ptr[int32](1),
			MiniBatchSize:             to.Ptr[int64](1),
			Model: &armmachinelearningservices.PartialIDAssetReference{
				ReferenceType: to.Ptr(armmachinelearningservices.ReferenceTypeID),
				AssetID:       to.Ptr("string"),
			},
			OutputAction:   to.Ptr(armmachinelearningservices.BatchOutputActionSummaryOnly),
			OutputFileName: to.Ptr("string"),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			RetrySettings: &armmachinelearningservices.PartialBatchRetrySettings{
				MaxRetries: to.Ptr[int32](1),
				Timeout:    to.Ptr("PT5M"),
			},
		},
		SKU: &armmachinelearningservices.PartialSKU{
			Name:     to.Ptr("string"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("string"),
			Size:     to.Ptr("string"),
			Tier:     to.Ptr(armmachinelearningservices.SKUTierFree),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BatchDeploymentData = armmachinelearningservices.BatchDeploymentData{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearningservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearningservices.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
	// 			"string": &armmachinelearningservices.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearningservices.BatchDeploymentDetails{
	// 		Description: to.Ptr("string"),
	// 		CodeConfiguration: &armmachinelearningservices.CodeConfiguration{
	// 			CodeID: to.Ptr("string"),
	// 			ScoringScript: to.Ptr("string"),
	// 		},
	// 		EnvironmentID: to.Ptr("string"),
	// 		EnvironmentVariables: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Compute: to.Ptr("string"),
	// 		ErrorThreshold: to.Ptr[int32](1),
	// 		LoggingLevel: to.Ptr(armmachinelearningservices.BatchLoggingLevelInfo),
	// 		MaxConcurrencyPerInstance: to.Ptr[int32](1),
	// 		MiniBatchSize: to.Ptr[int64](1),
	// 		Model: &armmachinelearningservices.IDAssetReference{
	// 			ReferenceType: to.Ptr(armmachinelearningservices.ReferenceTypeID),
	// 			AssetID: to.Ptr("string"),
	// 		},
	// 		OutputAction: to.Ptr(armmachinelearningservices.BatchOutputActionSummaryOnly),
	// 		OutputFileName: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armmachinelearningservices.DeploymentProvisioningStateSucceeded),
	// 		Resources: &armmachinelearningservices.ResourceConfiguration{
	// 			InstanceCount: to.Ptr[int32](1),
	// 			InstanceType: to.Ptr("string"),
	// 			Properties: map[string]any{
	// 				"string": map[string]any{
	// 					"1e5e1cf9-b0ea-4cf6-9764-e750bf85c10a": nil,
	// 				},
	// 			},
	// 		},
	// 		RetrySettings: &armmachinelearningservices.BatchRetrySettings{
	// 			MaxRetries: to.Ptr[int32](1),
	// 			Timeout: to.Ptr("PT5M"),
	// 		},
	// 	},
	// 	SKU: &armmachinelearningservices.SKU{
	// 		Name: to.Ptr("string"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("string"),
	// 		Size: to.Ptr("string"),
	// 		Tier: to.Ptr(armmachinelearningservices.SKUTierFree),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/createOrUpdate.json
func ExampleBatchDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewBatchDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", armmachinelearningservices.BatchDeploymentData{
		Location: to.Ptr("string"),
		Tags:     map[string]*string{},
		Identity: &armmachinelearningservices.ManagedServiceIdentity{
			Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
			UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
				"string": &armmachinelearningservices.UserAssignedIdentity{},
			},
		},
		Kind: to.Ptr("string"),
		Properties: &armmachinelearningservices.BatchDeploymentDetails{
			Description: to.Ptr("string"),
			CodeConfiguration: &armmachinelearningservices.CodeConfiguration{
				CodeID:        to.Ptr("string"),
				ScoringScript: to.Ptr("string"),
			},
			EnvironmentID: to.Ptr("string"),
			EnvironmentVariables: map[string]*string{
				"string": to.Ptr("string"),
			},
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Compute:                   to.Ptr("string"),
			ErrorThreshold:            to.Ptr[int32](1),
			LoggingLevel:              to.Ptr(armmachinelearningservices.BatchLoggingLevelInfo),
			MaxConcurrencyPerInstance: to.Ptr[int32](1),
			MiniBatchSize:             to.Ptr[int64](1),
			Model: &armmachinelearningservices.IDAssetReference{
				ReferenceType: to.Ptr(armmachinelearningservices.ReferenceTypeID),
				AssetID:       to.Ptr("string"),
			},
			OutputAction:   to.Ptr(armmachinelearningservices.BatchOutputActionSummaryOnly),
			OutputFileName: to.Ptr("string"),
			Resources: &armmachinelearningservices.ResourceConfiguration{
				InstanceCount: to.Ptr[int32](1),
				InstanceType:  to.Ptr("string"),
				Properties: map[string]any{
					"string": map[string]any{
						"cd3c37dc-2876-4ca4-8a54-21bd7619724a": nil,
					},
				},
			},
			RetrySettings: &armmachinelearningservices.BatchRetrySettings{
				MaxRetries: to.Ptr[int32](1),
				Timeout:    to.Ptr("PT5M"),
			},
		},
		SKU: &armmachinelearningservices.SKU{
			Name:     to.Ptr("string"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("string"),
			Size:     to.Ptr("string"),
			Tier:     to.Ptr(armmachinelearningservices.SKUTierFree),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BatchDeploymentData = armmachinelearningservices.BatchDeploymentData{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearningservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearningservices.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearningservices.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearningservices.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
	// 			"string": &armmachinelearningservices.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearningservices.BatchDeploymentDetails{
	// 		Description: to.Ptr("string"),
	// 		CodeConfiguration: &armmachinelearningservices.CodeConfiguration{
	// 			CodeID: to.Ptr("string"),
	// 			ScoringScript: to.Ptr("string"),
	// 		},
	// 		EnvironmentID: to.Ptr("string"),
	// 		EnvironmentVariables: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Compute: to.Ptr("string"),
	// 		ErrorThreshold: to.Ptr[int32](1),
	// 		LoggingLevel: to.Ptr(armmachinelearningservices.BatchLoggingLevelInfo),
	// 		MaxConcurrencyPerInstance: to.Ptr[int32](1),
	// 		MiniBatchSize: to.Ptr[int64](1),
	// 		Model: &armmachinelearningservices.IDAssetReference{
	// 			ReferenceType: to.Ptr(armmachinelearningservices.ReferenceTypeID),
	// 			AssetID: to.Ptr("string"),
	// 		},
	// 		OutputAction: to.Ptr(armmachinelearningservices.BatchOutputActionSummaryOnly),
	// 		OutputFileName: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armmachinelearningservices.DeploymentProvisioningStateSucceeded),
	// 		Resources: &armmachinelearningservices.ResourceConfiguration{
	// 			InstanceCount: to.Ptr[int32](1),
	// 			InstanceType: to.Ptr("string"),
	// 			Properties: map[string]any{
	// 				"string": map[string]any{
	// 					"4939850d-8eae-4343-8566-0826259a2ad1": nil,
	// 				},
	// 			},
	// 		},
	// 		RetrySettings: &armmachinelearningservices.BatchRetrySettings{
	// 			MaxRetries: to.Ptr[int32](1),
	// 			Timeout: to.Ptr("PT5M"),
	// 		},
	// 	},
	// 	SKU: &armmachinelearningservices.SKU{
	// 		Name: to.Ptr("string"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("string"),
	// 		Size: to.Ptr("string"),
	// 		Tier: to.Ptr(armmachinelearningservices.SKUTierFree),
	// 	},
	// }
}
