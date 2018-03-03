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

// An executable pipeline command.
type PipelineCommand struct {

	// The executable command.
	Command string `json:"command,omitempty"`

	// The name of the command.
	Name string `json:"name,omitempty"`

	// The range in the log that contains the execution output of this command.
	LogRange *PipelineLogRange `json:"log_range,omitempty"`
}