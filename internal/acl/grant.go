package acl

// Predefined grants to simplify configuration.
var (
	GrantFullAccess   = Grant{FullAccess: true, AccessAll: true, AccessLibrary: true, ActionCreate: true, ActionUpdate: true, ActionDelete: true, ActionDownload: true, ActionShare: true, ActionRate: true, ActionReact: true, ActionManage: true, ActionSubscribe: true}
	GrantSearchShared = Grant{AccessShared: true, ActionSearch: true, ActionView: true, ActionDownload: true}
	GrantSubscribeAll = Grant{AccessAll: true, ActionSubscribe: true}
	GrantSubscribeOwn = Grant{AccessOwn: true, ActionSubscribe: true}

	// Custom, family role related permissions.
	// In order to like and react to a photo, the following permissions are needed in addition to the read-only ones:
	//   - manage: used determine whether the "like/favorite" button will be shown
	//   - update: required to be able to "like" a photo
	//   - update: required to be able to use the experimental "react" feature
	GrantLoginOnly      = Grant{AccessOwn: true}
	GrantChangePassword = Grant{ActionUpdate: true}
	GrantReadOnly       = GrantSearchShared.Plus(Grant{AccessLibrary: true})
	GrantReadOnlyReact  = GrantReadOnly.Plus(Grant{ActionReact: true, ActionManage: true, ActionUpdate: true})
)

// Grant represents permissions granted or denied.
type Grant map[Permission]bool

// Allow checks whether the permission is granted.
func (grant Grant) Allow(perm Permission) bool {
	if result, ok := grant[perm]; ok {
		return result
	} else if result, ok = grant[FullAccess]; ok {
		return result
	}

	return false
}

// Plus creates a new grant by adding up all permissions.
func (grant Grant) Plus(updated Grant) Grant {
	merged := make(Grant)
	for k, v := range grant {
		merged[k] = v
	}
	for k, v := range updated {
		merged[k] = v
	}
	return merged
}
