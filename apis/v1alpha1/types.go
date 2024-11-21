// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Indicates that an Amazon S3 canned ACL should be set to control ownership
// of stored query results. When Athena stores query results in Amazon S3, the
// canned ACL is set with the x-amz-acl request header. For more information
// about S3 Object Ownership, see Object Ownership settings (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html#object-ownership-overview)
// in the Amazon S3 User Guide.
type ACLConfiguration struct {
	S3ACLOption *string `json:"s3ACLOption,omitempty"`
}

// Contains the application runtime IDs and their supported DPU sizes.
type ApplicationDPUSizes struct {
	ApplicationRuntimeID *string `json:"applicationRuntimeID,omitempty"`
}

// Provides information about an Athena query error. The AthenaError feature
// provides standardized error information to help you understand failed queries
// and take steps after a query failure occurs. AthenaError includes an ErrorCategory
// field that specifies whether the cause of the failed query is due to system
// error, user error, or other error.
type AthenaError struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// Contains statistics for a notebook calculation.
type CalculationStatistics struct {
	Progress *string `json:"progress,omitempty"`
}

// Contains information about the status of a notebook calculation.
type CalculationStatus struct {
	CompletionDateTime *metav1.Time `json:"completionDateTime,omitempty"`
	StateChangeReason  *string      `json:"stateChangeReason,omitempty"`
	SubmissionDateTime *metav1.Time `json:"submissionDateTime,omitempty"`
}

// Summary information for a notebook calculation.
type CalculationSummary struct {
	Description *string `json:"description,omitempty"`
}

// Contains the submission time of a single allocation request for a capacity
// reservation and the most recent status of the attempted allocation.
type CapacityAllocation struct {
	StatusMessage *string `json:"statusMessage,omitempty"`
}

// Contains metadata for a column in a table.
type Column struct {
	Name *string `json:"name,omitempty"`
}

// Information about the columns in a query execution result.
type ColumnInfo struct {
	CatalogName *string `json:"catalogName,omitempty"`
	Label       *string `json:"label,omitempty"`
	Name        *string `json:"name,omitempty"`
	SchemaName  *string `json:"schemaName,omitempty"`
	TableName   *string `json:"tableName,omitempty"`
	Type        *string `json:"type_,omitempty"`
}

// Specifies the customer managed KMS key that is used to encrypt the user's
// data stores in Athena. When an Amazon Web Services managed key is used, this
// value is null. This setting does not apply to Athena SQL workgroups.
type CustomerContentEncryptionConfiguration struct {
	KMSKey *string `json:"kmsKey,omitempty"`
}

// Contains information about a data catalog in an Amazon Web Services account.
//
// In the Athena console, data catalogs are listed as "data sources" on the
// Data sources page under the Data source name column.
type DataCatalog struct {
	Description *string `json:"description,omitempty"`
}

// Contains metadata information for a database in a data catalog.
type Database struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// If query and calculation results are encrypted in Amazon S3, indicates the
// encryption option used (for example, SSE_KMS or CSE_KMS) and key information.
type EncryptionConfiguration struct {
	EncryptionOption *string `json:"encryptionOption,omitempty"`
	KMSKey           *string `json:"kmsKey,omitempty"`
}

// The Athena engine version for running queries, or the PySpark engine version
// for running sessions.
type EngineVersion struct {
	EffectiveEngineVersion *string `json:"effectiveEngineVersion,omitempty"`
	SelectedEngineVersion  *string `json:"selectedEngineVersion,omitempty"`
}

// Specifies whether the workgroup is IAM Identity Center supported.
type IdentityCenterConfiguration struct {
	EnableIdentityCenter      *bool   `json:"enableIdentityCenter,omitempty"`
	IdentityCenterInstanceARN *string `json:"identityCenterInstanceARN,omitempty"`
}

// A query, where QueryString contains the SQL statements that make up the query.
type NamedQuery struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	QueryString *string `json:"queryString,omitempty"`
	WorkGroup   *string `json:"workGroup,omitempty"`
}

// Contains metadata for notebook, including the notebook name, ID, workgroup,
// and time created.
type NotebookMetadata struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	WorkGroup        *string      `json:"workGroup,omitempty"`
}

// Contains the notebook session ID and notebook session creation time.
type NotebookSessionSummary struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
}

// The name and last modified time of the prepared statement.
type PreparedStatementSummary struct {
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	StatementName    *string      `json:"statementName,omitempty"`
}

