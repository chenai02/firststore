package handle

import(
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)
var ctx = context.Background()
var redisdb *redis.Client

func redis_init(){
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func SetCurrentUser(username string){
	redis_init()
	err := redisdb.Ping(ctx).Err()
	if err != nil {
		fmt.Println("connect redis failed")
	}
	err = redisdb.Set(ctx, "currentuser", username, 100*time.Minute).Err()
	if err != nil {
		fmt.Println("Set data to redis failed")
	}
	//fmt.Println("currentuser:",username)
}
func GetCurrentUser()string{
	redis_init()
	err := redisdb.Ping(ctx).Err()
	if err != nil {
		fmt.Println("connect redis failed")
	}
	err = redisdb.Get(ctx,"currentuser").Err()
	if err != nil {
		fmt.Println("Get data from redis failed")
	}
	return redisdb.Get(ctx,"currentuser").Val()

}
