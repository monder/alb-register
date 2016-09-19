#!/bin/sh

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
export VERSION=v0.1.0

go build -ldflags '-extldflags "-static"'

DIR=$(mktemp -d)
echo $DIR
echo '
{
    "acKind": "ImageManifest",
    "acVersion": "0.8.5",
    "annotations": [
        {
            "name": "authors",
            "value": "Aleksejs Sinicins <monder@monder.cc>"
        }
    ],
    "app": {
        "exec": [
            "/bin/alb-register"
        ],
        "group": "0",
        "user": "0"
    },
    "labels": [
        {
            "name": "version",
            "value": "'"${VERSION}"'"
        },
        {
            "name": "arch",
            "value": "'"${GOARCH}"'"
        },
        {
            "name": "os",
            "value": "'"${GOOS}"'"
        }
    ],
    "name": "monder.cc/alb-register"
}        
' > $DIR/manifest
mkdir -p $DIR/rootfs/bin/ $DIR/rootfs/etc/ssl/certs/
cp alb-register $DIR/rootfs/bin/
curl -L -o $DIR/rootfs/etc/ssl/certs/ca-certificates.crt https://curl.haxx.se/ca/cacert.pem
tar -zcvf alb-register-${VERSION}-${GOOS}-${GOARCH}.aci -C $DIR -s '/.//' . 
rm -rf $DIR
