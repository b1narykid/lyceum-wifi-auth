HSE WiFi AP authentication scripts and utils
============================================

# Synopsis

HSE facilities usually have two wireless networks with SSIDs *hse.ru* and *HSE*.
Both of them use Captive Portal for authentication, but the last one requires
credentials (username:*hseguest*/password:*hsepassword*) each hour or so.
This project contains utils for automated *HSE* Captive Portal authentication.

### HTTP request example

__TBD__

Note that their Captive Portal rejects requests without queries below:
- `Submit=Submit` query for auth,
- `Logout=Logout` for logout,
- `buttonClicked=4` for auth,

### Expirements (TODO)

- Investigate how `hse.ru` authenicates users.
- DNS port may not be restricted.

# License

Copyright (c) 2016-2018 Ivan Trubach &lt;dev at b1nary.tk&gt;

Licensed and published under the MIT license.
See [license file] for more information.

## Warning

Any use of this software is entirely at your own risk.

[license file]: LICENSE.txt
