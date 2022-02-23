from typing import List, Optional
from rightchain.domain.models.file import File


class RepositoryStatus:
    def __init__(self, files: List[File], copyrightInfo: Optional[str], commit:str) -> None:
        self.files = list(files)
        self.copyrightInfo = copyrightInfo
        self.commit = commit
