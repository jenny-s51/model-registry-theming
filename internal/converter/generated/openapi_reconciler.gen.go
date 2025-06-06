// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	"fmt"
	converter "github.com/kubeflow/model-registry/internal/converter"
	openapi "github.com/kubeflow/model-registry/pkg/openapi"
)

type OpenAPIReconcilerImpl struct{}

func (c *OpenAPIReconcilerImpl) UpdateExistingDocArtifact(source converter.OpenapiUpdateWrapper[openapi.DocArtifact]) (openapi.DocArtifact, error) {
	openapiDocArtifact := converter.InitWithExisting(source)
	var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
	if source.Update != nil {
		pMapStringOpenapiMetadataValue = source.Update.CustomProperties
	}
	if pMapStringOpenapiMetadataValue != nil {
		var mapStringOpenapiMetadataValue map[string]openapi.MetadataValue
		if (*pMapStringOpenapiMetadataValue) != nil {
			mapStringOpenapiMetadataValue = make(map[string]openapi.MetadataValue, len((*pMapStringOpenapiMetadataValue)))
			for key, value := range *pMapStringOpenapiMetadataValue {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
		}
		openapiDocArtifact.CustomProperties = &mapStringOpenapiMetadataValue
	}
	var pString *string
	if source.Update != nil {
		pString = source.Update.Description
	}
	if pString != nil {
		xstring := *pString
		openapiDocArtifact.Description = &xstring
	}
	var pString2 *string
	if source.Update != nil {
		pString2 = source.Update.ExternalId
	}
	if pString2 != nil {
		xstring2 := *pString2
		openapiDocArtifact.ExternalId = &xstring2
	}
	var pString3 *string
	if source.Update != nil {
		pString3 = source.Update.Uri
	}
	if pString3 != nil {
		xstring3 := *pString3
		openapiDocArtifact.Uri = &xstring3
	}
	var pOpenapiArtifactState *openapi.ArtifactState
	if source.Update != nil {
		pOpenapiArtifactState = source.Update.State
	}
	if pOpenapiArtifactState != nil {
		openapiArtifactState, err := c.openapiArtifactStateToOpenapiArtifactState(*pOpenapiArtifactState)
		if err != nil {
			return openapiDocArtifact, fmt.Errorf("error setting field State: %w", err)
		}
		openapiDocArtifact.State = &openapiArtifactState
	}
	return openapiDocArtifact, nil
}
func (c *OpenAPIReconcilerImpl) UpdateExistingInferenceService(source converter.OpenapiUpdateWrapper[openapi.InferenceService]) (openapi.InferenceService, error) {
	openapiInferenceService := converter.InitWithExisting(source)
	var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
	if source.Update != nil {
		pMapStringOpenapiMetadataValue = source.Update.CustomProperties
	}
	if pMapStringOpenapiMetadataValue != nil {
		var mapStringOpenapiMetadataValue map[string]openapi.MetadataValue
		if (*pMapStringOpenapiMetadataValue) != nil {
			mapStringOpenapiMetadataValue = make(map[string]openapi.MetadataValue, len((*pMapStringOpenapiMetadataValue)))
			for key, value := range *pMapStringOpenapiMetadataValue {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
		}
		openapiInferenceService.CustomProperties = &mapStringOpenapiMetadataValue
	}
	var pString *string
	if source.Update != nil {
		pString = source.Update.Description
	}
	if pString != nil {
		xstring := *pString
		openapiInferenceService.Description = &xstring
	}
	var pString2 *string
	if source.Update != nil {
		pString2 = source.Update.ExternalId
	}
	if pString2 != nil {
		xstring2 := *pString2
		openapiInferenceService.ExternalId = &xstring2
	}
	var pString3 *string
	if source.Update != nil {
		pString3 = source.Update.ModelVersionId
	}
	if pString3 != nil {
		xstring3 := *pString3
		openapiInferenceService.ModelVersionId = &xstring3
	}
	var pString4 *string
	if source.Update != nil {
		pString4 = source.Update.Runtime
	}
	if pString4 != nil {
		xstring4 := *pString4
		openapiInferenceService.Runtime = &xstring4
	}
	var pOpenapiInferenceServiceState *openapi.InferenceServiceState
	if source.Update != nil {
		pOpenapiInferenceServiceState = source.Update.DesiredState
	}
	if pOpenapiInferenceServiceState != nil {
		openapiInferenceServiceState, err := c.openapiInferenceServiceStateToOpenapiInferenceServiceState(*pOpenapiInferenceServiceState)
		if err != nil {
			return openapiInferenceService, fmt.Errorf("error setting field DesiredState: %w", err)
		}
		openapiInferenceService.DesiredState = &openapiInferenceServiceState
	}
	return openapiInferenceService, nil
}
func (c *OpenAPIReconcilerImpl) UpdateExistingModelArtifact(source converter.OpenapiUpdateWrapper[openapi.ModelArtifact]) (openapi.ModelArtifact, error) {
	openapiModelArtifact := converter.InitWithExisting(source)
	var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
	if source.Update != nil {
		pMapStringOpenapiMetadataValue = source.Update.CustomProperties
	}
	if pMapStringOpenapiMetadataValue != nil {
		var mapStringOpenapiMetadataValue map[string]openapi.MetadataValue
		if (*pMapStringOpenapiMetadataValue) != nil {
			mapStringOpenapiMetadataValue = make(map[string]openapi.MetadataValue, len((*pMapStringOpenapiMetadataValue)))
			for key, value := range *pMapStringOpenapiMetadataValue {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
		}
		openapiModelArtifact.CustomProperties = &mapStringOpenapiMetadataValue
	}
	var pString *string
	if source.Update != nil {
		pString = source.Update.Description
	}
	if pString != nil {
		xstring := *pString
		openapiModelArtifact.Description = &xstring
	}
	var pString2 *string
	if source.Update != nil {
		pString2 = source.Update.ExternalId
	}
	if pString2 != nil {
		xstring2 := *pString2
		openapiModelArtifact.ExternalId = &xstring2
	}
	var pString3 *string
	if source.Update != nil {
		pString3 = source.Update.ModelFormatName
	}
	if pString3 != nil {
		xstring3 := *pString3
		openapiModelArtifact.ModelFormatName = &xstring3
	}
	var pString4 *string
	if source.Update != nil {
		pString4 = source.Update.StorageKey
	}
	if pString4 != nil {
		xstring4 := *pString4
		openapiModelArtifact.StorageKey = &xstring4
	}
	var pString5 *string
	if source.Update != nil {
		pString5 = source.Update.StoragePath
	}
	if pString5 != nil {
		xstring5 := *pString5
		openapiModelArtifact.StoragePath = &xstring5
	}
	var pString6 *string
	if source.Update != nil {
		pString6 = source.Update.ModelFormatVersion
	}
	if pString6 != nil {
		xstring6 := *pString6
		openapiModelArtifact.ModelFormatVersion = &xstring6
	}
	var pString7 *string
	if source.Update != nil {
		pString7 = source.Update.ServiceAccountName
	}
	if pString7 != nil {
		xstring7 := *pString7
		openapiModelArtifact.ServiceAccountName = &xstring7
	}
	var pString8 *string
	if source.Update != nil {
		pString8 = source.Update.ModelSourceKind
	}
	if pString8 != nil {
		xstring8 := *pString8
		openapiModelArtifact.ModelSourceKind = &xstring8
	}
	var pString9 *string
	if source.Update != nil {
		pString9 = source.Update.ModelSourceClass
	}
	if pString9 != nil {
		xstring9 := *pString9
		openapiModelArtifact.ModelSourceClass = &xstring9
	}
	var pString10 *string
	if source.Update != nil {
		pString10 = source.Update.ModelSourceGroup
	}
	if pString10 != nil {
		xstring10 := *pString10
		openapiModelArtifact.ModelSourceGroup = &xstring10
	}
	var pString11 *string
	if source.Update != nil {
		pString11 = source.Update.ModelSourceId
	}
	if pString11 != nil {
		xstring11 := *pString11
		openapiModelArtifact.ModelSourceId = &xstring11
	}
	var pString12 *string
	if source.Update != nil {
		pString12 = source.Update.ModelSourceName
	}
	if pString12 != nil {
		xstring12 := *pString12
		openapiModelArtifact.ModelSourceName = &xstring12
	}
	var pString13 *string
	if source.Update != nil {
		pString13 = source.Update.Uri
	}
	if pString13 != nil {
		xstring13 := *pString13
		openapiModelArtifact.Uri = &xstring13
	}
	var pOpenapiArtifactState *openapi.ArtifactState
	if source.Update != nil {
		pOpenapiArtifactState = source.Update.State
	}
	if pOpenapiArtifactState != nil {
		openapiArtifactState, err := c.openapiArtifactStateToOpenapiArtifactState(*pOpenapiArtifactState)
		if err != nil {
			return openapiModelArtifact, fmt.Errorf("error setting field State: %w", err)
		}
		openapiModelArtifact.State = &openapiArtifactState
	}
	return openapiModelArtifact, nil
}
func (c *OpenAPIReconcilerImpl) UpdateExistingModelVersion(source converter.OpenapiUpdateWrapper[openapi.ModelVersion]) (openapi.ModelVersion, error) {
	openapiModelVersion := converter.InitWithExisting(source)
	var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
	if source.Update != nil {
		pMapStringOpenapiMetadataValue = source.Update.CustomProperties
	}
	if pMapStringOpenapiMetadataValue != nil {
		var mapStringOpenapiMetadataValue map[string]openapi.MetadataValue
		if (*pMapStringOpenapiMetadataValue) != nil {
			mapStringOpenapiMetadataValue = make(map[string]openapi.MetadataValue, len((*pMapStringOpenapiMetadataValue)))
			for key, value := range *pMapStringOpenapiMetadataValue {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
		}
		openapiModelVersion.CustomProperties = &mapStringOpenapiMetadataValue
	}
	var pString *string
	if source.Update != nil {
		pString = source.Update.Description
	}
	if pString != nil {
		xstring := *pString
		openapiModelVersion.Description = &xstring
	}
	var pString2 *string
	if source.Update != nil {
		pString2 = source.Update.ExternalId
	}
	if pString2 != nil {
		xstring2 := *pString2
		openapiModelVersion.ExternalId = &xstring2
	}
	var pOpenapiModelVersionState *openapi.ModelVersionState
	if source.Update != nil {
		pOpenapiModelVersionState = source.Update.State
	}
	if pOpenapiModelVersionState != nil {
		openapiModelVersionState, err := c.openapiModelVersionStateToOpenapiModelVersionState(*pOpenapiModelVersionState)
		if err != nil {
			return openapiModelVersion, fmt.Errorf("error setting field State: %w", err)
		}
		openapiModelVersion.State = &openapiModelVersionState
	}
	var pString3 *string
	if source.Update != nil {
		pString3 = source.Update.Author
	}
	if pString3 != nil {
		xstring3 := *pString3
		openapiModelVersion.Author = &xstring3
	}
	return openapiModelVersion, nil
}
func (c *OpenAPIReconcilerImpl) UpdateExistingRegisteredModel(source converter.OpenapiUpdateWrapper[openapi.RegisteredModel]) (openapi.RegisteredModel, error) {
	openapiRegisteredModel := converter.InitWithExisting(source)
	var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
	if source.Update != nil {
		pMapStringOpenapiMetadataValue = source.Update.CustomProperties
	}
	if pMapStringOpenapiMetadataValue != nil {
		var mapStringOpenapiMetadataValue map[string]openapi.MetadataValue
		if (*pMapStringOpenapiMetadataValue) != nil {
			mapStringOpenapiMetadataValue = make(map[string]openapi.MetadataValue, len((*pMapStringOpenapiMetadataValue)))
			for key, value := range *pMapStringOpenapiMetadataValue {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
		}
		openapiRegisteredModel.CustomProperties = &mapStringOpenapiMetadataValue
	}
	var pString *string
	if source.Update != nil {
		pString = source.Update.Description
	}
	if pString != nil {
		xstring := *pString
		openapiRegisteredModel.Description = &xstring
	}
	var pString2 *string
	if source.Update != nil {
		pString2 = source.Update.ExternalId
	}
	if pString2 != nil {
		xstring2 := *pString2
		openapiRegisteredModel.ExternalId = &xstring2
	}
	var pString3 *string
	if source.Update != nil {
		pString3 = source.Update.Owner
	}
	if pString3 != nil {
		xstring3 := *pString3
		openapiRegisteredModel.Owner = &xstring3
	}
	var pOpenapiRegisteredModelState *openapi.RegisteredModelState
	if source.Update != nil {
		pOpenapiRegisteredModelState = source.Update.State
	}
	if pOpenapiRegisteredModelState != nil {
		openapiRegisteredModelState, err := c.openapiRegisteredModelStateToOpenapiRegisteredModelState(*pOpenapiRegisteredModelState)
		if err != nil {
			return openapiRegisteredModel, fmt.Errorf("error setting field State: %w", err)
		}
		openapiRegisteredModel.State = &openapiRegisteredModelState
	}
	return openapiRegisteredModel, nil
}
func (c *OpenAPIReconcilerImpl) UpdateExistingServeModel(source converter.OpenapiUpdateWrapper[openapi.ServeModel]) (openapi.ServeModel, error) {
	openapiServeModel := converter.InitWithExisting(source)
	var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
	if source.Update != nil {
		pMapStringOpenapiMetadataValue = source.Update.CustomProperties
	}
	if pMapStringOpenapiMetadataValue != nil {
		var mapStringOpenapiMetadataValue map[string]openapi.MetadataValue
		if (*pMapStringOpenapiMetadataValue) != nil {
			mapStringOpenapiMetadataValue = make(map[string]openapi.MetadataValue, len((*pMapStringOpenapiMetadataValue)))
			for key, value := range *pMapStringOpenapiMetadataValue {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
		}
		openapiServeModel.CustomProperties = &mapStringOpenapiMetadataValue
	}
	var pString *string
	if source.Update != nil {
		pString = source.Update.Description
	}
	if pString != nil {
		xstring := *pString
		openapiServeModel.Description = &xstring
	}
	var pString2 *string
	if source.Update != nil {
		pString2 = source.Update.ExternalId
	}
	if pString2 != nil {
		xstring2 := *pString2
		openapiServeModel.ExternalId = &xstring2
	}
	var pOpenapiExecutionState *openapi.ExecutionState
	if source.Update != nil {
		pOpenapiExecutionState = source.Update.LastKnownState
	}
	if pOpenapiExecutionState != nil {
		openapiExecutionState, err := c.openapiExecutionStateToOpenapiExecutionState(*pOpenapiExecutionState)
		if err != nil {
			return openapiServeModel, fmt.Errorf("error setting field LastKnownState: %w", err)
		}
		openapiServeModel.LastKnownState = &openapiExecutionState
	}
	return openapiServeModel, nil
}
func (c *OpenAPIReconcilerImpl) UpdateExistingServingEnvironment(source converter.OpenapiUpdateWrapper[openapi.ServingEnvironment]) (openapi.ServingEnvironment, error) {
	openapiServingEnvironment := converter.InitWithExisting(source)
	var pMapStringOpenapiMetadataValue *map[string]openapi.MetadataValue
	if source.Update != nil {
		pMapStringOpenapiMetadataValue = source.Update.CustomProperties
	}
	if pMapStringOpenapiMetadataValue != nil {
		var mapStringOpenapiMetadataValue map[string]openapi.MetadataValue
		if (*pMapStringOpenapiMetadataValue) != nil {
			mapStringOpenapiMetadataValue = make(map[string]openapi.MetadataValue, len((*pMapStringOpenapiMetadataValue)))
			for key, value := range *pMapStringOpenapiMetadataValue {
				mapStringOpenapiMetadataValue[key] = c.openapiMetadataValueToOpenapiMetadataValue(value)
			}
		}
		openapiServingEnvironment.CustomProperties = &mapStringOpenapiMetadataValue
	}
	var pString *string
	if source.Update != nil {
		pString = source.Update.Description
	}
	if pString != nil {
		xstring := *pString
		openapiServingEnvironment.Description = &xstring
	}
	var pString2 *string
	if source.Update != nil {
		pString2 = source.Update.ExternalId
	}
	if pString2 != nil {
		xstring2 := *pString2
		openapiServingEnvironment.ExternalId = &xstring2
	}
	return openapiServingEnvironment, nil
}
func (c *OpenAPIReconcilerImpl) openapiArtifactStateToOpenapiArtifactState(source openapi.ArtifactState) (openapi.ArtifactState, error) {
	var openapiArtifactState openapi.ArtifactState
	switch source {
	case openapi.ARTIFACTSTATE_ABANDONED:
		openapiArtifactState = openapi.ARTIFACTSTATE_ABANDONED
	case openapi.ARTIFACTSTATE_DELETED:
		openapiArtifactState = openapi.ARTIFACTSTATE_DELETED
	case openapi.ARTIFACTSTATE_LIVE:
		openapiArtifactState = openapi.ARTIFACTSTATE_LIVE
	case openapi.ARTIFACTSTATE_MARKED_FOR_DELETION:
		openapiArtifactState = openapi.ARTIFACTSTATE_MARKED_FOR_DELETION
	case openapi.ARTIFACTSTATE_PENDING:
		openapiArtifactState = openapi.ARTIFACTSTATE_PENDING
	case openapi.ARTIFACTSTATE_REFERENCE:
		openapiArtifactState = openapi.ARTIFACTSTATE_REFERENCE
	case openapi.ARTIFACTSTATE_UNKNOWN:
		openapiArtifactState = openapi.ARTIFACTSTATE_UNKNOWN
	default:
		return openapiArtifactState, fmt.Errorf("unexpected enum element: %v", source)
	}
	return openapiArtifactState, nil
}
func (c *OpenAPIReconcilerImpl) openapiExecutionStateToOpenapiExecutionState(source openapi.ExecutionState) (openapi.ExecutionState, error) {
	var openapiExecutionState openapi.ExecutionState
	switch source {
	case openapi.EXECUTIONSTATE_CACHED:
		openapiExecutionState = openapi.EXECUTIONSTATE_CACHED
	case openapi.EXECUTIONSTATE_CANCELED:
		openapiExecutionState = openapi.EXECUTIONSTATE_CANCELED
	case openapi.EXECUTIONSTATE_COMPLETE:
		openapiExecutionState = openapi.EXECUTIONSTATE_COMPLETE
	case openapi.EXECUTIONSTATE_FAILED:
		openapiExecutionState = openapi.EXECUTIONSTATE_FAILED
	case openapi.EXECUTIONSTATE_NEW:
		openapiExecutionState = openapi.EXECUTIONSTATE_NEW
	case openapi.EXECUTIONSTATE_RUNNING:
		openapiExecutionState = openapi.EXECUTIONSTATE_RUNNING
	case openapi.EXECUTIONSTATE_UNKNOWN:
		openapiExecutionState = openapi.EXECUTIONSTATE_UNKNOWN
	default:
		return openapiExecutionState, fmt.Errorf("unexpected enum element: %v", source)
	}
	return openapiExecutionState, nil
}
func (c *OpenAPIReconcilerImpl) openapiInferenceServiceStateToOpenapiInferenceServiceState(source openapi.InferenceServiceState) (openapi.InferenceServiceState, error) {
	var openapiInferenceServiceState openapi.InferenceServiceState
	switch source {
	case openapi.INFERENCESERVICESTATE_DEPLOYED:
		openapiInferenceServiceState = openapi.INFERENCESERVICESTATE_DEPLOYED
	case openapi.INFERENCESERVICESTATE_UNDEPLOYED:
		openapiInferenceServiceState = openapi.INFERENCESERVICESTATE_UNDEPLOYED
	default:
		return openapiInferenceServiceState, fmt.Errorf("unexpected enum element: %v", source)
	}
	return openapiInferenceServiceState, nil
}
func (c *OpenAPIReconcilerImpl) openapiMetadataValueToOpenapiMetadataValue(source openapi.MetadataValue) openapi.MetadataValue {
	var openapiMetadataValue openapi.MetadataValue
	openapiMetadataValue.MetadataBoolValue = c.pOpenapiMetadataBoolValueToPOpenapiMetadataBoolValue(source.MetadataBoolValue)
	openapiMetadataValue.MetadataDoubleValue = c.pOpenapiMetadataDoubleValueToPOpenapiMetadataDoubleValue(source.MetadataDoubleValue)
	openapiMetadataValue.MetadataIntValue = c.pOpenapiMetadataIntValueToPOpenapiMetadataIntValue(source.MetadataIntValue)
	openapiMetadataValue.MetadataProtoValue = c.pOpenapiMetadataProtoValueToPOpenapiMetadataProtoValue(source.MetadataProtoValue)
	openapiMetadataValue.MetadataStringValue = c.pOpenapiMetadataStringValueToPOpenapiMetadataStringValue(source.MetadataStringValue)
	openapiMetadataValue.MetadataStructValue = c.pOpenapiMetadataStructValueToPOpenapiMetadataStructValue(source.MetadataStructValue)
	return openapiMetadataValue
}
func (c *OpenAPIReconcilerImpl) openapiModelVersionStateToOpenapiModelVersionState(source openapi.ModelVersionState) (openapi.ModelVersionState, error) {
	var openapiModelVersionState openapi.ModelVersionState
	switch source {
	case openapi.MODELVERSIONSTATE_ARCHIVED:
		openapiModelVersionState = openapi.MODELVERSIONSTATE_ARCHIVED
	case openapi.MODELVERSIONSTATE_LIVE:
		openapiModelVersionState = openapi.MODELVERSIONSTATE_LIVE
	default:
		return openapiModelVersionState, fmt.Errorf("unexpected enum element: %v", source)
	}
	return openapiModelVersionState, nil
}
func (c *OpenAPIReconcilerImpl) openapiRegisteredModelStateToOpenapiRegisteredModelState(source openapi.RegisteredModelState) (openapi.RegisteredModelState, error) {
	var openapiRegisteredModelState openapi.RegisteredModelState
	switch source {
	case openapi.REGISTEREDMODELSTATE_ARCHIVED:
		openapiRegisteredModelState = openapi.REGISTEREDMODELSTATE_ARCHIVED
	case openapi.REGISTEREDMODELSTATE_LIVE:
		openapiRegisteredModelState = openapi.REGISTEREDMODELSTATE_LIVE
	default:
		return openapiRegisteredModelState, fmt.Errorf("unexpected enum element: %v", source)
	}
	return openapiRegisteredModelState, nil
}
func (c *OpenAPIReconcilerImpl) pOpenapiMetadataBoolValueToPOpenapiMetadataBoolValue(source *openapi.MetadataBoolValue) *openapi.MetadataBoolValue {
	var pOpenapiMetadataBoolValue *openapi.MetadataBoolValue
	if source != nil {
		var openapiMetadataBoolValue openapi.MetadataBoolValue
		openapiMetadataBoolValue.BoolValue = (*source).BoolValue
		openapiMetadataBoolValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataBoolValue = &openapiMetadataBoolValue
	}
	return pOpenapiMetadataBoolValue
}
func (c *OpenAPIReconcilerImpl) pOpenapiMetadataDoubleValueToPOpenapiMetadataDoubleValue(source *openapi.MetadataDoubleValue) *openapi.MetadataDoubleValue {
	var pOpenapiMetadataDoubleValue *openapi.MetadataDoubleValue
	if source != nil {
		var openapiMetadataDoubleValue openapi.MetadataDoubleValue
		openapiMetadataDoubleValue.DoubleValue = (*source).DoubleValue
		openapiMetadataDoubleValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataDoubleValue = &openapiMetadataDoubleValue
	}
	return pOpenapiMetadataDoubleValue
}
func (c *OpenAPIReconcilerImpl) pOpenapiMetadataIntValueToPOpenapiMetadataIntValue(source *openapi.MetadataIntValue) *openapi.MetadataIntValue {
	var pOpenapiMetadataIntValue *openapi.MetadataIntValue
	if source != nil {
		var openapiMetadataIntValue openapi.MetadataIntValue
		openapiMetadataIntValue.IntValue = (*source).IntValue
		openapiMetadataIntValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataIntValue = &openapiMetadataIntValue
	}
	return pOpenapiMetadataIntValue
}
func (c *OpenAPIReconcilerImpl) pOpenapiMetadataProtoValueToPOpenapiMetadataProtoValue(source *openapi.MetadataProtoValue) *openapi.MetadataProtoValue {
	var pOpenapiMetadataProtoValue *openapi.MetadataProtoValue
	if source != nil {
		var openapiMetadataProtoValue openapi.MetadataProtoValue
		openapiMetadataProtoValue.Type = (*source).Type
		openapiMetadataProtoValue.ProtoValue = (*source).ProtoValue
		openapiMetadataProtoValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataProtoValue = &openapiMetadataProtoValue
	}
	return pOpenapiMetadataProtoValue
}
func (c *OpenAPIReconcilerImpl) pOpenapiMetadataStringValueToPOpenapiMetadataStringValue(source *openapi.MetadataStringValue) *openapi.MetadataStringValue {
	var pOpenapiMetadataStringValue *openapi.MetadataStringValue
	if source != nil {
		var openapiMetadataStringValue openapi.MetadataStringValue
		openapiMetadataStringValue.StringValue = (*source).StringValue
		openapiMetadataStringValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataStringValue = &openapiMetadataStringValue
	}
	return pOpenapiMetadataStringValue
}
func (c *OpenAPIReconcilerImpl) pOpenapiMetadataStructValueToPOpenapiMetadataStructValue(source *openapi.MetadataStructValue) *openapi.MetadataStructValue {
	var pOpenapiMetadataStructValue *openapi.MetadataStructValue
	if source != nil {
		var openapiMetadataStructValue openapi.MetadataStructValue
		openapiMetadataStructValue.StructValue = (*source).StructValue
		openapiMetadataStructValue.MetadataType = (*source).MetadataType
		pOpenapiMetadataStructValue = &openapiMetadataStructValue
	}
	return pOpenapiMetadataStructValue
}
