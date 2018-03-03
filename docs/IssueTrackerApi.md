# \IssueTrackerApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesUsernameRepoSlugComponentsComponentIdGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugComponentsComponentIdGet) | **Get** /repositories/{username}/{repo_slug}/components/{component_id} | 
[**RepositoriesUsernameRepoSlugComponentsGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugComponentsGet) | **Get** /repositories/{username}/{repo_slug}/components | 
[**RepositoriesUsernameRepoSlugIssuesGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesGet) | **Get** /repositories/{username}/{repo_slug}/issues | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsGet) | **Get** /repositories/{username}/{repo_slug}/issues/{issue_id}/attachments | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPathDelete**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPathDelete) | **Delete** /repositories/{username}/{repo_slug}/issues/{issue_id}/attachments/{path} | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPathGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPathGet) | **Get** /repositories/{username}/{repo_slug}/issues/{issue_id}/attachments/{path} | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPost**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPost) | **Post** /repositories/{username}/{repo_slug}/issues/{issue_id}/attachments | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdCommentsCommentIdGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdCommentsCommentIdGet) | **Get** /repositories/{username}/{repo_slug}/issues/{issue_id}/comments/{comment_id} | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdCommentsGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdCommentsGet) | **Get** /repositories/{username}/{repo_slug}/issues/{issue_id}/comments | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdDelete**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdDelete) | **Delete** /repositories/{username}/{repo_slug}/issues/{issue_id} | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdGet) | **Get** /repositories/{username}/{repo_slug}/issues/{issue_id} | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdVoteDelete**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdVoteDelete) | **Delete** /repositories/{username}/{repo_slug}/issues/{issue_id}/vote | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdVoteGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdVoteGet) | **Get** /repositories/{username}/{repo_slug}/issues/{issue_id}/vote | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdVotePut**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdVotePut) | **Put** /repositories/{username}/{repo_slug}/issues/{issue_id}/vote | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdWatchDelete**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdWatchDelete) | **Delete** /repositories/{username}/{repo_slug}/issues/{issue_id}/watch | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdWatchGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdWatchGet) | **Get** /repositories/{username}/{repo_slug}/issues/{issue_id}/watch | 
[**RepositoriesUsernameRepoSlugIssuesIssueIdWatchPut**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesIssueIdWatchPut) | **Put** /repositories/{username}/{repo_slug}/issues/{issue_id}/watch | 
[**RepositoriesUsernameRepoSlugIssuesPost**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugIssuesPost) | **Post** /repositories/{username}/{repo_slug}/issues | 
[**RepositoriesUsernameRepoSlugMilestonesGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugMilestonesGet) | **Get** /repositories/{username}/{repo_slug}/milestones | 
[**RepositoriesUsernameRepoSlugMilestonesMilestoneIdGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugMilestonesMilestoneIdGet) | **Get** /repositories/{username}/{repo_slug}/milestones/{milestone_id} | 
[**RepositoriesUsernameRepoSlugVersionsGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugVersionsGet) | **Get** /repositories/{username}/{repo_slug}/versions | 
[**RepositoriesUsernameRepoSlugVersionsVersionIdGet**](IssueTrackerApi.md#RepositoriesUsernameRepoSlugVersionsVersionIdGet) | **Get** /repositories/{username}/{repo_slug}/versions/{version_id} | 


# **RepositoriesUsernameRepoSlugComponentsComponentIdGet**
> Component RepositoriesUsernameRepoSlugComponentsComponentIdGet(ctx, username, repoSlug, componentId)


Returns the specified issue tracker component object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **componentId** | **int32**| The component&#39;s id | 

### Return type

[**Component**](component.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugComponentsGet**
> PaginatedComponents RepositoriesUsernameRepoSlugComponentsGet(ctx, username, repoSlug)


Returns the components that have been defined in the issue tracker.  This resource is only available on repositories that have the issue tracker enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

[**PaginatedComponents**](paginated_components.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesGet**
> PaginatedIssues RepositoriesUsernameRepoSlugIssuesGet(ctx, username, repoSlug)


Returns the issues in the issue tracker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

[**PaginatedIssues**](paginated_issues.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsGet**
> PaginatedIssueAttachments RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsGet(ctx, username, repoSlug, issueId)


Returns all attachments for this issue.  This returns the files' meta data. This does not return the files' actual contents.  The files are always ordered by their upload date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **issueId** | **int32**| The issue&#39;s id | 

### Return type

[**PaginatedIssueAttachments**](paginated_issue_attachments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPathDelete**
> RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPathDelete(ctx, username, path, issueId, repoSlug)


Deletes an attachment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **path** | **string**|  | 
  **issueId** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPathGet**
> RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPathGet(ctx, username, path, issueId, repoSlug)


Returns the contents of the specified file attachment.  Note that this endpoint does not return a JSON response, but instead returns a redirect pointing to the actual file that in turn will return the raw contents.  The redirect URL contains a one-time token that has a limited lifetime. As a result, the link should not be persisted, stored, or shared.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **path** | **string**|  | 
  **issueId** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPost**
> RepositoriesUsernameRepoSlugIssuesIssueIdAttachmentsPost(ctx, username, repoSlug, issueId)


Upload new issue attachments.  To upload files, perform a `multipart/form-data` POST containing one or more file fields.  When a file is uploaded with the same name as an existing attachment, then the existing file will be replaced.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **issueId** | **int32**| The issue&#39;s id | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdCommentsCommentIdGet**
> ModelError RepositoriesUsernameRepoSlugIssuesIssueIdCommentsCommentIdGet(ctx, username, commentId, issueId, repoSlug)


Returns the specified issue comment object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **commentId** | **string**|  | 
  **issueId** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdCommentsGet**
> ModelError RepositoriesUsernameRepoSlugIssuesIssueIdCommentsGet(ctx, username, issueId, repoSlug)


Returns all comments that were made on the specified issue.  The default sorting is oldest to newest and can be overridden with the `sort` query parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **issueId** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdDelete**
> Issue RepositoriesUsernameRepoSlugIssuesIssueIdDelete(ctx, username, issueId, repoSlug)


Deletes the specified issue. This requires write access to the repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **issueId** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

[**Issue**](issue.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdGet**
> Issue RepositoriesUsernameRepoSlugIssuesIssueIdGet(ctx, username, issueId, repoSlug)


Returns the specified issue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **issueId** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

[**Issue**](issue.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdVoteDelete**
> ModelError RepositoriesUsernameRepoSlugIssuesIssueIdVoteDelete(ctx, username, repoSlug, issueId)


Retract your vote.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **issueId** | **int32**| The issue&#39;s id | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdVoteGet**
> ModelError RepositoriesUsernameRepoSlugIssuesIssueIdVoteGet(ctx, username, repoSlug, issueId)


Check whether the authenticated user has voted for this issue. A 204 status code indicates that the user has voted, while a 404 implies they haven't.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **issueId** | **int32**| The issue&#39;s id | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdVotePut**
> ModelError RepositoriesUsernameRepoSlugIssuesIssueIdVotePut(ctx, username, repoSlug, issueId)


Vote for this issue.  To cast your vote, do an empty PUT. The 204 status code indicates that the operation was successful.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **issueId** | **int32**| The issue&#39;s id | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdWatchDelete**
> ModelError RepositoriesUsernameRepoSlugIssuesIssueIdWatchDelete(ctx, username, repoSlug, issueId)


Stop watching this issue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **issueId** | **int32**| The issue&#39;s id | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdWatchGet**
> ModelError RepositoriesUsernameRepoSlugIssuesIssueIdWatchGet(ctx, username, repoSlug, issueId)


Indicated whether or not the authenticated user is watching this issue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **issueId** | **int32**| The issue&#39;s id | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesIssueIdWatchPut**
> ModelError RepositoriesUsernameRepoSlugIssuesIssueIdWatchPut(ctx, username, repoSlug, issueId)


Start watching this issue.  To start watching this issue, do an empty PUT. The 204 status code indicates that the operation was successful.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **issueId** | **int32**| The issue&#39;s id | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugIssuesPost**
> Issue RepositoriesUsernameRepoSlugIssuesPost(ctx, username, repoSlug, body)


Creates a new issue.  This call requires authentication. Private repositories or private issue trackers require the caller to authenticate with an account that has appropriate authorisation.  The authenticated user is used for the issue's `reporter` field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **body** | [**Issue**](Issue.md)| The new issue. Note that the only required element is &#x60;title&#x60;. All other elements can be omitted from the body. | 

### Return type

[**Issue**](issue.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugMilestonesGet**
> PaginatedMilestones RepositoriesUsernameRepoSlugMilestonesGet(ctx, username, repoSlug)


Returns the milestones that have been defined in the issue tracker.  This resource is only available on repositories that have the issue tracker enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

[**PaginatedMilestones**](paginated_milestones.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugMilestonesMilestoneIdGet**
> Milestone RepositoriesUsernameRepoSlugMilestonesMilestoneIdGet(ctx, username, repoSlug, milestoneId)


Returns the specified issue tracker milestone object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **milestoneId** | **int32**| The milestone&#39;s id | 

### Return type

[**Milestone**](milestone.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugVersionsGet**
> PaginatedVersions RepositoriesUsernameRepoSlugVersionsGet(ctx, username, repoSlug)


Returns the versions that have been defined in the issue tracker.  This resource is only available on repositories that have the issue tracker enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 

### Return type

[**PaginatedVersions**](paginated_versions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesUsernameRepoSlugVersionsVersionIdGet**
> Version RepositoriesUsernameRepoSlugVersionsVersionIdGet(ctx, username, repoSlug, versionId)


Returns the specified issue tracker version object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **username** | **string**|  | 
  **repoSlug** | **string**|  | 
  **versionId** | **int32**| The version&#39;s id | 

### Return type

[**Version**](version.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

