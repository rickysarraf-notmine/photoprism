# Ubuntu 23.10 (Mantic Minotaur)
FROM photoprism/develop:231127-mantic

## Alternative Environments:
# FROM photoprism/develop:armv7    # ARMv7 (32bit)
# FROM photoprism/develop:lunar    # Ubuntu 23.04 (Lunar Lobster)
# FROM photoprism/develop:jammy    # Ubuntu 22.04 LTS (Jammy Jellyfish)
# FROM photoprism/develop:impish   # Ubuntu 21.10 (Impish Indri)
# FROM photoprism/develop:bookworm # Debian 12 (Bookworm)
# FROM photoprism/develop:bullseye # Debian 11 (Bullseye)
# FROM photoprism/develop:buster   # Debian 10 (Buster)

# Set default working directory.
WORKDIR "/go/src/github.com/photoprism/photoprism"

# Copy source to image.
COPY . .
COPY --chown=root:root /scripts/dist/ /scripts/
