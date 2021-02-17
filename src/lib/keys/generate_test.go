package keys

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"

	"github.com/commnerd/SpiderWeb/src/lib/config"
	"github.com/commnerd/SpiderWeb/src/lib/util"
)

func TestMain(m *testing.M) {
	conf := config.GetConfig()
	publicFile := conf.GetString("id_rsa_pub_path")
	privateFile := conf.GetString("id_rsa_path")
	log.SetOutput(ioutil.Discard)

	m.Run()

	removeFiles(publicFile, privateFile)
}

func TestGenerate(t *testing.T) {
	privKey, pubKey := Generate()

	privateKey, err := ssh.ParseRawPrivateKey([]byte(privKey))
	if err != nil {
		log.Fatal(err)
	}

	signer, err := ssh.NewSignerFromKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := getPubKey(t, pubKey)

	assert.True(t, reflect.DeepEqual(signer.PublicKey(), publicKey))
}

func TestWriteToFile(t *testing.T) {
	conf := config.GetConfig()

	privKey, pubKey := Generate()
	WriteToFile(privKey, pubKey)

	publicFile := conf.GetString("id_rsa_pub_path")
	privateFile := conf.GetString("id_rsa_path")

	assert.True(t, util.FileExists(publicFile))
	assert.True(t, util.FileExists(privateFile))
}

func TestGenerateOverwrite(t *testing.T) {
	conf := config.GetConfig()
	privateFile := conf.GetString("id_rsa_path")

	privKey, pubKey := Generate()
	WriteToFile(privKey, pubKey)

	dat1, err := ioutil.ReadFile(privateFile)
	check(err)

	privKey, pubKey = Generate()
	WriteToFile(privKey, pubKey)

	dat2, err := ioutil.ReadFile(privateFile)
	check(err)

	assert.NotEqual(t, dat1, dat2)
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

func getPubKey(t *testing.T, pubKey string) ssh.PublicKey {
	parsedKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(pubKey))
	if err != nil {
		t.Errorf("ERROR! %s", err)
	}
	return parsedKey
}
