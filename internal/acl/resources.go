package acl

// Resources specifies granted permissions by Resource and Role.
var Resources = ACL{
	ResourceFiles: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourcePhotos: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnlyReact,
		RoleClient:  GrantFullAccess,
		RoleVisitor: Grant{AccessShared: true, ActionView: true, ActionDownload: true},
	},
	ResourceVideos: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleClient:  GrantFullAccess,
		RoleVisitor: Grant{AccessShared: true, ActionView: true, ActionDownload: true},
	},
	ResourceAlbums: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: GrantSearchShared,
	},
	ResourceFolders: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: GrantSearchShared,
		RoleClient:  GrantFullAccess,
	},
	ResourcePlaces: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily: GrantReadOnly,
		RoleVisitor: GrantViewShared,
		RoleClient:  GrantFullAccess,
	},
	ResourceCalendar: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: GrantSearchShared,
		RoleClient:  GrantFullAccess,
	},
	ResourceMoments: Roles{
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
	ResourceFavorites: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleClient: GrantFullAccess,
	},
	ResourceLabels: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantReadOnly,
		RoleClient: GrantFullAccess,
	},
	ResourceLogs: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceSettings: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleVisitor: Grant{AccessOwn: true, ActionView: true},
	},
	ResourceFeedback: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantLoginOnly,
	},
	ResourceShares: Roles{
		RoleAdmin: GrantFullAccess,
		RoleFamily:  GrantReadOnly,
	},
	ResourceConfig: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleClient:  GrantViewOwn,
		RoleDefault: GrantViewOwn,
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
	ResourceWebDAV: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantFullAccess,
	},
	ResourceMetrics: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantViewAll,
	},
	ResourceDefault: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleClient: GrantNone,
	},
}
