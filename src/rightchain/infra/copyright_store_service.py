import json
import os
import pathlib
from typing import Any, Dict


class CopyrightStoreService:
    def __init__(self) -> None:
        self.right_dir = "copyrightstore"

        self.wait_dir = f"{self.right_dir}/waiting"
        self.packaged_dir = f"{self.right_dir}/packaged"
        self.index_file = f"{self.right_dir}/index.json"
        self.salt_file = f"{self.right_dir}/salt.json"

    def create_dir(self, path: str):
        fileinfo = pathlib.Path(path)
        if fileinfo.exists():
            assert fileinfo.is_dir()
        else:
            fileinfo.mkdir()

    def writeJsonToFile(self, filename: str, obj: Any) -> None:
        content = json.dumps(obj, ensure_ascii=True, indent=4)
        with open(filename, "w", encoding="utf8") as f:
            f.write(content)

    def readJsonFromFile(self, filename: str) -> Any:
        with open(filename, "r", encoding="utf8") as f:
            return json.load(f)

    def ensure_all_dir_created(self):
        self.create_dir(self.right_dir)
        self.create_dir(self.wait_dir)
        self.create_dir(self.packaged_dir)

    def WriteIndexFile(self, obj: Dict) -> None:
        self.writeJsonToFile(self.index_file, obj)

    def WriteSaltFile(self, obj: Dict) -> None:
        content = json.dumps(obj, ensure_ascii=True, indent=4)
        assert isinstance(content, str)
        with open(self.salt_file, "r", encoding="utf8") as f:
            f.write(content)

    def ReadIndexFile(self) -> Dict[str, Any]:
        return self.readJsonFromFile(self.index_file)

    def WriteWaitFile(self, commit: str, obj: Dict) -> None:
        filename = os.path.join(self.wait_dir, commit)
        self.writeJsonToFile(filename, obj)

