from typing import List
from pathlib import Path
import pathspec

from rightchain.infra.copyright_store_service import CopyrightStoreService


class FileListerService:
    def __init__(self, copyrightStoreService: CopyrightStoreService) -> None:
        self.copyrightStoreService = copyrightStoreService
        pass

    def list_all_files_in_workdir(self) -> List[str]:
        def formatPath(x: Path) -> str:
            x = x.absolute()
            x = x.relative_to(Path(".").absolute())
            s = str(x).replace("\\", "/")

            return s

        files = self.getFilesRecursively(Path("."), True)

        files = [formatPath(p) for p in files]

        return files

    def getFilesRecursively(self, path: Path, fileOnly: bool) -> List[Path]:
        items: List[Path] = list(path.rglob("*"))

        if fileOnly:
            items = [p for p in items if p.is_file()]

        return items

    def list_included_file_in_workdir(self, exclude_spec_lines: List[str]):
        spec = pathspec.PathSpec.from_lines(
            pathspec.patterns.GitWildMatchPattern, exclude_spec_lines
        )

        all_files = self.list_all_files_in_workdir()
        included: List[str] = []

        for file in all_files:
            if not spec.match_file(file):
                included.append(file)

        return included

    def list_files_with_gitignore_and_rightignore(self) -> List[str]:
        exclude_spec_lines = []
        exclude_spec_lines.append(".git")
        exclude_spec_lines.append(self.copyrightStoreService.right_dir)

        for ignorefile in [".gitignore", ".rightignore"]:
            if Path(ignorefile).exists():
                print(f"using ignore file: {ignorefile}")

                with open(ignorefile, "r", encoding="utf8") as f:
                    lines = f.readlines()

                exclude_spec_lines += lines

        files = self.list_included_file_in_workdir(exclude_spec_lines)
        return files

    def TryReadCopyrightInfo(self) -> "str|None":
        copyright_info = "copyright-info"
        if Path(copyright_info).exists():
            print(f"using copyright info: {copyright_info}")

            with open(copyright_info, "r", encoding="utf8") as f:
                info = f.read()
                return info
        else:
            return None


# files = list_files_with_gitignore_and_rightignore()
# print(len(files))

# for file in files:
#     print(file)
