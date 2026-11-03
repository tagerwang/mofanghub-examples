"""最小示例：pip install openai 后直接跑"""
from openai import OpenAI

client = OpenAI(
    api_key="你的魔方Hub令牌",
    base_url="https://mofanghub.cn/v1",
)
resp = client.chat.completions.create(
    model="glm-5.3-flash",
    messages=[{"role": "user", "content": "你好"}],
)
print(resp.choices[0].message.content)
print("用量:", resp.usage)
