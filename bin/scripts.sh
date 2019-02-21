BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )";
ROOT_DIR="$(dirname "$BIN_DIR")";
SRC_DIR="$BIN_DIR/src";
echo $SRC_DIR;

build () {
    echo 'build goserver'
    docker build -t goserver -f Dockerfile .
}

start () {
    echo 'starting goserver'
    echo $SRC_DIR;
    docker run --rm -p 3001:3001 -it --name goserver --mount source=$SRC_DIR,destination=/src,type=bind goserver
}