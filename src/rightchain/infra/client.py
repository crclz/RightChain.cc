import os
import requests


class RightClient:
    def __init__(self, base_url="https://rightchain.cc") -> None:
        self.base_url = base_url
        self.sess = requests.session()

    def login(self, username, password):
        res = self.sess.post(self.base_url+"/api/access/login", json={
            "username": username,
            "password": password,
        })

        if not res.ok:
            raise Exception(f"登录失败. 响应内容：{res.text}")

    def create_record(self, name, text):
        res = self.sess.post(self.base_url + "/api/records", json={
            "name": name, "text": text
        })

        assert res.ok
        return res.json()['id']

    def get_record(self, id):
        res = self.sess.get(self.base_url + f"/api/records/{id}")
        assert res.ok, res.text
        return res.json()

    def outofbox_create_record(self, hashstr):
        assert len(hashstr) <= 64

        res = self.sess.post(self.base_url+'/api/out-of-box/create-record', json={
            'hash':  hashstr
        })

        assert res.ok
        return res.json()['token']

    def outofbox_get_record(self, token):
        res = self.sess.get(self.base_url+'/api/out-of-box/get-record', params={
            'token': token
        })

        assert res.ok

        return res.json()


def login_using_env_and_get_client():
    try:
        username = os.environ["RIGHT_USERNAME"]
        password = os.environ["RIGHT_PASSWORD"]
    except Exception:
        print(f"请设置环境变量：RIGHT_USERNAME, RIGHT_PASSWORD")
        raise

    client = RightClient()

    client.login(username, password)

    return client
