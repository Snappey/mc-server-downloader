## mc-server-downloader
___
Simple library to help install a minecraft server version (release or snapshot)

``go get -u github.com/Snappey/mc-server-downloader``

## Example
___

```go
package awesomeProject

import (
    serverDownloader "github.com/Snappey/mc-server-downloader"
)

func main() {
    downloader, err := serverDownloader.CreateDownloader(serverDownloader.ManifestV2)
    if err != nil {
        panic(err)
    }

    err = downloader.DownloadLatestRelease("./downloads")
    if err != nil {
        panic(err)
    }
}
```