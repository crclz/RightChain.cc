# 如何验证版权

## 太长不看版

通过不断使用 `sha256(s)` 和 `字符串连接` ，我们最终会获得一个哈希值。这个哈希值会被登记在区块链上。

同时，不断使用 `sha256(s)` 和 `字符串连接` 的过程，也是建立一棵树的过程。如果树根在某时某刻被登记在区块链上，那么从数学上来说，树叶也被登记了。

本文介绍了如何使用用户的文章文件、导出的相关信息来构建这棵树的详细流程。

---


```python
import hashlib
def sha256(s):  
    h = hashlib.sha256()   
    h.update(s.encode('utf8'))   
    b = h.hexdigest().lower() 
    return b
```


这是一条示例记录：

```
text = "I:ecde115cdca96a1a181db6b9e88c1e1f6e208ca77916627cc8a9fa02b39d0692"
hash = "8822981f04030acac0c4dc635ce2ee0d0b1d7fc3dbde5197f1debd7b8bbf75dc"
txid = "950be274423b3138a4abfa111d4c847252d71ed7d51ce7168fbef1c6167547b1"

recipe = [ [ 1, "e351ac1c9420b30b37c28be10af53ec6bbb124b551681e0038f0de3cbaf8f046" ], [ 0, "ea5b7a39bb57ed609c139780acd2c102d9f09e16644278e984f535d5de9b2683" ], [ 0, "0418e2cea8ef594e6bc05ab3c3fe1fae0db09a2981d32b291992ddc2191075fe" ] ] 

root_hash = "3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0 " 
```


```python
text = "I:ecde115cdca96a1a181db6b9e88c1e1f6e208ca77916627cc8a9fa02b39d0692"
hash = "8822981f04030acac0c4dc635ce2ee0d0b1d7fc3dbde5197f1debd7b8bbf75dc"
txid = "950be274423b3138a4abfa111d4c847252d71ed7d51ce7168fbef1c6167547b1"

recipe = [ [ 1, "e351ac1c9420b30b37c28be10af53ec6bbb124b551681e0038f0de3cbaf8f046" ], 
          [ 0, "ea5b7a39bb57ed609c139780acd2c102d9f09e16644278e984f535d5de9b2683" ],
          [ 0, "0418e2cea8ef594e6bc05ab3c3fe1fae0db09a2981d32b291992ddc2191075fe" ] ] 

root_hash = "3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0" 
```

首先，这里面的信息是有重复的信息。重复的信息是`hash`和`root_hash`。虽然信息有重复，但是依然推荐用户保存这些重复的信息。

`hash`和`root_hash`均可由其他的数据计算得出：

`hash = sha256(text)`

`root_hash = build_root_hash(hash, recipe)`


```python
# 验证 hash = sha256(text)

print(hash)
print(sha256(text))

print(hash == sha256(text))
```

    8822981f04030acac0c4dc635ce2ee0d0b1d7fc3dbde5197f1debd7b8bbf75dc
    8822981f04030acac0c4dc635ce2ee0d0b1d7fc3dbde5197f1debd7b8bbf75dc
    True
    

这样，就可以理解`hash = sha256(text)`.

`build_root_hash`过程稍微有点复杂。在介绍它之前，我们先对证明的链条进行一个总览

## 证明链条

首先，小明写了一篇文章`article`。`article`保存在本地，

然后，小明计算`article`文件的sha256，得到`fd2e1ad2b...`

然后，小明登录网站，将`text`填写为`fd2e1ad2b...`。如果小明的`article`里面没有作者信息，小明可能还会在`text`的头部附加上"本文章由小明原创！"字样，得到` 本文章由小明原创！fd2e1ad2b...`

通过`sha256(text)`，又可以得出 `hash`

通过`build_root_hash(hash, recipe)`，可以得到`root_hash`。`build_root_hash`是一个只含有连接和sha256的过程。

`root_hash`被记录在区块链上（目前是BCH）。

## 理清思路

通过计算小明本地保存的`article`的sha256值，可以建立`article`与`text`的联系。然后，通过计算`sha256(text)`，可以与`hash`建立联系。通过`build_root_hash(hash, recipe)`，可以与`root_hash`建立联系。`root_hash`又与区块链建立联系。


## 理论基础
如果怀疑这个算法是否能够保证版权的正确性，以下理论基础可以帮助你证明这个算法：
1. 摘要的版权 能推导出 原文的版权。（因为sha256碰撞很难，并且尚未找到任何一例）
2. 原文的版权 能推导出 部分的版权。（你写了一篇文章，如果拿一个段落出来，你同样拥有那个段落的版权）

*注：以上的“版权”，更准确来说，是“某人在某时发表了某些内容”*

## 验证其他部分

我们还剩下2个部分尚待验证：
1. 通过`build_root_hash(hash, recipe)`，建立与`root_hash`的关联
2. `root_hash`在区块链上的验证

### `root_hash`与区块链
首先来看`root_hash`在区块链上的验证。`root_hash`在区块链上的交易id就是`txid`。

`txid`对应的交易中如果包含`root_hash`，那么就可以建立关联。

通过区块链浏览器，我们可以查询某个交易的信息。


```python
print("root_hash:", root_hash)
print("txid:     ", txid)
print()

print("https://blockchair.com/bitcoin-cash/transaction/"+txid)
```

    root_hash: 3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0
    txid:      950be274423b3138a4abfa111d4c847252d71ed7d51ce7168fbef1c6167547b1
    
    https://blockchair.com/bitcoin-cash/transaction/950be274423b3138a4abfa111d4c847252d71ed7d51ce7168fbef1c6167547b1
    

通过上面的链接，我们发现，`OP_RETURN`的值为:\
`j@3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0`
而我们的`root_hash`为`3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0`。移除`OP_RETURN`前面的2个字符，就可以得到我们的`root_hash`.

```
j@3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0
  3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0
```

当然，通过api，我们也可以验证：


```python
import requests

res = requests.get('https://api.blockchair.com/bitcoin-cash/dashboards/transaction/'+txid)
```


```python
tx = list(res.json()['data'].values())[0]
script_hex = tx['outputs'][0]["script_hex"]
info = bytes.fromhex(script_hex).decode("utf8")

print(info)
print(root_hash)

print(info[2:] == root_hash)
```

    j@3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0
    3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0
    True
    

## build_root_hash

`root_hash = build_root_hash(hash, recipe)`

算法如下：


```python
# begin build_root_hash

x = hash

for r, hash_str in recipe:
    x = sha256(hash_str+x) if r else sha256(x+hash_str)

# end

print("recipe:")
for z in recipe:
    print(z)
print()
print(x)
print(root_hash)
    
print(x == root_hash)
```

    recipe:
    [1, 'e351ac1c9420b30b37c28be10af53ec6bbb124b551681e0038f0de3cbaf8f046']
    [0, 'ea5b7a39bb57ed609c139780acd2c102d9f09e16644278e984f535d5de9b2683']
    [0, '0418e2cea8ef594e6bc05ab3c3fe1fae0db09a2981d32b291992ddc2191075fe']
    
    3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0
    3904da8c4887e7d9cdc2bd6223a270b1332ce41f087662b325f995009fe803e0
    True
    

从上面的三行代码来看，`build_root_hash`的过程只包含字符串连接 和 sha256 的操作。

通过上面的2条数学基础，我们可以简单证明算法的有效性。

---

至此，一个例子的讲解完成了。

