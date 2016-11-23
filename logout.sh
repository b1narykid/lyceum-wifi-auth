#!/bin/sh

CURL_BIN="$(which curl)"
if [ -z "$CURL_BIN" ]; then exit 0; fi

opt() {
	printf "$1=$2"
}

ACTION_URL="https://wlc38.hse.ru/logout.html"

# AP rejects requests without this option
OPT_LOGOUT="$(opt Logout "Logout")"

$CURL_BIN -L "$ACTION_URL" --data "$OPT_LOGOUT"

exit
