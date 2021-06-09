import pathlib
import os

right_dir = 'copyrightstore'

wait_dir = f"{right_dir}/waiting"
packaged_dir = f"{right_dir}/packaged"
index_file = f'{right_dir}/index.json'


def create_dir(path):
    fileinfo = pathlib.Path(path)
    if fileinfo.exists():
        assert fileinfo.is_dir()
    else:
        fileinfo.mkdir()


def ensure_all_dir_created():
    create_dir(right_dir)
    create_dir(wait_dir)
    create_dir(packaged_dir)
