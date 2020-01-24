# Manager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The name of the resource. | [optional] [readonly] 
**Name** | **string** | The name of the resource. | [readonly] 
**UUID** | **string** |  | [optional] 
**ServiceEntryPointUUID** | **string** |  | [optional] 
**OdataType** | **string** | The type of a resource. | [readonly] 
**OdataId** | **string** | The unique identifier for a resource. | [readonly] 
**OdataContext** | **string** | The OData description of a payload. | [optional] [readonly] 
**RedfishCopyright** | **string** | redfish copyright | [optional] 
**Model** | Pointer to **string** |  | [optional] [readonly] 
**ManagerType** | [**ManagerType**](ManagerType.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**DateTime** | Pointer to **string** |  | [optional] 
**DateTimeLocalOffset** | Pointer to **string** | The time offset from UTC that the DateTime property is set to in format: +06:00 . | [optional] 
**Description** | Pointer to **string** | description | [optional] [readonly] 
**EthernetInterfaces** | [**IdRef**](idRef.md) |  | [optional] 
**FirmwareVersion** | Pointer to **string** |  | [optional] [readonly] 
**Links** | [**ManagerLinks**](ManagerLinks.md) |  | [optional] 
**PowerState** | [**PowerState**](PowerState.md) |  | [optional] 
**VirtualMedia** | [**IdRef**](idRef.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


