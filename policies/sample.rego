package sample

# service assignments
# group_roles := {
# 	"admin": ["admin"],
# 	"quality_head_design": ["quality_head_design"],
# }

# service-whitelist assignments
# service_whitelists := {
# 	"admin": [
# 		{"action": "view_all", "object": "design"},
# 		{"action": "edit", "object": "design"},
# 	],
# 	"quality_head_design": [
# 		{"action": "view_all", "object": "design"},
# 	],
# }

service_whitelists := {
	"service-a": ["service-b"],
	"sevice-b": ["service-c"],
}

default allow = false

allow {
	# lookup the list of roles for the user
	# roles := group_roles[input.user[_]]

	# for each role in that list
	# r := roles[_]

	# lookup the permissions list for role r
	# permissions := role_permissions[r]
	whitelists := service_whitelists[input.audience]

	# for each permission
	whitelist := whitelists[_]

	# check if the permission granted to r matches the user's request
	whitelist == input.issuer
}
