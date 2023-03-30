# PortainerTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Mandatory container/stack fields Template Identifier | [optional] [default to null]
**AdministratorOnly** | **bool** | Whether the template should be available to administrators only | [optional] [default to null]
**Categories** | **[]string** | A list of categories associated to the template | [optional] [default to null]
**Command** | **string** | The command that will be executed in a container template | [optional] [default to null]
**Description** | **string** | Description of the template | [optional] [default to null]
**Env** | [**[]PortainerTemplateEnv**](portainer.TemplateEnv.md) | A list of environment(endpoint) variables used during the template deployment | [optional] [default to null]
**Hostname** | **string** | Container hostname | [optional] [default to null]
**Image** | **string** | Mandatory container fields Image associated to a container template. Mandatory for a container template | [optional] [default to null]
**Interactive** | **bool** | Whether the container should be started in interactive mode (-i -t equivalent on the CLI) | [optional] [default to null]
**Labels** | [**[]PortainerPair**](portainer.Pair.md) | Container labels | [optional] [default to null]
**Logo** | **string** | URL of the template&#39;s logo | [optional] [default to null]
**Name** | **string** | Optional stack/container fields Default name for the stack/container to be used on deployment | [optional] [default to null]
**Network** | **string** | Name of a network that will be used on container deployment if it exists inside the environment(endpoint) | [optional] [default to null]
**Note** | **string** | A note that will be displayed in the UI. Supports HTML content | [optional] [default to null]
**Platform** | **string** | Platform associated to the template. Valid values are: &#39;linux&#39;, &#39;windows&#39; or leave empty for multi-platform | [optional] [default to null]
**Ports** | **[]string** | A list of ports exposed by the container | [optional] [default to null]
**Privileged** | **bool** | Whether the container should be started in privileged mode | [optional] [default to null]
**Registry** | **string** | Optional container fields The URL of a registry associated to the image for a container template | [optional] [default to null]
**Repository** | [***PortainerTemplateRepository**](portainer.TemplateRepository.md) | Mandatory stack fields | [optional] [default to null]
**RestartPolicy** | **string** | Container restart policy | [optional] [default to null]
**StackFile** | **string** | Mandatory Edge stack fields Stack file used for this template | [optional] [default to null]
**Title** | **string** | Title of the template | [optional] [default to null]
**Type_** | **int32** | Template type. Valid values are: 1 (container), 2 (Swarm stack) or 3 (Compose stack) | [optional] [default to null]
**Volumes** | [**[]PortainerTemplateVolume**](portainer.TemplateVolume.md) | A list of volumes used during the container template deployment | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


