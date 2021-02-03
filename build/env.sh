#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
gitdir="$workspace/src/github.com/2miners"
if [ ! -L "$gitdir/stratum-ping" ]; then
    mkdir -p "$gitdir"
    cd "$gitdir"
    ln -s ../../../../../. stratum-ping
    cd "$root"
fi

# Set up the environment to use the workspace.
# Also add Godeps workspace so we build using canned dependencies.
GOPATH="$workspace"
GOBIN="$PWD/build/bin"
export GOPATH GOBIN 

# Run the command inside the workspace.
cd "$gitdir/stratum-ping"
PWD="$gitdir/stratum-ping"

# Launch the arguments with the configured environment.
exec "$@"
