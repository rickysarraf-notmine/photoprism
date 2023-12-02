# PhotoPrism® Installation Packages

As an alternative to our [Docker images](https://docs.photoprism.app/getting-started/docker-compose/), you can use the *tar.gz* archives available at [**dl.photoprism.app/pkg/linux/**](https://dl.photoprism.app/pkg/linux/) to install PhotoPrism on compatible Linux distributions without building it from source:

- <https://dl.photoprism.app/pkg/linux/amd64.tar.gz>
- <https://dl.photoprism.app/pkg/linux/arm64.tar.gz>

Since these packages need to be set up manually and do not include the system dependencies required to make use of all the features, we recommend that **only experienced users** choose this installation method.

Also note that the minimum required glibc version is 2.35, so for example Ubuntu 22.04 and Debian Bookworm will work with these binaries, but older Linux distributions may not be compatible.

## Usage

You can download and install PhotoPrism in `/opt/photoprism` by running the following commands:

```
sudo mkdir -p /opt/photoprism
cd /opt/photoprism
wget -c https://dl.photoprism.app/pkg/linux/amd64.tar.gz -O - | sudo tar -xz
sudo ln -sf /opt/photoprism/bin/photoprism /usr/local/bin/photoprism
photoprism --version
```

If your server has an ARM-based CPU, make sure to install `arm64.tar.gz` instead of `amd64.tar.gz` when using the commands above. Both are linked to the latest available build.

## Updates

To update your installation, please stop all running PhotoPrism instances and replace the contents of the installation directory, e.g. `/opt/photoprism`, with the new version.

## Dependencies

In order to use all PhotoPrism features and have [full file format support](https://www.photoprism.app/kb/file-formats), additional distribution packages must be installed manually as they are not included in the tar.gz archive, for example exiftool, darktable, rawtherapee, [libheif](https://dl.photoprism.app/dist/libheif/README.html), imagemagick, ffmpeg, libavcodec-extra, mariadb, sqlite3, and tzdata.

For details on the packages installed in our official Docker images, see <https://github.com/photoprism/photoprism/tree/develop/docker/develop>.

## Configuration

Run `photoprism --help` in a terminal to get an [overview of the command flags and environment variables](https://docs.photoprism.app/getting-started/config-options/) available for configuration. Their current values can be displayed with the `photoprism config` command.

If no explicit *originals*, *import* and/or *assets* path has been configured, a list of [default directory paths](https://github.com/photoprism/photoprism/blob/develop/pkg/fs/dirs.go) will be searched and the first existing directory will be used for the respective path.

Global config defaults can be defined in a `/etc/photoprism/defaults.yml` file (see below). When specifying paths, `~` is supported as a placeholder for the current user's home directory, e.g. `~/Pictures`. Relative paths can also be specified via `./pathname`.

Please keep in mind that any changes to the config options, either [through the UI](https://docs.photoprism.app/user-guide/settings/advanced/), [config files](https://docs.photoprism.app/getting-started/config-files/), or by [setting environment variables](https://docs.photoprism.app/getting-started/config-options/), require a restart to take effect.

### `defaults.yml`

Global config defaults, including the config and storage paths to use, can optionally be [set with a `defaults.yml` file](https://dl.photoprism.app/pkg/linux/defaults.yml) in the `/etc/photoprism` directory (requires root privileges). A custom filename for loading the defaults can be specified with the `PHOTOPRISM_DEFAULTS_YAML` environment variable or the `--defaults-yaml` command flag.

A `defaults.yml` file affects all users and should only contain values for options for which you want to set a global default, e.g.:

```yaml
ConfigPath: "~/.config/photoprism"
StoragePath: "~/.photoprism"
OriginalsPath: "~/Pictures"
ImportPath: "/media"
AdminUser: "admin"
AdminPassword: "insecure"
AuthMode: "password"
DatabaseDriver: "sqlite"
JpegQuality: 85
DetectNSFW: false
UploadNSFW: true
```

For a list of supported options and their names, see <https://docs.photoprism.app/getting-started/config-files/#config-options>.

When specifying values, make sure that the data type is the [same as in the documentation](https://docs.photoprism.app/getting-started/config-files/#config-options), e.g. *bool* values must be either `true` or `false` and *int* values must be whole numbers without any quotes like in the example above.

### `options.yml`

Default config values can be overridden by values [specified in an `options.yml` file](https://dl.photoprism.app/pkg/linux/options.yml) as well as with command flags and environment variables. To load values from an existing `options.yml` file, you can specify its storage path (excluding the filename) by setting the `ConfigPath` option in your `defaults.yml` file, using the `--config-path` command flag, or with the `PHOTOPRISM_CONFIG_PATH` environment variable.

The values in an `options.yml` file are not global and can be used to customize individual instances e.g. based on the default values in a `defaults.yml` file. Both files allow you to set any of the [supported options](https://docs.photoprism.app/getting-started/config-files/#config-options).

## Documentation

For detailed information on specific features and related resources, see our [Knowledge Base](https://www.photoprism.app/kb), or check the [User Guide](https://docs.photoprism.app/user-guide/) for help [navigating the user interface](https://docs.photoprism.app/user-guide/navigate/), a [complete list of config options](https://docs.photoprism.app/getting-started/config-options/), and [other installation methods](https://docs.photoprism.app/getting-started/):

- [PhotoPrism® User Guide](https://docs.photoprism.app/user-guide/)
- [PhotoPrism® Developer Guide](https://docs.photoprism.app/developer-guide/)
- [PhotoPrism® Knowledge Base](https://www.photoprism.app/kb)

## Getting Support

If you need help installing our software at home, you are welcome to post your question in [GitHub Discussions](https://link.photoprism.app/discussions) or ask in our [Community Chat](https://link.photoprism.app/chat). Common problems can be quickly diagnosed and solved using our [Troubleshooting Checklists](https://docs.photoprism.app/getting-started/troubleshooting/). [Silver, Gold, and Platinum](https://link.photoprism.app/membership) members are also welcome to email us for technical support and advice.

[View Support Options ›](https://www.photoprism.app/kb/getting-support)
