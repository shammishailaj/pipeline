# CreateUpdateDeploymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | **string** | Version of the deployment. If not specified, the latest version is used. | [optional] 
**Package** | **string** | The chart content packaged by &#x60;helm package&#x60;. If specified chart version is ignored. | [optional] 
**Namespace** | **string** |  | [optional] 
**ReleaseName** | **string** |  | [optional] 
**ReuseValues** | **bool** |  | [optional] 
**Values** | [**map[string]interface{}**](map[string]interface{}.md) | current values of the deployment | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


