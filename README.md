# Scripts and utils for HSE Wireless Access Points

## Synopsis

HSE facilities usually have two wireless networks with *HSE* and *hse.ru* SSIDs.
Both of them use captive portal for authentication, but the former requires
credentials input each hour or so.  This project contains utils for automated
HSE captive portal authentication.

# Captive Portal Query

## HSE

### Login

Request URL `https://wlc38.hse.ru/login.html` works fine.  Though I've seen other URLs (e.g. `wlc37`).

- `Submit`: required, value is `Submit`
- `buttonClicked`: required, value is `4`
- `username`: required, value is `hseguest`
- `password`: required, value is `hsepassword`
- `redirect_url`: optional, redirect URL (defunct)

### Logout

Request URL is `https://wlc38.hse.ru/logout.html`.

- `Logout`: required, value is `Logout`.

## hse.ru

### Login

Request URL is `https://cp.hse.ru:8443/hse-captive-portal/login`.

- `username`: required, user's name
- `password`: required, user's password
- `usergroup`: required, user's group. Well, afaik defaults to hse.ru but dgaf.
   Values:
   - 1: hse.ru
   - 2: edu.hse.ru
- `cookie-enabled`: perhaps the server sends a cookie in response.  Corresponds
  to the *Запомнить*/*Remember me* option on auth web page.

### Logout

Unfortunately, I didn't find the request URL.  Looks like even the web interface
does not support logout.


# Legal notice

## License

No, it's public domain.

## Warning

Any use of this software is entirely at your own risk.
