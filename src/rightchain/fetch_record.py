from .client import RightClient, login_using_env_and_get_client
import json
import os
from . import dir_config as dc

from .utils import *


def get_waiting_items():
    """
    returns list of (commit_hash, record_id)
    """
    items = []
    for name in os.listdir(dc.wait_dir):
        commit_hash = name

        json_info = read_json_file(os.path.join(dc.wait_dir, name))
        token = json_info['token']

        items.append((commit_hash, token))

    return items


def is_packaged(record: dict):
    return record['transactionId'] is not None


def do_fetch():

    client = RightClient()

    for commit_hash, token in get_waiting_items():

        record = client.outofbox_get_record(token)

        if is_packaged(record):
            write_json_to_file(record, path=os.path.join(
                dc.packaged_dir, commit_hash))
            os.remove(os.path.join(dc.wait_dir, commit_hash))
            print("packaged:", commit_hash)
        else:
            print("still not packaged:", commit_hash)
            # print(record)
