package main

import (
	"github.com/bboortz/highlander-api/pkg"
)

func main() {
	// initialize config with default settings
	c := pkg.Conf{
		Port:          ":8443",
		ReadTimeout:   5,
		WriteTimeout:  10,
		TLSCipher:     "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256|TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
		TLSMinVersion: "VersionTLS12",
		TLSCert:       "./examples/certs/cert.pem",
		TLSKey:        "./examples/certs/key.pem",
	}

	// load the config
	c.LoadConf("highlander.yaml")

	pkg.LogStartup(c)
	pkg.ListenAPI(c)
}
