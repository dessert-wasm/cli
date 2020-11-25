package cmd

const currentVersionCLI = "dessert-cli v3.0"

const (
	logLoggingYouIn        = "logging you in..."
	logLoggingYouOut       = "logging you out..."
	logInSuccess           = "logged in"
	logOutSuccess          = "logged out"
	logPublishingToDessert = "publishing your work to the dessert platform"
	logPublishingSuccess   = "success, your work was published on dessert !"
	logMakingFolderDReady  = "making your folder dessert-ready"
	logSuccessHappyDessert = "success, happy dessert !"
	logTokenSaved          = "token saved to dessert.yml"
	logIsCore              = "project set to 'core'"
	logIsConnector         = "project set to 'connector'"
)

const (
	warnDessertConfig404 = "dessert.yml not found"
	warnPublishNPM       = "don't forget to npm publish !"
)

const (
	errPackageJSONNeeded = "package.json needed to continue"
	errDessertYMLNeeded  = "dessert.yml needed to continue"
	errPrompt            = "prompt error"
	errLogin             = "login failed"
	errGetToken          = "getting your token"
	errWriteToken        = "writing your token"
	errAddDessertKeyword = "adding dessert key to package.json"
	errBadServerResp     = "bad server response"
	errLoggingOut        = "logging you out"
	errGettingToken      = "getting token from dessert_config.yml"
	errPublishingModule  = "publishing your module"
)

const (
	fatalCookieJar = "could not create cursed cookie jar"
)
