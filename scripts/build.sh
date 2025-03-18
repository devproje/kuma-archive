INPUT_ARGS=$1
LDFLAGS="-X main.version=$(jq -r .version package.json) -X main.hash=$(git rev-parse --short=7 HEAD) -X main.branch=$(git rev-parse --abbrev-ref HEAD)"
FLAGS=""

if [[ $INPUT_ARGS == "--api-only" ]]; then
	FLAGS="--api-only"
fi

if [[ $INPUT_ARGS == "--run" ]]; then
	INPUT_ARGS_2=$2
	if [[ $INPUT_ARGS_2 == "--api-only" ]]; then
		FLAGS="--api-only"
	fi

	go run -ldflags "${LDFLAGS}" ./app.go daemon -d ${FLAGS}
	exit 0
fi

go build -ldflags "${LDFLAGS}" -o kuma-archive ${FLAGS}
