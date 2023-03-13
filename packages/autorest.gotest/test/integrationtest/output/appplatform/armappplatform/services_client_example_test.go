//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_Get.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "myResourceGroup", "myservice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armappplatform.ServiceResource{
	// 	Name: to.Ptr("myservice"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armappplatform.ClusterResourceProperties{
	// 		NetworkProfile: &armappplatform.NetworkProfile{
	// 			OutboundIPs: &armappplatform.NetworkProfileOutboundIPs{
	// 				PublicIPs: []*string{
	// 					to.Ptr("20.39.3.173"),
	// 					to.Ptr("40.64.67.13")},
	// 				},
	// 				RequiredTraffics: []*armappplatform.RequiredTraffic{
	// 					{
	// 						Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 						IPs: []*string{
	// 							to.Ptr("20.62.211.25"),
	// 							to.Ptr("52.188.47.226")},
	// 							Port: to.Ptr[int32](443),
	// 							Protocol: to.Ptr("TCP"),
	// 						},
	// 						{
	// 							Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 							IPs: []*string{
	// 								to.Ptr("20.62.211.25"),
	// 								to.Ptr("52.188.47.226")},
	// 								Port: to.Ptr[int32](1194),
	// 								Protocol: to.Ptr("UDP"),
	// 							},
	// 							{
	// 								Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 								IPs: []*string{
	// 									to.Ptr("20.62.211.25"),
	// 									to.Ptr("52.188.47.226")},
	// 									Port: to.Ptr[int32](9000),
	// 									Protocol: to.Ptr("TCP"),
	// 							}},
	// 						},
	// 						ProvisioningState: to.Ptr(armappplatform.ProvisioningStateSucceeded),
	// 						ServiceID: to.Ptr("12345678abcd1234abcd12345678abcd"),
	// 					},
	// 					SKU: &armappplatform.SKU{
	// 						Name: to.Ptr("S0"),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_CreateOrUpdate.json
