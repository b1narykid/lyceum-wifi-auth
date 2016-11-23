#!/bin/sh

CURL_BIN="$(which curl)"
if [ -z "$CURL_BIN" ]; then exit 0; fi

opt() {
	printf "$1=$2"
}

ACTION_URL="https://wlc38.hse.ru/login.html"

REDIRECT="http://www.hse.ru"

USERNAME="hseguest"
PASSWORD="hsepassword"

# AP rejects requests without this option
OPT_FMSUBMIT="$(opt Submit        "Submit")"
# AP will not authorize user without this option
OPT_BTNCLCKD="$(opt buttonClicked "4")"
# Looks like `redirect_url` is optional (yay, options are optional)
OPT_REDIRECT="$(opt redirect_url  "$REDIRECT")"
# -- Credentials --
OPT_USERNAME="$(opt username      "$USERNAME")"
OPT_PASSWORD="$(opt password      "$PASSWORD")"

$CURL_BIN -L "$ACTION_URL" --data "$OPT_USERNAME&$OPT_PASSWORD&$OPT_REDIRECT&$OPT_BTNCLCKD&$OPT_FMSUBMIT"

exit
