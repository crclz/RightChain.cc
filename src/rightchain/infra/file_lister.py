from typing import List
import pathspec
import os

from rightchain.infra.copyright_store_service import CopyrightStoreService


class FileListerService:
    def __init__(self, copyrightStoreService: CopyrightStoreService) -> None:
        self.copyrightStoreService = copyrightStoreService
        pass

    def list_all_files_in_workdir(self) -> List[str]:
        all_files: List[str] = []

        for dirpath, dirs, files in os.walk("."):
            for name in files:
                filename = os.path.join(dirpath, name)
                filename = filename.replace("\\", "/")
                all_files.append(filename)

        return all_files

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

    def list_files_with_gitignore_and_rightignore(self):
        exclude_spec_lines = []
        exclude_spec_lines.append(".git")
        exclude_spec_lines.append(self.copyrightStoreService.right_dir)

        for ignorefile in [".gitignore", ".rightignore"]:
            if os.path.exists(ignorefile):
                print(f"using ignore file: {ignorefile}")

                with open(ignorefile, "r", encoding="utf8") as f:
                    lines = f.readlines()

                exclude_spec_lines += lines

        files = self.list_included_file_in_workdir(exclude_spec_lines)
        return files

    def TryReadCopyrightInfo(self) -> "str|None":
        copyright_info = "copyright-info"
        if os.path.exists(copyright_info):
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
