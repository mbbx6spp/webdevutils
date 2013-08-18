# webdevutils

This will contain a set of simple web development utilities that are portable
and packaged across all \*NIX based OSes for the purpose of making web
development less annoying. Not forcing backend developers to install NodeJS
just because their frontend developers use it. And, of course, not forcing
frontend developers to download the whole effing internet with Maven just
because the backend developers build with Java or whatever.

## Purpose

The goal is to build _really_ simple command line tools that:
* allow the two parts of development to evolve separately;
* not forcing both toolsets to be installed in both frontend and backend
development environments, creating a Frankenstein environment to maintain.
* provide packages for these command line tools for the main \*NIX OSes used
for development including: Linux (deb and RPM support), OSX (Homebrew),
FreeBSD (ports).

If you have additional suggestions for packages to provide, submit an
[issue](https://github.com/mbbx6spp/webdevutils/issues).

## Contents

At the moment I only have two command line tools in mind:
* `staticserver` - a command line tool that can statically serve the
contents of a local directory on a specific port. Hey, I did say
_really_ simple earlier and I meant it.
* `autocompiler` - a command line tool that can listen for filesystem
events and autocompile source files in the directory tree it is
listening on to a destination using whatever command you want.

The use case for `staticserver` are:
* Allowing backend developers to serve compile assets already built and
generated by front end devs to drive their API endpoints.
* Allowing frontend developers to serve static JSON/XML canned responses
to fake a REST API for their frontend app until it is ready.

There will be more to come, I promise.

## License

Released under the BSD 3-clause license:
http://opensource.org/licenses/BSD-3-Clause

See LICENSE file for more information.

## Developers

So far just me, but feel free to submit pull requests and add your name to
this file:

* [Susan Potter](https://github.com/mbbx6spp)

