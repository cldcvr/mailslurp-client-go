# UpdateInboxOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Optional description of an inbox for labelling purposes | [optional] 
**ExpiresAt** | [**time.Time**](time.Time.md) | When, if ever, will the inbox expire and be deleted. If null then this inbox is permanent and the emails in it won&#39;t be deleted. Timestamp passed as string. | [optional] 
**Favourite** | **bool** | Is the inbox favourited | [optional] 
**Name** | **string** | Optional name of the inbox. Displayed in the dashboard for easier search | [optional] 
**Tags** | **[]string** | Tags that inbox has been tagged with | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


