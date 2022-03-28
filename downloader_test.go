package mcserverdownloader

import (
    "os"
    "path"
    "testing"
)

func TestCreateDownloader(t *testing.T) {
    type args struct {
        manifestUrl string
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {
            name: "Basic_v1",
            args: args{
                manifestUrl: ManifestV1,
            },
            wantErr: false,
        },
        {
            name: "Basic_v2",
            args: args{
                manifestUrl: ManifestV2,
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            i, err := CreateDownloader(tt.args.manifestUrl)
            if (err != nil) != tt.wantErr {
                t.Errorf("CreateDownloader() error = %v, wantErr %v", err, tt.wantErr)
            } else {
                t.Logf("fetched %d versions from %s", len(i.manifest.Versions), i.manifestUrl)
            }
        })
    }
}

func TestDownloader_Download(t *testing.T) {
    type fields struct {
        manifestUrl string
    }
    type args struct {
        location  string
        versionId string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            name: "Download_Specific_Version",
            fields: fields{
                manifestUrl: ManifestV2,
            },
            args: args{
                location:  path.Join(os.TempDir(), "mc-downloader-tests"),
                versionId: "1.18.1",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            i, err := CreateDownloader(tt.fields.manifestUrl)
            if err != nil {
                t.Errorf("CreateDownloader() error = %v", err)
            }

            if err := i.Download(tt.args.location, tt.args.versionId); (err != nil) != tt.wantErr {
                t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestDownloader_DownloadLatestSnapshot(t *testing.T) {
    type fields struct {
        manifestUrl string
        manifest    Manifest
    }
    type args struct {
        location string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            name: "Download_Snapshot",
            fields: fields{
                manifestUrl: ManifestV2,
            },
            args: args{
                location: path.Join(os.TempDir(), "mc-downloader-tests"),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            i, err := CreateDownloader(tt.fields.manifestUrl)
            if err != nil {
                t.Errorf("CreateDownloader() error = %v", err)
            }

            if err := i.DownloadLatestSnapshot(tt.args.location); (err != nil) != tt.wantErr {
                t.Errorf("DownloadLatestSnapshot() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestDownloader_DownloadLatestRelease(t *testing.T) {
    type fields struct {
        manifestUrl string
        manifest    Manifest
    }
    type args struct {
        location string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            name: "Download_Release",
            fields: fields{
                manifestUrl: ManifestV2,
            },
            args: args{
                location: path.Join(os.TempDir(), "mc-downloader-tests"),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            i, err := CreateDownloader(tt.fields.manifestUrl)
            if err != nil {
                t.Errorf("CreateDownloader() error = %v", err)
            }

            if err := i.DownloadLatestRelease(tt.args.location); (err != nil) != tt.wantErr {
                t.Errorf("DownloadLatestRelease() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
