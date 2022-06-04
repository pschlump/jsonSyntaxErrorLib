# Tools

## Overview

1. gen64ByteRandomValue - JWT tokens use a HMAC encryption/validation - this generates a random key for the HMAC of the correct length.
1. jsonCheck - Parse a JSON file and verify that it has certain keys, or that the reverse that certain keys are not present.
1. jsonStatus - check the "status" field in a JSON response to see if it is success.
1. jwtParse - Parse a JWT token and extract it's contents (relies on having the HMAC key)

<!--

## Other Useful Tools

1. `~/bin/acc`  - command line tool for reading QR codes and generating one time passwords.  From:
`~/go/src/github.com/pschlump/htotp_acc`.
2. `../qrcode` - a tool to generate QR codes in `.png` or `.svg` format.
3. qr-decode - a tool to read a `.png` QR code and extract the text.
3. wget - web client
4. curl - web client

-->
