import hashlib
import json


def read_json_file(filename):
    return json.load(open(filename, 'r', encoding='utf8'))


def write_json_to_file(obj, path):
    with open(path, 'w', encoding='utf8') as f:
        json.dump(obj, f, ensure_ascii=False, indent=4)


def file_sha256(filename):
    with open(filename, 'rb') as f:
        h = hashlib.sha256()
        h.update(f.read())
        return h.hexdigest().lower()
