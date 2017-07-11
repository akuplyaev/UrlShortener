package storage
import (
	"github.com/go-redis/redis"
	"crypto/md5"
	"encoding/hex"
)
var store = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})
//создаем  shortUrl и ищем его в хранилище,не нашли-добавляем туда
func FindShortUrl(url string )string  {
	shortUrl:=GenerateUrl(url)
	_,err:=store.Get(shortUrl).Result()
	if err!=nil{
		store.Set(shortUrl,url,0)
		return shortUrl
	}
	return shortUrl
}
func FindLongUrl(shortUrl string ) string {
	longUrl,err:=store.Get(shortUrl).Result()
	if err!=nil{
		longUrl=""
		return longUrl
	}
	return longUrl

}
//создаем новый url на основе хеша...толково ничего придумать не смог
func GenerateUrl(url string) string {
	hash := md5.New()
	hash.Write([]byte(url))
	shortUrl:=hex.EncodeToString(hash.Sum(nil))[0:8]
	return  shortUrl
}