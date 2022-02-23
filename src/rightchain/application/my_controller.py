import json
import os
from typing import Any
from rightchain.domain.services.indexing import IndexingService
from rightchain.infra.client import RightClient
from rightchain.infra.copyright_store_service import CopyrightStoreService


class MyController:
    def __init__(
        self,
        indexingService: IndexingService,
        copyrightStoreService: CopyrightStoreService,
        rightClient: RightClient,
    ) -> None:
        self.indexingService = indexingService
        self.copyrightStoreService = copyrightStoreService
        self.rightClient = rightClient

    def CreateIndex(self) -> None:
        repoStatus = self.indexingService.GetRepositoryStatus()
        saltInfo = self.indexingService.GenerateSaltInfo(repoStatus)
        indexInfo = self.indexingService.GenerateIndexInfo(repoStatus)

        self.copyrightStoreService.WriteSaltFile(saltInfo)
        self.copyrightStoreService.WriteIndexFile(indexInfo)

        print(f"CreateIndex ok")

    def PushIndex(self) -> None:
        indexFileHash = self.indexingService.GetSha256OfIndexFile()

        commit = self.copyrightStoreService.ReadIndexFile()["previousCommit"]
        assert isinstance(commit, str)

        token = self.rightClient.outofbox_create_record(indexFileHash)

        # responsibility transfer
        self.copyrightStoreService.WriteWaitFile(commit, {"token": token})

        print(f"Pushed and create waiting file: {commit}")

    def Fetch(self) -> None:
        def is_packaged(record: dict):
            return record["transactionId"] is not None

        for commit, token in self.copyrightStoreService.GetWaitingItems():

            record = self.rightClient.outofbox_get_record(token)

            if is_packaged(record):
                self.copyrightStoreService.WritePackaged(commit, record)
                self.copyrightStoreService.RemoveWaitFile(commit)
                print("packaged:", commit)
            else:
                print("still not packaged:", commit)
