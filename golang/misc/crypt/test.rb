require "openssl"

nonce = "100"
url = "https://coincheck.com/api/ec/buttons"
body = "hoge=foo&bar=boo"
message = nonce + url + body
secret = "foo"
p OpenSSL::HMAC.hexdigest(OpenSSL::Digest.new("sha256"), secret, message)
