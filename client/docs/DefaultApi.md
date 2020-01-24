# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EjectVirtualMedia**](DefaultApi.md#EjectVirtualMedia) | **Post** /redfish/v1/Managers/{managerId}/VirtualMedia/{virtualMediaId}/Actions/VirtualMedia.EjectMedia | 
[**FirmwareInventory**](DefaultApi.md#FirmwareInventory) | **Get** /redfish/v1/UpdateService/FirmwareInventory | 
[**FirmwareInventoryDownloadImage**](DefaultApi.md#FirmwareInventoryDownloadImage) | **Post** /redfish/v1/UpdateService/FirmwareInventory | 
[**GetManager**](DefaultApi.md#GetManager) | **Get** /redfish/v1/Managers/{managerId} | 
[**GetManagerVirtualMedia**](DefaultApi.md#GetManagerVirtualMedia) | **Get** /redfish/v1/Managers/{managerId}/VirtualMedia/{virtualMediaId} | 
[**GetRoot**](DefaultApi.md#GetRoot) | **Get** /redfish/v1 | 
[**GetSystem**](DefaultApi.md#GetSystem) | **Get** /redfish/v1/Systems/{systemId} | 
[**GetTask**](DefaultApi.md#GetTask) | **Get** /redfish/v1/TaskService/Tasks/{taskId} | 
[**InsertVirtualMedia**](DefaultApi.md#InsertVirtualMedia) | **Post** /redfish/v1/Managers/{managerId}/VirtualMedia/{virtualMediaId}/Actions/VirtualMedia.InsertMedia | 
[**ListManagerVirtualMedia**](DefaultApi.md#ListManagerVirtualMedia) | **Get** /redfish/v1/Managers/{managerId}/VirtualMedia | 
[**ListManagers**](DefaultApi.md#ListManagers) | **Get** /redfish/v1/Managers | 
[**ListSystems**](DefaultApi.md#ListSystems) | **Get** /redfish/v1/Systems | 
[**ResetSystem**](DefaultApi.md#ResetSystem) | **Post** /redfish/v1/Systems/{ComputerSystemId}/Actions/ComputerSystem.Reset | 
[**SetSystem**](DefaultApi.md#SetSystem) | **Patch** /redfish/v1/Systems/{systemId} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Get** /redfish/v1/UpdateService | 
[**UpdateServiceSimpleUpdate**](DefaultApi.md#UpdateServiceSimpleUpdate) | **Post** /redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate | 



## EjectVirtualMedia

> RedfishError EjectVirtualMedia(ctx, managerId, virtualMediaId, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string**| ID of resource | 
**virtualMediaId** | **string**| ID of resource | 
**body** | **map[string]interface{}**|  | 

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


## FirmwareInventory

> Collection FirmwareInventory(ctx, )



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


## FirmwareInventoryDownloadImage

> RedfishError FirmwareInventoryDownloadImage(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FirmwareInventoryDownloadImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FirmwareInventoryDownloadImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **softwareImage** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**RedfishError**](RedfishError.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/formdata
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManager

> Manager GetManager(ctx, managerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string**| ID of resource | 

### Return type

[**Manager**](Manager.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagerVirtualMedia

> VirtualMedia GetManagerVirtualMedia(ctx, managerId, virtualMediaId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string**| ID of resource | 
**virtualMediaId** | **string**| ID of resource | 

### Return type

[**VirtualMedia**](VirtualMedia.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetTask

> Task GetTask(ctx, taskId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Task ID | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertVirtualMedia

> RedfishError InsertVirtualMedia(ctx, managerId, virtualMediaId, insertMediaRequestBody)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string**| ID of resource | 
**virtualMediaId** | **string**| ID of resource | 
**insertMediaRequestBody** | [**InsertMediaRequestBody**](InsertMediaRequestBody.md)|  | 

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


## ListManagerVirtualMedia

> Collection ListManagerVirtualMedia(ctx, managerId)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managerId** | **string**| ID of resource | 

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


## SetSystem

> ComputerSystem SetSystem(ctx, systemId, computerSystem)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string**| ID of resource | 
**computerSystem** | [**ComputerSystem**](ComputerSystem.md)|  | 

### Return type

[**ComputerSystem**](ComputerSystem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> UpdateService UpdateService(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UpdateService**](UpdateService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceSimpleUpdate

> RedfishError UpdateServiceSimpleUpdate(ctx, simpleUpdateRequestBody)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simpleUpdateRequestBody** | [**SimpleUpdateRequestBody**](SimpleUpdateRequestBody.md)|  | 

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

