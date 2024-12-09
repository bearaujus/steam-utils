package model

const (
	LibraryAutoUpdateAlwaysKeepAll = "0"
	LibraryAutoUpdateOnlyOnLaunch  = "1"
)

var LibraryAutoUpdate = map[string]string{
	LibraryAutoUpdateAlwaysKeepAll: "Keep all games updated automatically",
	LibraryAutoUpdateOnlyOnLaunch:  "Only update a game when you launch it",
}

var LibraryAutoUpdateR = getReversedKeyValueMap(LibraryAutoUpdate)

const (
	LibraryBackgroundDownloadsFollowGlobal = "0"
	LibraryBackgroundDownloadsAlwaysAllow  = "1"
	LibraryBackgroundDownloadsNeverAllow   = "2"
)

var LibraryBackgroundDownloads = map[string]string{
	LibraryBackgroundDownloadsFollowGlobal: "Use the global Steam settings for background downloads",
	LibraryBackgroundDownloadsAlwaysAllow:  "Always allow background downloads to run",
	LibraryBackgroundDownloadsNeverAllow:   "Never allow background downloads to run",
}

var LibraryBackgroundDownloadsR = getReversedKeyValueMap(LibraryBackgroundDownloads)

func getReversedKeyValueMap(src map[string]string) map[string]string {
	ret := make(map[string]string, len(src))
	for k, v := range src {
		ret[v] = k
	}
	return ret
}
