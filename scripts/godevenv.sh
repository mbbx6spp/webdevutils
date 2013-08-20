#!/usr/bin/env bash

LOCKFILE=$PWD/.bootstrap.lock
[ -f $LOCKFILE ] && return

cleanup() {
  test -f $LOCKFILE && rm $LOCKFILE;
}

exit_trace() {
  local lineno=$1;
  local cmd=$2;
  cleanup;
  which logger>/dev/null && logger -p notice "$0: exit on line $lineno with command $cmd";
}

recommend_installing_johnny_deps() {
  echo "WARNING: It is recommended you install johnny_deps to setup "
  echo "your Go dev environment for this project or you may use the "
  echo "incorrect revision of dependencies."
  echo

  PLATFORM=$(uname -s)
  if [ "$PLATFORM" == "Darwin" ]; then
    echo "You can download an unofficial Homebrew formula for OSX here:"
    echo "    https://gist.github.com/mbbx6spp/6096071"
    echo
    echo "You can run the following to install"
    echo "    wget -O /usr/local/Library/Formula/johnny_deps.rb \\"
    echo "      https://gist.github.com/mbbx6spp/6096071/raw/johnny_deps.rb && \\"
    echo "    brew install -HEAD johnny_deps"
  else
    echo "Visit https://github.com/VividCortex/johnny-deps and follow build instructions"
  fi;
  echo
  echo "Getting dependencies with \`go get\` now"
}

when_go_installed() {
  which johnny_deps>/dev/null && johnny_deps || \
    recommend_installing_johnny_deps
}

touch $LOCKFILE

which go>/dev/null
exit_code=$?
if [ "$exit_code" == 0 ]; then
  export GOPATH=$(pwd)
  echo "Set GOPATH to $GOPATH"

  when_go_installed
else
  echo "You must install Go locally."
  echo
fi;

cleanup
