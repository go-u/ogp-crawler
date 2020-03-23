package crawler

import (
	"github.com/dghubble/go-twitter/twitter"
	"log"
)

func (c *Crawler) Crawl(tweet twitter.Tweet) {
	// cacheされているツイートは重複処理を避けるため何もしない
	hasCache, err := c.MemStoreUsecase.HasCacheTweet(tweet.IDStr)
	if err != nil || hasCache {
		log.Println(err, hasCache)
		return
	}

	// 将来の重複処理を避けるため処理するtweetをキャッシュ
	err = c.MemStoreUsecase.AddCacheTweet(tweet)
	if err != nil {
		log.Println(err)
		return
	}

	// link先のOGP情報をcache済みの場合はcrawl処理をスキップ
	// (複数のツイートが同じリンク先をシェアしている場合に該当)
	link := tweet.Entities.Urls[0].ExpandedURL
	ogp, err := c.MemStoreUsecase.GetCacheOgp(link)
	if err != nil {
		log.Println(err)
		return
	}

	// link先のOGP情報をcacheしていない場合は探索＆取得
	if ogp == nil {
		// ogpを探索
		ogp, err = c.fetchOgp(tweet)
		if err != nil {
			log.Println(err)
			return
		}
		// 探索したogpを保存
		err = c.MemStoreUsecase.AddCacheOgp(ogp)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// サンプルとしてOGP情報を記録
	// DB容量節約のためFQDN毎に60分に1度のみ記録(config.ymlで変更可能)
	isRateLimitedRecordOgp, err := c.MemStoreUsecase.IsRateLimitedRecordOgp(ogp.FQDN)
	if err != nil {
		log.Println(err)
		return
	}
	if !isRateLimitedRecordOgp {
		// 探索したogpをサンプルOGPに追加
		err := c.OgpUsecase.RecordSample(ogp)
		if err != nil {
			log.Println(err)
			return
		}

		// rate limitを追加
		err = c.MemStoreUsecase.AddRateLimitRecordOgp(ogp.FQDN, c.RateLimitDurationOgp)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// stat(統計情報を更新)
	err = c.StatUsecase.Record(ogp)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
