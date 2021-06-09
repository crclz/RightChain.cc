from . import dir_config
import pathspec
import os


def list_all_files_in_workdir():
    all_files = []

    for dirpath, dirs, files in os.walk('.'):
        for name in files:
            filename = os.path.join(dirpath, name)
            filename = filename.replace("\\", "/")
            all_files.append(filename)

    return all_files


def list_included_file_in_workdir(exclude_spec_lines: "list[str]"):
    spec = pathspec.PathSpec.from_lines(
        pathspec.patterns.GitWildMatchPattern, exclude_spec_lines)

    all_files = list_all_files_in_workdir()
    included = []

    for file in all_files:
        if not spec.match_file(file):
            included.append(file)

    return included


def list_files_with_gitignore_and_rightignore():
    exclude_spec_lines = []
    exclude_spec_lines.append(".git")
    exclude_spec_lines.append(dir_config.right_dir)

    for ignorefile in ['.gitignore', '.rightignore']:
        if os.path.exists(ignorefile):
            print(f"using ignore file: {ignorefile}")

            with open(ignorefile, 'r', encoding='utf8') as f:
                lines = f.readlines()

            exclude_spec_lines += lines

    files = list_included_file_in_workdir(exclude_spec_lines)
    return files


# files = list_files_with_gitignore_and_rightignore()
# print(len(files))

# for file in files:
#     print(file)
