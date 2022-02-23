import argparse

from rightchain.application.my_controller import MyController
from rightchain.domain.services.indexing import IndexingService
from rightchain.infra.client import RightClient
from rightchain.infra.copyright_store_service import CopyrightStoreService
from rightchain.infra.file_lister import FileListerService
from rightchain.infra.git_service import GitService

if __name__ == "__main__":
    gitService = GitService()
    rightClient = RightClient()
    copyrightStoreService = CopyrightStoreService()
    fileListerService = FileListerService(copyrightStoreService)
    indexingService = IndexingService(
        fileListerService, copyrightStoreService, gitService
    )
    controller = MyController(indexingService, copyrightStoreService, rightClient)

    parser = argparse.ArgumentParser(
        prog="rccc", description="Commandline tool for copyright on blockchain"
    )

    # parser.add_argument('command', choices=[
    #                     'create-index', 'push-index', 'index', 'fetch', 'test-login'], help="action to perform")

    parser.add_argument(
        "command",
        choices=["create-index", "push-index", "index", "fetch"],
        help="action to perform",
    )

    args = parser.parse_args()

    command = args.command

    if command == "create-index":
        controller.CreateIndex()
        print("Create index complete.")
    elif command == "push-index":
        controller.PushIndex()
        print("Pushed index to server.")
    elif command == "index":
        controller.CreateIndex()
        controller.PushIndex()
    elif command == "fetch":
        controller.Fetch()
        print("Fetch from server complete")

    raise Exception("Not reached")
