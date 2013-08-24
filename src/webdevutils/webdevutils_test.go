package webdevutils

import (
	. "launchpad.net/gocheck"
	"testing"
  "os"
  "path"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type WebdevutilsSuite struct{
  rootdir string
  crtfile string
  keyfile string
}

var _ = Suite(&WebdevutilsSuite{})

func (s *WebdevutilsSuite) SetUpSuite(c *C) {
  confdir := os.Getenv("CONFIGDIR")
  s.crtfile = path.Join(confdir, "server.crt")
  s.keyfile = path.Join(confdir, "server.pem")
  s.rootdir = c.MkDir()
}

func (s *WebdevutilsSuite) TearDownSuite(c *C) {
  // TODO
}

func (s *WebdevutilsSuite) TestStaticServerTLSInvalidKeyPair(c *C) {
	err := StaticServerTLS(":9000", "crap.crt", "crap.pem", s.rootdir)
	c.Check(err, Not(IsNil))
  c.Check(err.Error(), Matches, ".*no such file or directory")
}

func (s *WebdevutilsSuite) TestStaticServerTLSInvalidRootDir(c *C) {
  err := StaticServerTLS(":9000", s.crtfile, s.keyfile, "crapdir")
  c.Check(err, Not(IsNil))
  c.Check(err.Error(), Matches, "stat crapdir: no such file or directory")
}

func (s *WebdevutilsSuite) TestStaticServerTLSPrivilegedPort(c *C) {
  err := StaticServerTLS(":443", s.crtfile, s.keyfile, s.rootdir)
  c.Check(err, Not(IsNil))
  c.Check(err.Error(), Matches, ".*:443:.*?permission denied$")
}

func (s *WebdevutilsSuite) TestStaticServerInvalidRootDir(c *C) {
  err := StaticServer(":9000", "crapdir")
  c.Check(err, Not(IsNil))
  c.Check(err.Error(), Matches, "stat crapdir: no such file or directory")
}

func (s *WebdevutilsSuite) TestStaticServerPrivilegedPort(c *C) {
  err := StaticServer(":80", s.rootdir)
  c.Check(err, Not(IsNil))
  c.Check(err.Error(), Matches, ".*:80:.*?permission denied$")
}
