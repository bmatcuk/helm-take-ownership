#!/usr/bin/env bash

# some ideas borrowed from: https://github.com/nouney/helm-gcs/blob/0.1.1/install.sh

if type "curl" >/dev/null 2>&1; then
  GET_LATEST="curl -s"
  DOWNLOAD="curl -sSLO"
elif type "wget" >/dev/null 2>&1; then
  GET_LATEST="wget -q -O -"
  DOWNLOAD="wget -q"
else
  echo "curl or wget are required to install!"
  exit 1
fi

fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
    echo "Failed to install helm-take-ownership"
    echo "\tSee https://github.com/bmatcuk/helm-take-ownership"
  fi
  exit $result
}

trap "fail_trap" EXIT
set -e

os=`uname -s`
arch=`uname -m`
url=`$GET_LATEST https://api.github.com/repos/bmatcuk/helm-take-ownership/releases/latest | awk '/browser_download_url/ { print $2 }' | sed 's/"//g' | grep ${os}_${arch}`
filename=`echo $url | rev | cut -d '/' -f 1 | rev`

echo "* Downloading $url"
$DOWNLOAD $url

echo "* Extracting $filename"
rm -rf bin && mkdir bin && tar xzvf $filename -C bin >/dev/null && rm -f $filename
echo "* helm-take-ownership installed!\n"
bin/helm-take-ownership --help
