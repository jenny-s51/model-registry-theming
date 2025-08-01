-- Seed TypeProperty table
INSERT INTO "TypeProperty" (type_id, name, data_type)
SELECT (SELECT id FROM "Type" WHERE name = 'kf.RegisteredModel'), 'description', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.RegisteredModel'), 'owner', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.RegisteredModel'), 'state', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelVersion'), 'author', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelVersion'), 'description', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelVersion'), 'model_name', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelVersion'), 'state', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelVersion'), 'version', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.DocArtifact'), 'description', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelArtifact'), 'description', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelArtifact'), 'model_format_name', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelArtifact'), 'model_format_version', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelArtifact'), 'service_account_name', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelArtifact'), 'storage_key', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ModelArtifact'), 'storage_path', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ServingEnvironment'), 'description', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.InferenceService'), 'description', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.InferenceService'), 'desired_state', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.InferenceService'), 'model_version_id', 1 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.InferenceService'), 'registered_model_id', 1 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.InferenceService'), 'runtime', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.InferenceService'), 'serving_environment_id', 1 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ServeModel'), 'description', 3 UNION ALL
SELECT (SELECT id FROM "Type" WHERE name = 'kf.ServeModel'), 'model_version_id', 1; 