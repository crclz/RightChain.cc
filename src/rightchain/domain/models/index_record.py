from typing import Any


class IndexRecord:
    def __init__(self, commit: str, token: "str|None", recordInfo: Any) -> None:
        self.commit = commit
        self.token = token
        self.recordInfo: Any = recordInfo

    @property
    def IsWaiting(self) -> bool:
        return self.token is not None

    def UpdateInfo(self, recordInfo) -> None:
        assert self.IsWaiting
        assert recordInfo is not None
        assert isinstance(recordInfo, dict)

        self.token = None
        self.recordInfo = recordInfo

