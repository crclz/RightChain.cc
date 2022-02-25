from unittest.mock import Mock

from pathlib import Path
from rightchain.infra.file_lister import FileListerService
import os


def test_list_all_files_in_workdir_happy():
    # arrange
    lister = FileListerService(Mock())
    basePath = Path(__file__).parent / "sample"

    # act
    savedWorkDir = os.getcwd()
    os.chdir(basePath)
    files = lister.list_all_files_in_workdir()
    os.chdir(savedWorkDir)

    # assert

    assert set(files) == {"1.hello", "2/3.hello"}
