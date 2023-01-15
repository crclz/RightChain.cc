# rightchain.cc

## 前言

rightchain 借助**区块链**保护你的**版权**。以下的群体皆可以受益于本应用——

- 你习惯在博客网站撰写自己的**博客**
- 你习惯将博客文章储存在`git`仓库，并习惯使用hugo、hexo、jekyll等框架生成博客页面。
    - 这种情况下，你的**每一个版本**的文章都可以被记录版权
- 你是**科研工作者**，想要在**不公布内容**的情况下，为每一个版本的**论文**建立优先权，并且保护每一天产生的**实验数据**
- 你是时间旅行者，或者你拥有预言能力。区块链的性质将为你的预言发布日期提供证明
- 作为普通人的你在心中暗暗立下誓言。等到目标实现的那一天，你希望能够找到很久以前誓言的证明

## 应用优势

rightchain 是一个公益性质的业余作品，具备以下商业应用没有的优势：

- 免费：完全免费
- 方便：仅仅一条命令，便可为仓库里的文章、代码、数据的**当前版本**(commit)建立进行登记
- 自由：版权证明与本网站低耦合。每12个小时，记录会统一打包到区块链上，当用户导出记录后，即使网站停止服务，用户也无任何影响。
- 数学有效：区块链拥有数学上的绝对证明效力。（当前使用BCH）。

## Installation

- 请认准消息发布页：**本仓库**。  
- 服务网址：https://rightchain.cc
- 遇到问题，请首先确认已经及时更新命令行工具

安装命令行工具：
1. 安装 git
2. 安装 golang：https://go.dev
3. 设置国内go代理（https://goproxy.cn）：
    ```bash
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct
    ```
4. 安装本工具的客户端：
   ```bash
   go install github.com/crclz/rightchain.cc
   ```

## Getting Started

### snapshot 命令
1. 在git仓库根目录下打开命令行。这个git仓库里面是你的文章或者其他知识成果。
2. 通过 `git log --pretty=format:%H -1` 命令来检查git仓库里面至少有一个提交
3. 运行 `rightchain.cc snapshot`. 本命令会做以下几个事情:
   - 在`copyrightstore`目录下生成`snapshot.json`，这个文件包含了被记录的文件。被gitignore的文件不会出现在里面。
   - 基于`snapshot.json`里的hash建立RecipeTree。将树的根节点输出上传到rightchain服务
   - 将RecipeTree、fetch凭据等其他信息储存在`copyrightstore/unpackaged`的新文件里面
4. 立刻进行 git commit ，将刚刚的新文件及时提交，commit前不要修改仓库文件，以免让`snapshot.json`不准确

从业务角度看，snapshot命令基于仓库的所有文件的hash生成RecipeTree，再将RecipeTree的根节点的输出登记到服务器上，从而等效地将整个仓库登记在服务器上。

### fetch 命令

服务器上的记录会定期写入区块链（目前是12h）。所以12h后，用户需要从服务器上获取记录，然后生成完整的RecipeTree。

1. 等待服务器将记录定期写入区块链
2. 运行`rightchain.cc fetch`。本命令的作用：
   - 遍历`copyrightstore/unpackaged`的文件，向服务器查询是否被打包到区块链
   - 对于已经打包到区块链的文件，将它的RecipeTree和服务器返回的信息结合，构造出新的RecipeTree。这棵树的根节点的输出值可以在transactionId这个交易里面找到
   - 将新RecipeTree和其他信息写入`copyrightstore/packaged`

被写入`copyrightstore/packaged`的记录，就彻底和rightchain服务无关了，可以脱离于服务器发挥作用。


### proof 命令

当我们需要证明某（几）个文件的版权，我们需要这两个东西
- 关注的文件列表
- RecipeTree

其中，RecipeTree需要满足：
- 根节点的输出被登记到了区块链上
- 对于“关注的文件列表”中的每一个文件，它的hash都能在树中找到

步骤:
1. 从仓库的提交记录，或者最新的`copyrightstore/packaged`里面进行寻找，找出想要证明的仓库版本。由于同一文件在历史上存在多次修改和登记，所以你需要找到你想要的版本（commit）。这里记commit为$version
2. 在checkout到\$version前，拷贝`copyrightstore/packaged/$version.json`到其他地方，例如桌面。
3. git checkout $version
4. 将刚刚的\$version.json拷贝到`copyrightstore/packaged/$version.json`
5. 运行`rightchain.cc proof --filenames some-file.txt`
   - `--filenames`参数可被传递多次，例如`rightchain.cc proof --filenames a.txt --filenames b.txt`
   - `--trycrlf`参数是可选参数，是为了解决编辑器、git配置的CRLF设置不一致导致hash无法对上\$version.json里面的hash。开启此参数，程序核对hash失败后，就会自动尝试LF版本的和CRLF版本的文件。
6. 程序会生成`rightchain.proof.$timestamp`文件夹，建议不要提交到git。这个文件夹里面包含了：刚刚输入的文件（文件名是hash）、精简后的RecipeTree和相关transaction信息。


## 原理
[fundamental](fundamental.md)


