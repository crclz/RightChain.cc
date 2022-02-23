class File:
    def __init__(self, filename: str, salt: str, hash: str) -> None:
        self.filename = filename
        self.salt = salt
        self.hash = hash
