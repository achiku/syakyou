#!/usr/bin/env python
# -*- coding: utf-8 -*-
import hmac
import hashlib
import base64


if __name__ == '__main__':
    nonce = "100"
    url = "https://coincheck.com/api/ec/buttons"
    body = "hoge=foo&bar=boo"
    message = nonce + url + body
    secret = "foo"
    dig = hmac.new(secret, msg=message, digestmod=hashlib.sha256).hexdigest()
    print(dig)
    # 9d6bbf1897701ec73bc80c395e98778756cb009164a949d62aaa6bb99e1223e1
