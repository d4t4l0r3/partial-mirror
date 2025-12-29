# partial-mirror

This software is aimed at hosting partial mirrors for any linux distro (and potentially any other system using similar HTTP(S) repositories).
The server acts as a regular mirror, but only syncs a selected list of packages rather than the entire repository, to preserve bandwidth and storage space.

Configurable is

- the upstream mirror to sync from (and fallbacks)
- the list of packages to keep synced
- package retention (how many past versions to keep & how long to keep them)
