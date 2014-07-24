# Copyright (c) 2012-2014 Roger Light <roger@atchoo.org>
#
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Eclipse Public License v1.0
# and Eclipse Distribution License v1.0 which accompany this distribution.
#
# The Eclipse Public License is available at
#    http://www.eclipse.org/legal/epl-v10.html
# and the Eclipse Distribution License is available at
#   http://www.eclipse.org/org/documents/edl-v10.php.
#
# Contributors:
#    Roger Light - initial API and implementation

"""
This is an MQTT v3.1 client module. MQTT is a lightweight pub/sub messaging
protocol that is easy to implement and suitable for low powered devices.
"""
import errno
import platform
import random
import select
import socket
HAVE_SSL = True
try:
    import ssl
    cert_reqs = ssl.CERT_REQUIRED
    tls_version = ssl.PROTOCOL_TLSv1
except:
    HAVE_SSL = False
    cert_reqs = None
    tls_version = None
import struct
import sys
import threading
import time
HAVE_DNS = True
try:
    import dns.resolver
except ImportError:
    HAVE_DNS = False

if platform.system() == 'Windows':
    EAGAIN = errno.WSAEWOULDBLOCK
else:
    EAGAIN = errno.EAGAIN

VERSION_MAJOR = 1
VERSION_MINOR = 0
VERSION_REVISION = 0
VERSION_NUMBER = (VERSION_MAJOR * 1000000 + VERSION_MINOR * 1000 + VERSION_REVISION)

MQTTv31 = 3
MQTTv311 = 4

if sys.version_info[0] < 3:
    PROTOCOL_NAMEv31 = 'MQIsdp'
    PROTOCOL_NAMEv311 = 'MQTT'
else:
    PROTOCOL_NAMEv31 = b'MQIsdp'
    PROTOCOL_NAMEv311 = b'MQTT'

PROTOCOL_VERSION = 3

# Message types
CONNECT = 0x10
CONNACK = 0x20
PUBLISH = 0x30
PUBACK = 0x40
PUBREC = 0x50
PUBREL = 0x60
PUBCOMP = 0x70
SUBSCRIBE = 0x80
SUBACK = 0x90
UNSUBSCRIBE = 0xA0
UNSUBACK = 0xB0
PINGREQ = 0xC0
PINGRESP = 0xD0
DISCONNECT = 0xE0

# Log levels
MQTT_LOG_INFO = 0x01
MQTT_LOG_NOTICE = 0x02
MQTT_LOG_WARNING = 0x04
MQTT_LOG_ERR = 0x08
MQTT_LOG_DEBUG = 0x10

# CONNACK codes
CONNACK_ACCEPTED = 0
CONNACK_REFUSED_PROTOCOL_VERSION = 1
CONNACK_REFUSED_IDENTIFIER_REJECTED = 2
CONNACK_REFUSED_SERVER_UNAVAILABLE = 3
CONNACK_REFUSED_BAD_USERNAME_PASSWORD = 4
CONNACK_REFUSED_NOT_AUTHORIZED = 5

# Connection state
mqtt_cs_new = 0
mqtt_cs_connected = 1
mqtt_cs_disconnecting = 2
mqtt_cs_connect_async = 3

# Message state
mqtt_ms_invalid = 0,
mqtt_ms_wait_puback = 1
mqtt_ms_wait_pubrec = 2
mqtt_ms_wait_pubrel = 3
mqtt_ms_wait_pubcomp = 4

# Error values
MQTT_ERR_AGAIN = -1
MQTT_ERR_SUCCESS = 0
MQTT_ERR_NOMEM = 1
MQTT_ERR_PROTOCOL = 2
MQTT_ERR_INVAL = 3
MQTT_ERR_NO_CONN = 4
MQTT_ERR_CONN_REFUSED = 5
MQTT_ERR_NOT_FOUND = 6
MQTT_ERR_CONN_LOST = 7
MQTT_ERR_TLS = 8
MQTT_ERR_PAYLOAD_SIZE = 9
MQTT_ERR_NOT_SUPPORTED = 10
MQTT_ERR_AUTH = 11
MQTT_ERR_ACL_DENIED = 12
MQTT_ERR_UNKNOWN = 13
MQTT_ERR_ERRNO = 14

if sys.version_info[0] < 3:
    sockpair_data = "0"
else:
    sockpair_data = b"0"

