// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/auth/users/users.proto

package users

import (
	request "github.com/chef/automate/components/automate-gateway/api/auth/users/request"
	policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.users.UsersMgmt/GetUsers", "auth:users", "read", "GET", "/auth/users", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.users.UsersMgmt/GetUserByUsername", "auth:users:{username}", "read", "GET", "/auth/users/{username}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Username); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "username":
					return m.Username
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.users.UsersMgmt/CreateUser", "auth:users", "create", "POST", "/auth/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateUser); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "email":
					return m.Email
				case "username":
					return m.Username
				case "password":
					return m.Password
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.users.UsersMgmt/DeleteUserByUsername", "auth:users:{username}", "delete", "DELETE", "/auth/users/{username}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Username); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "username":
					return m.Username
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.users.UsersMgmt/UpdateUser", "auth:users:{username}", "update", "PUT", "/auth/users/{username}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateUser); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "email":
					return m.Email
				case "name":
					return m.Name
				case "password":
					return m.Password
				case "username":
					return m.Username
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.users.UsersMgmt/UpdateSelf", "users:{username}", "update", "PUT", "/users/{username}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateSelf); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "password":
					return m.Password
				case "username":
					return m.Username
				case "previous_password":
					return m.PreviousPassword
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.users.UsersMgmt/GetUser", "auth:users:{email}", "read", "", "", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Email); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "email":
					return m.Email
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.users.UsersMgmt/DeleteUser", "auth:users:{email}", "delete", "", "", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Email); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "email":
					return m.Email
				default:
					return ""
				}
			})
		}
		return ""
	})
}