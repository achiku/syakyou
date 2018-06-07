import os

if __name__ == '__main__':
    s = os.getenv('SHELL', 'empty_shell')
    print(s)
    s = os.getenv('HOGE', 'empty_hoge')
    print(s)
