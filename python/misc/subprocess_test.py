# -*- coding: utf-8 -*-
import subprocess


if __name__ == '__main__':
    # Run the command described by args. Wait for command to complete, then return the returncode attribute.
    ret1 = subprocess.call(['echo', '$HOME'])
    print(ret1)
    ret2 = subprocess.call('echo $HOME', shell=True)
    print(ret2)
    ret3 = subprocess.call('exit 1', shell=True)
    print(ret3)

    ret4 = subprocess.check_call(['ls', '-l'])
    print(ret4)

    return_output = subprocess.check_output(['ls', '-l'])
    print(type(return_output))
    print(return_output)
