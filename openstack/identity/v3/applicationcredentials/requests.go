package applicationcredentials

import (
	"time"

	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToApplicationCredentialListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Name filters the response by an application credential name
	Name string `q:"name"`
}

// ToApplicationCredentialListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToApplicationCredentialListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the ApplicationCredentials to which the current token has access.
func List(client *gophercloud.ServiceClient, userID string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client, userID)
	if opts != nil {
		query, err := opts.ToApplicationCredentialListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ApplicationCredentialPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details on a single user, by ID.
func Get(client *gophercloud.ServiceClient, userID string, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, userID, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToApplicationCredentialCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to create an application credential.
type CreateOpts struct {
	// The name of the application credential.
	Name string `json:"name,omitempty" required:"true"`
	// A description of the application credential’s purpose.
	Description string `json:"description,omitempty"`
	// A flag indicating whether the application credential may be used for creation or destruction of other application credentials or trusts.
	// Defaults to false
	Unrestricted bool `json:"unrestricted"`
	// The secret for the application credential, either generated by the server or provided by the user.
	// This is only ever shown once in the response to a create request. It is not stored nor ever shown again.
	// If the secret is lost, a new application credential must be created.
	Secret string `json:"secret,omitempty"`
	// A list of one or more roles that this application credential has associated with its project.
	// A token using this application credential will have these same roles.
	Roles []Role `json:"roles,omitempty"`
	// A list of access rules objects.
	AccessRules []AccessRule `json:"access_rules,omitempty"`
	// The expiration time of the application credential, if one was specified.
	ExpiresAt *time.Time `json:"-"`
}

// ToApplicationCredentialCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToApplicationCredentialCreateMap() (map[string]interface{}, error) {
	parent := "application_credential"
	b, err := gophercloud.BuildRequestBody(opts, parent)
	if err != nil {
		return nil, err
	}

	if opts.ExpiresAt != nil {
		if v, ok := b[parent].(map[string]interface{}); ok {
			v["expires_at"] = opts.ExpiresAt.Format(gophercloud.RFC3339MilliNoZ)
		}
	}

	return b, nil
}

// Create creates a new ApplicationCredential.
func Create(client *gophercloud.ServiceClient, userID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToApplicationCredentialCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client, userID), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete deletes an application credential.
func Delete(client *gophercloud.ServiceClient, userID string, id string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, userID, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListAccessRules enumerates the AccessRules to which the current user has access.
func ListAccessRules(client *gophercloud.ServiceClient, userID string) pagination.Pager {
	url := listAccessRulesURL(client, userID)
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return AccessRulePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// GetAccessRule retrieves details on a single access rule by ID.
func GetAccessRule(client *gophercloud.ServiceClient, userID string, id string) (r GetAccessRuleResult) {
	resp, err := client.Get(getAccessRuleURL(client, userID, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// DeleteAccessRule deletes an access rule.
func DeleteAccessRule(client *gophercloud.ServiceClient, userID string, id string) (r DeleteResult) {
	resp, err := client.Delete(deleteAccessRuleURL(client, userID, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
