#!/usr/bin/env bash
# 最小调用示例：改两行即可跑（key + base_url 已指向魔方Hub）
curl https://mofanghub.cn/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $MOFANGHUB_API_KEY" \
  -d '{
    "model": "glm-5.3-flash",
    "messages": [{"role": "user", "content": "你好"}]
  }'
