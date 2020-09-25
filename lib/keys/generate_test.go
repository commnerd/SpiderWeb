package keys

import (
	"github.com/stretchr/testify/assert"
	"../config"
	"io/ioutil"
	"../util"
	"testing"
	"log"
	"os"
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
	conf := config.GetConfig()

	Generate()

	publicFile := conf.GetString("id_rsa_pub_path")
	privateFile := conf.GetString("id_rsa_path")

	assert.True(t, util.FileExists(publicFile))
	assert.True(t, util.FileExists(privateFile))
}

func TestGenerateOverwrite(t *testing.T) {
	conf := config.GetConfig()
	privateFile := conf.GetString("id_rsa_path")

	Generate()

	dat1, err := ioutil.ReadFile(privateFile)
	check(err)

	Generate()

	dat2, err := ioutil.ReadFile(privateFile)
	check(err)

	assert.NotEqual(t, dat1, dat2)
}

func removeFiles(files ...string) {
	for _, file := range files {
		err := os.Remove(file)
		check(err)
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}