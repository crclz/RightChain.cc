import pathlib
from unittest.mock import Mock

from rightchain.domain.services.indexing import IndexingService


def test_fileSha256():
    # arrange
    indexingService = IndexingService(Mock(), Mock(), Mock())

    expectedSha256 = "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4"

    file = pathlib.Path(__file__).parent / "sample.txt"

    # act
    result = indexingService.fileSha256(file.__str__())

    # assert
    assert expectedSha256 == result


def test_stringSha256():
    # arrange
    indexingService = IndexingService(Mock(), Mock(), Mock())

    expectedSha256 = "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"

    # act
    result = indexingService.stringSha256("Hello, World!")

    # assert
    assert expectedSha256 == result


def test_generateSalt():
    # arrange
    indexingService = IndexingService(Mock(), Mock(), Mock())

    # act
    result = indexingService.generateSalt(5)

    # assert
    assert isinstance(result, str)
    assert 10 == len(result)

