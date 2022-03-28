package mcserverdownloader

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path"
)

type Downloader struct {
    manifestUrl string
    manifest    Manifest
}

const (
    ManifestV2 = "https://launchermeta.mojang.com/mc/game/version_manifest_v2.json"
    ManifestV1 = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
)

func CreateDownloader(manifestUrl string) (Downloader, error) {
    installer := Downloader{
        manifestUrl: manifestUrl,
    }

    if err := installer.load(); err != nil {
        return installer, err
    }

    return installer, nil
}

func (i Downloader) DownloadLatestSnapshot(installLocation string) error {
    return i.Download(installLocation, i.manifest.Latest.Snapshot)
}

func (i Downloader) DownloadLatestRelease(installLocation string) error {
    return i.Download(installLocation, i.manifest.Latest.Release)
}

func (i Downloader) Download(installLocation string, versionId string) error {
    version, err := i.fetchVersionDetails(versionId)
    if err != nil {
        return err
    }

    return i.fetchServer(installLocation, version)
}

func (i *Downloader) load() error {
    res, err := http.Get(i.manifestUrl)
    if err != nil {
        return err
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err
    }

    err = json.Unmarshal(body, &i.manifest)
    if err != nil {
        return err
    }

    return nil
}

func (i *Downloader) fetchServer(installLocation string, versionDetails Version) error {
    res, err := http.Get(versionDetails.Downloads.Server.Url)
    if err != nil {
        return err
    }

    if res.StatusCode != http.StatusOK {
        return fmt.Errorf("error fetching version, url=%s", versionDetails.Downloads.Server.Url)
    }

    serverBytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return err
    }

    err = os.MkdirAll(installLocation, 0755)
    if err != nil {
        return err
    }

    fileName := fmt.Sprintf("server_%s.jar", versionDetails.Id)
    return os.WriteFile(path.Join(installLocation, fileName), serverBytes, os.ModeAppend)
}

func (i *Downloader) fetchVersionDetails(versionId string) (Version, error) {
    var version ManifestVersion
    for _, manifestVersion := range i.manifest.Versions {
        if manifestVersion.Id == versionId {
            version = manifestVersion
            break
        }
    }
    if version == (ManifestVersion{}) {
        return Version{}, fmt.Errorf("failed to find version, version=%s", versionId)
    }

    res, err := http.Get(version.Url)
    if err != nil {
        return Version{}, err
    }

    if res.StatusCode != http.StatusOK {
        return Version{}, fmt.Errorf("error fetching version, url=%s", version.Url)
    }

    versionBytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return Version{}, err
    }

    var versionDetails Version
    err = json.Unmarshal(versionBytes, &versionDetails)
    if err != nil {
        return Version{}, err
    }

    return versionDetails, nil
}
