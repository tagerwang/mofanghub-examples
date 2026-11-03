// 最小示例：npm i openai 后 node examples/node.mjs
import OpenAI from "openai";

const client = new OpenAI({
  apiKey: process.env.MOFANGHUB_API_KEY ?? "你的魔方Hub令牌",
  baseURL: "https://mofanghub.cn/v1",
});
const resp = await client.chat.completions.create({
  model: "glm-5.3-flash",
  messages: [{ role: "user", content: "你好" }],
});
console.log(resp.choices[0].message.content);
