import { $gettext } from "common/vm";

// All user role with their display name.
export const Roles = () => {
  return {
    admin: $gettext("Admin"),
    user: $gettext("User"),
    viewer: $gettext("Viewer"),
    contributor: $gettext("Contributor"),
    guest: $gettext("Guest"),
    client: $gettext("Client"),
    visitor: $gettext("Visitor"),
    "": $gettext("Unauthorized"),
  };
};

// AuthProviders maps known auth providers to their display name.
export const Providers = () => {
  return {
    "": $gettext("Default"),
    default: $gettext("Default"),
    local: $gettext("Local"),
    client: $gettext("Client"),
    client_credentials: $gettext("Client Credentials"),
    application: $gettext("Application"),
    access_token: $gettext("Access Token"),
    password: $gettext("Local"),
    ldap: $gettext("LDAP/AD"),
    link: $gettext("Link"),
    token: $gettext("Link"),
    none: $gettext("None"),
  };
};

// AuthMethods maps known auth methods to their display name.
export const Methods = () => {
  return {
    "": $gettext("Default"),
    default: $gettext("Default"),
    session: $gettext("Session"),
    personal: $gettext("Personal"),
    client: $gettext("Client"),
    access_token: $gettext("Access Token"),
    oauth2: "OAuth2",
    totp: "TOTP/2FA",
    oidc: "OIDC",
  };
};
