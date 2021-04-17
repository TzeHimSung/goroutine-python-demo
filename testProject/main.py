# sleep 15 seconds and print

import sys
import time

f = open('output.txt', 'a')
sys.stdout = f
sys.stderr = f


def main():
    localtime = time.asctime(time.localtime(time.time()))
    print('current time is', localtime)
    time.sleep(15)
    print('program finished')
    localtime = time.asctime(time.localtime(time.time()))
    print('current time is', localtime)


if __name__ == '__main__':
    main()
