version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Downloading and installing helm-kill ${version} ..."

url=""
if [ "$(uname)" = "Darwin" ]; then
    url="https://github.com/rollerd/helm-kill-plugin/releases/download/${version}/helm-kill_${version}_darwin_amd64.tar.gz"
fi

echo $url

mkdir -p "bin"

# Download with curl if possible.
if [ -x "$(which curl 2>/dev/null)" ]; then
    curl -sSL "${url}" -o "bin/${version}.tar.gz"
else
    wget -q "${url}" -O "bin/${version}.tar.gz"
fi

tar xzf "bin/${version}.tar.gz" -C "bin/"
