import hashlib
import json
from typing import Any, Dict, List
from rightchain.domain.models.file import File
from rightchain.domain.models.repo_status import RepositoryStatus
from rightchain.infra.copyright_store_service import CopyrightStoreService
from rightchain.infra.file_lister import FileListerService
import secrets

from rightchain.infra.git_service import GitService


class IndexingService:
    def __init__(
        self,
        fileListerService: FileListerService,
        copyrightStoreService: CopyrightStoreService,
        gitService: GitService,
    ) -> None:
        self.fileListerService = fileListerService
        self.copyrightStoreService = copyrightStoreService
        self.gitService = gitService

    def fileSha256(self, filename: str) -> str:
        with open(filename, "rb") as f:
            h = hashlib.sha256()
            h.update(f.read())
            return h.hexdigest().lower()

    def stringSha256(self, s: str) -> str:
        h = hashlib.sha256()
        h.update(s.encode("utf8"))
        return h.hexdigest().lower()

    def generateSalt(self, nbytes: int) -> str:
        return secrets.token_hex(nbytes).lower()

    def getFileEntity(self, filename: str) -> File:
        salt: str = self.generateSalt(30)
        hash = self.stringSha256(self.fileSha256(filename) + salt)
        return File(filename, salt, hash)

    def GetRepositoryStatus(self) -> RepositoryStatus:
        filenames = self.fileListerService.list_files_with_gitignore_and_rightignore()

        files: List[File] = [self.getFileEntity(p) for p in filenames]
        copyrightInfo = self.fileListerService.TryReadCopyrightInfo()
        commit = self.gitService.GetPreviousCommitHash()

        return RepositoryStatus(files, copyrightInfo, commit)

    def GenerateSaltInfo(self, status: RepositoryStatus):
        return {
            "warning": "This file should keep secret. When a file is needed, only the salt of it should be exposed",
            "hint": "for a file, sha256(sha256(content_of_file)+salt) = sha256_in_index_json",
            "data": [{"filename": p.filename, "salt": p.salt} for p in status.files],
        }

    def GenerateIndexInfo(self, status: RepositoryStatus):
        return {
            "previousCommit": status.commit,
            "copyrightInfo": status.copyrightInfo,
            "data": [p.hash for p in status.files],
        }

    def GetSha256OfIndexFile(self) -> str:
        return self.fileSha256(self.copyrightStoreService.index_file)

