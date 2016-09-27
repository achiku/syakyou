require "openssl"

nonce = "100"
url = "https://coincheck.com/api/ec/buttons"
body = "hoge=foo&bar=boo"
message = nonce + url + body
secret = "foo"
p OpenSSL::HMAC.hexdigest(OpenSSL::Digest.new("sha256"), secret, message)
# "9d6bbf1897701ec73bc80c395e98778756cb009164a949d62aaa6bb99e1223e1"
