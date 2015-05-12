# -*- coding: utf-8 -*-
import json
import subprocess


if __name__ == '__main__':
    cmd = 'echo {\"ClusterId\":\"j-xxxxxxxxxx\"}'
    cmdl = cmd.split(' ')
    subprocess.call(cmdl)

    output = subprocess.check_output(cmdl)
    data = json.loads(output)
    print data, type(data)
    print data['ClusterId']
