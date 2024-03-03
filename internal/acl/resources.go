package acl

// Resources specifies granted permissions by Resource and Role.
var Resources = ACL{
	ResourceFiles: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceFolders: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnlyReact,
		RoleFamily:  GrantReadOnly,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: GrantSearchShared,
		RoleClient:  GrantFullAccess,
	},
	ResourceShares: Roles{
		RoleAdmin: GrantFullAccess,
		RoleFamily:  GrantReadOnly,
	},
	ResourcePhotos: GrantDefaults,
	ResourceVideos: GrantDefaults,
	ResourceFavorites: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleClient: GrantFullAccess,
	},
	ResourceAlbums: GrantDefaults,
	ResourceMoments: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: GrantSearchShared,
		RoleClient:  GrantFullAccess,
	},
	ResourceCalendar: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: GrantSearchShared,
		RoleClient:  GrantFullAccess,
	},
	ResourcePeople: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantReadOnly,
		RoleClient: GrantFullAccess,
	},
	ResourcePlaces: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily: GrantReadOnly,
		RoleVisitor: GrantViewShared,
		RoleClient:  GrantFullAccess,
	},
	ResourceLabels: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantReadOnly,
		RoleClient: GrantFullAccess,
	},
	ResourceConfig: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleClient:  GrantViewOwn,
		RoleDefault: GrantViewOwn,
	},
	ResourceSettings: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleVisitor: Grant{AccessOwn: true, ActionView: true},
	},
	ResourceServices: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourcePassword: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantChangePassword,
	},
	ResourceUsers: Roles{
		RoleAdmin: Grant{AccessAll: true, AccessOwn: true, ActionView: true, ActionCreate: true, ActionUpdate: true, ActionDelete: true, ActionSubscribe: true},
	},
	ResourceLogs: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceWebDAV: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceMetrics: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantViewAll,
	},
	ResourceFeedback: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantLoginOnly,
	},
	ResourceDefault: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantNone,
	},
}
