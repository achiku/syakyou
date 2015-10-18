# -*- coding: utf-8 -*-
# http://makina-corpus.com/blog/metier/2015/python-http-server-with-the-new-async-await-syntax
import asyncio

async def slow_operation(n):
    await asyncio.sleep(1)
    print("Slow operation {} complete".format(n))

async def main():
    await asyncio.wait([
        slow_operation(1),
        slow_operation(2),
        slow_operation(3),
        slow_operation(4),
    ])

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
