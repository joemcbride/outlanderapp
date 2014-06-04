package main

import (
    "encoding/json"
    "time"
)

var downloadName = "Outlander"
var downloadNotes = ""
var downloadUrl = "https://github.com/joemcbride/outlander-osx/releases/download"
var releaseNotesUrl = "https://github.com/joemcbride/outlander-osx/releases/tag"
var downloadVersion = "v0.3"
var downloadFile = "Outlander.app.zip"

type Version struct {
    Name string
    Notes string
    ReleaseNotesUrl string
    Version string
    PublishDate time.Time
    Url string
}

type Response map[string]interface{}

func (r Response) String() (s string) {
    b, err := json.Marshal(r)
    if err != nil {
            s = ""
            return
    }
    s = string(b)
    return
}

func version_data () (Version) {
    return Version {
        Name: downloadName + " " + downloadVersion,
        Notes: downloadNotes,
        ReleaseNotesUrl: releaseNotesUrl + "/" + downloadVersion,
        Version: downloadVersion,
        PublishDate: time.Now(),
        Url: version_download(),
    }
}

func version_response (ver Version) (Response) {

	dt, _ := ISO8601FullFromTime(ver.PublishDate)

	return Response {
        "url": ver.Url,
        "name": ver.Name,
        "notes": ver.Notes,
        "pub_date": dt,
    }
}

func version_download () (string) {
    return downloadUrl + "/" + downloadVersion + "/" + downloadFile
}