# Inbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | When was the inbox created | [optional] 
**Description** | **string** | Optional description of an inbox for labelling purposes | [optional] 
**EmailAddress** | **string** | The inbox&#39;s email address. Send an email to this address and the inbox will receive and store it for you. To retrieve the email use the Inbox and Email Controller endpoints. | [optional] 
**ExpiresAt** | **string** | When, if ever, will the inbox expire and be deleted. If null then this inbox is permanent and the emails in it won&#39;t be deleted. Timestamp passed as string. | [optional] 
**Favourite** | **bool** | Is the inbox favourited | [optional] 
**Id** | **string** | ID of the inbox | [optional] 
**Name** | **string** | Optional name of the inbox. Displayed in the dashboard for easier search | [optional] 
**Tags** | **[]string** | Tags that inbox has been tagged with | [optional] 
**UserId** | **string** | ID of user that inbox belongs to | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


