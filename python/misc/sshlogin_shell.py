# -*- coding: utf-8 -*-
import subprocess
import sys

if __name__ == '__main__':
    hostname = sys.argv[1]
    username = sys.argv[2]
    keyfile_path = sys.argv[3]

    cmd = 'ssh {}@{} -i {}'.format(username, hostname, keyfile_path)
    subprocess.call(cmd, shell=True)
