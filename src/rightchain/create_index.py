from . import file_lister
import pathlib
from . import dir_config as dc
from .client import RightClient, login_using_env_and_get_client
import json
import os
import hashlib
import subprocess
from .utils import *


def get_previous_commit_hash():
    cmd = 'git log --pretty=format:%H -1'
    outp = subprocess.check_output(cmd)
    return outp.decode('utf8')


def unix_style(s: str):
    """ change \\ to / """
    return s.replace("\\", '/')


def create_empty_file(path):
    pathlib.Path(path).touch()


def get_sha256_of_files(files: "list[str]"):
    """
    returns list of {name: ..., hash: ...}
    """
    data = []

    for file in files:
        hash = file_sha256(file)
        data.append({
            "name": file,
            "hash": hash,
        })

    return data


def try_read_copyright_info():
    copyright_info = "copyright-info"
    if os.path.exists(copyright_info):
        print(f'using copyright info: {copyright_info}')

        with open(copyright_info, 'r', encoding='utf8') as f:
            info = f.read()
            return info
    else:
        return None


def create_index():
    files = file_lister.list_files_with_gitignore_and_rightignore()
    data = get_sha256_of_files(files)

    data = [p['hash'] for p in data] # only hash for filename privacy

    # write to index file

    prev_commit = get_previous_commit_hash()  # return "" when no prev commit

    index_content = {
        'previousCommit': prev_commit,
        'copyrightInfo': try_read_copyright_info(),
        "data": data
    }

    # save index content to index file
    write_json_to_file(index_content, dc.index_file)

    print(f"Index created: {dc.index_file}")


def push_index():
    # post and get record id

    client = RightClient()

    hash_of_index_file = file_sha256(dc.index_file)

    prev_commit = read_json_file(dc.index_file)['previousCommit']

    token = client.outofbox_create_record(hash_of_index_file)

    wait_record_file = os.path.join(dc.wait_dir, prev_commit)

    write_json_to_file({"token": token}, wait_record_file)

    print(f"Pushed and create waiting file: {wait_record_file}")
