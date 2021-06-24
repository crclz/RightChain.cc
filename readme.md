# RightChain

**域名（rightchain.cc）暂时不可用，浏览器端暂时不开放，现在暂时在等待备案审核。命令行端可以正常提供服务。**

请认准消息发布页：**本仓库**。

## 前言


本人目前大三。本服务是一个公益性质的课余作品。

<!-- 服务所在的网址，是 https://RightChain.cc -->

<!-- Right = copyright = 版权。 Chain = BlockChain = 区块链。 -->

欢迎在ISSUE区提出任何问题，包括但不限于：
- bug反馈、功能建议
- 文档建议、勘误
- 使用感受

## 简介

rightchain 借助区块链保护你的版权。以下的群体皆可以受益于本应用——

- 你习惯在博客网站撰写自己的**博客**
- 你习惯将博客文章储存在`git`仓库，并习惯使用hugo、hexo、jekyll等框架生成博客页面。
    - 这种情况下，你的**每一个版本**的文章都可以被记录版权
- 你是**科研工作者**，想要在**不公布内容**的情况下，为每一个版本的**论文**建立优先权，并且保护每一天产生的**实验数据**
- 你是时间旅行者，或者你拥有预言能力。区块链的性质将为你的预言发布日期提供证明。
- 你是普通人，在心中暗暗立下誓言。等到目标实现的那一天，你希望能够找到很久以前立下誓言的证明。

## 应用优势

现在这个服务的状态，是一个公益性质的课余作品。尽管这是一个半成品，本服务还是具备以下商业应用没有的优势：

- 免费。当然目前也没有付费项目，即使付费也是“将12小时缩短为1小时”这种服务。
- 方便。仅仅一条命令，便可为仓库里的文章、代码、数据的当前版本建立保护，时间成本低。
- 自由。不捆绑用户，版权证明与本网站低耦合。每12个小时，记录会统一打包到区块链上。当你导出记录后，即使网站停止服务，也无法对你造成任何影响。
- 拥有数学上的绝对证明效力。区块链不可篡改、很难消亡（当前是BCH）。

## 通过浏览器访问网站【暂不开放】

**注意，网站正在备案等待过程中，此途径暂时不开放。**

用户可以选择通过浏览器来为自己作品登记。也可以选择通过命令行的自动化工具为多个作品登记。首先我们来介绍浏览器的方式。

1. 登记。首先，用户需要在本地计算目标文章文件的sha256。详细的方法可以搜索引擎。7zip或者powershell都提供计算sha256的方法。获取sha256后，用户需要在网站的“创建记录”页面按照提示进行填写信息。
2. 等待打包。网站隔一段时间会统一打包，将必要的信息写到区块链上。首页会显示“下次打包时间”。上链后，你的作品就算真正登记了。现在的等待的平均时间是6小时（`12h/2`）
3. 导出信息。网站将信息打包完成、上链后，用户可以在“记录详情页面”导出相关信息保存至本地。

以下是用户需要保存的两样东西。保存了这些东西，用户就和本服务无关了，只和区块链有关，**不怕网站挂掉或者停止服务**。在需要证明版权的时候，这两样就行了：

- 从网站导出的内容
- 原始文件
    



## 通过命令行工具保护仓库里的文章

命令行工具将很好地契合git仓库，为你的**每一个版本**的**每一篇**文章、数据提供保护，同时自动化登记的流程，省时省力。版权信息也能够以文件的形式伴随仓库，不会为用户增加任何负担，用户无需额外寻找地方进行储存。

[用户手册：通过命令行工具保护仓库里的文章](./readme-tool.md)

## 延伸阅读

数学上的证明是否成立？`recipe`和其他的字段是什么意思？

[implementation-details](implementation-details.md)

[validate](./validate.md)
