package acl

// Resources specifies granted permissions by Resource and Role.
var Resources = ACL{
	ResourceFiles: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourcePhotos: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnlyReact,
		RoleVisitor: Grant{AccessShared: true, ActionView: true, ActionDownload: true},
	},
	ResourceVideos: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
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
	},
	ResourcePlaces: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: Grant{AccessShared: true, ActionView: true, ActionDownload: true},
	},
	ResourceCalendar: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: GrantSearchShared,
	},
	ResourceMoments: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleFamily:  GrantReadOnly,
		RoleVisitor: GrantSearchShared,
	},
	ResourcePeople: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantReadOnly,
	},
	ResourceFavorites: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantReadOnly,
	},
	ResourceLabels: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantReadOnly,
	},
	ResourceLogs: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourceSettings: Roles{
		RoleAdmin:   GrantFullAccess,
		RoleVisitor: Grant{AccessOwn: true, ActionView: true},
	},
	ResourceFeedback: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourcePassword: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantChangePassword,
	},
	ResourceShares: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourceServices: Roles{
		RoleAdmin: GrantFullAccess,
	},
	ResourceUsers: Roles{
		RoleAdmin: Grant{AccessAll: true, AccessOwn: true, ActionView: true, ActionCreate: true, ActionUpdate: true, ActionDelete: true, ActionSubscribe: true},
	},
	ResourceConfig: Roles{
		RoleAdmin:  GrantFullAccess,
		RoleFamily: GrantLoginOnly,
	},
	ResourceDefault: Roles{
		RoleAdmin: GrantFullAccess,
	},
}
