## Service Feature

## Follow Service

```
service tiktokFollowService {
  rpc Ping(PingReq) returns (BaseRsp){};
  rpc RelationAction(RelationActionReq) returns (BaseRsp){};
  rpc RelationQuery(RelationQueryReq) returns (RelationQueryRsp){};
  rpc FollowList(UserListReq) returns (UserListRsp){};
  rpc FollowerList(UserListReq) returns (UserListRsp){};
  rpc FriendList(UserListReq) returns (UserListRsp){};
}
```

**Follow Service 提供了六个接口，**

* `Ping` 测试服务连接状态
* `RelationAction` 为用户间添加关系数据，操作类型由 `RelationActionReq` 中的 `action_type` 字段决定，若为 1，则关注；若为 2，则取关；其他情况返回错误。
* `RelationQuery` 查询用户间关系，结果由 `relation_code` 定义，有四个状态：
  * `relation_code` 为 1，未关注。
  * `relation_code` 为 2，已关注。
  * `relation_code` 为 3，对方为用户粉丝。
  * `relation_code` 为 4，好友，互相关注。
* `FollowList`，`FollowerList`，`FriendList` 三个接口分别关注列表，粉丝列表，好友列表。

**在关系服务中，主要的难点在于如果提高关系服务的性能。因为关系服务是一个频繁使用的服务，并伴随着大量的更新和查询。所以对该服务进行合理的设计是有必要的。**

**在本模块中我将 **`Follow` 结构体设计如下，

```
type (
Follow struct {
gorm.Model
UserId   int64 `json:"user_id,omitempty" gorm:"column:user_id;index"`
ToUserId int64 `json:"to_user_id,omitempty" gorm:"column:to_user_id;index"`
}
)
```

**即存储用户与目标用户的 **`ID`，以及基础的表字段包括 `ID`, `CreateAt`, `UpdateAt`, `DeleteAt`。

**通过查询用户为 **`UserId` 的数据行返回关注列表，通过查询用户为 `ToUserId` 的数据行返回粉丝列表，通过去上面两表的交集获取用户好友列表。

**当然这里还有许多优化空间，但是由于时间关系编码者无法在这里进行实现，将思路叙述如下。**

**由于用户与用户之间是一个多对多的关系，那么实际上我们可以使用 Redis 处理这样的关系，为每个用户在 Redis 中创建一个 hash 表。hash 表内包含该用户关注的所有用户 ID。这样会对处理器带来更大的负担，但是可以显著提高查询速度。与此同时，由于 Redis 并不是一个持久化数据库，有数据丢失的危险，需要设立定时任务将 Redis 中的信息更新到用户 Follow 表之中。除此之外，为了应对高峰期的压力，可以在 Redis 外再接一层 MQ 进行削峰处理，保护服务器不会因为大量的请求宕机。**



There are five apis in the follow service.
