import os
from . import dir_config
from .dir_config import ensure_all_dir_created
from .fetch_record import do_fetch
from .create_index import create_index, push_index
from . import client as api_client
import sys
import argparse

if __name__ == "__main__":

    parser = argparse.ArgumentParser(
        prog="rccc",
        description='Command line tool for copyright on blockchain')

    # parser.add_argument('command', choices=[
    #                     'create-index', 'push-index', 'index', 'fetch', 'test-login'], help="action to perform")

    parser.add_argument('command', choices=[
        'create-index', 'push-index', 'index', 'fetch'], help="action to perform")

    args = parser.parse_args()

    command = args.command

    # in case git not remembering empty folders
    ensure_all_dir_created()

    # if command == 'test-login':
    #     api_client.login_using_env_and_get_client()
    #     print("Login test success!")
    #     exit(0)

    if command == 'create-index':
        create_index()
        print("Create index complete.")
        exit(0)

    if command == 'push-index':
        push_index()
        # print("Pushed index to server.")
        exit(0)

    if command == 'index':
        create_index()
        push_index()
        exit(0)

    if command == 'fetch':
        do_fetch()
        print("Fetch from server complete")
        exit(0)

    raise Exception("Not reached")
