package sshd

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"
	_ "../config"
	_ "io/ioutil"
	_ "../util"
	"testing"
	"net"
	_ "log"
	_ "os"
)

func TestInit(t *testing.T) {
	c := &mockConnMeta{}
	perms, err := Config.PasswordCallback(c, []byte("bar"))

	assert.IsType(t, &ssh.ServerConfig{}, Config)
	assert.Equal(t, perms, (*ssh.Permissions)(nil))
	assert.Equal(t, err, nil)
}

type mockAddr struct{}

func (m *mockAddr) Network() string {
	return "tcp"
}

func (m *mockAddr) String() string {
	return "localhost:22"
}

type mockConnMeta struct{}

// User returns the user ID for this connection.
func (m *mockConnMeta) User() string {
	return "foo"
}
// SessionID returns the session hash, also denoted by H.
func (m *mockConnMeta) SessionID() []byte {
	return []byte("blah")
}

// ClientVersion returns the client's version string as hashed
// into the session ID.
func (m *mockConnMeta) ClientVersion() []byte {
	return []byte("uber")
}

// ServerVersion returns the server's version string as hashed
// into the session ID.
func (m *mockConnMeta) ServerVersion() []byte {
	return []byte("biz")
}

// RemoteAddr returns the remote address for this connection.
func (m *mockConnMeta) RemoteAddr() net.Addr {
	return &mockAddr{}
}

// LocalAddr returns the local address for this connection.
func (m *mockConnMeta) LocalAddr() net.Addr {
	return &mockAddr{}
}