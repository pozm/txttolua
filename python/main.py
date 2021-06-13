import pathlib
import sys


def main():
    try:
        pathTo = sys.argv[1]
    except IndexError:
        print('I require a path to be put in as an argument.')
        return
    path = pathlib.Path(pathTo)
    if not path.exists():
        print('Sorry, the path you have entered does not exist!')
        return
    print(path.rename(path.stem + '.lua'))


main()  # same thing as js
