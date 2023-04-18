//go:build test_unit

/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package iotcoremqtt

import (
	"testing"
	"time"

	"github.com/nuclio/nuclio/pkg/processor/trigger/mqtt"

	nucliozap "github.com/nuclio/zap"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	trigger iotcoremqtt
}

func (suite *TestSuite) SetupTest() {
	suite.trigger = iotcoremqtt{
		AbstractTrigger: &mqtt.AbstractTrigger{},
	}
	logger, err := nucliozap.NewNuclioZapTest("test")
	suite.Require().NoError(err)
	suite.trigger.Logger = logger.GetChild("iotcore")
	suite.trigger.configuration = &Configuration{}
}

func (suite *TestSuite) TestSetEmptyParametersNoChange() {

	// This is just an example of user's private key.
	// generated by https://travistidwell.com/jsencrypt/demo
	privateKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQB84lRLxsHTvsKGJCdjP5X7jFX53QVbBr7usOWSUWaDpTjhXaru
OvbQUMX7Z327ua+795Tbj61OrWPz5klVHm9NE8QYR9qbh8fdf+5MvSC1SuVX9eSA
/m04+6/Hj7s9pgo+EVgxP+pn84phf8jWB4aFAY6o/krM6ElhPsBxx2tmzPWrkVnJ
ccs846IyGkK/dijRQPiSjDJIWB8NXzKsNfSzkyUCJB1qKpscir/Ahs0Gmihztl7P
fCKdP8eN0aoXcQ2K18t8fpgk4lG6pU6gIplQJkwMtxogWRSjklxJ+Za2awTyKIeD
qxijsxt5L+kcEG4NGpTDgibi9KrMUYIkV+5fAgMBAAECggEAba2EZOyRC/sL8i1M
XlDY0wxD8eZXrsL06sJ3TJKq2tE/bEYMOZ+VZgyaZBpXBYiluwpMNKwxg9ip4OmN
2/CXxXSnkr+AUXnlYTXavpCXakD1vNOatGM+333DuhsUUadCdZJcBxzgLB1qmghV
BQuk66kbdnWmkeErmPk4oHSIccYjs10/8Pw8jnpjPqrm1d58ZLZzMaycMBCAwR4q
VmU6nex6qRA0+d8Og/DEm+1VxoSk/dy44dSui/OzVHEya6lz70UBf6hey4iJK6KL
i+axeP2Ec52n/ihT9MhTCAykAiw6G3VEa1duE2mmLnefEBZQc9SKxNw+77H7e+hb
tOqFEQKBgQD1mA42khXEm7flSYYsg7t9lcP0ciBntFpB7IET1C9CVObp/gcgw6gE
8Z3j80mi1Dw2mWKSaDHmyVw/D+7yQFwXTZSCbEdn8/TErVVPqvnORZyjcFYQYNlC
EwZjbLJExhDEq17aQ1WKLycmn3BbvVI9oRa2rLDZu/IHJX5KvkPMVwKBgQCCLPC6
qDrH6JEOBs8sVLFrue8+eMyMcbtdtePvpov0yH03Dh3pcLADq/HZ2t375o6xxYot
jtiMEkozHcM7pnouc46VKD40biZixnX6oeevYcWTs4nADcmVcEJagkmlYBVzqbac
F/QgVpQyDSJX95ImQPKKc5oy7JbbFIemjTapOQKBgQCzzuXpEj+ZuKCE4LW5daEZ
q0LSf5Q2GRdT2MIQMHOBTwPZIUE6vcUQCY4dzIuHCXgkSVyf8GVIoPhGu3WoK3LB
JO2sJ3aIJ1Z3gKhLMdS/Lrwl9SMtzpqCA8fTl0tViuXP99/0UQQZrbguUOFEaXIC
6SPmDr1UTIRAszSpqG+e9QKBgCqZmGINQcdABZBIjC3evX0aiP+xuobPhViCeMhp
gW2m2stUlFdbqE5bS7dWl8Siy9nDYpfMInOcXKnjuIthzKQ87tFDLTAtR+SVO/C8
YTyUy3qti4vNN0XvSdeiwYUcL4j9ZiQo9pxKmQ7UG4QcIbjhEj5a3ICDyk6Bpm7L
0bKJAoGAVJTA/d+uAOg8oPtmFbBqoHYFCgXs1OKZ0MmtDD+oTX+PL1UOhjQyxz68
LXermZHjPBHTE5xfioNxGnUBI48wEqhOxj8n7WLQMK6v3h5CPmhAjme6A/iRDjqu
NSoWKZp3rjklHIznbbSiBuRtuT47bYcpo/3IOFwKv3F3rr7zAXg=
-----END RSA PRIVATE KEY-----`

	suite.trigger.configuration.PrivateKey.Contents = privateKey
	suite.trigger.configuration.ProjectID = "something"
	suite.trigger.configuration.jwtRefreshInterval = 1 * time.Second
	issuedAt := time.Now()

	token, err := suite.trigger.createJWT(issuedAt)
	suite.Require().NoError(err)

	suite.Require().NotEmpty(token)

}

func TestRabbitMQIOTCoreSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
