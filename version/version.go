package version

import "log"

// -----------------------------------------------------------------------------

/**
 * -ldflags "-X 'main.Version=v${VERSION}'"
 */
var Version = "development"

/**
 * -ldflags "-X 'main.Executable=v${EXECUTABLE}'"
 */
var Executable = "app"

// -----------------------------------------------------------------------------

/**
 * -ldflags "-X 'main.BuildNumber=v${BUILD_NUMBER}'"
 */
var BuildNumber = "0"

/**
 * -ldflags "-X 'main.BuildNumber=v${BUILD_BRANCH}'"
 */
var BuildBranch string

/**
 * -ldflags "-X 'main.BuildHash=v${BUILD_HASH}'"
 */
var BuildHash string

/**
 * -ldflags "-X 'main.BuildTime=v${BUILD_TIME}'"
 */
var BuildTime string

/**
 * -ldflags "-X 'main.BuildUser=v${BUILD_USER}'"
 */
var BuildUser string

// -----------------------------------------------------------------------------

func BuildPrintln() {
	log.Println("Version:\t\t", Version)
	log.Println("Executable:\t\t", Executable)
	log.Println("Build:\t\t", BuildNumber)
	log.Println("Branch:\t\t", BuildBranch)
	log.Println("Hash:\t\t", BuildHash)
	log.Println("Time:\t\t", BuildTime)
	log.Println("User:\t\t", BuildUser)
}

// -----------------------------------------------------------------------------
// end.
