package config

import (
	"github.com/mkboudreau/loggo"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strconv"
	"testing"
)

func TestBasicMailConfigConstruction(t *testing.T) {
	testLogger := loggo.DefaultLevelLogger()
	testLogger.SetInfo()

	Convey("Site Parsing Tests", t, func() {
		logger.SetTrace()

		Convey("builds a smtp with passed in values", func() {
			user := "myuser"
			pass := "mypass"
			host := "myhost"
			port := 0xCADE
			smtpconfig := NewSmtpConfigWithValues(host, port, user, pass)

			So(smtpconfig.Host, ShouldEqual, host)
			So(smtpconfig.Port, ShouldEqual, port)
			So(smtpconfig.User, ShouldEqual, user)
			So(smtpconfig.Pass, ShouldEqual, pass)

		})
		Convey("builds a smtp server from environment", func() {
			user := "aaaa"
			pass := "bbbb"
			host := "cccc"
			port := 1232

			if isEnvironmentSet() {
				testLogger.Info("Environment Is Set")
				smtpconfig := NewSmtpConfigFromEnvironment()

				testLogger.Debug("host:", smtpconfig.Host, "port:", smtpconfig.Port, "user:", smtpconfig.User, "pass:", smtpconfig.Pass)
				So(smtpconfig.Host, ShouldNotBeNil)
				So(smtpconfig.Port, ShouldNotBeNil)
				So(smtpconfig.User, ShouldNotBeNil)
				So(smtpconfig.Pass, ShouldNotBeNil)
			} else {
				testLogger.Info("Environment Not Set... Setting")
				os.Setenv(ENV_SMTPHOST, host)
				os.Setenv(ENV_SMTPPORT, strconv.Itoa(port))
				os.Setenv(ENV_SMTPUSER, user)
				os.Setenv(ENV_SMTPPASS, pass)

				smtpconfig := NewSmtpConfigFromEnvironment()

				So(smtpconfig.Host, ShouldEqual, host)
				So(smtpconfig.Port, ShouldEqual, port)
				So(smtpconfig.User, ShouldEqual, user)
				So(smtpconfig.Pass, ShouldEqual, pass)
			}

		})

	})
}

func isEnvironmentSet() bool {
	host := os.Getenv(ENV_SMTPHOST)
	return host != ""
}
