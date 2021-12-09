# Creeper

Creeper 基于 RUST Meilisearch 的 轻量级日志分析服务  (ELK 替代品)  


## 可视化

- Meilisearch web ui: http://127.0.0.1:7700/

## 部署:

`docker-compose up -d`

## 配置ENV参数:
- `ListenAddr` 当前服务监听地址                          `default: 0.0.0.0:8745`
- `Token` 配置 Auth Token  (空则 无鉴权)
- `MeilisearchAddr` Meilisearch地址                    
- `MeilisearchToken`  Meilisearch token
- `FlashSec`   插入 Meilisearch 刷新时间 sec            `default: 3`
- `FlashSize`  插入 Meilisearch 刷新大小                `default: 1000`
- `MaxFlashPoolSize`   插入 Meilisearch 最大刷新线程数   `default: 100` 

## Api 文档

- 约定:

参数返回200 正确 其他均 错误

- 鉴权  (当 `Token` 参数配置后 必传)

Set Header  Key: `token`

- 获取所有 index

GET `/api/v1/index`  

- 插入log

POST  `/api/v1/log`

JSON:
``` 
{
    "index": "插入那个index中 没有就自动创建",
    "message": "具体消息"
}
```

- 删除 index

POST `/api/v1/del_index`  

JSON:
``` 
{
    "index": "需要删除的index"
}
```

- 日志瘦身 (删除老大日志)

POST `/api/v1/log_slimming`

JSON:
``` 
{
    "index": "需要瘦身的index",
    "retention_days": int  需要保留最近日志的天数
}
```

- 日志 查询 

POST `/api/v1/search`

JSON:
``` 
{
    "index": "需要查询到index",
    "key_word": "关键字" (可选  默认最新数据 倒排),
    "offset": int (可选),, 
    "limit": int defait: 500 (可选),,
    "start_time": int 起始时间 时间戳  (可选),
    "end_time": int 结束时间 时间戳 (可选),
}
```
