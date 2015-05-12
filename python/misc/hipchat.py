# -*- coding: utf-8 -*-
import json
import requests


if __name__ == '__main__':
    token = 'your_room_token'
    room_id = 'your_room_id'
    notification_url = 'https://api.hipchat.com/v2/room/{0}/notification'.format(room_id)
    msg = '''
    <pre>hello from Python script!</pre>
    <a href="https://google.com">google link</a>
    '''
    payload = {
        'message': msg,
        'color': 'yellow',
        'notify': 'true',
        'message_format': 'html',
    }
    headers = {
        'Authorization': 'Bearer {}'.format(token),
        'Accept': 'application/json',
        'Content-type': 'application/json',
    }

    res = requests.post(
        notification_url,
        json.dumps(payload),
        headers=headers
    )
    print res
