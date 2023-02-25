package commands

import (
	"fmt"

	"github.com/dustin/go-humanize/english"
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/query"
	"github.com/photoprism/photoprism/pkg/report"
	"github.com/photoprism/photoprism/pkg/txt"
)

// UsersListCommand configures the command name, flags, and action.
var UsersListCommand = cli.Command{
	Name:   "ls",
	Usage:  "Displays existing user accounts",
	Flags:  report.CliFlags,
	Action: usersListAction,
}

// usersListAction displays existing user accounts.
func usersListAction(ctx *cli.Context) error {
	return CallWithDependencies(ctx, func(conf *config.Config) error {
		cols := []string{"UID", "Name", "User", "Email", "Role", "Super Admin", "Web Login", "WebDAV", "Attributes", "Created At"}

		// Fetch users from database.
		users := query.RegisteredUsers()
		rows := make([][]string, len(users))

		// Show log message.
		log.Infof("found %s", english.Plural(len(users), "user", "users"))

		// Display report.
		for i, user := range users {
			rows[i] = []string{
				user.UID(),
				user.FullName(),
				user.Login(),
				user.Email(),
				user.AclRole().String(),
				report.Bool(user.SuperAdmin, report.Yes, report.No),
				report.Bool(user.CanLogIn(), report.Enabled, report.Disabled),
				report.Bool(user.CanUseWebDAV(), report.Enabled, report.Disabled),
				user.Attr(),
				txt.TimeStamp(&user.CreatedAt),
			}
		}

		result, err := report.RenderFormat(rows, cols, report.CliFormat(ctx))

		fmt.Printf("\n%s\n", result)

		return err
	})
}
