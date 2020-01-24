# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] [readonly] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] [readonly] 
**OdataId** | **string** | The name of the resource. | [readonly] 
**OdataType** | **string** | The type of a resource. | [readonly] 
**Description** | Pointer to **string** | description | [optional] [readonly] 
**EndTime** | [**time.Time**](time.Time.md) | The date-time stamp that the task was last completed. | [optional] [readonly] 
**HidePayload** | **bool** | Indicates that the contents of the Payload should be hidden from view after the Task has been created.  When set to True, the Payload object will not be returned on GET. | [optional] [readonly] 
**Id** | **string** | The name of the resource. | [readonly] 
**Messages** | [**[]Message**](Message.md) | This is an array of messages associated with the task. | [optional] 
**Name** | **string** | The name of the resource. | [readonly] 
**Oem** | **string** | This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections. | [optional] 
**Payload** | [**Payload**](Payload.md) |  | [optional] 
**StartTime** | [**time.Time**](time.Time.md) | The date-time stamp that the task was last started. | [optional] [readonly] 
**TaskMonitor** | **string** | The URI of the Task Monitor for this task. | [optional] [readonly] 
**TaskState** | [**TaskState**](TaskState.md) |  | [optional] 
**TaskStatus** | [**Health**](Health.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


