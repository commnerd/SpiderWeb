package reverse_tunnel

import (
  "github.com/stretchr/testify/assert"
  "golang.org/x/crypto/ssh"
  "../config"
  "../keys"
  "../util"
  "testing"
  "time"
  "net"
  "os"
  "io"
)

type mockConn struct{
  input []byte
  output []byte
  readIndex int64
  writeIndex int64
}

type mockAddr struct {
  str string
}

func (a *mockAddr) Network() string {
  return "tcp"
}

func (a *mockAddr) String() string {
  return a.str
}

func (c mockConn) Read(b []byte) (n int, err error) {
  if c.readIndex >= int64(len(c.output)) {
    err = io.EOF
    return
  }

  n = copy(b, c.output[c.readIndex:])
  c.readIndex += int64(n)
  return
}

func (c mockConn) Write(b []byte) (n int, err error) {
  if c.writeIndex >= int64(len(c.input)) {
    err = io.EOF
    return
  }

  n = copy(c.input[c.writeIndex:], b)
  c.writeIndex += int64(n)
  return
}

func (c mockConn) Close() error {
  return nil
}

func (c mockConn) LocalAddr() net.Addr {
  return &mockAddr{
    str: "local",
  }
}

func (c mockConn) RemoteAddr() net.Addr {
  return &mockAddr{
    str: "remote",
  }
}

func (c mockConn) SetDeadline(t time.Time) error {
  return nil
}

func (c mockConn) SetReadDeadline(t time.Time) error {
  return nil
}

func (c mockConn) SetWriteDeadline(t time.Time) error {
  return nil
}

func TestPublicKeyFile(t *testing.T) {
  privKey, pubKey := keys.Generate()
  keys.WriteToFile(privKey, pubKey)
  res := publicKeyFile(config.GetString("id_rsa_path"))
  _, ok := res.(ssh.AuthMethod)
  assert.True(t, ok)
  removeFiles(config.GetString("id_rsa_pub_path"), config.GetString("id_rsa_path"))
}

func TestMockConnectionAsReader(t *testing.T) {
  out := make([]byte, 17)
  conn := mockConn{
    output: []byte("test.allTheThings"),
    readIndex: 0,
    input: make([]byte, 0),
    writeIndex: 0,
  }
  conn.Read(out)
  assert.Equal(t, []byte("test.allTheThings"), out)
}

func TestMockConnectionAsWriter(t *testing.T) {
  in := []byte("test.allTheThings")
  conn := mockConn{
    output: in,
    readIndex: 0,
    input: make([]byte, 17),
    writeIndex: 0,
  }
  conn.Write(in)
  assert.Equal(t, []byte("test.allTheThings"), conn.input)
}

func TestHandleClient(t *testing.T) {
  local := mockConn{
    output: []byte("test.allTheThings"),
    readIndex: 0,
    input: make([]byte, 24),
    writeIndex: 0,
  }
  remote := mockConn{
    output: []byte("allTheThings.were.tested"),
    readIndex: 0,
    input: make([]byte, 17),
    writeIndex: 0,
  }

  go func() {
    handleClient(local, remote)
  }()

  chDone <- true

  assert.Equal(t, []byte("allTheThings.were.tested"), local.input)
  assert.Equal(t, []byte("test.allTheThings"), remote.input)
}

func mockDial(network, address string, config *ssh.ClientConfig) (*ssh.Client, error) {
  return &ssh.Client{}, nil
}

func removeFiles(files ...string) {
	for _, file := range files {
		if util.FileExists(file) {
			err := os.Remove(file)
			check(err)
		}
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