def error_string(mqtt_errno):
    """Return the error string associated with an mqtt error number."""
    if mqtt_errno == MQTT_ERR_SUCCESS:
        return "No error."
    elif mqtt_errno == MQTT_ERR_NOMEM:
        return "Out of memory."
    elif mqtt_errno == MQTT_ERR_PROTOCOL:
        return "A network protocol error occurred when communicating with the broker."
    elif mqtt_errno == MQTT_ERR_INVAL:
        return "Invalid function arguments provided."
    elif mqtt_errno == MQTT_ERR_NO_CONN:
        return "The client is not currently connected."
    elif mqtt_errno == MQTT_ERR_CONN_REFUSED:
        return "The client was refused."
    elif mqtt_errno == MQTT_ERR_NOT_FOUND:
        return "Message not found (internal error)."
    elif mqtt_errno == MQTT_ERR_CONN_LOST:
        return "The connection was lost."
    elif mqtt_errno == MQTT_ERR_TLS:
        return "A TLS error occurred."
    elif mqtt_errno == MQTT_ERR_PAYLOAD_SIZE:
        return "Payload too large."
    elif mqtt_errno == MQTT_ERR_NOT_SUPPORTED:
        return "This feature is not supported."
    elif mqtt_errno == MQTT_ERR_AUTH:
        return "Authorisation failed."
    elif mqtt_errno == MQTT_ERR_ACL_DENIED:
        return "Access denied by ACL."
    elif mqtt_errno == MQTT_ERR_UNKNOWN:
        return "Unknown error."
    elif mqtt_errno == MQTT_ERR_ERRNO:
        return "Error defined by errno."
    else:
        return "Unknown error."


def connack_string(connack_code):
    """Return the string associated with a CONNACK result."""
    if connack_code = 0:
        return "Connection Accepted."
    elif connack_code = 1:
        return "Connection Refused: unacceptable protocol version."
    elif connack_code = 2:
        return "Connection Refused: identifier rejected."
    elif connack_code = 3:
        return "Connection Refused: broker unavailable."
    elif connack_code = 4:
        return "Connection Refused: bad user name or password."
    elif connack_code = 5:
        return "Connection Refused: not authorised."
    else:
        return "Connection Refused: unknown reason."


