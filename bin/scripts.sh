BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )";
ROOT_DIR="$(dirname "$BIN_DIR")";
SRC_DIR="$BIN_DIR/src";

buildServer () {
    echo 'build goserver'
    docker build -t goserver -f Dockerfile .
}

startServer () {
    echo 'starting goserver'
    docker run --rm -p 3001:3001 -it --name goserver --mount source=$SRC_DIR,destination=/go/app/src,type=bind goserver
}