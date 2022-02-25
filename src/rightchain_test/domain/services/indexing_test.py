import pathlib
from unittest.mock import Mock

from rightchain.domain.services.indexing import IndexingService


def test_fileSha256():
    # arrange
    indexingService = IndexingService(Mock(), Mock(), Mock())

    expectedSha256 = "03AC674216F3E15C761EE1A5E255F067953623C8B388B4459E13F978D7C846F4"
    expectedSha256 = expectedSha256.lower()
    
    file = pathlib.Path(__file__).parent / 'sample.txt'

    # act
    result = indexingService.fileSha256(file.__str__())

    # assert
    assert expectedSha256 == result
