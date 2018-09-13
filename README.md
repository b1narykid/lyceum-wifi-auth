# Scripts and utils for HSE Wireless Access Points

## Synopsis

HSE facilities usually have two wireless networks with *HSE* and *hse.ru* SSIDs.
Both of them use Captive Portal for authentication, but the former requires
credentials input (username:*hseguest*|password:*hsepassword*) each hour or so.
This project contains utils for automated HSE Captive Portal authentication.

Note that their Captive Portal rejects requests without the following queries:
- `Submit=Submit&buttonClicked=4` for auth
- `Logout=Logout` for logout

### Expirements

- Investigate how `hse.ru` authenicates users.
- DNS port may not be restricted.

## License

No, it's public domain.

## Warning

Any use of this software is entirely at your own risk.