func ExampleServicesClient_BeginCreateOrUpdate_servicesCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", armappplatform.ServiceResource{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armappplatform.ClusterResourceProperties{},
		SKU: &armappplatform.SKU{
			Name: to.Ptr("S0"),
			Tier: to.Ptr("Standard"),
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
	// res.ServiceResource = armappplatform.ServiceResource{
	// 	Name: to.Ptr("myservice"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armappplatform.ClusterResourceProperties{
	// 		NetworkProfile: &armappplatform.NetworkProfile{
	// 			OutboundIPs: &armappplatform.NetworkProfileOutboundIPs{
	// 				PublicIPs: []*string{
	// 					to.Ptr("20.39.3.173"),
	// 					to.Ptr("40.64.67.13")},
	// 				},
	// 				RequiredTraffics: []*armappplatform.RequiredTraffic{
	// 					{
	// 						Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 						IPs: []*string{
	// 							to.Ptr("20.62.211.25"),
	// 							to.Ptr("52.188.47.226")},
	// 							Port: to.Ptr[int32](443),
	// 							Protocol: to.Ptr("TCP"),
	// 						},
	// 						{
	// 							Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 							IPs: []*string{
	// 								to.Ptr("20.62.211.25"),
	// 								to.Ptr("52.188.47.226")},
	// 								Port: to.Ptr[int32](1194),
	// 								Protocol: to.Ptr("UDP"),
	// 							},
	// 							{
	// 								Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 								IPs: []*string{
	// 									to.Ptr("20.62.211.25"),
	// 									to.Ptr("52.188.47.226")},
	// 									Port: to.Ptr[int32](9000),
	// 									Protocol: to.Ptr("TCP"),
	// 							}},
	// 						},
	// 						ProvisioningState: to.Ptr(armappplatform.ProvisioningStateSucceeded),
	// 						ServiceID: to.Ptr("12345678abcd1234abcd12345678abcd"),
	// 					},
	// 					SKU: &armappplatform.SKU{
	// 						Name: to.Ptr("S0"),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_CreateOrUpdate_VNetInjection.json
func ExampleServicesClient_BeginCreateOrUpdate_servicesCreateOrUpdateVNetInjection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", armappplatform.ServiceResource{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armappplatform.ClusterResourceProperties{
			NetworkProfile: &armappplatform.NetworkProfile{
				AppNetworkResourceGroup:            to.Ptr("my-app-network-rg"),
				AppSubnetID:                        to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/apps"),
				ServiceCidr:                        to.Ptr("10.8.0.0/16,10.244.0.0/16,10.245.0.1/16"),
				ServiceRuntimeNetworkResourceGroup: to.Ptr("my-service-runtime-network-rg"),
				ServiceRuntimeSubnetID:             to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/serviceRuntime"),
			},
		},
		SKU: &armappplatform.SKU{
			Name: to.Ptr("S0"),
			Tier: to.Ptr("Standard"),
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
	// res.ServiceResource = armappplatform.ServiceResource{
	// 	Name: to.Ptr("myservice"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armappplatform.ClusterResourceProperties{
	// 		NetworkProfile: &armappplatform.NetworkProfile{
	// 			AppNetworkResourceGroup: to.Ptr("my-app-network-rg"),
	// 			AppSubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/apps"),
	// 			OutboundIPs: &armappplatform.NetworkProfileOutboundIPs{
	// 				PublicIPs: []*string{
	// 					to.Ptr("40.64.67.13")},
	// 				},
	// 				RequiredTraffics: []*armappplatform.RequiredTraffic{
	// 					{
	// 						Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 						IPs: []*string{
	// 							to.Ptr("20.62.211.25"),
	// 							to.Ptr("52.188.47.226")},
	// 							Port: to.Ptr[int32](443),
	// 							Protocol: to.Ptr("TCP"),
	// 						},
	// 						{
	// 							Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 							IPs: []*string{
	// 								to.Ptr("20.62.211.25"),
	// 								to.Ptr("52.188.47.226")},
	// 								Port: to.Ptr[int32](1194),
	// 								Protocol: to.Ptr("UDP"),
	// 							},
	// 							{
	// 								Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 								IPs: []*string{
	// 									to.Ptr("20.62.211.25"),
	// 									to.Ptr("52.188.47.226")},
	// 									Port: to.Ptr[int32](9000),
	// 									Protocol: to.Ptr("TCP"),
	// 							}},
	// 							ServiceCidr: to.Ptr("10.8.0.0/16,10.244.0.0/16,10.245.0.1/16"),
	// 							ServiceRuntimeNetworkResourceGroup: to.Ptr("my-service-runtime-network-rg"),
	// 							ServiceRuntimeSubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/serviceRuntime"),
	// 						},
	// 						ProvisioningState: to.Ptr(armappplatform.ProvisioningStateSucceeded),
	// 						ServiceID: to.Ptr("12345678abcd1234abcd12345678abcd"),
	// 					},
	// 					SKU: &armappplatform.SKU{
	// 						Name: to.Ptr("S0"),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_Delete.json
func ExampleServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginDelete(ctx, "myResourceGroup", "myservice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_Update.json
func ExampleServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginUpdate(ctx, "myResourceGroup", "myservice", armappplatform.ServiceResource{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armappplatform.ClusterResourceProperties{},
		SKU: &armappplatform.SKU{
			Name: to.Ptr("S0"),
			Tier: to.Ptr("Standard"),
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
	// res.ServiceResource = armappplatform.ServiceResource{
	// 	Name: to.Ptr("myservice"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Properties: &armappplatform.ClusterResourceProperties{
	// 		NetworkProfile: &armappplatform.NetworkProfile{
	// 			OutboundIPs: &armappplatform.NetworkProfileOutboundIPs{
	// 				PublicIPs: []*string{
	// 					to.Ptr("20.39.3.173"),
	// 					to.Ptr("40.64.67.13")},
	// 				},
	// 				RequiredTraffics: []*armappplatform.RequiredTraffic{
	// 					{
	// 						Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 						IPs: []*string{
	// 							to.Ptr("20.62.211.25"),
	// 							to.Ptr("52.188.47.226")},
	// 							Port: to.Ptr[int32](443),
	// 							Protocol: to.Ptr("TCP"),
	// 						},
	// 						{
	// 							Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 							IPs: []*string{
	// 								to.Ptr("20.62.211.25"),
	// 								to.Ptr("52.188.47.226")},
	// 								Port: to.Ptr[int32](1194),
	// 								Protocol: to.Ptr("UDP"),
	// 							},
	// 							{
	// 								Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
	// 								IPs: []*string{
	// 									to.Ptr("20.62.211.25"),
	// 									to.Ptr("52.188.47.226")},
	// 									Port: to.Ptr[int32](9000),
	// 									Protocol: to.Ptr("TCP"),
	// 							}},
	// 						},
	// 						ProvisioningState: to.Ptr(armappplatform.ProvisioningStateSucceeded),
	// 						ServiceID: to.Ptr("12345678abcd1234abcd12345678abcd"),
	// 					},
	// 					SKU: &armappplatform.SKU{
	// 						Name: to.Ptr("S0"),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_ListTestKeys.json
func ExampleServicesClient_ListTestKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().ListTestKeys(ctx, "myResourceGroup", "myservice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TestKeys = armappplatform.TestKeys{
	// 	Enabled: to.Ptr(true),
	// 	PrimaryKey: to.Ptr("<primaryKey>"),
	// 	PrimaryTestEndpoint: to.Ptr("<primaryTestEndpoint>"),
	// 	SecondaryKey: to.Ptr("<secondaryKey>"),
	// 	SecondaryTestEndpoint: to.Ptr("<secondaryTestEndpoint>"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_RegenerateTestKey.json
func ExampleServicesClient_RegenerateTestKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().RegenerateTestKey(ctx, "myResourceGroup", "myservice", armappplatform.RegenerateTestKeyRequestPayload{
		KeyType: to.Ptr(armappplatform.TestKeyTypePrimary),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TestKeys = armappplatform.TestKeys{
	// 	Enabled: to.Ptr(true),
	// 	PrimaryKey: to.Ptr("<primaryKey>"),
	// 	PrimaryTestEndpoint: to.Ptr("<primaryTestEndpoint>"),
	// 	SecondaryKey: to.Ptr("<secondaryKey>"),
	// 	SecondaryTestEndpoint: to.Ptr("<secondaryTestEndpoint>"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_DisableTestEndpoint.json
func ExampleServicesClient_DisableTestEndpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewServicesClient().DisableTestEndpoint(ctx, "myResourceGroup", "myservice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_EnableTestEndpoint.json
func ExampleServicesClient_EnableTestEndpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().EnableTestEndpoint(ctx, "myResourceGroup", "myservice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TestKeys = armappplatform.TestKeys{
	// 	Enabled: to.Ptr(true),
	// 	PrimaryKey: to.Ptr("<primaryKey>"),
	// 	PrimaryTestEndpoint: to.Ptr("<primaryTestEndpoint>"),
	// 	SecondaryKey: to.Ptr("<secondaryKey>"),
	// 	SecondaryTestEndpoint: to.Ptr("<secondaryTestEndpoint>"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_CheckNameAvailability.json
func ExampleServicesClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().CheckNameAvailability(ctx, "eastus", armappplatform.NameAvailabilityParameters{
		Name: to.Ptr("myservice"),
		Type: to.Ptr("Microsoft.AppPlatform/Spring"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NameAvailability = armappplatform.NameAvailability{
	// 	Message: to.Ptr("The name is already used."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr("AlreadyExists"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_ListBySubscription.json
func ExampleServicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListBySubscriptionPager(nil)
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
		// page.ServiceResourceList = armappplatform.ServiceResourceList{
		// 	Value: []*armappplatform.ServiceResource{
		// 		{
		// 			Name: to.Ptr("myservice"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Properties: &armappplatform.ClusterResourceProperties{
		// 				NetworkProfile: &armappplatform.NetworkProfile{
		// 					OutboundIPs: &armappplatform.NetworkProfileOutboundIPs{
		// 						PublicIPs: []*string{
		// 							to.Ptr("20.39.3.173"),
		// 							to.Ptr("40.64.67.13")},
		// 						},
		// 						RequiredTraffics: []*armappplatform.RequiredTraffic{
		// 							{
		// 								Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 								IPs: []*string{
		// 									to.Ptr("20.62.211.25"),
		// 									to.Ptr("52.188.47.226")},
		// 									Port: to.Ptr[int32](443),
		// 									Protocol: to.Ptr("TCP"),
		// 								},
		// 								{
		// 									Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 									IPs: []*string{
		// 										to.Ptr("20.62.211.25"),
		// 										to.Ptr("52.188.47.226")},
		// 										Port: to.Ptr[int32](1194),
		// 										Protocol: to.Ptr("UDP"),
		// 									},
		// 									{
		// 										Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 										IPs: []*string{
		// 											to.Ptr("20.62.211.25"),
		// 											to.Ptr("52.188.47.226")},
		// 											Port: to.Ptr[int32](9000),
		// 											Protocol: to.Ptr("TCP"),
		// 									}},
		// 								},
		// 								ProvisioningState: to.Ptr(armappplatform.ProvisioningStateSucceeded),
		// 								ServiceID: to.Ptr("12345678abcd1234abcd12345678abcd"),
		// 							},
		// 							SKU: &armappplatform.SKU{
		// 								Name: to.Ptr("S0"),
		// 								Tier: to.Ptr("Standard"),
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_List.json
func ExampleServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListPager("myResourceGroup", nil)
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
		// page.ServiceResourceList = armappplatform.ServiceResourceList{
		// 	Value: []*armappplatform.ServiceResource{
		// 		{
		// 			Name: to.Ptr("myservice"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Properties: &armappplatform.ClusterResourceProperties{
		// 				NetworkProfile: &armappplatform.NetworkProfile{
		// 					OutboundIPs: &armappplatform.NetworkProfileOutboundIPs{
		// 						PublicIPs: []*string{
		// 							to.Ptr("20.39.3.173"),
		// 							to.Ptr("40.64.67.13")},
		// 						},
		// 						RequiredTraffics: []*armappplatform.RequiredTraffic{
		// 							{
		// 								Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 								IPs: []*string{
		// 									to.Ptr("20.62.211.25"),
		// 									to.Ptr("52.188.47.226")},
		// 									Port: to.Ptr[int32](443),
		// 									Protocol: to.Ptr("TCP"),
		// 								},
		// 								{
		// 									Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 									IPs: []*string{
		// 										to.Ptr("20.62.211.25"),
		// 										to.Ptr("52.188.47.226")},
		// 										Port: to.Ptr[int32](1194),
		// 										Protocol: to.Ptr("UDP"),
		// 									},
		// 									{
		// 										Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 										IPs: []*string{
		// 											to.Ptr("20.62.211.25"),
		// 											to.Ptr("52.188.47.226")},
		// 											Port: to.Ptr[int32](9000),
		// 											Protocol: to.Ptr("TCP"),
		// 									}},
		// 								},
		// 								ProvisioningState: to.Ptr(armappplatform.ProvisioningStateSucceeded),
		// 								ServiceID: to.Ptr("12345678abcd1234abcd12345678abcd"),
		// 							},
		// 							SKU: &armappplatform.SKU{
		// 								Name: to.Ptr("S0"),
		// 								Tier: to.Ptr("Standard"),
		// 							},
		// 					}},
		// 				}
	}
}
