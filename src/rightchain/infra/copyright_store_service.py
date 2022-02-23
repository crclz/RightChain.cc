import json
import os
import pathlib
from typing import Any, Dict, List, Tuple


class CopyrightStoreService:
    def __init__(self) -> None:
        self.right_dir = "copyrightstore"

        self.wait_dir = f"{self.right_dir}/waiting"
        self.packaged_dir = f"{self.right_dir}/packaged"
        self.index_file = f"{self.right_dir}/index.json"
        self.salt_file = f"{self.right_dir}/salt.json"

        self.ensure_all_dir_created()

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
        self.writeJsonToFile(self.salt_file, obj)

    def ReadIndexFile(self) -> Dict[str, Any]:
        return self.readJsonFromFile(self.index_file)

    def WriteWaitFile(self, commit: str, obj: Dict) -> None:
        filename = os.path.join(self.wait_dir, commit)
        self.writeJsonToFile(filename, obj)

    def GetWaitingItems(self) -> List[Tuple[str, str]]:
        """
        returns list of (commit_hash, token)
        """
        items: List[Tuple[str, str]] = []
        for name in os.listdir(self.wait_dir):
            commit_hash = name
            filename = os.path.join(self.wait_dir, name)
            json_info = self.readJsonFromFile(filename)

            token = json_info["token"]

            items.append((commit_hash, token))

        return items

    def WritePackaged(self, commit: str, obj: Any) -> None:
        self.writeJsonToFile(os.path.join(self.packaged_dir, commit), obj)

    def RemoveWaitFile(self, commit: str) -> None:
        filename = os.path.join(self.wait_dir, commit)
        os.remove(filename)
