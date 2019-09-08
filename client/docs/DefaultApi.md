# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRoot**](DefaultApi.md#GetRoot) | **Get** /redfish/v1 | 
[**GetSystem**](DefaultApi.md#GetSystem) | **Get** /redfish/v1/Systems/{systemId} | 
[**ListManagers**](DefaultApi.md#ListManagers) | **Get** /redfish/v1/Managers | 
[**ListSystems**](DefaultApi.md#ListSystems) | **Get** /redfish/v1/Systems | 
[**ResetSystem**](DefaultApi.md#ResetSystem) | **Post** /redfish/v1/Systems/{ComputerSystemId}/Actions/ComputerSystem.Reset | 



## GetRoot

> Root GetRoot(ctx, )


### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Root**](Root.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystem

> ComputerSystem GetSystem(ctx, systemId)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string**| ID of resource | 

### Return type

[**ComputerSystem**](ComputerSystem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListManagers

> Collection ListManagers(ctx, )


### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystems

> Collection ListSystems(ctx, )


### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSystem

> RedfishError ResetSystem(ctx, computerSystemId, resetRequestBody)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**computerSystemId** | **string**|  | 
**resetRequestBody** | [**ResetRequestBody**](ResetRequestBody.md)|  | 

### Return type

[**RedfishError**](RedfishError.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

