import json
import string
from itertools import groupby
from typing import List

TYPE_NONE = 0
TYPE_WORD = 1
TYPE_SPACE = 2
TYPE_PUNCT = 3
TYPE_OTHER = 4

class TextWrap:
    def __init__(self, text: str, type: int):
        self.text = text
        self.type = type
    def __repr__(self):
        return f"('{self.text}', {self.type})"
    def __str__(self):
        return self.__repr__()

class _():
    def __init__(self):
        self.dict_filename = "dictionary.json"
        try:
            with open(self.dict_filename, "r") as dictionary:
                self.dictionary = json.load(dictionary)
        except:
            self.dictionary = {}

    def dorepl(self):
        while True:
            text = input("> ")
            if text in ("exit", "quit"):
                exit(0)
            self.process_text(self.split_text(text))

    def dofile(self, filename: str):
        try:
            f = open(filename, "r")
        except Exception as e:
            print(e)
            exit(1)

        self.process_text(self.split_text(f.read()))
        f.close()

    def split_text(self, english_text: str) -> List[TextWrap]:
        processed: List[str] = []
        accumulator: str = ""
        last_type = TYPE_NONE

        def subprocess(id: int):
            nonlocal accumulator, last_type
            if last_type != TYPE_NONE and last_type != id:
                processed.append(TextWrap(accumulator, last_type))
                accumulator = ""
            last_type = id

        for char in english_text + chr(27):
            if char.isalnum():
                subprocess(TYPE_WORD)
                accumulator += char
            elif char.isspace():
                subprocess(TYPE_SPACE)
                accumulator += char
            elif char in string.punctuation:
                subprocess(TYPE_PUNCT)
                accumulator += char
            else:
                subprocess(TYPE_OTHER)
        return processed

    def process_text(self, process_text: List[TextWrap]):
        print(process_text)

if __name__ == '__main__':
    _ = _()
    _.dorepl()