// A prepared SQL statement for use with Athena.
type PreparedStatement_SDK struct {
	Description      *string      `json:"description,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	QueryStatement   *string      `json:"queryStatement,omitempty"`
	StatementName    *string      `json:"statementName,omitempty"`
	WorkGroupName    *string      `json:"workGroupName,omitempty"`
}

// Information about a single instance of a query execution.
type QueryExecution struct {
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion *EngineVersion `json:"engineVersion,omitempty"`
	Query         *string        `json:"query,omitempty"`
	// Specifies whether Amazon S3 access grants are enabled for query results.
	QueryResultsS3AccessGrantsConfiguration *QueryResultsS3AccessGrantsConfiguration `json:"queryResultsS3AccessGrantsConfiguration,omitempty"`
	// The location in Amazon S3 where query and calculation results are stored
	// and the encryption option, if any, used for query and calculation results.
	// These are known as "client-side settings". If workgroup settings override
	// client-side settings, then the query uses the workgroup settings.
	ResultConfiguration *ResultConfiguration `json:"resultConfiguration,omitempty"`
	SubstatementType    *string              `json:"substatementType,omitempty"`
	WorkGroup           *string              `json:"workGroup,omitempty"`
}

// The amount of data scanned during the query execution and the amount of time
// that it took to execute, and the type of statement that was run.
type QueryExecutionStatistics struct {
	DataManifestLocation *string `json:"dataManifestLocation,omitempty"`
}

// The completion date, current state, submission time, and state change reason
// (if applicable) for the query execution.
type QueryExecutionStatus struct {
	CompletionDateTime *metav1.Time `json:"completionDateTime,omitempty"`
	StateChangeReason  *string      `json:"stateChangeReason,omitempty"`
	SubmissionDateTime *metav1.Time `json:"submissionDateTime,omitempty"`
}

// Specifies whether Amazon S3 access grants are enabled for query results.
type QueryResultsS3AccessGrantsConfiguration struct {
	AuthenticationType    *string `json:"authenticationType,omitempty"`
	CreateUserLevelPrefix *bool   `json:"createUserLevelPrefix,omitempty"`
	EnableS3AccessGrants  *bool   `json:"enableS3AccessGrants,omitempty"`
}

// Stage statistics such as input and output rows and bytes, execution time
// and stage state. This information also includes substages and the query stage
// plan.
type QueryStage struct {
	State *string `json:"state,omitempty"`
}

// Stage plan information such as name, identifier, sub plans, and remote sources.
type QueryStagePlanNode struct {
	Identifier *string `json:"identifier,omitempty"`
	Name       *string `json:"name,omitempty"`
}

// The location in Amazon S3 where query and calculation results are stored
// and the encryption option, if any, used for query and calculation results.
// These are known as "client-side settings". If workgroup settings override
// client-side settings, then the query uses the workgroup settings.
type ResultConfiguration struct {
	// Indicates that an Amazon S3 canned ACL should be set to control ownership
	// of stored query results. When Athena stores query results in Amazon S3, the
	// canned ACL is set with the x-amz-acl request header. For more information
	// about S3 Object Ownership, see Object Ownership settings (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html#object-ownership-overview)
	// in the Amazon S3 User Guide.
	ACLConfiguration *ACLConfiguration `json:"aclConfiguration,omitempty"`
	// If query and calculation results are encrypted in Amazon S3, indicates the
	// encryption option used (for example, SSE_KMS or CSE_KMS) and key information.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
	ExpectedBucketOwner     *string                  `json:"expectedBucketOwner,omitempty"`
	OutputLocation          *string                  `json:"outputLocation,omitempty"`
}

// The information about the updates in the query results, such as output location
// and encryption configuration for the query results.
type ResultConfigurationUpdates struct {
	// Indicates that an Amazon S3 canned ACL should be set to control ownership
	// of stored query results. When Athena stores query results in Amazon S3, the
	// canned ACL is set with the x-amz-acl request header. For more information
	// about S3 Object Ownership, see Object Ownership settings (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html#object-ownership-overview)
	// in the Amazon S3 User Guide.
	ACLConfiguration *ACLConfiguration `json:"aclConfiguration,omitempty"`
	// If query and calculation results are encrypted in Amazon S3, indicates the
	// encryption option used (for example, SSE_KMS or CSE_KMS) and key information.
	EncryptionConfiguration       *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
	ExpectedBucketOwner           *string                  `json:"expectedBucketOwner,omitempty"`
	OutputLocation                *string                  `json:"outputLocation,omitempty"`
	RemoveACLConfiguration        *bool                    `json:"removeACLConfiguration,omitempty"`
	RemoveEncryptionConfiguration *bool                    `json:"removeEncryptionConfiguration,omitempty"`
	RemoveExpectedBucketOwner     *bool                    `json:"removeExpectedBucketOwner,omitempty"`
	RemoveOutputLocation          *bool                    `json:"removeOutputLocation,omitempty"`
}

// Contains session configuration information.
type SessionConfiguration struct {
	// If query and calculation results are encrypted in Amazon S3, indicates the
	// encryption option used (for example, SSE_KMS or CSE_KMS) and key information.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
	ExecutionRole           *string                  `json:"executionRole,omitempty"`
	WorkingDirectory        *string                  `json:"workingDirectory,omitempty"`
}

// Contains information about the status of a session.
type SessionStatus struct {
	EndDateTime          *metav1.Time `json:"endDateTime,omitempty"`
	IdleSinceDateTime    *metav1.Time `json:"idleSinceDateTime,omitempty"`
	LastModifiedDateTime *metav1.Time `json:"lastModifiedDateTime,omitempty"`
	StartDateTime        *metav1.Time `json:"startDateTime,omitempty"`
	StateChangeReason    *string      `json:"stateChangeReason,omitempty"`
}

// Contains summary information about a session.
type SessionSummary struct {
	Description *string `json:"description,omitempty"`
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion   *EngineVersion `json:"engineVersion,omitempty"`
	NotebookVersion *string        `json:"notebookVersion,omitempty"`
}

// Contains metadata for a table.
type TableMetadata struct {
	Name *string `json:"name,omitempty"`
}

// A label that you assign to a resource. Athena resources include workgroups,
// data catalogs, and capacity reservations. Each tag consists of a key and
// an optional value, both of which you define. For example, you can use tags
// to categorize Athena resources by purpose, owner, or environment. Use a consistent
// set of tag keys to make it easier to search and filter the resources in your
// account. For best practices, see Tagging Best Practices (https://docs.aws.amazon.com/whitepapers/latest/tagging-best-practices/tagging-best-practices.html).
// Tag keys can be from 1 to 128 UTF-8 Unicode characters, and tag values can
// be from 0 to 256 UTF-8 Unicode characters. Tags can use letters and numbers
// representable in UTF-8, and the following characters: + - = . _ : / @. Tag
// keys and values are case-sensitive. Tag keys must be unique per resource.
// If you specify more than one tag, separate them by commas.
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// The name of a prepared statement that could not be returned.
type UnprocessedPreparedStatementName struct {
	StatementName *string `json:"statementName,omitempty"`
}

// The configuration of the workgroup, which includes the location in Amazon
// S3 where query and calculation results are stored, the encryption option,
// if any, used for query and calculation results, whether the Amazon CloudWatch
// Metrics are enabled for the workgroup and whether workgroup settings override
// query settings, and the data usage limits for the amount of data scanned
// per query or per workgroup. The workgroup settings override is specified
// in EnforceWorkGroupConfiguration (true/false) in the WorkGroupConfiguration.
// See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
type WorkGroupConfiguration struct {
	AdditionalConfiguration    *string `json:"additionalConfiguration,omitempty"`
	BytesScannedCutoffPerQuery *int64  `json:"bytesScannedCutoffPerQuery,omitempty"`
	// Specifies the customer managed KMS key that is used to encrypt the user's
	// data stores in Athena. When an Amazon Web Services managed key is used, this
	// value is null. This setting does not apply to Athena SQL workgroups.
	CustomerContentEncryptionConfiguration *CustomerContentEncryptionConfiguration `json:"customerContentEncryptionConfiguration,omitempty"`
	EnableMinimumEncryptionConfiguration   *bool                                   `json:"enableMinimumEncryptionConfiguration,omitempty"`
	EnforceWorkGroupConfiguration          *bool                                   `json:"enforceWorkGroupConfiguration,omitempty"`
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion *EngineVersion `json:"engineVersion,omitempty"`
	ExecutionRole *string        `json:"executionRole,omitempty"`
	// Specifies whether the workgroup is IAM Identity Center supported.
	IdentityCenterConfiguration     *IdentityCenterConfiguration `json:"identityCenterConfiguration,omitempty"`
	PublishCloudWatchMetricsEnabled *bool                        `json:"publishCloudWatchMetricsEnabled,omitempty"`
	// Specifies whether Amazon S3 access grants are enabled for query results.
	QueryResultsS3AccessGrantsConfiguration *QueryResultsS3AccessGrantsConfiguration `json:"queryResultsS3AccessGrantsConfiguration,omitempty"`
	RequesterPaysEnabled                    *bool                                    `json:"requesterPaysEnabled,omitempty"`
	// The location in Amazon S3 where query and calculation results are stored
	// and the encryption option, if any, used for query and calculation results.
	// These are known as "client-side settings". If workgroup settings override
	// client-side settings, then the query uses the workgroup settings.
	ResultConfiguration *ResultConfiguration `json:"resultConfiguration,omitempty"`
}

// The configuration information that will be updated for this workgroup, which
// includes the location in Amazon S3 where query and calculation results are
// stored, the encryption option, if any, used for query results, whether the
// Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup
// settings override the client-side settings, and the data usage limit for
// the amount of bytes scanned per query, if it is specified.
type WorkGroupConfigurationUpdates struct {
	AdditionalConfiguration    *string `json:"additionalConfiguration,omitempty"`
	BytesScannedCutoffPerQuery *int64  `json:"bytesScannedCutoffPerQuery,omitempty"`
	// Specifies the customer managed KMS key that is used to encrypt the user's
	// data stores in Athena. When an Amazon Web Services managed key is used, this
	// value is null. This setting does not apply to Athena SQL workgroups.
	CustomerContentEncryptionConfiguration *CustomerContentEncryptionConfiguration `json:"customerContentEncryptionConfiguration,omitempty"`
	EnableMinimumEncryptionConfiguration   *bool                                   `json:"enableMinimumEncryptionConfiguration,omitempty"`
	EnforceWorkGroupConfiguration          *bool                                   `json:"enforceWorkGroupConfiguration,omitempty"`
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion                   *EngineVersion `json:"engineVersion,omitempty"`
	ExecutionRole                   *string        `json:"executionRole,omitempty"`
	PublishCloudWatchMetricsEnabled *bool          `json:"publishCloudWatchMetricsEnabled,omitempty"`
	// Specifies whether Amazon S3 access grants are enabled for query results.
	QueryResultsS3AccessGrantsConfiguration      *QueryResultsS3AccessGrantsConfiguration `json:"queryResultsS3AccessGrantsConfiguration,omitempty"`
	RemoveBytesScannedCutoffPerQuery             *bool                                    `json:"removeBytesScannedCutoffPerQuery,omitempty"`
	RemoveCustomerContentEncryptionConfiguration *bool                                    `json:"removeCustomerContentEncryptionConfiguration,omitempty"`
	RequesterPaysEnabled                         *bool                                    `json:"requesterPaysEnabled,omitempty"`
	// The information about the updates in the query results, such as output location
	// and encryption configuration for the query results.
	ResultConfigurationUpdates *ResultConfigurationUpdates `json:"resultConfigurationUpdates,omitempty"`
}

// The summary information for the workgroup, which includes its name, state,
// description, and the date and time it was created.
type WorkGroupSummary struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	Description  *string      `json:"description,omitempty"`
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion                *EngineVersion `json:"engineVersion,omitempty"`
	IdentityCenterApplicationARN *string        `json:"identityCenterApplicationARN,omitempty"`
	Name                         *string        `json:"name,omitempty"`
	State                        *string        `json:"state,omitempty"`
}

// A workgroup, which contains a name, description, creation time, state, and
// other configuration, listed under WorkGroup$Configuration. Each workgroup
// enables you to isolate queries for you or your group of users from other
// queries in the same account, to configure the query results location and
// the encryption configuration (known as workgroup settings), to enable sending
// query metrics to Amazon CloudWatch, and to establish per-query data usage
// control limits for all queries in a workgroup. The workgroup settings override
// is specified in EnforceWorkGroupConfiguration (true/false) in the WorkGroupConfiguration.
// See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
type WorkGroup_SDK struct {
	// The configuration of the workgroup, which includes the location in Amazon
	// S3 where query and calculation results are stored, the encryption option,
	// if any, used for query and calculation results, whether the Amazon CloudWatch
	// Metrics are enabled for the workgroup and whether workgroup settings override
	// query settings, and the data usage limits for the amount of data scanned
	// per query or per workgroup. The workgroup settings override is specified
	// in EnforceWorkGroupConfiguration (true/false) in the WorkGroupConfiguration.
	// See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
	Configuration                *WorkGroupConfiguration `json:"configuration,omitempty"`
	CreationTime                 *metav1.Time            `json:"creationTime,omitempty"`
	Description                  *string                 `json:"description,omitempty"`
	IdentityCenterApplicationARN *string                 `json:"identityCenterApplicationARN,omitempty"`
	Name                         *string                 `json:"name,omitempty"`
	State                        *string                 `json:"state,omitempty"`
}
