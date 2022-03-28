package main

import (
    "time"
)

type ManifestVersion struct {
    Id              string    `json:"id"`
    Type            string    `json:"type"`
    Url             string    `json:"url"`
    Time            time.Time `json:"time"`
    ReleaseTime     time.Time `json:"releaseTime"`
    Sha1            string    `json:"sha1,omitempty"`
    ComplianceLevel int       `json:"complianceLevel,omitempty"`
}

type ManifestLatest struct {
    Release  string `json:"release"`
    Snapshot string `json:"snapshot"`
}

type Manifest struct {
    Latest   ManifestLatest    `json:"latest"`
    Versions []ManifestVersion `json:"versions"`
}

type Version struct {
    Id          string                     `json:"id"`
    Type        string                     `json:"type"`
    Time        time.Time                  `json:"time"`
    ReleaseTime time.Time                  `json:"releaseTime"`
    Downloads   VersionDownloads           `json:"downloads"`
}

type VersionDownloads struct {
    Client         VersionDownloadDetails `json:"client"`
    ClientMappings VersionDownloadDetails `json:"client_mappings"`
    Server         VersionDownloadDetails `json:"server"`
    ServerMappings VersionDownloadDetails `json:"server_mappings"`
}

type VersionDownloadDetails struct {
    Size int    `json:"size"`
    Url  string `json:"url"`
    Sha1 string `json:"sha1"`
}
