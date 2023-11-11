package acl

// Events specifies granted permissions by event channel and Role.
var Events = ACL{
	ResourceDefault: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantSubscribeOwn,
	},
	ChannelUser: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantSubscribeOwn,
		RoleVisitor: GrantSubscribeOwn,
	},
	ChannelSession: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantSubscribeOwn,
		RoleVisitor: GrantSubscribeOwn,
	},
}
