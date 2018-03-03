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
	"time"
)

type Commitstatus struct {

	Type_ string `json:"type"`

	Links *CommitstatusLinks `json:"links,omitempty"`

	// The commit status' id.
	Uuid string `json:"uuid,omitempty"`

	// An identifier for the status that's unique to         its type (current \"build\" is the only supported type) and the vendor,         e.g. BB-DEPLOY
	Key string `json:"key,omitempty"`

	//  The name of the ref that pointed to this commit at the time the status object was created. Note that this the ref may since have moved off of the commit. This optional field can be useful for build systems whose build triggers and configuration are branch-dependent (e.g. a Pipeline build). It is legitimate for this field to not be set, or even apply (e.g. a static linting job).
	Refname string `json:"refname,omitempty"`

	// A URL linking back to the vendor or build system, for providing more information about whatever process produced this status. Accepts context variables `repository` and `commit` that Bitbucket will evaluate at runtime whenever at runtime. For example, one could use https://foo.com/builds/{repository.full_name} which Bitbucket will turn into https://foo.com/builds/foo/bar at render time.
	Url string `json:"url,omitempty"`

	// Provides some indication of the status of this commit
	State string `json:"state,omitempty"`

	// An identifier for the build itself, e.g. BB-DEPLOY-1
	Name string `json:"name,omitempty"`

	// A description of the build (e.g. \"Unit tests in Bamboo\")
	Description string `json:"description,omitempty"`

	CreatedOn time.Time `json:"created_on,omitempty"`

	UpdatedOn time.Time `json:"updated_on,omitempty"`
}