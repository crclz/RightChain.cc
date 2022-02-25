from typing import Any, Dict
import requests


class RightClient:
    def __init__(self, base_url="https://rightchain.cc") -> None:
        self.base_url = base_url
        self.sess = requests.session()

    def outofbox_create_record(self, hashstr: str) -> str:
        """
        return the token
        """

        assert len(hashstr) <= 64

        res = self.sess.post(
            self.base_url + "/api/out-of-box/create-record", json={"hash": hashstr}
        )

        assert res.ok
        return res.json()["token"]

    def outofbox_get_record(self, token: str) -> Dict[str, Any]:
        res = self.sess.get(
            self.base_url + "/api/out-of-box/get-record", params={"token": token}
        )

        assert res.ok

        return res.json()

