package private_key

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/rs/zerolog/log"
)

func ReadPrivateKeyFile(privateKeyPath string) *rsa.PrivateKey {
	// Open private key file
	privateKeyFile, err := os.Open(privateKeyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("on reading private key file")
	}
	defer privateKeyFile.Close()

	// Allocate Memory
	pemfileinfo, _ := privateKeyFile.Stat()
	pembytes := make([]byte, pemfileinfo.Size())

	// Load file into memory
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		log.Fatal().Err(err).Msg("on reading private key data")
	}

	// Decode byte into private key object
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyImported, err := x509.ParsePKCS8PrivateKey(data.Bytes)
	if err != nil {
		log.Fatal().Err(err).Msg("on parsing private key")
	}
	return privateKeyImported.(*rsa.PrivateKey)
}
