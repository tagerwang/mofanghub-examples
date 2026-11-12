# MoFangHub Examples

> 一个 key，调用全部主流国产模型。人民币按量计费，倍率全量公开。

MoFangHub（[mofanghub.cn](https://mofanghub.cn)）是国产大模型 API 聚合分发平台。本仓库提供最小可跑接入示例与上游官方价格快照数据。

## 30 秒接入（OpenAI 兼容协议）

```bash
pip install openai
```

```python
from openai import OpenAI

client = OpenAI(
    api_key="你的魔方Hub令牌",          # mofanghub.cn 注册后在控制台创建
    base_url="https://mofanghub.cn/v1"  # 只需改这两行
)

resp = client.chat.completions.create(
    model="glm-5.3-flash",
    messages=[{"role": "user", "content": "你好"}]
)
print(resp.choices[0].message.content)
```

## 目录

```
examples/   四语言最小示例（curl / Python / Node / Go）
data/       上游官方价格快照（JSON，含核验日期）
```

## 价格快照说明

`data/` 内所有数字为**上游官方渠道公开标价的快照**，核验日期见文件内字段。
快照仅作选型参考，调用前请以各官方页面当日展示为准（价格可能调整）。

## 相关链接

- 官网：https://mofanghub.cn
- 技术博客：https://blog.mofanghub.cn（模型选型实测与计费拆解）
  - [混元 Hy4 接入决策：定价 6/18 元与 5 折折扣下的场景账](https://blog.mofanghub.cn/hunyuan-hy4-adoption-pricing-2026-11/)
  - [GLM Coding Plan vs API 按量：一张表算清临界用量](https://blog.mofanghub.cn/glm-coding-plan-vs-api-usage-break-even/)
  - [DeepSeek 联网搜索 API 计费规则：5 元/千次](https://blog.mofanghub.cn/deepseek-search-api-pricing-2026-11/)
- 更新日志：https://mofanghub.cn/about

## License

MIT
