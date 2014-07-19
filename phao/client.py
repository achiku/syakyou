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
