import subprocess


class GitService:
    def __init__(self) -> None:
        pass

    def GetPreviousCommitHash(self) -> str:
        cmd = "git log --pretty=format:%H -1"
        outp = subprocess.check_output(cmd)
        return outp.decode("utf8")

