/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type BranchrestrictionsApiService service


/* BranchrestrictionsApiService 
 Returns a paginated list of all branch restrictions on the repository.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username 
 @param repoSlug 
 @return PaginatedBranchrestrictions*/
func (a *BranchrestrictionsApiService) RepositoriesUsernameRepoSlugBranchRestrictionsGet(ctx context.Context, username string, repoSlug string) (PaginatedBranchrestrictions,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PaginatedBranchrestrictions
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/branch-restrictions"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BranchrestrictionsApiService 
 Deletes an existing branch restriction rule.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username 
 @param repoSlug 
 @param id The restriction rule&#39;s id
 @return */
func (a *BranchrestrictionsApiService) RepositoriesUsernameRepoSlugBranchRestrictionsIdDelete(ctx context.Context, username string, repoSlug string, id string) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/branch-restrictions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* BranchrestrictionsApiService 
 Returns a specific branch restriction rule.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username 
 @param repoSlug 
 @param id The restriction rule&#39;s id
 @return Branchrestriction*/
func (a *BranchrestrictionsApiService) RepositoriesUsernameRepoSlugBranchRestrictionsIdGet(ctx context.Context, username string, repoSlug string, id string) (Branchrestriction,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Branchrestriction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/branch-restrictions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BranchrestrictionsApiService 
 Updates an existing branch restriction rule.  Fields not present in the request body are ignored.  See [&#x60;POST&#x60;](../../branch-restrictions#post) for details.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username 
 @param repoSlug 
 @param id The restriction rule&#39;s id
 @param body The new version of the existing rule
 @return Branchrestriction*/
func (a *BranchrestrictionsApiService) RepositoriesUsernameRepoSlugBranchRestrictionsIdPut(ctx context.Context, username string, repoSlug string, id string, body Branchrestriction) (Branchrestriction,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Branchrestriction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/branch-restrictions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* BranchrestrictionsApiService 
 Creates a new branch restriction rule for a repository.  &#x60;kind&#x60; describes what will be restricted. Allowed values are: &#x60;push&#x60;, &#x60;force&#x60;, &#x60;delete&#x60;, and &#x60;restrict_merges&#x60;.  Different kinds of branch restrictions have different requirements:  * &#x60;push&#x60; and &#x60;restrict_merges&#x60; require &#x60;users&#x60; and &#x60;groups&#x60; to be   specified. Empty lists are allowed, in which case permission is   denied for everybody. * &#x60;force&#x60; can not be specified in a Mercurial repository.  &#x60;pattern&#x60; is used to determine which branches will be restricted.  A &#x60;&#39;*&#39;&#x60; in &#x60;pattern&#x60; will expand to match zero or more characters, and every other character matches itself. For example, &#x60;&#39;foo*&#39;&#x60; will match &#x60;&#39;foo&#39;&#x60; and &#x60;&#39;foobar&#39;&#x60;, but not &#x60;&#39;barfoo&#39;&#x60;. &#x60;&#39;*&#39;&#x60; will match all branches.  &#x60;users&#x60; and &#x60;groups&#x60; are lists of user names and group names.  &#x60;kind&#x60; and &#x60;pattern&#x60; must be unique within a repository; adding new users or groups to an existing restriction should be done via &#x60;PUT&#x60;.  Note that branch restrictions with overlapping patterns are allowed, but the resulting behavior may be surprising.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username 
 @param repoSlug 
 @param body The new rule
 @return Branchrestriction*/
func (a *BranchrestrictionsApiService) RepositoriesUsernameRepoSlugBranchRestrictionsPost(ctx context.Context, username string, repoSlug string, body Branchrestriction) (Branchrestriction,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Branchrestriction
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/branch-restrictions"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

