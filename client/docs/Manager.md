# Manager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The name of the resource. | [optional] 
**Name** | **string** | The name of the resource. | 
**UUID** | **string** |  | [optional] 
**ServiceEntryPointUUID** | **string** |  | [optional] 
**OdataType** | **string** | The type of a resource. | 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**RedfishCopyright** | **string** | redfish copyright | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ManagerType** | [**ManagerType**](ManagerType.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**DateTime** | Pointer to **string** |  | [optional] 
**DateTimeLocalOffset** | Pointer to **string** | The time offset from UTC that the DateTime property is set to in format: +06:00 . | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**EthernetInterfaces** | [**IdRef**](idRef.md) |  | [optional] 
**FirmwareVersion** | Pointer to **string** |  | [optional] 
**Links** | [**ManagerLinks**](ManagerLinks.md) |  | [optional] 
**PowerState** | [**PowerState**](PowerState.md) |  | [optional] 
**VirtualMedia** | [**IdRef**](idRef.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