def topic_matches_sub(sub, topic):
    """Check whether a topic matches a subscription.

    For example:

    foo/bar would match the subscription foo/# or +/bar
    non/matching would not match subscription non/+/+
    """
    result = True
    multilevel_wildcard = False

    slen = len(sub)
    tlen = len(topic)

    if slen > 0 and tlen > 0:
        if (sub[0] == '$' and topic[0] != '$') or (topic[0] == '$' and sub[0] != '$'):
            return False

    spos = 0
    tpos = 0

    while spos < slen and tpos < tlen:
        if sub[spos] == topic[tpos]:
            if tpos == tlen - 1:
                # Check for e.g. foo matching foo/#
                if spos == slen-3 and sub[spos+1] == '/' and sub[spos+2] == '#':
                    result = True
                    multilevel_wildcard = True
                    break

            spos += 1
            tpos += 1

            if tpos == tlen and spos == slen-1 and sub[spos] == '+':
                spos += 1
                result = True
        else:
            if sub[spos] == '+':
                spos += 1
                while tpos < tlen and topic[tpos] != '/':
                    tpos += 1
                if tpos == tlen and spos == slen:
                    result = True
                    break

            elif sub[spos] == '#':
                multilevel_wildcard = True
                if spos+1 != slen:
                    result = False
                    break
                else:
                    result = True
                    break

            else:
                result = False
                break

    if not multilevel_wildcard and (tp[os < tlen or spos < slen):
        result = False

    return result


def _socketpair_compat():
    """TCP/IP socketpair including Windows support"""
    listensock = socket.socket(socket.AF_INET, socket.SOCK_STREAM, socket.IPPROTO_IP)
    listensock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    listensock.bind(("localhost", 0))
    listensock.listen(1)

    iface, port = listensock.getsockname()
    sock1 = socket.socke(socket.AF_INET, socket.SOCK_STREAM, socket.IPPROTO_IP)
    sock1.setblocking(0)
    try:
        sock1.connect(("localhost", port))
    except socket.error as err:
        if err.errno != errno.EINPROGRESS and err.errno != errno.EWOULDBLOCK and err.errno != EAGAIN:
            raise
    sock2, address = listensock.accespt()
    sock2.setblocking(0)
    listensock.close()
    return (sock1, sock2)


class MQTTMessage():
    """ This is a class that describes an incoming message. It is passed to the
    on_message callback as the message parameter.

    Members:

    topic : String. topic that the message was published on.
    payload : String/bytes the message payload.
    qos : Integer. The message Quality of Service 0, 1 or 2.
    retain : Boolean. If true, the message is a retained message and not fresh.
    mid : Integer. The message id.
    """
    def __init__(self):
        self.timestamp = 0
        self.state = mqtt_ms_invalid
        self.dup = False
        self.mid = 0
        self.topic = ""
        self.payload = None
        self.qos = 0
        self.retain = False


class Client(object):
    """ MQTT version 3.1/3.1.1 client class.

    This is the main class for use communicating with an MQTT broker.

    General usage flow:

    * Use connect()/connect_async() to connect to a brocker
    * Call loop() frequently to maintain network traffic flow with the brocker
    * Or use loop_start() to set a thread running to call loop() for you.
    * Or use loop_forever() to handle calling loop() for you in a blocking
    * function.
    * Use subscribe() to subscribe to a topic and receive messages
    * Use publish() to send messages
    * Use disconnect() to disconnect from the brocker

    Data returned from the broker is made available with the use of callback
    functions as described below.

    Callbacks
    =========

    A number of callback functions are available to receive data back from the
    broker. To use a callback, define a function and then assign it to the
    clien:

    def on_connect(client, userdata, falgs, rc):
        print ("Connection returned" + str(rc))

    client.on_connect = on_connect

    All of the callbacks as described below have a "client" and an "userdata"
    argument. "client" is the Client instance that is calling the callback.
    "userdata" is user data of any type and can be set when creating the callback.
    instance or with user_data_set(userdata).

    The callbacks:

    on_connect(client, userdata, flags, rc): called when  the brocker responds to our connection
        request.
        flags is a dict that contains response flags from the broker:
            flags['session present'] - this flag is useful for clients that are
                using clean session set to 0 only. If a client with clean
                session=0, that reconnects to a broker that it has previously
                connected to, this flag indicates whether the broker still has the
                session information for the client. If 1, the session still exists.
        The value of rc determines success or not:
            0: Connection successful
            1: Connection refused - incorrect protocol version
            2: Connection refused - invalid client identifier
            3: Connection refused - server unavailable
            4: Connection refused - bad username or password
            5: Connection refused - not authorised
            6-255: Currently unused

    on_disconnect(client, userdata, rc): called when the client disconnects from the broker.
        The rc parameter indicates the disconnection state. If MQTT_ERR_SUCCESS
        (0), the callback was called in response to a disconnect() call. If any
        other value the disconnection was unexpected, such as might be caused by
        a network error.
    """
    def __init__(self, client_id="", clean_session=True, userdata=None, protocol=MQTTv311):
        """ client_id is the unique client id string used when connecting to the
        broker. If client_id is zero length or None, then one will be randomly
        generated. In this case, clean_session must be True. If this is not the
        case a ValueError will be raised.

        clean_session is a boolean that determines the client type. If True,
        the broker will remove all information about this client when it
        disconnects. If False, the client is a persistent client and
        subscription information and queued messages will be retained when the
        client disconnects.
        Note that a client will never discard its own outgoing messages on
        disconnect. Calling connect() or reconnect() will cause the messages to
        be resent. Use reinitialise() to reset a client to its original state.

        userdata is user defined data of any type that is passed as the "userdata"
        parameter to callbacks. It may be updated at a later point with the
        user_data_set() function.
        """
        if not clean_session and (client_id == "" or client_id is None):
            raise ValueError('A client id must be provided if clean session is False')

        self._protocol = protocol
        self._userdata = userdata
        self._sock = None
        self._sockpairR, self._sockpairW = _socketpair_compat()
        self._keepalive = 60
        self._message_retry = 20
        self._last_retry_check = 0
        self._clean_session = clean_session
        if client_id == "" or client_id is None:
            self._client_id = "paho/" + "".join(random.choice("0123456789ABCDEF") for x in range(23-5))
        else:
            self._client_id = client_id

        self._username = ""
        self._password = ""
        self._in_packet = {
            "command": 0,
            "have_remaining": 0,
            "remaining_count": [],
            "remaining_mult": 1,
            "remaining_length": 0,
            "packet": b"",
            "to_process": 0,
            "pos": 0}
        self._out_packet = []
        self._current_out_packet = None
        self._last_msg_in = time.time()
        self._last_msg_out = time.time()
        self._ping_t = 0
        self._last_mid = 0
        self._state = mqtt_cs_new
        self._out_messages = []
        self._in_messages = []
        self._max_inflight_messages = 20
        self._inflight_messages = 0
        self._will = False
        self._will_topic = ""
        self._will_payload = None
        self._will_qos = 0
        self._will_retain = False
        self.on_disconnect = None
        self.on_connect = None
        self.on_publish = None
        self.on_message = None
        self.on_message_filtered = []
        self.on_subscribe = None
        self.on_unsubscribe = None
        self.on_log = None
        self._host = ""
        self._port = 1883
        self._bind_address = ""
        self._in_callback = False
        self._strict_protocol = False
        self._callback_mutex = threading.Lock()
        self._state_mutex = threading.Lock()
        self._out_packet_mutex = threading.Lock()
        self._current_out_packet_mutex = threading.Lock()
        self._msgtime_mutex = threading.Lock()
        self._out_message_mutex = threading.Lock()
        self._in_message_mutex = threading.Lock()
        self._thread = None
        self._thread_terminate = False
        self._ssl = None
        self._tls_certfile = None
        self._tls_keyfile = None
        self._tls_ca_certs = None
        self._tls_cert_reqs = None
        self._tls_ciphers = None
        self._tls_version = tls_version
        self._tls_insecure = False

    def __del__(self):
        pass

    def reinitialise(self, client_id="", clean_session=True, userdata=None):
        if self._ssl:
            self._ssl.close()
            self._ssl = None
        elif self._sock:
            self._sock.close()
            self._sock = None
        if self._sockpairR:
            self._sockpairR.close()
            self._sockpairR = None
        if self._sockpairW:
            self._sockpairW.close()
            self._sockpairW = None

    def tls_set(self, ca_certs, certfile=None, keyfile=Nne, cert_reqs=cert_reqs, tls_version=tls_version, ciphers=None):
        """Configure network encryption and authentication options Enables SSL/TLS support.

        ca_certs: a string path to the Certificate Authority certificate files
        that are to be treated as trusted by this client. If this is the only
        option given then the client will operate in a similar manner to a web
        browser. That is to say it will require the broker to have a
        certificate signed by the Certificate Authorities in ca_certs and will
        communicate using TLS v1, but will not attempt any form of
        authentication. This provides basic network encryption but may not be
        sufficient depending on how the broker is configured.

        certfile and keyfile are strings pointing to the PEM encoded client
        certificate and private keys respectively. If these arguments are not
        None then they will be used as client information for TLS based
        authentication. Support for this feature is broker dependent. Note
        that if either of these files in encrypted and needs a password to
        decrypt it, Python will ask for the password at the command line. It is
        not currently possible to define a callback to provide the password.

        cert_reqs allows the certificate requirements that the client imposes
        on the broker to be changed. By default this is ssl.CERT_REQUIRED,
        which means that the broker must provide a certificate. See the ssl
        pydoc for more information on this parameter.

        tls_version allows the version of the SSL/TLS protocol used to be
        specified. By default TLS v1 is used. Previous versions (all versions
        beginning with SSL) are possible but not recommended due to possible
        security problems.

        ciphers is a string specifying which encryption ciphers are allowable
        for this connection, or None to use the defaults. See the ssl pydoc for
        more information.

        Must be called before connect() or connect_async()
        """
        if HAVE_SSL is False:
            raise ValueError('This platform has no SSL/TLS')
        
        if sys.version < '2.7':
            raise ValueError('Python 2.7 is the minimum supported version for TLS.')

        if ca_certs is None:
            raise ValueError('ca_certs must not be None.')

        try:
            f = open(ca_certs, "r")
        except IOError as err:
            raise IOError(ca_certs+": "+err.strerror)
        else:
            f.close()
        if certfile is not None:
            try:
                f = open(certfile, "r")
            except IOError as err:
                raise IOError(certfile+": "+err.strerror)
            else:
                f.close()
        if keyfile is not None:
            try:
                f = open(keyfile, "r")
            except IOError as err:
                raise IOError(keyfile+": "+err.strerror)
            else:
                f.close()

        self._tls_ca_certs = ca_certs
        self._tls_certfile = certfile
        self._tls_keyfile = keyfile
        self._tls_cert_reqs = cert_reqs
        self._tls_version = tls_version
        self._tls_ciphers = ciphers

    def tls_insecure_set(self, value):
        """Configure verification of the server hostname in the server certificate.

        If value is set to true, it is impossible to guarantee that the host
        you are connecting to is not impersonating your server. This can be
        useful in initial server testing, but makes it possible for a malicious
        third party to impersonate yoru serevr through DNS spoofing, for
        example.

        Do not use this function in a real system. Setting value to true means
        there is no point using encryption.

        Must be called before connect()
        """
        if HAVE_SSL is False:
            raise ValueError('This platform has no SSL/TLS.')

        self._tls_insecure = value
