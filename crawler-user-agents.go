package main

import (
	"encoding/json"
	"sync"
)

type CrawlerPattern struct {
	Pattern      string   `json:"pattern"`
	AdditionDate string   `json:"addition_date"`
	Instances    []string `json:"instances"`
	Url          string   `json:"url"`
}

var parseOnce sync.Once
var crawlerPatterns []CrawlerPattern

func getCrawlerPatterns() []CrawlerPattern {
	parseOnce.Do(func() {
		crawlerPatterns = parseCrawlerPatterns()
	})

	return crawlerPatterns
}

func parseCrawlerPatterns() []CrawlerPattern {
	crawlers := `
		[
		  {
			"pattern": "Googlebot\\/",
			"url": "http://www.google.com/bot.html",
			"instances": [
			  "Googlebot/2.1 (+http://www.google.com/bot.html)",
			  "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Safari/537.36"
			]
		  }
		,
		  {
			"pattern": "Googlebot-Mobile",
			"instances": [
			  "DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_1 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
			  "Nokia6820/2.0 (4.83) Profile/MIDP-1.0 Configuration/CLDC-1.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
			  "SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)"
			]
		  }
		,
		  {
			"pattern": "Googlebot-Image",
			"instances": [
			  "Googlebot-Image/1.0"
			]
		  }
		,
		  {
			"pattern": "Googlebot-News",
			"instances": [
			  "Googlebot-News"
			]
		  }
		,
		  {
			"pattern": "Googlebot-Video",
			"instances": [
			  "Googlebot-Video/1.0"
			]
		  }
		,
		  {
			"pattern": "AdsBot-Google([^-]|$)",
			"url": "https://support.google.com/webmasters/answer/1061943?hl=en",
			"instances": [
			  "AdsBot-Google (+http://www.google.com/adsbot.html)"
			]
		  }
		,
		  {
			"pattern": "AdsBot-Google-Mobile",
			"addition_date": "2017/08/21",
			"url": "https://support.google.com/adwords/answer/2404197",
			"instances": [
			  "AdsBot-Google-Mobile-Apps",
			  "Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1 (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)"
			]
		  }
		,
		  {
			"pattern": "Feedfetcher-Google",
			"addition_date": "2018/06/27",
			"url": "https://support.google.com/webmasters/answer/178852",
			"instances": [
			  "Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; 1 subscribers; feed-id=728742641706423)"
			]
		  }
		,
		  {
			"pattern": "Mediapartners-Google",
			"url": "https://support.google.com/webmasters/answer/1061943?hl=en",
			"instances": [
			  "Mediapartners-Google",
			  "Mozilla/5.0 (compatible; MSIE or Firefox mutant; not on Windows server;) Daumoa/4.0 (Following Mediapartners-Google)",
			  "Mozilla/5.0 (iPhone; U; CPU iPhone OS 10_0 like Mac OS X; en-us) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A5297c Safari/602.1 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)",
			  "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_1 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)"
			]
		  }
		,
		  {
			"pattern": "Mediapartners \\(Googlebot\\)",
			"addition_date": "2017/08/08",
			"url": "https://support.google.com/webmasters/answer/1061943?hl=en",
			"instances": []
		  }
		,
		  {
			"pattern": "APIs-Google",
			"addition_date": "2017/08/08",
			"url": "https://support.google.com/webmasters/answer/1061943?hl=en",
			"instances": [
			  "APIs-Google (+https://developers.google.com/webmasters/APIs-Google.html)"
			]
		  }
		,
		  {
			"pattern": "bingbot",
			"url": "http://www.bing.com/bingbot.htm",
			"instances": [
			  "Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 530) like Gecko (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (compatible; adidxbot/2.0;  http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (compatible; bingbot/2.0;  http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm",
			  "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) SitemapProbe",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; adidxbot/2.0;  http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0;  http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 (seoanalyzer; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
			  "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) Safari/537.36"
			]
		  }
		,
		  {
			"pattern": "Slurp",
			"url": "http://help.yahoo.com/help/us/ysearch/slurp",
			"instances": [
			  "Mozilla/5.0 (compatible; Yahoo! Slurp/3.0; http://help.yahoo.com/help/us/ysearch/slurp)",
			  "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
			  "Mozilla/5.0 (compatible; Yahoo! Slurp China; http://misc.yahoo.com.cn/help.html)"
			]
		  }
		,
		  {
			"pattern": "[wW]get",
			"instances": [
			  "WGETbot/1.0 (+http://wget.alanreed.org)",
			  "Wget/1.14 (linux-gnu)",
			  "Wget/1.20.3 (linux-gnu)"
			]
		  }
		,
		  {
			"pattern": "LinkedInBot",
			"instances": [
			  "LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/3.1 +http://www.linkedin.com)",
			  "LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/4.3 +http://www.linkedin.com)",
			  "LinkedInBot/1.0 (compatible; Mozilla/5.0; Apache-HttpClient +http://www.linkedin.com)"
			]
		  }
		,
		  {
			"pattern": "Python-urllib",
			"instances": [
			  "Python-urllib/1.17",
			  "Python-urllib/2.5",
			  "Python-urllib/2.6",
			  "Python-urllib/2.7",
			  "Python-urllib/3.1",
			  "Python-urllib/3.2",
			  "Python-urllib/3.3",
			  "Python-urllib/3.4",
			  "Python-urllib/3.5",
			  "Python-urllib/3.6",
			  "Python-urllib/3.7"
			]
		  }
		,
		  {
			"pattern": "python-requests",
			"addition_date": "2018/05/27",
			"instances": [
			  "python-requests/2.9.2",
			  "python-requests/2.11.1",
			  "python-requests/2.18.4",
			  "python-requests/2.19.1",
			  "python-requests/2.20.0",
			  "python-requests/2.21.0",
			  "python-requests/2.22.0"
			]
		  }
		,
		  {
			"pattern": "libwww-perl",
			"instances": [
			  "2Bone_LinkChecker/1.0 libwww-perl/6.03",
			  "2Bone_LinkChkr/1.0 libwww-perl/6.03",
			  "amibot - http://www.amidalla.de - tech@amidalla.com libwww-perl/5.831"
			]
		  }
		,
		  {
			"pattern": "httpunit",
			"instances": [
			  "httpunit/1.x"
			]
		  }
		,
		  {
			"pattern": "nutch",
			"instances": [
			  "NutchCVS/0.7.1 (Nutch; http://lucene.apache.org/nutch/bot.html; nutch-agent@lucene.apache.org)",
			  "istellabot-nutch/Nutch-1.10"
			]
		  }
		,
		  {
			"pattern": "Go-http-client",
			"addition_date": "2016/03/26",
			"url": "https://golang.org/pkg/net/http/",
			"instances": [
			  "Go-http-client/1.1",
			  "Go-http-client/2.0"
			]
		  }
		,
		  {
			"pattern": "phpcrawl",
			"addition_date": "2012-09/17",
			"url": "http://phpcrawl.cuab.de/",
			"instances": [
			  "phpcrawl"
			]
		  }
		,
		  {
			"pattern": "msnbot",
			"url": "http://search.msn.com/msnbot.htm",
			"instances": [
			  "adidxbot/1.1 (+http://search.msn.com/msnbot.htm)",
			  "adidxbot/2.0 (+http://search.msn.com/msnbot.htm)",
			  "librabot/1.0 (+http://search.msn.com/msnbot.htm)",
			  "librabot/2.0 (+http://search.msn.com/msnbot.htm)",
			  "msnbot-NewsBlogs/2.0b (+http://search.msn.com/msnbot.htm)",
			  "msnbot-UDiscovery/2.0b (+http://search.msn.com/msnbot.htm)",
			  "msnbot-media/1.0 (+http://search.msn.com/msnbot.htm)",
			  "msnbot-media/1.1 (+http://search.msn.com/msnbot.htm)",
			  "msnbot-media/2.0b (+http://search.msn.com/msnbot.htm)",
			  "msnbot/1.0 (+http://search.msn.com/msnbot.htm)",
			  "msnbot/1.1 (+http://search.msn.com/msnbot.htm)",
			  "msnbot/2.0b (+http://search.msn.com/msnbot.htm)",
			  "msnbot/2.0b (+http://search.msn.com/msnbot.htm).",
			  "msnbot/2.0b (+http://search.msn.com/msnbot.htm)._"
			]
		  }
		,
		  {
			"pattern": "jyxobot",
			"instances": []
		  }
		,
		  {
			"pattern": "FAST-WebCrawler",
			"instances": [
			  "FAST-WebCrawler/3.6/FirstPage (atw-crawler at fast dot no;http://fast.no/support/crawler.asp)",
			  "FAST-WebCrawler/3.7 (atw-crawler at fast dot no; http://fast.no/support/crawler.asp)",
			  "FAST-WebCrawler/3.7/FirstPage (atw-crawler at fast dot no;http://fast.no/support/crawler.asp)",
			  "FAST-WebCrawler/3.8"
			]
		  }
		,
		  {
			"pattern": "FAST Enterprise Crawler",
			"instances": [
			  "FAST Enterprise Crawler 6 / Scirus scirus-crawler@fast.no; http://www.scirus.com/srsapp/contactus/",
			  "FAST Enterprise Crawler 6 used by Schibsted (webcrawl@schibstedsok.no)"
			]
		  }
		,
		  {
			"pattern": "BIGLOTRON",
			"instances": [
			  "BIGLOTRON (Beta 2;GNU/Linux)"
			]
		  }
		,
		  {
			"pattern": "Teoma",
			"instances": [
			  "Mozilla/2.0 (compatible; Ask Jeeves/Teoma; +http://sp.ask.com/docs/about/tech_crawling.html)",
			  "Mozilla/2.0 (compatible; Ask Jeeves/Teoma; +http://about.ask.com/en/docs/about/webmasters.shtml)"
			],
			"url": "http://about.ask.com/en/docs/about/webmasters.shtml"
		  }
		,
		  {
			"pattern": "convera",
			"instances": [
			  "ConveraCrawler/0.9e (+http://ews.converasearch.com/crawl.htm)"
			],
			"url": "http://ews.converasearch.com/crawl.htm"
		  }
		,
		  {
			"pattern": "seekbot",
			"instances": [
			  "Seekbot/1.0 (http://www.seekbot.net/bot.html) RobotsTxtFetcher/1.2"
			],
			"url": "http://www.seekbot.net/bot.html"
		  }
		,
		  {
			"pattern": "Gigabot",
			"instances": [
			  "Gigabot/1.0",
			  "Gigabot/2.0 (http://www.gigablast.com/spider.html)"
			],
			"url": "http://www.gigablast.com/spider.html"
		  }
		,
		  {
			"pattern": "Gigablast",
			"instances": [
			  "GigablastOpenSource/1.0"
			],
			"url": "https://github.com/gigablast/open-source-search-engine"
		  }
		,
		  {
			"pattern": "exabot",
			"instances": [
			  "Mozilla/5.0 (compatible; Alexabot/1.0; +http://www.alexa.com/help/certifyscan; certifyscan@alexa.com)",
			  "Mozilla/5.0 (compatible; Exabot PyExalead/3.0; +http://www.exabot.com/go/robot)",
			  "Mozilla/5.0 (compatible; Exabot-Images/3.0; +http://www.exabot.com/go/robot)",
			  "Mozilla/5.0 (compatible; Exabot/3.0 (BiggerBetter); +http://www.exabot.com/go/robot)",
			  "Mozilla/5.0 (compatible; Exabot/3.0; +http://www.exabot.com/go/robot)",
			  "Mozilla/5.0 (compatible; Exabot/3.0;  http://www.exabot.com/go/robot)"
			]
		  }
		,
		  {
			"pattern": "ia_archiver",
			"instances": [
			  "ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)",
			  "ia_archiver-web.archive.org"
			]
		  }
		,
		  {
			"pattern": "GingerCrawler",
			"instances": [
			  "GingerCrawler/1.0 (Language Assistant for Dyslexics; www.gingersoftware.com/crawler_agent.htm; support at ginger software dot com)"
			]
		  }
		,
		  {
			"pattern": "webmon ",
			"instances": []
		  }
		,
		  {
			"pattern": "HTTrack",
			"instances": [
			  "Mozilla/4.5 (compatible; HTTrack 3.0x; Windows 98)"
			]
		  }
		,
		  {
			"pattern": "grub.org",
			"instances": [
			  "Mozilla/4.0 (compatible; grub-client-0.3.0; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.0.4; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.0.5; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.0.6; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.0.7; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.1.1; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.2.1; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.3.1; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.3.7; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.4.3; Crawl your own stuff with http://grub.org)",
			  "Mozilla/4.0 (compatible; grub-client-1.5.3; Crawl your own stuff with http://grub.org)"
			]
		  }
		,
		  {
			"pattern": "UsineNouvelleCrawler",
			"instances": []
		  }
		,
		  {
			"pattern": "antibot",
			"instances": []
		  }
		,
		  {
			"pattern": "netresearchserver",
			"instances": []
		  }
		,
		  {
			"pattern": "speedy",
			"instances": [
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Speedy Spider (http://www.entireweb.com/about/search_tech/speedy_spider/)",
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Speedy Spider for SpeedyAds (http://www.entireweb.com/about/search_tech/speedy_spider/)",
			  "Mozilla/5.0 (compatible; Speedy Spider; http://www.entireweb.com/about/search_tech/speedy_spider/)",
			  "Speedy Spider (Entireweb; Beta/1.2; http://www.entireweb.com/about/search_tech/speedyspider/)",
			  "Speedy Spider (http://www.entireweb.com/about/search_tech/speedy_spider/)"
			]
		  }
		,
		  {
			"pattern": "fluffy",
			"instances": []
		  }
		,
		  {
			"pattern": "findlink",
			"instances": [
			  "findlinks/1.0 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.3-beta8 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.3-beta9 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.5-beta7 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.6-beta1 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.6-beta1 (+http://wortschatz.uni-leipzig.de/findlinks/; YaCy 0.1; yacy.net)",
			  "findlinks/1.1.6-beta2 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.6-beta3 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.6-beta4 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.6-beta5 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/1.1.6-beta6 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.0 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.0.1 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.0.2 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.0.4 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.0.5 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.0.9 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.1 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.1.3 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.1.5 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.2 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.5 (+http://wortschatz.uni-leipzig.de/findlinks/)",
			  "findlinks/2.6 (+http://wortschatz.uni-leipzig.de/findlinks/)"
			]
		  }
		,
		  {
			"pattern": "msrbot",
			"instances": []
		  }
		,
		  {
			"pattern": "panscient",
			"instances": [
			  "panscient.com"
			]
		  }
		,
		  {
			"pattern": "yacybot",
			"instances": [
			  "yacybot (/global; amd64 FreeBSD 10.3-RELEASE; java 1.8.0_77; GMT/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 FreeBSD 10.3-RELEASE-p7; java 1.7.0_95; GMT/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 FreeBSD 9.2-RELEASE-p10; java 1.7.0_65; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 2.6.32-042stab093.4; java 1.7.0_65; Etc/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 2.6.32-042stab094.8; java 1.7.0_79; America/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 2.6.32-042stab108.8; java 1.7.0_91; America/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 2.6.32-042stab111.11; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 2.6.32-042stab116.1; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 2.6.32-573.3.1.el6.x86_64; java 1.7.0_85; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 3.10.0-229.4.2.el7.x86_64; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 3.10.0-229.4.2.el7.x86_64; java 1.8.0_45; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.10.0-229.7.2.el7.x86_64; java 1.8.0_45; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.10.0-327.22.2.el7.x86_64; java 1.7.0_101; Etc/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.11.10-21-desktop; java 1.7.0_51; America/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.12.1; java 1.7.0_65; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-042stab093.4; java 1.7.0_79; Europe/de) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-042stab093.4; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-45-generic; java 1.7.0_75; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 3.13.0-61-generic; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-74-generic; java 1.7.0_91; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-83-generic; java 1.7.0_95; Europe/de) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-83-generic; java 1.7.0_95; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-85-generic; java 1.7.0_101; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-85-generic; java 1.7.0_95; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.13.0-88-generic; java 1.7.0_101; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.14-0.bpo.1-amd64; java 1.7.0_55; Europe/de) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.14.32-xxxx-grs-ipv6-64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 3.14.32-xxxx-grs-ipv6-64; java 1.8.0_111; Europe/de) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_111; Europe/de) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; America/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_79; Europe/de) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_91; Europe/de) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_95; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.8.0_111; Europe/en) http://yacy.net/bot.html",
			  "yacybot (/global; amd64 Linux 3.16-0.bpo.2-amd64; java 1.7.0_65; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 3.19.0-15-generic; java 1.8.0_45-internal; Europe/de) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 3.2.0-4-amd64; java 1.7.0_65; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 3.2.0-4-amd64; java 1.7.0_67; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 4.4.0-57-generic; java 9-internal; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Windows 8.1 6.3; java 1.7.0_55; Europe/de) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Windows 8 6.2; java 1.7.0_55; Europe/de) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 5.2.8-Jinsol; java 12.0.2; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 5.2.9-Jinsol; java 12.0.2; Europe/en) http://yacy.net/bot.html",
			  "yacybot (-global; amd64 Linux 5.2.11-Jinsol; java 12.0.2; Europe/en) http://yacy.net/bot.html"
			]
		  }
		,
		  {
			"pattern": "AISearchBot",
			"instances": []
		  }
		,
		  {
			"pattern": "ips-agent",
			"instances": [
			  "BlackBerry9000/4.6.0.167 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/102 ips-agent",
			  "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.12; ips-agent) Gecko/20050922 Fedora/1.0.7-1.1.fc4 Firefox/1.0.7",
			  "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.1.3; ips-agent) Gecko/20090824 Fedora/1.0.7-1.1.fc4  Firefox/3.5.3",
			  "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.24; ips-agent) Gecko/20111107 Ubuntu/10.04 (lucid) Firefox/3.6.24",
			  "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:14.0; ips-agent) Gecko/20100101 Firefox/14.0.1"
			]
		  }
		,
		  {
			"pattern": "tagoobot",
			"instances": []
		  }
		,
		  {
			"pattern": "MJ12bot",
			"instances": [
			  "MJ12bot/v1.2.0 (http://majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.2.1; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.2.3; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.2.4; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.2.5; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.3.0; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.3.1; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.3.2; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.3.3; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.0; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.1; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.2; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.3; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.4 (domain ownership verifier); http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.4; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.5; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.6; http://mj12bot.com/)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.7; http://mj12bot.com/)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.7; http://www.majestic12.co.uk/bot.php?+)",
			  "Mozilla/5.0 (compatible; MJ12bot/v1.4.8; http://mj12bot.com/)"
			]
		  }
		,
		  {
			"pattern": "woriobot",
			"instances": [
			  "Mozilla/5.0 (compatible; woriobot +http://worio.com)",
			  "Mozilla/5.0 (compatible; woriobot support [at] zite [dot] com +http://zite.com)"
			]
		  }
		,
		  {
			"pattern": "yanga",
			"instances": [
			  "Yanga WorldSearch Bot v1.1/beta (http://www.yanga.co.uk/)"
			]
		  }
		,
		  {
			"pattern": "buzzbot",
			"instances": [
			  "Buzzbot/1.0 (Buzzbot; http://www.buzzstream.com; buzzbot@buzzstream.com)"
			]
		  }
		,
		  {
			"pattern": "mlbot",
			"instances": [
			  "MLBot (www.metadatalabs.com/mlbot)"
			]
		  }
		,
		  {
			"pattern": "YandexBot",
			"url": "http://yandex.com/bots",
			"instances": [
			  "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)"
			],
			"addition_date": "2015/04/14"
		  }
		,
		  {
			"pattern": "YandexImages",
			"url": "http://yandex.com/bots",
			"instances": [
			  "Mozilla/5.0 (compatible; YandexImages/3.0; +http://yandex.com/bots)"
			],
			"addition_date": "2015/04/14"
		  }
		,
		  {
			"pattern": "YandexAccessibilityBot",
			"url": "http://yandex.com/bots",
			"instances": [
			  "Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots"
			],
			"addition_date": "2019/03/01"
		  }
		,
		  {
			"pattern": "YandexMobileBot",
			"url": "https://yandex.com/support/webmaster/robot-workings/check-yandex-robots.xml#robot-in-logs",
			"instances": [
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4 (compatible; YandexMobileBot/3.0; +http://yandex.com/bots)"
			],
			"addition_date": "2016/12/01"
		  }
		,
		  {
			"pattern": "purebot",
			"addition_date": "2010/01/19",
			"instances": []
		  }
		,
		  {
			"pattern": "Linguee Bot",
			"addition_date": "2010/01/26",
			"url": "http://www.linguee.com/bot",
			"instances": [
			  "Linguee Bot (http://www.linguee.com/bot)",
			  "Linguee Bot (http://www.linguee.com/bot; bot@linguee.com)"
			]
		  }
		,
		  {
			"pattern": "CyberPatrol",
			"addition_date": "2010/02/11",
			"url": "http://www.cyberpatrol.com/cyberpatrolcrawler.asp",
			"instances": [
			  "CyberPatrol SiteCat Webbot (http://www.cyberpatrol.com/cyberpatrolcrawler.asp)"
			]
		  }
		,
		  {
			"pattern": "voilabot",
			"addition_date": "2010/05/18",
			"instances": [
			  "Mozilla/5.0 (Windows NT 5.1; U; Win64; fr; rv:1.8.1) VoilaBot BETA 1.2 (support.voilabot@orange-ftgroup.com)",
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.8.1) VoilaBot BETA 1.2 (support.voilabot@orange-ftgroup.com)"
			]
		  }
		,
		  {
			"pattern": "Baiduspider",
			"addition_date": "2010/07/15",
			"url": "http://www.baidu.jp/spider/",
			"instances": [
			  "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
			  "Mozilla/5.0 (compatible; Baiduspider-render/2.0; +http://www.baidu.com/search/spider.html)"
			]
		  }
		,
		  {
			"pattern": "citeseerxbot",
			"addition_date": "2010/07/17",
			"instances": []
		  }
		,
		  {
			"pattern": "spbot",
			"addition_date": "2010/07/31",
			"url": "http://www.seoprofiler.com/bot",
			"instances": [
			  "Mozilla/5.0 (compatible; spbot/1.0; +http://www.seoprofiler.com/bot/ )",
			  "Mozilla/5.0 (compatible; spbot/1.1; +http://www.seoprofiler.com/bot/ )",
			  "Mozilla/5.0 (compatible; spbot/1.2; +http://www.seoprofiler.com/bot/ )",
			  "Mozilla/5.0 (compatible; spbot/2.0.1; +http://www.seoprofiler.com/bot/ )",
			  "Mozilla/5.0 (compatible; spbot/2.0.2; +http://www.seoprofiler.com/bot/ )",
			  "Mozilla/5.0 (compatible; spbot/2.0.3; +http://www.seoprofiler.com/bot/ )",
			  "Mozilla/5.0 (compatible; spbot/2.0.4; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/2.0; +http://www.seoprofiler.com/bot/ )",
			  "Mozilla/5.0 (compatible; spbot/2.1; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/3.0; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/3.1; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.1; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.2; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.3; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.4; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.5; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.6; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.7; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.7; +https://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.8; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0.9; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0a; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.0b; +http://www.seoprofiler.com/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.1.0; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.2.0; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.3.0; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.4.0; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.4.1; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/4.4.2; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/5.0.1; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/5.0.2; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/5.0.3; +http://OpenLinkProfiler.org/bot )",
			  "Mozilla/5.0 (compatible; spbot/5.0; +http://OpenLinkProfiler.org/bot )"
			]
		  }
		,
		  {
			"pattern": "twengabot",
			"addition_date": "2010/08/03",
			"url": "http://www.twenga.com/bot.html",
			"instances": []
		  }
		,
		  {
			"pattern": "postrank",
			"addition_date": "2010/08/03",
			"url": "http://www.postrank.com",
			"instances": [
			  "PostRank/2.0 (postrank.com)",
			  "PostRank/2.0 (postrank.com; 1 subscribers)"
			]
		  }
		,
		  {
			"pattern": "TurnitinBot",
			"addition_date": "2010/09/26",
			"url": "http://www.turnitin.com",
			"instances": [
			  "TurnitinBot (https://turnitin.com/robot/crawlerinfo.html)"
			]
		  }
		,
		  {
			"pattern": "scribdbot",
			"addition_date": "2010/09/28",
			"url": "http://www.scribd.com",
			"instances": []
		  }
		,
		  {
			"pattern": "page2rss",
			"addition_date": "2010/10/07",
			"url": "http://www.page2rss.com",
			"instances": [
			  "Mozilla/5.0 (compatible;  Page2RSS/0.7; +http://page2rss.com/)"
			]
		  }
		,
		  {
			"pattern": "sitebot",
			"addition_date": "2010/12/15",
			"url": "http://www.sitebot.org",
			"instances": [
			  "Mozilla/5.0 (compatible; Whoiswebsitebot/0.1; +http://www.whoiswebsite.net)"
			]
		  }
		,
		  {
			"pattern": "linkdex",
			"addition_date": "2011/01/06",
			"url": "http://www.linkdex.com",
			"instances": [
			  "Mozilla/5.0 (compatible; linkdexbot/2.0; +http://www.linkdex.com/about/bots/)",
			  "Mozilla/5.0 (compatible; linkdexbot/2.0; +http://www.linkdex.com/bots/)",
			  "Mozilla/5.0 (compatible; linkdexbot/2.1; +http://www.linkdex.com/about/bots/)",
			  "Mozilla/5.0 (compatible; linkdexbot/2.1; +http://www.linkdex.com/bots/)",
			  "Mozilla/5.0 (compatible; linkdexbot/2.2; +http://www.linkdex.com/bots/)",
			  "linkdex.com/v2.0",
			  "linkdexbot/Nutch-1.0-dev (http://www.linkdex.com/; crawl at linkdex dot com)"
			]
		  }
		,
		  {
			"pattern": "Adidxbot",
			"url": "http://onlinehelp.microsoft.com/en-us/bing/hh204496.aspx",
			"instances": []
		  }
		,
		  {
			"pattern": "ezooms",
			"addition_date": "2011/04/27",
			"url": "http://www.phpbb.com/community/viewtopic.php?f=64&t=935605&start=450#p12948289",
			"instances": [
			  "Mozilla/5.0 (compatible; Ezooms/1.0; ezooms.bot@gmail.com)"
			]
		  }
		,
		  {
			"pattern": "dotbot",
			"addition_date": "2011/04/27",
			"instances": [
			  "Mozilla/5.0 (compatible; DotBot/1.1; http://www.opensiteexplorer.org/dotbot, help@moz.com)",
			  "dotbot"
			]
		  }
		,
		  {
			"pattern": "Mail.RU_Bot",
			"addition_date": "2011/04/27",
			"instances": [
			  "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/help/robots)",
			  "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/",
			  "Mozilla/5.0 (compatible; Mail.RU_Bot/2.0; +http://go.mail.ru/",
			  "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/Robots/2.0; +http://go.mail.ru/help/robots)"
			]
		  }
		,
		  {
			"pattern": "discobot",
			"addition_date": "2011/05/03",
			"url": "http://discoveryengine.com/discobot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; discobot/1.0; +http://discoveryengine.com/discobot.html)",
			  "Mozilla/5.0 (compatible; discobot/2.0; +http://discoveryengine.com/discobot.html)",
			  "mozilla/5.0 (compatible; discobot/1.1; +http://discoveryengine.com/discobot.html)"
			]
		  }
		,
		  {
			"pattern": "heritrix",
			"addition_date": "2011/06/21",
			"url": "https://github.com/internetarchive/heritrix3/wiki",
			"instances": [
			  "Mozilla/5.0 (compatible; heritrix/1.12.1 +http://www.webarchiv.cz)",
			  "Mozilla/5.0 (compatible; heritrix/1.12.1b +http://netarkivet.dk/website/info.html)",
			  "Mozilla/5.0 (compatible; heritrix/1.14.2 +http://rjpower.org)",
			  "Mozilla/5.0 (compatible; heritrix/1.14.2 +http://www.webarchiv.cz)",
			  "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://archive.org)",
			  "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://www.accelobot.com)",
			  "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://www.webarchiv.cz)",
			  "Mozilla/5.0 (compatible; heritrix/1.14.3.r6601 +http://www.buddybuzz.net/yptrino)",
			  "Mozilla/5.0 (compatible; heritrix/1.14.4 +http://parsijoo.ir)",
			  "Mozilla/5.0 (compatible; heritrix/1.14.4 +http://www.exif-search.com)",
			  "Mozilla/5.0 (compatible; heritrix/2.0.2 +http://aihit.com)",
			  "Mozilla/5.0 (compatible; heritrix/2.0.2 +http://seekda.com)",
			  "Mozilla/5.0 (compatible; heritrix/3.0.0-SNAPSHOT-20091120.021634 +http://crawler.archive.org)",
			  "Mozilla/5.0 (compatible; heritrix/3.1.0-RC1 +http://boston.lti.cs.cmu.edu/crawler_12/)",
			  "Mozilla/5.0 (compatible; heritrix/3.1.1 +http://places.tomtom.com/crawlerinfo)",
			  "Mozilla/5.0 (compatible; heritrix/3.1.1 +http://www.mixdata.com)",
			  "Mozilla/5.0 (compatible; heritrix/3.1.1; UniLeipzigASV +http://corpora.informatik.uni-leipzig.de/crawler_faq.html)",
			  "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.crim.ca)",
			  "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.exif-search.com)",
			  "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.mixdata.com)",
			  "Mozilla/5.0 (compatible; heritrix/3.3.0-SNAPSHOT-20160309-0050; UniLeipzigASV +http://corpora.informatik.uni-leipzig.de/crawler_faq.html)",
			  "Mozilla/5.0 (compatible; sukibot_heritrix/3.1.1 +http://suki.ling.helsinki.fi/eng/webmasters.html)"
			]
		  }
		,
		  {
			"pattern": "findthatfile",
			"addition_date": "2011/06/21",
			"url": "http://www.findthatfile.com/",
			"instances": []
		  }
		,
		  {
			"pattern": "europarchive.org",
			"addition_date": "2011/06/21",
			"url": "",
			"instances": [
			  "Mozilla/5.0 (compatible; MSIE 7.0 +http://www.europarchive.org)"
			]
		  }
		,
		  {
			"pattern": "NerdByNature.Bot",
			"addition_date": "2011/07/12",
			"url": "http://www.nerdbynature.net/bot",
			"instances": [
			  "Mozilla/5.0 (compatible; NerdByNature.Bot; http://www.nerdbynature.net/bot)"
			]
		  }
		,
		  {
			"pattern": "sistrix crawler",
			"addition_date": "2011/08/02",
			"instances": []
		  }
		,
		  {
			"pattern": "Ahrefs(Bot|SiteAudit)",
			"addition_date": "2011/08/28",
			"instances": [
			  "Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)",
			  "Mozilla/5.0 (compatible; AhrefsSiteAudit/6.1; +http://ahrefs.com/robot/)",
			  "Mozilla/5.0 (compatible; AhrefsBot/5.2; News; +http://ahrefs.com/robot/)",
			  "Mozilla/5.0 (compatible; AhrefsBot/5.2; +http://ahrefs.com/robot/)",
			  "Mozilla/5.0 (compatible; AhrefsSiteAudit/5.2; +http://ahrefs.com/robot/)",
			  "Mozilla/5.0 (compatible; AhrefsBot/6.1; News; +http://ahrefs.com/robot/)"
			]
		  }
		,
		  {
			"pattern": "fuelbot",
			"addition_date": "2018/06/28",
			"instances": [
			  "fuelbot"
			]
		  }
		,
		  {
			"pattern": "CrunchBot",
			"addition_date": "2018/06/28",
			"instances": [
			  "CrunchBot/1.0 (+http://www.leadcrunch.com/crunchbot)"
			]
		  }
		,
		  {
			"pattern": "IndeedBot",
			"addition_date": "2018/06/28",
			"instances": [
			  "Mozilla/5.0 (Windows NT 6.1; rv:38.0) Gecko/20100101 Firefox/38.0 (IndeedBot 1.1)"
			]
		  }
		,
		  {
			"pattern": "mappydata",
			"addition_date": "2018/06/28",
			"instances": [
			  "Mozilla/5.0 (compatible; Mappy/1.0; +http://mappydata.net/bot/)"
			]
		  }
		,
		  {
			"pattern": "woobot",
			"addition_date": "2018/06/28",
			"instances": [
			  "woobot"
			]
		  }
		,
		  {
			"pattern": "ZoominfoBot",
			"addition_date": "2018/06/28",
			"instances": [
			  "ZoominfoBot (zoominfobot at zoominfo dot com)"
			]
		  }
		,
		  {
			"pattern": "PrivacyAwareBot",
			"addition_date": "2018/06/28",
			"instances": [
			  "Mozilla/5.0 (compatible; PrivacyAwareBot/1.1; +http://www.privacyaware.org)"
			]
		  }
		,
		  {
			"pattern": "Multiviewbot",
			"addition_date": "2018/06/28",
			"instances": [
			  "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Multiviewbot"
			]
		  }
		,
		  {
			"pattern": "SWIMGBot",
			"addition_date": "2018/06/28",
			"instances": [
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 SWIMGBot"
			]
		  }
		,
		  {
			"pattern": "Grobbot",
			"addition_date": "2018/06/28",
			"instances": [
			  "Mozilla/5.0 (compatible; Grobbot/2.2; +https://grob.it)"
			]
		  }
		,
		  {
			"pattern": "eright",
			"addition_date": "2018/06/28",
			"instances": [
			  "Mozilla/5.0 (compatible; eright/1.0; +bot@eright.com)"
			]
		  }
		,
		  {
			"pattern": "Apercite",
			"addition_date": "2018/06/28",
			"instances": [
			  "Mozilla/5.0 (compatible; Apercite; +http://www.apercite.fr/robot/index.html)"
			]
		  }
		,
		  {
			"pattern": "semanticbot",
			"addition_date": "2018/06/28",
			"instances": [
			  "semanticbot",
			  "semanticbot (info@semanticaudience.com)"
			]
		  }
		,
		  {
			"pattern": "Aboundex",
			"addition_date": "2011/09/28",
			"url": "http://www.aboundex.com/crawler/",
			"instances": [
			  "Aboundex/0.2 (http://www.aboundex.com/crawler/)",
			  "Aboundex/0.3 (http://www.aboundex.com/crawler/)"
			]
		  }
		,
		  {
			"pattern": "domaincrawler",
			"addition_date": "2011/10/21",
			"instances": [
			  "CipaCrawler/3.0 (info@domaincrawler.com; http://www.domaincrawler.com/www.example.com)"
			]
		  }
		,
		  {
			"pattern": "wbsearchbot",
			"addition_date": "2011/12/21",
			"url": "http://www.warebay.com/bot.html",
			"instances": []
		  }
		,
		  {
			"pattern": "summify",
			"addition_date": "2012/01/04",
			"url": "http://summify.com",
			"instances": [
			  "Summify (Summify/1.0.1; +http://summify.com)"
			]
		  }
		,
		  {
			"pattern": "CCBot",
			"addition_date": "2012/02/05",
			"url": "http://www.commoncrawl.org/bot.html",
			"instances": [
			  "CCBot/2.0 (http://commoncrawl.org/faq/)",
			  "CCBot/2.0 (https://commoncrawl.org/faq/)"
			]
		  }
		,
		  {
			"pattern": "edisterbot",
			"addition_date": "2012/02/25",
			"instances": []
		  }
		,
		  {
			"pattern": "seznambot",
			"addition_date": "2012/03/14",
			"instances": [
			  "Mozilla/5.0 (compatible; SeznamBot/3.2-test1-1; +http://napoveda.seznam.cz/en/seznambot-intro/)",
			  "Mozilla/5.0 (compatible; SeznamBot/3.2-test1; +http://napoveda.seznam.cz/en/seznambot-intro/)",
			  "Mozilla/5.0 (compatible; SeznamBot/3.2-test2; +http://napoveda.seznam.cz/en/seznambot-intro/)",
			  "Mozilla/5.0 (compatible; SeznamBot/3.2-test4; +http://napoveda.seznam.cz/en/seznambot-intro/)",
			  "Mozilla/5.0 (compatible; SeznamBot/3.2; +http://napoveda.seznam.cz/en/seznambot-intro/)"
			]
		  }
		,
		  {
			"pattern": "ec2linkfinder",
			"addition_date": "2012/03/22",
			"instances": [
			  "ec2linkfinder"
			]
		  }
		,
		  {
			"pattern": "gslfbot",
			"addition_date": "2012/04/03",
			"instances": []
		  }
		,
		  {
			"pattern": "aiHitBot",
			"addition_date": "2012/04/16",
			"instances": [
			  "Mozilla/5.0 (compatible; aiHitBot/2.9; +https://www.aihitdata.com/about)"
			]
		  }
		,
		  {
			"pattern": "intelium_bot",
			"addition_date": "2012/05/07",
			"instances": []
		  }
		,
		  {
			"pattern": "facebookexternalhit",
			"addition_date": "2012/05/07",
			"instances": [
			  "facebookexternalhit/1.0 (+http://www.facebook.com/externalhit_uatext.php)",
			  "facebookexternalhit/1.1",
			  "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)"
			],
			"url": "https://developers.facebook.com/docs/sharing/webmasters/crawler/"
		  }
		,
		  {
			"pattern": "Yeti",
			"addition_date": "2012/05/07",
			"url": "http://naver.me/bot",
			"instances": [
			  "Mozilla/5.0 (compatible; Yeti/1.1; +http://naver.me/bot)"
			]
		  }
		,
		  {
			"pattern": "RetrevoPageAnalyzer",
			"addition_date": "2012/05/07",
			"instances": [
			  "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; RetrevoPageAnalyzer; +http://www.retrevo.com/content/about-us)"
			]
		  }
		,
		  {
			"pattern": "lb-spider",
			"addition_date": "2012/05/07",
			"instances": []
		  }
		,
		  {
			"pattern": "Sogou",
			"addition_date": "2012/05/13",
			"url": "http://www.sogou.com/docs/help/webmasters.htm#07",
			"instances": [
			  "Sogou News Spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)",
			  "Sogou Pic Spider/3.0(+http://www.sogou.com/docs/help/webmasters.htm#07)",
			  "Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)"
			]
		  }
		,
		  {
			"pattern": "lssbot",
			"addition_date": "2012/05/15",
			"instances": []
		  }
		,
		  {
			"pattern": "careerbot",
			"addition_date": "2012/05/23",
			"url": "http://www.career-x.de/bot.html",
			"instances": []
		  }
		,
		  {
			"pattern": "wotbox",
			"addition_date": "2012/06/12",
			"url": "http://www.wotbox.com",
			"instances": [
			  "Wotbox/2.0 (bot@wotbox.com; http://www.wotbox.com)",
			  "Wotbox/2.01 (+http://www.wotbox.com/bot/)"
			]
		  }
		,
		  {
			"pattern": "wocbot",
			"addition_date": "2012/07/25",
			"url": "http://www.wocodi.com/crawler",
			"instances": []
		  }
		,
		  {
			"pattern": "ichiro",
			"addition_date": "2012/08/28",
			"url": "http://help.goo.ne.jp/help/article/1142",
			"instances": [
			  "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)",
			  "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
			  "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo;+http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
			  "DoCoMo/2.0 P900i(c100;TB;W24H11)(compatible; ichiro/mobile goo;+http://help.goo.ne.jp/door/crawler.html)",
			  "DoCoMo/2.0 P901i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/door/crawler.html)",
			  "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)",
			  "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo; +http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
			  "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo;+http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
			  "ichiro/2.0 (http://help.goo.ne.jp/door/crawler.html)",
			  "ichiro/2.0 (ichiro@nttr.co.jp)",
			  "ichiro/3.0 (http://help.goo.ne.jp/door/crawler.html)",
			  "ichiro/3.0 (http://help.goo.ne.jp/help/article/1142)",
			  "ichiro/3.0 (http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
			  "ichiro/4.0 (http://help.goo.ne.jp/door/crawler.html)",
			  "ichiro/5.0 (http://help.goo.ne.jp/door/crawler.html)"
			]
		  }
		,
		  {
			"pattern": "DuckDuckBot",
			"addition_date": "2012/09/19",
			"url": "http://duckduckgo.com/duckduckbot.html",
			"instances": [
			  "DuckDuckBot/1.0; (+http://duckduckgo.com/duckduckbot.html)",
			  "DuckDuckBot/1.1; (+http://duckduckgo.com/duckduckbot.html)",
			  "Mozilla/5.0 (compatible; DuckDuckBot-Https/1.1; https://duckduckgo.com/duckduckbot)",
			  "'Mozilla/5.0 (compatible; DuckDuckBot-Https/1.1; https://duckduckgo.com/duckduckbot)'"
			]
		  }
		,
		  {
			"pattern": "lssrocketcrawler",
			"addition_date": "2012/09/24",
			"instances": []
		  }
		,
		  {
			"pattern": "drupact",
			"addition_date": "2012/09/27",
			"url": "http://www.arocom.de/drupact",
			"instances": [
			  "drupact/0.7; http://www.arocom.de/drupact"
			]
		  }
		,
		  {
			"pattern": "webcompanycrawler",
			"addition_date": "2012/10/03",
			"instances": []
		  }
		,
		  {
			"pattern": "acoonbot",
			"addition_date": "2012/10/07",
			"url": "http://www.acoon.de/robot.asp",
			"instances": []
		  }
		,
		  {
			"pattern": "openindexspider",
			"addition_date": "2012/10/26",
			"url": "http://www.openindex.io/en/webmasters/spider.html",
			"instances": []
		  }
		,
		  {
			"pattern": "gnam gnam spider",
			"addition_date": "2012/10/31",
			"instances": []
		  }
		,
		  {
			"pattern": "web-archive-net.com.bot",
			"instances": []
		  }
		,
		  {
			"pattern": "backlinkcrawler",
			"addition_date": "2013/01/04",
			"instances": []
		  }
		,
		  {
			"pattern": "coccoc",
			"addition_date": "2013/01/04",
			"url": "http://help.coccoc.vn/",
			"instances": [
			  "Mozilla/5.0 (compatible; coccoc/1.0; +http://help.coccoc.com/)",
			  "Mozilla/5.0 (compatible; coccoc/1.0; +http://help.coccoc.com/searchengine)",
			  "Mozilla/5.0 (compatible; coccocbot-image/1.0; +http://help.coccoc.com/searchengine)",
			  "Mozilla/5.0 (compatible; coccocbot-web/1.0; +http://help.coccoc.com/searchengine)",
			  "Mozilla/5.0 (compatible; image.coccoc/1.0; +http://help.coccoc.com/)",
			  "Mozilla/5.0 (compatible; imagecoccoc/1.0; +http://help.coccoc.com/)",
			  "Mozilla/5.0 (compatible; imagecoccoc/1.0; +http://help.coccoc.com/searchengine)",
			  "coccoc",
			  "coccoc/1.0 ()",
			  "coccoc/1.0 (http://help.coccoc.com/)",
			  "coccoc/1.0 (http://help.coccoc.vn/)"
			]
		  }
		,
		  {
			"pattern": "integromedb",
			"addition_date": "2013/01/10",
			"url": "http://www.integromedb.org/Crawler",
			"instances": [
			  "www.integromedb.org/Crawler"
			]
		  }
		,
		  {
			"pattern": "content crawler spider",
			"addition_date": "2013/01/11",
			"instances": []
		  }
		,
		  {
			"pattern": "toplistbot",
			"addition_date": "2013/02/05",
			"instances": []
		  }
		,
		  {
			"pattern": "it2media-domain-crawler",
			"addition_date": "2013/03/12",
			"instances": [
			  "it2media-domain-crawler/1.0 on crawler-prod.it2media.de",
			  "it2media-domain-crawler/2.0"
			]
		  }
		,
		  {
			"pattern": "ip-web-crawler.com",
			"addition_date": "2013/03/22",
			"instances": []
		  }
		,
		  {
			"pattern": "siteexplorer.info",
			"addition_date": "2013/05/01",
			"instances": [
			  "Mozilla/5.0 (compatible; SiteExplorer/1.0b; +http://siteexplorer.info/)",
			  "Mozilla/5.0 (compatible; SiteExplorer/1.1b; +http://siteexplorer.info/Backlink-Checker-Spider/)"
			]
		  }
		,
		  {
			"pattern": "elisabot",
			"addition_date": "2013/06/27",
			"instances": []
		  }
		,
		  {
			"pattern": "proximic",
			"addition_date": "2013/09/12",
			"url": "http://www.proximic.com/info/spider.php",
			"instances": [
			  "Mozilla/5.0 (compatible; proximic; +http://www.proximic.com)",
			  "Mozilla/5.0 (compatible; proximic; +http://www.proximic.com/info/spider.php)"
			]
		  }
		,
		  {
			"pattern": "changedetection",
			"addition_date": "2013/09/13",
			"url": "http://www.changedetection.com/bot.html",
			"instances": [
			  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1;  http://www.changedetection.com/bot.html )"
			]
		  }
		,
		  {
			"pattern": "arabot",
			"addition_date": "2013/10/09",
			"instances": []
		  }
		,
		  {
			"pattern": "WeSEE:Search",
			"addition_date": "2013/11/18",
			"instances": [
			  "WeSEE:Search",
			  "WeSEE:Search/0.1 (Alpha, http://www.wesee.com/en/support/bot/)"
			]
		  }
		,
		  {
			"pattern": "niki-bot",
			"addition_date": "2014/01/01",
			"instances": []
		  }
		,
		  {
			"pattern": "CrystalSemanticsBot",
			"addition_date": "2014/02/17",
			"url": "http://www.crystalsemantics.com/user-agent/",
			"instances": []
		  }
		,
		  {
			"pattern": "rogerbot",
			"addition_date": "2014/02/28",
			"url": "http://moz.com/help/pro/what-is-rogerbot-",
			"instances": [
			  "Mozilla/5.0 (compatible; rogerBot/1.0; UrlCrawler; http://www.seomoz.org/dp/rogerbot)",
			  "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+partager@moz.com)",
			  "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+shiny@moz.com)",
			  "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-wherecat@moz.com",
			  "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-wherecat@moz.com)",
			  "rogerbot/1.0 (http://www.moz.com/dp/rogerbot, rogerbot-crawler@moz.com)",
			  "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler+shiny@seomoz.org)",
			  "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler@seomoz.org)",
			  "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-wherecat@moz.com)",
			  "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr2-crawler-05@moz.com)",
			  "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr4-crawler-11@moz.com)",
			  "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr4-crawler-15@moz.com)",
			  "rogerbot/1.2 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+phaser-testing-crawler-01@moz.com)"
			]
		  }
		,
		  {
			"pattern": "360Spider",
			"addition_date": "2014/03/14",
			"url": "http://needs-be.blogspot.co.uk/2013/02/how-to-block-spider360.html",
			"instances": [
			  "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1; 360Spider",
			  "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1; 360Spider(compatible; HaosouSpider; http://www.haosou.com/help/help_3_2.html)",
			  "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 QIHU 360SE; 360Spider",
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; )  Firefox/1.5.0.11; 360Spider",
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11)  Firefox/1.5.0.11; 360Spider",
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Firefox/1.5.0.11 360Spider;",
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Gecko/20070312 Firefox/1.5.0.11; 360Spider",
			  "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider",
			  "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider(compatible; HaosouSpider; http://www.haosou.com/help/help_3_2.html)",
			  "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36; 360Spider"
			]
		  }
		,
		  {
			"pattern": "psbot",
			"addition_date": "2014/03/31",
			"url": "http://www.picsearch.com/bot.html",
			"instances": [
			  "psbot-image (+http://www.picsearch.com/bot.html)",
			  "psbot-page (+http://www.picsearch.com/bot.html)",
			  "psbot/0.1 (+http://www.picsearch.com/bot.html)"
			]
		  }
		,
		  {
			"pattern": "InterfaxScanBot",
			"addition_date": "2014/03/31",
			"url": "http://scan-interfax.ru",
			"instances": []
		  }
		,
		  {
			"pattern": "CC Metadata Scaper",
			"addition_date": "2014/04/01",
			"url": "http://wiki.creativecommons.org/Metadata_Scraper",
			"instances": [
			  "CC Metadata Scaper http://wiki.creativecommons.org/Metadata_Scraper"
			]
		  }
		,
		  {
			"pattern": "g00g1e.net",
			"addition_date": "2014/04/01",
			"url": "http://www.g00g1e.net/",
			"instances": []
		  }
		,
		  {
			"pattern": "GrapeshotCrawler",
			"addition_date": "2014/04/01",
			"url": "http://www.grapeshot.co.uk/crawler.php",
			"instances": [
			  "Mozilla/5.0 (compatible; GrapeshotCrawler/2.0; +http://www.grapeshot.co.uk/crawler.php)"
			]
		  }
		,
		  {
			"pattern": "urlappendbot",
			"addition_date": "2014/05/10",
			"url": "http://www.profound.net/urlappendbot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; URLAppendBot/1.0; +http://www.profound.net/urlappendbot.html)"
			]
		  }
		,
		  {
			"pattern": "brainobot",
			"addition_date": "2014/06/24",
			"instances": []
		  }
		,
		  {
			"pattern": "fr-crawler",
			"addition_date": "2014/07/31",
			"instances": [
			  "Mozilla/5.0 (compatible; fr-crawler/1.1)"
			]
		  }
		,
		  {
			"pattern": "binlar",
			"addition_date": "2014/09/12",
			"instances": [
			  "binlar_2.6.3 binlar2.6.3@unspecified.mail",
			  "binlar_2.6.3 binlar_2.6.3@unspecified.mail",
			  "binlar_2.6.3 larbin2.6.3@unspecified.mail",
			  "binlar_2.6.3 phanendra_kalapala@McAfee.com",
			  "binlar_2.6.3 test@mgmt.mic"
			]
		  }
		,
		  {
			"pattern": "SimpleCrawler",
			"addition_date": "2014/09/12",
			"instances": [
			  "SimpleCrawler/0.1"
			]
		  }
		,
		  {
			"pattern": "Twitterbot",
			"addition_date": "2014/09/12",
			"url": "https://dev.twitter.com/cards/getting-started",
			"instances": [
			  "Twitterbot/0.1",
			  "Twitterbot/1.0"
			]
		  }
		,
		  {
			"pattern": "cXensebot",
			"addition_date": "2014/10/05",
			"instances": [
			  "cXensebot/1.1a"
			],
			"url": "http://www.cxense.com/bot.html"
		  }
		,
		  {
			"pattern": "smtbot",
			"addition_date": "2014/10/04",
			"instances": [
			  "Mozilla/5.0 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)",
			  "SMTBot (similartech.com/smtbot)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko)                 Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)",
			  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)",
			  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36 (compatible; SMTBot/1.0; http://www.similartech.com/smtbot)"
			],
			"url": "http://www.similartech.com/smtbot"
		  }
		,
		  {
			"pattern": "bnf.fr_bot",
			"addition_date": "2014/11/18",
			"url": "http://www.bnf.fr/fr/outils/a.dl_web_capture_robot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; bnf.fr_bot; +http://bibnum.bnf.fr/robot/bnf.html)",
			  "Mozilla/5.0 (compatible; bnf.fr_bot; +http://www.bnf.fr/fr/outils/a.dl_web_capture_robot.html)"
			]
		  }
		,
		  {
			"pattern": "A6-Indexer",
			"addition_date": "2014/12/05",
			"url": "http://www.a6corp.com/a6-web-scraping-policy/",
			"instances": [
			  "A6-Indexer"
			]
		  }
		,
		  {
			"pattern": "ADmantX",
			"addition_date": "2014/12/05",
			"url": "http://www.admantx.com",
			"instances": [
			  "ADmantX Platform Semantic Analyzer - ADmantX Inc. - www.admantx.com - support@admantx.com"
			]
		  }
		,
		  {
			"pattern": "Facebot",
			"url": "https://developers.facebook.com/docs/sharing/best-practices#crawl",
			"addition_date": "2014/12/30",
			"instances": [
			  "Facebot/1.0"
			]
		  }
		,
		  {
			"pattern": "OrangeBot\\/",
			"instances": [
			  "Mozilla/5.0 (compatible; OrangeBot/2.0; support.orangebot@orange.com"
			],
			"addition_date": "2015/01/12"
		  }
		,
		  {
			"pattern": "memorybot",
			"url": "http://mignify.com/bot.htm",
			"instances": [
			  "Mozilla/5.0 (compatible; memorybot/1.21.14 +http://mignify.com/bot.html)"
			],
			"addition_date": "2015/02/01"
		  }
		,
		  {
			"pattern": "AdvBot",
			"url": "http://advbot.net/bot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; AdvBot/2.0; +http://advbot.net/bot.html)"
			],
			"addition_date": "2015/02/01"
		  }
		,
		  {
			"pattern": "MegaIndex",
			"url": "https://www.megaindex.ru/?tab=linkAnalyze",
			"instances": [
			  "Mozilla/5.0 (compatible; MegaIndex.ru/2.0; +https://www.megaindex.ru/?tab=linkAnalyze)",
			  "Mozilla/5.0 (compatible; MegaIndex.ru/2.0; +http://megaindex.com/crawler)"
			],
			"addition_date": "2015/03/28"
		  }
		,
		  {
			"pattern": "SemanticScholarBot",
			"url": "https://www.semanticscholar.org/crawler",
			"instances": [
			  "SemanticScholarBot/1.0 (+http://s2.allenai.org/bot.html)",
			  "Mozilla/5.0 (compatible) SemanticScholarBot (+https://www.semanticscholar.org/crawler)"
			],
			"addition_date": "2015/03/28"
		  }
		,
		  {
			"pattern": "ltx71",
			"url": "http://ltx71.com/",
			"instances": [
			  "ltx71 - (http://ltx71.com/)"
			],
			"addition_date": "2015/04/04"
		  }
		,
		  {
			"pattern": "nerdybot",
			"url": "http://nerdybot.com/",
			"instances": [
			  "nerdybot"
			],
			"addition_date": "2015/04/05"
		  }
		,
		  {
			"pattern": "xovibot",
			"url": "http://www.xovibot.net/",
			"instances": [
			  "Mozilla/5.0 (compatible; XoviBot/2.0; +http://www.xovibot.net/)"
			],
			"addition_date": "2015/04/05"
		  }
		,
		  {
			"pattern": "BUbiNG",
			"url": "http://law.di.unimi.it/BUbiNG.html",
			"instances": [
			  "BUbiNG (+http://law.di.unimi.it/BUbiNG.html)"
			],
			"addition_date": "2015/04/06"
		  }
		,
		  {
			"pattern": "Qwantify",
			"url": "https://www.qwant.com/",
			"instances": [
			  "Mozilla/5.0 (compatible; Qwantify/2.0n; +https://www.qwant.com/)/*",
			  "Mozilla/5.0 (compatible; Qwantify/2.4w; +https://www.qwant.com/)/2.4w",
			  "Mozilla/5.0 (compatible; Qwantify/Bleriot/1.1; +https://help.qwant.com/bot)",
			  "Mozilla/5.0 (compatible; Qwantify/Bleriot/1.2.1; +https://help.qwant.com/bot)"
			],
			"addition_date": "2015/04/06"
		  }
		,
		  {
			"pattern": "archive.org_bot",
			"url": "http://www.archive.org/details/archive.org_bot",
			"depends_on": ["heritrix"],
			"instances": [
			  "Mozilla/5.0 (compatible; heritrix/3.1.1-SNAPSHOT-20120116.200628 +http://www.archive.org/details/archive.org_bot)",
			  "Mozilla/5.0 (compatible; archive.org_bot/heritrix-1.15.4 +http://www.archive.org)",
			  "Mozilla/5.0 (compatible; heritrix/3.3.0-SNAPSHOT-20140702-2247 +http://archive.org/details/archive.org_bot)",
			  "Mozilla/5.0 (compatible; archive.org_bot +http://www.archive.org/details/archive.org_bot)",
			  "Mozilla/5.0 (compatible; archive.org_bot +http://archive.org/details/archive.org_bot)",
			  "Mozilla/5.0 (compatible; special_archiver/3.1.1 +http://www.archive.org/details/archive.org_bot)"
			],
			"addition_date": "2015/04/14"
		  }
		,
		  {
			"pattern": "Applebot",
			"url": "http://www.apple.com/go/applebot",
			"addition_date": "2015/04/15",
			"instances": [
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1)",
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)",
			  "Mozilla/5.0 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B410 Safari/600.1.4 (Applebot/0.1; +http://www.apple.com/go/applebot)"
			]
		  }
		,
		  {
			"pattern": "TweetmemeBot",
			"url": "http://datasift.com/bot.html",
			"instances": [
			  "Mozilla/5.0 (TweetmemeBot/4.0; +http://datasift.com/bot.html) Gecko/20100101 Firefox/31.0"
			],
			"addition_date": "2015/04/15"
		  }
		,
		  {
			"pattern": "crawler4j",
			"url": "https://github.com/yasserg/crawler4j",
			"instances": [
			  "crawler4j (http://code.google.com/p/crawler4j/)",
			  "crawler4j (https://github.com/yasserg/crawler4j/)"
			],
			"addition_date": "2015/05/07"
		  }
		,
		  {
			"pattern": "findxbot",
			"url": "http://www.findxbot.com",
			"instances": [
			  "Mozilla/5.0 (compatible; Findxbot/1.0; +http://www.findxbot.com)"
			],
			"addition_date": "2015/05/07"
		  }
		,
		  {
			"pattern": "S[eE][mM]rushBot",
			"url": "http://www.semrush.com/bot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; SemrushBot-SA/0.97; +http://www.semrush.com/bot.html)",
			  "Mozilla/5.0 (compatible; SemrushBot-SI/0.97; +http://www.semrush.com/bot.html)",
			  "Mozilla/5.0 (compatible; SemrushBot/3~bl; +http://www.semrush.com/bot.html)",
			  "Mozilla/5.0 (compatible; SemrushBot/0.98~bl; +http://www.semrush.com/bot.html)",
			  "Mozilla/5.0 (compatible; SemrushBot-BA; +http://www.semrush.com/bot.html)",
			  "Mozilla/5.0 (compatible; SemrushBot/6~bl; +http://www.semrush.com/bot.html)",
			  "SEMrushBot"
			],
			"addition_date": "2015/05/26"
		  }
		,
		  {
			"pattern": "yoozBot",
			"url": "http://yooz.ir",
			"instances": [
			  "Mozilla/5.0 (compatible; yoozBot-2.2; http://yooz.ir; info@yooz.ir)"
			],
			"addition_date": "2015/05/26"
		  }
		,
		  {
			"pattern": "lipperhey",
			"url": "http://www.lipperhey.com/",
			"instances": [
			  "Mozilla/5.0 (compatible; Lipperhey Link Explorer; http://www.lipperhey.com/)",
			  "Mozilla/5.0 (compatible; Lipperhey SEO Service; http://www.lipperhey.com/)",
			  "Mozilla/5.0 (compatible; Lipperhey Site Explorer; http://www.lipperhey.com/)",
			  "Mozilla/5.0 (compatible; Lipperhey-Kaus-Australis/5.0; +https://www.lipperhey.com/en/about/)"
			],
			"addition_date": "2015/08/26"
		  }
		,
		  {
			"pattern": "Y!J",
			"url": "https://www.yahoo-help.jp/app/answers/detail/p/595/a_id/42716/~/%E3%82%A6%E3%82%A7%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%AB%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E3%81%99%E3%82%8B%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AE%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E3%82%A8%E3%83%BC%E3%82%B8%E3%82%A7%E3%83%B3%E3%83%88%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6",
			"instances": [
			  "Y!J-ASR/0.1 crawler (http://www.yahoo-help.jp/app/answers/detail/p/595/a_id/42716/)",
			  "Y!J-BRJ/YATS crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)",
			  "Y!J-PSC/1.0 crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)",
			  "Y!J-BRW/1.0 crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)",
			  "Mozilla/5.0 (iPhone; Y!J-BRY/YATSH crawler; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)",
			  "Mozilla/5.0 (compatible; Y!J SearchMonkey/1.0 (Y!J-AGENT; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html))"
			],
			"addition_date": "2015/05/26"
		  }
		,
		  {
			"pattern": "Domain Re-Animator Bot",
			"url": "http://domainreanimator.com",
			"instances": [
			  "Domain Re-Animator Bot (http://domainreanimator.com) - support@domainreanimator.com"
			],
			"addition_date": "2015/04/14"
		  }
		,
		  {
			"pattern": "AddThis",
			"url": "https://www.addthis.com",
			"instances": [
			  "AddThis.com robot tech.support@clearspring.com"
			],
			"addition_date": "2015/06/02"
		  }
		,
		  {
			"pattern": "Screaming Frog SEO Spider",
			"url": "http://www.screamingfrog.co.uk/seo-spider",
			"instances": [
			  "Screaming Frog SEO Spider/5.1"
			],
			"addition_date": "2016/01/08"
		  }
		,
		  {
			"pattern": "MetaURI",
			"url": "http://www.useragentstring.com/MetaURI_id_17683.php",
			"instances": [
			  "MetaURI API/2.0 +metauri.com"
			],
			"addition_date": "2016/01/02"
		  }
		,
		  {
			"pattern": "Scrapy",
			"url": "http://scrapy.org/",
			"instances": [
			  "Scrapy/1.0.3 (+http://scrapy.org)"
			],
			"addition_date": "2016/01/02"
		  }
		,
		  {
			"pattern": "Livelap[bB]ot",
			"url": "http://site.livelap.com/crawler",
			"instances": [
			  "LivelapBot/0.2 (http://site.livelap.com/crawler)",
			  "Livelapbot/0.1"
			],
			"addition_date": "2016/01/02"
		  }
		,
		  {
			"pattern": "OpenHoseBot",
			"url": "http://www.openhose.org/bot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; OpenHoseBot/2.1; +http://www.openhose.org/bot.html)"
			],
			"addition_date": "2016/01/02"
		  }
		,
		  {
			"pattern": "CapsuleChecker",
			"url": "http://www.capsulink.com/about",
			"instances": [
			  "CapsuleChecker (http://www.capsulink.com/)"
			],
			"addition_date": "2016/01/02"
		  }
		,
		  {
			"pattern": "collection@infegy.com",
			"url": "http://infegy.com/",
			"instances": [
			  "Mozilla/5.0 (compatible) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36 collection@infegy.com"
			],
			"addition_date": "2016/01/03"
		  }
		,
		  {
			"pattern": "IstellaBot",
			"url": "http://www.tiscali.it/",
			"instances": [
			  "Mozilla/5.0 (compatible; IstellaBot/1.23.15 +http://www.tiscali.it/)"
			],
			"addition_date": "2016/01/09"
		  }
		,
		  {
			"pattern": "DeuSu\\/",
			"addition_date": "2016/01/23",
			"url": "https://deusu.de/robot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; DeuSu/0.1.0; +https://deusu.org)",
			  "Mozilla/5.0 (compatible; DeuSu/5.0.2; +https://deusu.de/robot.html)"
			]
		  }
		,
		  {
			"pattern": "betaBot",
			"addition_date": "2016/01/23",
			"instances": []
		  }
		,
		  {
			"pattern": "Cliqzbot\\/",
			"addition_date": "2016/01/23",
			"url": "http://cliqz.com/company/cliqzbot",
			"instances": [
			  "Mozilla/5.0 (compatible; Cliqzbot/2.0; +http://cliqz.com/company/cliqzbot)",
			  "Cliqzbot/0.1 (+http://cliqz.com +cliqzbot@cliqz.com)",
			  "Cliqzbot/0.1 (+http://cliqz.com/company/cliqzbot)",
			  "Mozilla/5.0 (compatible; Cliqzbot/0.1 +http://cliqz.com/company/cliqzbot)",
			  "Mozilla/5.0 (compatible; Cliqzbot/1.0 +http://cliqz.com/company/cliqzbot)"
			]
		  }
		,
		  {
			"pattern": "MojeekBot\\/",
			"addition_date": "2016/01/23",
			"url": "https://www.mojeek.com/bot.html",
			"instances": [
			  "MojeekBot/0.2 (archi; http://www.mojeek.com/bot.html)",
			  "Mozilla/5.0 (compatible; MojeekBot/0.2; http://www.mojeek.com/bot.html#relaunch)",
			  "Mozilla/5.0 (compatible; MojeekBot/0.2; http://www.mojeek.com/bot.html)",
			  "Mozilla/5.0 (compatible; MojeekBot/0.5; http://www.mojeek.com/bot.html)",
			  "Mozilla/5.0 (compatible; MojeekBot/0.6; +https://www.mojeek.com/bot.html)",
			  "Mozilla/5.0 (compatible; MojeekBot/0.6; http://www.mojeek.com/bot.html)"
			]
		  }
		,
		  {
			"pattern": "netEstate NE Crawler",
			"addition_date": "2016/01/23",
			"url": "+http://www.website-datenbank.de/",
			"instances": [
			  "netEstate NE Crawler (+http://www.sengine.info/)",
			  "netEstate NE Crawler (+http://www.website-datenbank.de/)"
			]
		  }
		,
		  {
			"pattern": "SafeSearch microdata crawler",
			"addition_date": "2016/01/23",
			"url": "https://safesearch.avira.com",
			"instances": [
			  "SafeSearch microdata crawler (https://safesearch.avira.com, safesearch-abuse@avira.com)"
			]
		  }
		,
		  {
			"pattern": "Gluten Free Crawler\\/",
			"addition_date": "2016/01/23",
			"url": "http://glutenfreepleasure.com/",
			"instances": [
			  "Mozilla/5.0 (compatible; Gluten Free Crawler/1.0; +http://glutenfreepleasure.com/)"
			]
		  }
		,
		  {
			"pattern": "Sonic",
			"addition_date": "2016/02/08",
			"url": "http://www.yama.info.waseda.ac.jp/~crawler/info.html",
			"instances": [
			  "Mozilla/5.0 (compatible; RankSonicSiteAuditor/1.0; +https://ranksonic.com/ranksonic_sab.html)",
			  "Mozilla/5.0 (compatible; Sonic/1.0; http://www.yama.info.waseda.ac.jp/~crawler/info.html)",
			  "Mozzila/5.0 (compatible; Sonic/1.0; http://www.yama.info.waseda.ac.jp/~crawler/info.html)"
			]
		  }
		,
		  {
			"pattern": "Sysomos",
			"addition_date": "2016/02/08",
			"url": "http://www.sysomos.com",
			"instances": [
			  "Mozilla/5.0 (compatible; Sysomos/1.0; +http://www.sysomos.com/; Sysomos)"
			]
		  }
		,
		  {
			"pattern": "Trove",
			"addition_date": "2016/02/08",
			"url": "http://www.trove.com",
			"instances": []
		  }
		,
		  {
			"pattern": "deadlinkchecker",
			"addition_date": "2016/02/08",
			"url": "http://www.deadlinkchecker.com",
			"instances": [
			  "www.deadlinkchecker.com Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36",
			  "www.deadlinkchecker.com XMLHTTP/1.0",
			  "www.deadlinkchecker.com XMLHTTP/1.0 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36"
			]
		  }
		,
		  {
			"pattern": "Slack-ImgProxy",
			"addition_date": "2016/04/25",
			"url": "https://api.slack.com/robots",
			"instances": [
			  "Slack-ImgProxy (+https://api.slack.com/robots)",
			  "Slack-ImgProxy 0.59 (+https://api.slack.com/robots)",
			  "Slack-ImgProxy 0.66 (+https://api.slack.com/robots)",
			  "Slack-ImgProxy 1.106 (+https://api.slack.com/robots)",
			  "Slack-ImgProxy 1.138 (+https://api.slack.com/robots)",
			  "Slack-ImgProxy 149 (+https://api.slack.com/robots)"
			]
		  }
		,
		  {
			"pattern": "Embedly",
			"addition_date": "2016/04/25",
			"url": "http://support.embed.ly",
			"instances": [
			  "Embedly +support@embed.ly",
			  "Mozilla/5.0 (compatible; Embedly/0.2; +http://support.embed.ly/)",
			  "Mozilla/5.0 (compatible; Embedly/0.2; snap; +http://support.embed.ly/)"
			]
		  }
		,
		  {
			"pattern": "RankActiveLinkBot",
			"addition_date": "2016/06/20",
			"url": "https://rankactive.com/resources/rankactive-linkbot",
			"instances": [
			  "Mozilla/5.0 (compatible; RankActiveLinkBot; +https://rankactive.com/resources/rankactive-linkbot)"
			]
		  }
		,
		  {
			"pattern": "iskanie",
			"addition_date": "2016/09/02",
			"url": "http://www.iskanie.com",
			"instances": [
			  "iskanie (+http://www.iskanie.com)"
			]
		  }
		,
		  {
			"pattern": "SafeDNSBot",
			"addition_date": "2016/09/10",
			"url": "https://www.safedns.com/searchbot",
			"instances": [
			  "SafeDNSBot (https://www.safedns.com/searchbot)"
			]
		  }
		,
		  {
			"pattern": "SkypeUriPreview",
			"addition_date": "2016/10/10",
			"instances": [
			  "Mozilla/5.0 (Windows NT 6.1; WOW64) SkypeUriPreview Preview/0.5"
			]
		  }
		,
		  {
			"pattern": "Veoozbot",
			"addition_date": "2016/11/03",
			"url": "http://www.veooz.com/veoozbot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; Veoozbot/1.0; +http://www.veooz.com/veoozbot.html)"
			]
		  }
		,
		  {
			"pattern": "Slackbot",
			"addition_date": "2016/11/03",
			"url": "https://api.slack.com/robots",
			"instances": [
			  "Slackbot-LinkExpanding (+https://api.slack.com/robots)",
			  "Slackbot-LinkExpanding 1.0 (+https://api.slack.com/robots)",
			  "Slackbot 1.0 (+https://api.slack.com/robots)"
			]
		  }
		,
		  {
			"pattern": "redditbot",
			"addition_date": "2016/11/03",
			"url": "http://www.reddit.com/feedback",
			"instances": [
			  "Mozilla/5.0 (compatible; redditbot/1.0; +http://www.reddit.com/feedback)"
			]
		  }
		,
		  {
			"pattern": "datagnionbot",
			"addition_date": "2016/11/03",
			"url": "http://www.datagnion.com/bot.html",
			"instances": [
			  "datagnionbot (+http://www.datagnion.com/bot.html)"
			]
		  }
		,
		  {
			"pattern": "Google-Adwords-Instant",
			"addition_date": "2016/11/03",
			"url": "http://www.google.com/adsbot.html",
			"instances": [
			  "Google-Adwords-Instant (+http://www.google.com/adsbot.html)"
			]
		  }
		,
		  {
			"pattern": "adbeat_bot",
			"addition_date": "2016/11/04",
			"instances": [
			  "Mozilla/5.0 (compatible; adbeat_bot; +support@adbeat.com; support@adbeat.com)",
			  "adbeat_bot"
			]
		  }
		,
		  {
			"pattern": "WhatsApp",
			"addition_date": "2016/11/15",
			"url": "https://www.whatsapp.com/",
			"instances": [
			  "WhatsApp",
			  "WhatsApp/0.3.4479 N",
			  "WhatsApp/0.3.4679 N",
			  "WhatsApp/0.3.4941 N",
			  "WhatsApp/2.12.15/i",
			  "WhatsApp/2.12.16/i",
			  "WhatsApp/2.12.17/i",
			  "WhatsApp/2.12.449 A",
			  "WhatsApp/2.12.453 A",
			  "WhatsApp/2.12.510 A",
			  "WhatsApp/2.12.540 A",
			  "WhatsApp/2.12.548 A",
			  "WhatsApp/2.12.555 A",
			  "WhatsApp/2.12.556 A",
			  "WhatsApp/2.16.1/i",
			  "WhatsApp/2.16.13 A",
			  "WhatsApp/2.16.2/i",
			  "WhatsApp/2.16.42 A",
			  "WhatsApp/2.16.57 A",
			  "WhatsApp/2.19.92 i",
			  "WhatsApp/2.19.175 A",
			  "WhatsApp/2.19.244 A",
			  "WhatsApp/2.19.258 A",
			  "WhatsApp/2.19.308 A",
			  "WhatsApp/2.19.330 A"
			]
		  }
		,
		  {
			"pattern": "contxbot",
			"addition_date": "2017/02/25",
			"instances": [
			  "Mozilla/5.0 (compatible;contxbot/1.0)"
			]
		  }
		,
		  {
			"pattern": "pinterest.com.bot",
			"addition_date": "2017/03/03",
			"instances": [
			  "Mozilla/5.0 (compatible; Pinterestbot/1.0; +http://www.pinterest.com/bot.html)",
			  "Pinterest/0.2 (+http://www.pinterest.com/bot.html)"
			],
			"url": "http://www.pinterest.com/bot.html"
		  }
		,
		  {
			"pattern": "electricmonk",
			"addition_date": "2017/03/04",
			"instances": [
			  "Mozilla/5.0 (compatible; electricmonk/3.2.0 +https://www.duedil.com/our-crawler/)"
			],
			"url": "https://www.duedil.com/our-crawler/"
		  }
		,
		  {
			"pattern": "GarlikCrawler",
			"addition_date": "2017/03/18",
			"instances": [
			  "GarlikCrawler/1.2 (http://garlik.com/, crawler@garlik.com)"
			],
			"url": "http://garlik.com/"
		  }
		,
		  {
			"pattern": "BingPreview\\/",
			"addition_date": "2017/04/23",
			"url": "https://www.bing.com/webmaster/help/which-crawlers-does-bing-use-8c184ec0",
			"instances": [
			  "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko) BingPreview/1.0b",
			  "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0; BingPreview/1.0b) like Gecko",
			  "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0;  WOW64;  Trident/6.0;  BingPreview/1.0b)",
			  "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0;  WOW64;  Trident/5.0;  BingPreview/1.0b)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 BingPreview/1.0b"
			]
		  }
		,
		  {
			"pattern": "vebidoobot",
			"addition_date": "2017/05/08",
			"instances": [
			  "Mozilla/5.0 (compatible; vebidoobot/1.0; +https://blog.vebidoo.de/vebidoobot/"
			],
			"url": "https://blog.vebidoo.de/vebidoobot/"
		  }
		,
		  {
			"pattern": "FemtosearchBot",
			"addition_date": "2017/05/16",
			"instances": [
			  "Mozilla/5.0 (compatible; FemtosearchBot/1.0; http://femtosearch.com)"
			],
			"url": "http://femtosearch.com"
		  }
		,
		  {
			"pattern": "Yahoo Link Preview",
			"addition_date": "2017/06/28",
			"instances": [
			  "Mozilla/5.0 (compatible; Yahoo Link Preview; https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html)"
			],
			"url": "https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html"
		  }
		,
		  {
			"pattern": "MetaJobBot",
			"addition_date": "2017/08/16",
			"instances": [
			  "Mozilla/5.0 (compatible; MetaJobBot; http://www.metajob.de/crawler)"
			],
			"url": "http://www.metajob.de/the/crawler"
		  }
		,
		  {
			"pattern": "DomainStatsBot",
			"addition_date": "2017/08/16",
			"instances": [
			  "DomainStatsBot/1.0 (http://domainstats.io/our-bot)"
			],
			"url": "http://domainstats.io/our-bot"
		  }
		,
		  {
			"pattern": "mindUpBot",
			"addition_date": "2017/08/16",
			"instances": [
			  "mindUpBot (datenbutler.de)"
			],
			"url": "http://www.datenbutler.de/"
		  }
		,
		  {
			"pattern": "Daum\\/",
			"addition_date": "2017/08/16",
			"instances": [
			  "Mozilla/5.0 (compatible; Daum/4.1; +http://cs.daum.net/faq/15/4118.html?faqId=28966)"
			],
			"url": "http://cs.daum.net/faq/15/4118.html?faqId=28966"
		  }
		,
		  {
			"pattern": "Jugendschutzprogramm-Crawler",
			"addition_date": "2017/08/16",
			"instances": [
			  "Jugendschutzprogramm-Crawler; Info: http://www.jugendschutzprogramm.de"
			],
			"url": "http://www.jugendschutzprogramm.de"
		  }
		,
		  {
			"pattern": "Xenu Link Sleuth",
			"addition_date": "2017/08/19",
			"instances": [
			  "Xenu Link Sleuth/1.3.8"
			],
			"url": "http://home.snafu.de/tilman/xenulink.html"
		  }
		,
		  {
			"pattern": "Pcore-HTTP",
			"addition_date": "2017/08/19",
			"instances": [
			  "Pcore-HTTP/v0.40.3",
			  "Pcore-HTTP/v0.44.0"
			],
			"url": "https://bitbucket.org/softvisio/pcore/overview"
		  }
		,
		  {
			"pattern": "moatbot",
			"addition_date": "2017/09/16",
			"instances": [
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.111 Safari/537.36 moatbot",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4 moatbot"
			],
			"url": "https://moat.com"
		  }
		,
		  {
			"pattern": "KosmioBot",
			"addition_date": "2017/09/16",
			"instances": [
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.125 Safari/537.36 (compatible; KosmioBot/1.0; +http://kosm.io/bot.html)"
			],
			"url": "http://kosm.io/bot.html"
		  }
		,
		  {
			"pattern": "pingdom",
			"addition_date": "2017/09/16",
			"instances": [
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/59.0.3071.109 Chrome/59.0.3071.109 Safari/537.36 PingdomPageSpeed/1.0 (pingbot/2.0; +http://www.pingdom.com/)",
			  "Mozilla/5.0 (compatible; pingbot/2.0; +http://www.pingdom.com/)"
			],
			"url": "http://www.pingdom.com"
		  }
		,
		  {
			"pattern": "AppInsights",
			"addition_date": "2019/03/09",
			"instances": [
			  "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; AppInsights)"
			],
			"url": "https://docs.microsoft.com/en-us/azure/azure-monitor/app/app-insights-overview"
		  }
		,
		  {
			"pattern": "PhantomJS",
			"addition_date": "2017/09/18",
			"instances": [
			  "Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) PhantomJS/2.1.1 Safari/538.1 bl.uk_lddc_renderbot/2.0.0 (+ http://www.bl.uk/aboutus/legaldeposit/websites/websites/faqswebmaster/index.html)"
			],
			"url": "http://phantomjs.org/"
		  }
		,
		  {
			"pattern": "Gowikibot",
			"addition_date": "2017/10/26",
			"instances": [
			  "Mozilla/5.0 (compatible; Gowikibot/1.0; +http://www.gowikibot.com)"
			],
			"url": "http://www.gowikibot.com"
		  }
		,
		  {
			"pattern": "PiplBot",
			"addition_date": "2017/10/30",
			"instances": [
			  "PiplBot (+http://www.pipl.com/bot/)",
			  "Mozilla/5.0+(compatible;+PiplBot;+http://www.pipl.com/bot/)"
			],
			"url": "http://www.pipl.com/bot/"
		  }
		,
		  {
			"pattern": "Discordbot",
			"addition_date": "2017/09/22",
			"url": "https://discordapp.com",
			"instances": [
			  "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)"
			]
		  }
		,
		  {
			"pattern": "TelegramBot",
			"addition_date": "2017/10/01",
			"instances": [
			  "TelegramBot (like TwitterBot)"
			]
		  }
		,
		  {
			"pattern": "Jetslide",
			"addition_date": "2017/09/27",
			"url": "http://jetsli.de/crawler",
			"instances": [
			  "Mozilla/5.0 (compatible; Jetslide; +http://jetsli.de/crawler)"
			]
		  }
		,
		  {
			"pattern": "newsharecounts",
			"addition_date": "2017/09/30",
			"url": "http://newsharecounts.com/crawler",
			"instances": [
			  "Mozilla/5.0 (compatible; NewShareCounts.com/1.0; +http://newsharecounts.com/crawler)"
			]
		  }
		,
		  {
			"pattern": "James BOT",
			"addition_date": "2017/10/12",
			"url": "http://cognitiveseo.com/bot.html",
			"instances": [
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.6) Gecko/20070725 Firefox/2.0.0.6 - James BOT - WebCrawler http://cognitiveseo.com/bot.html"
			]
		  }
		,
		  {
			"pattern": "Bark[rR]owler",
			"addition_date": "2017/10/09",
			"url": "http://www.exensa.com/crawl",
			"instances": [
			  "Barkrowler/0.5.1 (experimenting / debugging - sorry for your logs ) http://www.exensa.com/crawl - admin@exensa.com -- based on BuBiNG",
			  "Barkrowler/0.7 (+http://www.exensa.com/crawl)",
			  "BarkRowler/0.7 (+http://www.exensa.com/crawling)",
			  "Barkrowler/0.9 (+http://www.exensa.com/crawl)"
			]
		  }
		,
		  {
			"pattern": "TinEye",
			"addition_date": "2017/10/14",
			"url": "http://www.tineye.com/crawler.html",
			"instances": [
			  "Mozilla/5.0 (compatible; TinEye-bot/1.31; +http://www.tineye.com/crawler.html)",
			  "TinEye/1.1 (http://tineye.com/crawler.html)"
			]
		  }
		,
		  {
			"pattern": "SocialRankIOBot",
			"addition_date": "2017/10/19",
			"url": "http://socialrank.io/about",
			"instances": [
			  "SocialRankIOBot; http://socialrank.io/about"
			]
		  }
		,
		  {
			"pattern": "trendictionbot",
			"addition_date": "2017/10/30",
			"url": "http://www.trendiction.de/bot",
			"instances": [
			  "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-GB; rv:1.0; trendictionbot0.5.0; trendiction search; http://www.trendiction.de/bot; please let us know of any problems; web at trendiction.com) Gecko/20071127 Firefox/3.0.0.11",
			  "Mozilla/5.0 (Windows NT 10.0; Win64; x64; trendictionbot0.5.0; trendiction search; http://www.trendiction.de/bot; please let us know of any problems; web at trendiction.com) Gecko/20170101 Firefox/67.0"
			]
		  }
		,
		  {
			"pattern": "Ocarinabot",
			"addition_date": "2017/09/27",
			"instances": [
			  "Ocarinabot"
			]
		  }
		,
		  {
			"pattern": "epicbot",
			"addition_date": "2017/10/31",
			"url": "http://www.epictions.com/epicbot",
			"instances": [
			  "Mozilla/5.0 (compatible; epicbot; +http://www.epictions.com/epicbot)"
			]
		  }
		,
		  {
			"pattern": "Primalbot",
			"addition_date": "2017/09/27",
			"url": "https://www.primal.com",
			"instances": [
			  "Mozilla/5.0 (compatible; Primalbot; +https://www.primal.com;)"
			]
		  }
		,
		  {
			"pattern": "DuckDuckGo-Favicons-Bot",
			"addition_date": "2017/10/06",
			"url": "http://duckduckgo.com",
			"instances": [
			  "Mozilla/5.0 (compatible; DuckDuckGo-Favicons-Bot/1.0; +http://duckduckgo.com)"
			]
		  }
		,
		  {
			"pattern": "GnowitNewsbot",
			"addition_date": "2017/10/30",
			"url": "http://www.gnowit.com",
			"instances": [
			  "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0 / GnowitNewsbot / Contact information at http://www.gnowit.com"
			]
		  }
		,
		  {
			"pattern": "Leikibot",
			"addition_date": "2017/09/24",
			"url": "http://www.leiki.com",
			"instances": [
			  "Mozilla/5.0 (Windows NT 6.3;compatible; Leikibot/1.0; +http://www.leiki.com)"
			]
		  }
		,
		  {
			"pattern": "LinkArchiver",
			"addition_date": "2017/09/24",
			"instances": [
			  "@LinkArchiver twitter bot"
			]
		  }
		,
		  {
			"pattern": "YaK\\/",
			"addition_date": "2017/09/25",
			"url": "http://linkfluence.com",
			"instances": [
			  "Mozilla/5.0 (compatible; YaK/1.0; http://linkfluence.com/; bot@linkfluence.com)"
			]
		  }
		,
		  {
			"pattern": "PaperLiBot",
			"addition_date": "2017/09/25",
			"url": "http://support.paper.li/entries/20023257-what-is-paper-li",
			"instances": [
			  "Mozilla/5.0 (compatible; PaperLiBot/2.1; http://support.paper.li/entries/20023257-what-is-paper-li)",
			  "Mozilla/5.0 (compatible; PaperLiBot/2.1; https://support.paper.li/entries/20023257-what-is-paper-li)"
		
			]
		  }
		,
		  {
			"pattern": "Digg Deeper",
			"addition_date": "2017/09/26",
			"url": "http://digg.com/about",
			"instances": [
			  "Digg Deeper/v1 (http://digg.com/about)"
			]
		  }
		,
		  {
			"pattern": "dcrawl",
			"addition_date": "2017/09/22",
			"instances": [
			  "dcrawl/1.0"
			]
		  }
		,
		  {
			"pattern": "Snacktory",
			"addition_date": "2017/09/23",
			"url": "https://github.com/karussell/snacktory",
			"instances": [
			  "Mozilla/5.0 (compatible; Snacktory; +https://github.com/karussell/snacktory)"
			]
		  }
		,
		  {
			"pattern": "AndersPinkBot",
			"addition_date": "2017/09/24",
			"url": "http://anderspink.com/bot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; AndersPinkBot/1.0; +http://anderspink.com/bot.html)"
			]
		  }
		,
		  {
			"pattern": "Fyrebot",
			"addition_date": "2017/09/22",
			"instances": [
			  "Fyrebot/1.0"
			]
		  }
		,
		  {
			"pattern": "EveryoneSocialBot",
			"addition_date": "2017/09/22",
			"url": "http://everyonesocial.com",
			"instances": [
			  "Mozilla/5.0 (compatible; EveryoneSocialBot/1.0; support@everyonesocial.com http://everyonesocial.com/)"
			]
		  }
		,
		  {
			"pattern": "Mediatoolkitbot",
			"addition_date": "2017/10/06",
			"url": "http://mediatoolkit.com",
			"instances": [
			  "Mediatoolkitbot (complaints@mediatoolkit.com)"
			]
		  }
		,
		  {
			"pattern": "Luminator-robots",
			"addition_date": "2017/09/22",
			"instances": [
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.13 (KHTML, like Gecko) Chrome/30.0.1599.66 Safari/537.13 Luminator-robots/2.0"
			]
		  }
		,
		  {
			"pattern": "ExtLinksBot",
			"addition_date": "2017/11/02",
			"url": "https://extlinks.com/Bot.html",
			"instances": [
			  "Mozilla/5.0 (compatible; ExtLinksBot/1.5 +https://extlinks.com/Bot.html)"
			]
		  }
		,
		  {
			"pattern": "SurveyBot",
			"addition_date": "2017/11/02",
			"instances": [
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.9.0.13) Gecko/2009073022 Firefox/3.5.2 (.NET CLR 3.5.30729) SurveyBot/2.3 (DomainTools)"
			]
		  }
		,
		  {
			"pattern": "NING\\/",
			"addition_date": "2017/11/02",
			"instances": [
			  "NING/1.0"
			]
		  }
		,
		  {
			"pattern": "okhttp",
			"addition_date": "2017/11/02",
			"instances": [
			  "okhttp/2.5.0",
			  "okhttp/2.7.5",
			  "okhttp/3.2.0",
			  "okhttp/3.5.0",
			  "okhttp/4.1.0"
			]
		  }
		,
		  {
			"pattern": "Nuzzel",
			"addition_date": "2017/11/02",
			"instances": [
			  "Nuzzel"
			]
		  }
		,
		  {
			"pattern": "omgili",
			"addition_date": "2017/11/02",
			"url": "http://omgili.com",
			"instances": [
			  "omgili/0.5 +http://omgili.com"
			]
		  }
		,
		  {
			"pattern": "PocketParser",
			"addition_date": "2017/11/02",
			"url": "https://getpocket.com/pocketparser_ua",
			"instances": [
			  "PocketParser/2.0 (+https://getpocket.com/pocketparser_ua)"
			]
		  }
		,
		  {
			"pattern": "YisouSpider",
			"addition_date": "2017/11/02",
			"instances": [
			  "YisouSpider",
			  "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 YisouSpider/5.0 Safari/537.36"
			]
		  }
		,
		  {
			"pattern": "um-LN",
			"addition_date": "2017/11/02",
			"instances": [
			  "Mozilla/5.0 (compatible; um-LN/1.0; mailto: techinfo@ubermetrics-technologies.com)"
			]
		  }
		,
		  {
			"pattern": "ToutiaoSpider",
			"addition_date": "2017/11/02",
			"url": "http://web.toutiao.com/media_cooperation/",
			"instances": [
			  "Mozilla/5.0 (compatible; ToutiaoSpider/1.0; http://web.toutiao.com/media_cooperation/;)"
			]
		  }
		,
		  {
			"pattern": "MuckRack",
			"addition_date": "2017/11/02",
			"url": "http://muckrack.com",
			"instances": [
			  "Mozilla/5.0 (compatible; MuckRack/1.0; +http://muckrack.com)"
			]
		  }
		,
		  {
			"pattern": "Jamie's Spider",
			"addition_date": "2017/11/02",
			"url": "http://jamiembrown.com/",
			"instances": [
			  "Jamie's Spider (http://jamiembrown.com/)"
			]
		  }
		,
		  {
			"pattern": "AHC\\/",
			"addition_date": "2017/11/02",
			"instances": [
			  "AHC/2.0"
			]
		  }
		,
		  {
			"pattern": "NetcraftSurveyAgent",
			"addition_date": "2017/11/02",
			"instances": [
			  "Mozilla/5.0 (compatible; NetcraftSurveyAgent/1.0; +info@netcraft.com)"
			]
		  }
		,
		  {
			"pattern": "Laserlikebot",
			"addition_date": "2017/11/02",
			"instances": [
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Laserlikebot/0.1)"
			]
		  }
		,
		  {
			"pattern": "^Apache-HttpClient",
			"addition_date": "2017/11/02",
			"instances": [
			  "Apache-HttpClient/4.2.3 (java 1.5)",
			  "Apache-HttpClient/4.2.5 (java 1.5)",
			  "Apache-HttpClient/4.3.1 (java 1.5)",
			  "Apache-HttpClient/4.3.3 (java 1.5)",
			  "Apache-HttpClient/4.3.5 (java 1.5)",
			  "Apache-HttpClient/4.4.1 (Java/1.8.0_65)",
			  "Apache-HttpClient/4.5.2 (Java/1.8.0_65)",
			  "Apache-HttpClient/4.5.2 (Java/1.8.0_151)",
			  "Apache-HttpClient/4.5.2 (Java/1.8.0_161)",
			  "Apache-HttpClient/4.5.2 (Java/1.8.0_181)",
			  "Apache-HttpClient/4.5.3 (Java/1.8.0_121)",
			  "Apache-HttpClient/4.5.3-SNAPSHOT (Java/1.8.0_152)",
			  "Apache-HttpClient/4.5.7 (Java/11.0.3)",
			  "Apache-HttpClient/4.5.10 (Java/1.8.0_201)"
			]
		  }
		,
		  {
			"pattern": "AppEngine-Google",
			"addition_date": "2017/11/02",
			"instances": [
			  "AppEngine-Google; (+http://code.google.com/appengine; appid: example)",
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36 AppEngine-Google; (+http://code.google.com/appengine; appid: s~feedly-nikon3)"
			]
		  }
		,
		  {
			"pattern": "Jetty",
			"addition_date": "2017/11/02",
			"instances": [
			  "Jetty/9.3.z-SNAPSHOT"
			]
		  }
		,
		  {
			"pattern": "Upflow",
			"addition_date": "2017/11/02",
			"instances": [
			  "Upflow/1.0"
			]
		  }
		,
		  {
			"pattern": "Thinklab",
			"addition_date": "2017/11/02",
			"url": "thinklab.com",
			"instances": [
			  "Thinklab (thinklab.com)"
			]
		  }
		,
		  {
			"pattern": "Traackr.com",
			"addition_date": "2017/11/02",
			"url": "Traackr.com",
			"instances": [
			  "Traackr.com"
			]
		  }
		,
		  {
			"pattern": "Twurly",
			"addition_date": "2017/11/02",
			"url": "http://twurly.org",
			"instances": [
			  "Ruby, Twurly v1.1 (http://twurly.org)"
			]
		  }
		,
		  {
			"pattern": "Mastodon",
			"addition_date": "2017/11/02",
			"instances": [
			  "http.rb/2.2.2 (Mastodon/1.5.1; +https://example-masto-instance.org/)"
			]
		  }
		,
		  {
			"pattern": "http_get",
			"addition_date": "2017/11/02",
			"instances": [
			  "http_get"
			]
		  }
		,
		  {
			"pattern": "DnyzBot",
			"addition_date": "2017/11/20",
			"instances": [
			  "Mozilla/5.0 (compatible; DnyzBot/1.0)"
			]
		  }
		,
		  {
			"pattern": "botify",
			"addition_date": "2018/02/01",
			"instances": [
			  "Mozilla/5.0 (compatible; botify; http://botify.com)"
			]
		  }
		,
		  {
			"pattern": "007ac9 Crawler",
			"addition_date": "2018/02/09",
			"instances": [
			  "Mozilla/5.0 (compatible; 007ac9 Crawler; http://crawler.007ac9.net/)"
			]
		  }
		,
		  {
			"pattern": "BehloolBot",
			"addition_date": "2018/02/09",
			"instances": [
			  "Mozilla/5.0 (compatible; BehloolBot/beta; +http://www.webeaver.com/bot)"
			]
		  }
		,
		  {
			"pattern": "BrandVerity",
			"addition_date": "2018/02/27",
			"instances": [
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:41.0) Gecko/20100101 Firefox/55.0 BrandVerity/1.0 (http://www.brandverity.com/why-is-brandverity-visiting-me)",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/11A465 Twitter for iPhone BrandVerity/1.0 (http://www.brandverity.com/why-is-brandverity-visiting-me)"
			],
			"url": "http://www.brandverity.com/why-is-brandverity-visiting-me"
		  }
		,
		  {
			"pattern": "check_http",
			"addition_date": "2018/02/09",
			"instances": [
			  "check_http/v2.2.1 (nagios-plugins 2.2.1)"
			]
		  }
		,
		  {
			"pattern": "BDCbot",
			"addition_date": "2018/02/09",
			"instances": [
			  "Mozilla/5.0 (Windows NT 6.1; compatible; BDCbot/1.0; +http://bigweb.bigdatacorp.com.br/faq.aspx) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.118 Safari/537.36",
			  "Mozilla/5.0 (Windows NT 10.0; Win64; x64; BDCbot/1.0; +http://bigweb.bigdatacorp.com.br/faq.aspx) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
			]
		  }
		,
		  {
			"pattern": "ZumBot",
			"addition_date": "2018/02/09",
			"instances": [
			  "Mozilla/5.0 (compatible; ZumBot/1.0; http://help.zum.com/inquiry)"
			]
		  }
		,
		  {
			"pattern": "EZID",
			"addition_date": "2018/02/09",
			"instances": [
			  "EZID (EZID link checker; https://ezid.cdlib.org/)"
			]
		  }
		,
		  {
			"pattern": "ICC-Crawler",
			"addition_date": "2018/02/28",
			"instances": [
			  "ICC-Crawler/2.0 (Mozilla-compatible; ; http://ucri.nict.go.jp/en/icccrawler.html)"
			],
			"url": "http://ucri.nict.go.jp/en/icccrawler.html"
		  }
		,
		  {
			"pattern": "ArchiveBot",
			"addition_date": "2018/02/28",
			"instances": [
			  "ArchiveTeam ArchiveBot/20170106.02 (wpull 2.0.2)"
			],
			"url": "https://github.com/ArchiveTeam/ArchiveBot"
		  }
		,
		  {
			"pattern": "^LCC ",
			"addition_date": "2018/02/28",
			"instances": [
			  "LCC (+http://corpora.informatik.uni-leipzig.de/crawler_faq.html)"
			],
			"url": "http://corpora.informatik.uni-leipzig.de/crawler_faq.html"
		  }
		,
		  {
			"pattern": "filterdb.iss.net\\/crawler",
			"addition_date": "2018/03/16",
			"instances": [
			  "Mozilla/5.0 (compatible; oBot/2.3.1; +http://filterdb.iss.net/crawler/)"
			],
			"url": "http://filterdb.iss.net/crawler/"
		  }
		,
		  {
			"pattern": "BLP_bbot",
			"addition_date": "2018/03/27",
			"instances": [
			  "BLP_bbot/0.1"
			]
		  }
		,
		  {
			"pattern": "BomboraBot",
			"addition_date": "2018/03/27",
			"instances": [
			  "Mozilla/5.0 (compatible; BomboraBot/1.0; +http://www.bombora.com/bot)"
			],
			"url": "http://www.bombora.com/bot"
		  }
		,
		  {
			"pattern": "Buck\\/",
			"addition_date": "2018/03/27",
			"instances": [
			  "Buck/2.2; (+https://app.hypefactors.com/media-monitoring/about.html)"
			],
			"url": "https://app.hypefactors.com/media-monitoring/about.html"
		  }
		,
		  {
			"pattern": "Companybook-Crawler",
			"addition_date": "2018/03/27",
			"instances": [
			  "Companybook-Crawler (+https://www.companybooknetworking.com/)"
			],
			"url": "https://www.companybooknetworking.com/"
		  }
		,
		  {
			"pattern": "Genieo",
			"addition_date": "2018/03/27",
			"instances": [
			  "Mozilla/5.0 (compatible; Genieo/1.0 http://www.genieo.com/webfilter.html)"
			],
			"url": "http://www.genieo.com/webfilter.html"
		  }
		,
		  {
			"pattern": "magpie-crawler",
			"addition_date": "2018/03/27",
			"instances": [
			  "magpie-crawler/1.1 (U; Linux amd64; en-GB; +http://www.brandwatch.net)"
			],
			"url": "http://www.brandwatch.net"
		  }
		,
		  {
			"pattern": "MeltwaterNews",
			"addition_date": "2018/03/27",
			"instances": [
			  "MeltwaterNews www.meltwater.com"
			],
			"url": "http://www.meltwater.com"
		  }
		,
		  {
			"pattern": "Moreover",
			"addition_date": "2018/03/27",
			"instances": [
			  "Mozilla/5.0 Moreover/5.1 (+http://www.moreover.com)"
			],
			"url": "http://www.moreover.com"
		  }
		,
		  {
			"pattern": "newspaper\\/",
			"addition_date": "2018/03/27",
			"instances": [
			  "newspaper/0.1.0.7",
			  "newspaper/0.2.5",
			  "newspaper/0.2.6",
			  "newspaper/0.2.8"
			]
		  }
		,
		  {
			"pattern": "ScoutJet",
			"addition_date": "2018/03/27",
			"instances": [
			  "Mozilla/5.0 (compatible; ScoutJet; +http://www.scoutjet.com/)"
			],
			"url": "http://www.scoutjet.com/"
		  }
		,
		  {
			"pattern": "(^| )sentry\\/",
			"addition_date": "2018/03/27",
			"instances": [
			  "sentry/8.22.0 (https://sentry.io)"
			],
			"url": "https://sentry.io"
		  }
		,
		  {
			"pattern": "StorygizeBot",
			"addition_date": "2018/03/27",
			"instances": [
			  "Mozilla/5.0 (compatible; StorygizeBot; http://www.storygize.com)"
			],
			"url": "http://www.storygize.com"
		  }
		,
		  {
			"pattern": "UptimeRobot",
			"addition_date": "2018/03/27",
			"instances": [
			  "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)"
			],
			"url": "http://www.uptimerobot.com/"
		  }
		,
		  {
			"pattern": "OutclicksBot",
			"addition_date": "2018/04/21",
			"instances": [
			  "OutclicksBot/2 +https://www.outclicks.net/agent/VjzDygCuk4ubNmg40ZMbFqT0sIh7UfOKk8s8ZMiupUR",
			  "OutclicksBot/2 +https://www.outclicks.net/agent/gIYbZ38dfAuhZkrFVl7sJBFOUhOVct6J1SvxgmBZgCe",
			  "OutclicksBot/2 +https://www.outclicks.net/agent/PryJzTl8POCRHfvEUlRN5FKtZoWDQOBEvFJ2wh6KH5J",
			  "OutclicksBot/2 +https://www.outclicks.net/agent/p2i4sNUh7eylJF1S6SGgRs5mP40ExlYvsr9GBxVQG6h"
			],
			"url": "https://www.outclicks.net"
		  }
		,
		  {
			"pattern": "seoscanners",
			"addition_date": "2018/05/27",
			"instances": [
			  "Mozilla/5.0 (compatible; seoscanners.net/1; +spider@seoscanners.net)"
			],
			"url": "http://www.seoscanners.net/"
		  }
		,
		  {
			"pattern": "Hatena",
			"addition_date": "2018/05/29",
			"instances": [
			  "Hatena Antenna/0.3",
			  "Hatena::Russia::Crawler/0.01",
			  "Hatena-Favicon/2 (http://www.hatena.ne.jp/faq/)",
			  "Hatena::Scissors/0.01",
			  "HatenaBookmark/4.0 (Hatena::Bookmark; Analyzer)",
			  "Hatena::Fetcher/0.01 (master) Furl/3.13"
			]
		  }
		,
		  {
			"pattern": "Google Web Preview",
			"addition_date": "2018/05/31",
			"instances": [
			  "Mozilla/5.0 (Linux; U; Android 2.3.4; generic) AppleWebKit/537.36 (KHTML, like Gecko; Google Web Preview) Version/4.0 Mobile Safari/537.36",
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko; Google Web Preview) Chrome/27.0.1453 Safari/537.36"
			]
		  }
		,
		  {
			"pattern": "MauiBot",
			"addition_date": "2018/06/06",
			"instances": [
			  "MauiBot (crawler.feedback+wc@gmail.com)"
			]
		  }
		,
		  {
			"pattern": "AlphaBot",
			"addition_date": "2018/05/27",
			"instances": [
			  "Mozilla/5.0 (compatible; AlphaBot/3.2; +http://alphaseobot.com/bot.html)"
			],
			"url": "http://alphaseobot.com/bot.html"
		  }
		,
		  {
			"pattern": "SBL-BOT",
			"addition_date": "2018/06/06",
			"instances": [
			  "SBL-BOT (http://sbl.net)"
			],
			"url": "http://sbl.net",
			"description" : "Bot of SoftByte BlackWidow"
		  }
		,
		  {
			"pattern": "IAS crawler",
			"addition_date": "2018/06/06",
			"instances": [
			  "IAS crawler (ias_crawler; http://integralads.com/site-indexing-policy/)"
			],
			"url": "http://integralads.com/site-indexing-policy/",
			"description" : "Bot of Integral Ad Science, Inc."
		  }
		,
		  {
			"pattern": "adscanner",
			"addition_date": "2018/06/24",
			"instances": [
			  "Mozilla/5.0 (compatible; adscanner/)"
			]
		  }
		,
		  {
			"pattern": "Netvibes",
			"addition_date": "2018/06/24",
			"instances": [
			  "Netvibes (crawler/bot; http://www.netvibes.com",
			  "Netvibes (crawler; http://www.netvibes.com)"
			],
			"url": "http://www.netvibes.com"
		  }
		,
		  {
			"pattern": "acapbot",
			"addition_date": "2018/06/27",
			"instances": [
			  "Mozilla/5.0 (compatible;acapbot/0.1;treat like Googlebot)",
			  "Mozilla/5.0 (compatible;acapbot/0.1.;treat like Googlebot)"
			]
		  }
		,
		  {
			"pattern": "Baidu-YunGuanCe",
			"addition_date": "2018/06/27",
			"instances": [
			  "Baidu-YunGuanCe-Bot(ce.baidu.com)",
			  "Baidu-YunGuanCe-SLABot(ce.baidu.com)",
			  "Baidu-YunGuanCe-ScanBot(ce.baidu.com)",
			  "Baidu-YunGuanCe-PerfBot(ce.baidu.com)",
			  "Baidu-YunGuanCe-VSBot(ce.baidu.com)"
			],
			"url": "https://ce.baidu.com/topic/topic20150908",
			"description": "Baidu Cloud Watch"
		  }
		,
		  {
			"pattern": "bitlybot",
			"addition_date": "2018/06/27",
			"instances": [
			  "bitlybot/3.0 (+http://bit.ly/)",
			  "bitlybot/2.0",
			  "bitlybot"
			],
			"url": "http://bit.ly/"
		  }
		,
		  {
			"pattern": "blogmuraBot",
			"addition_date": "2018/06/27",
			"instances": [
			  "blogmuraBot (+http://www.blogmura.com)"
			],
			"url": "http://www.blogmura.com",
			"description": "A blog ranking site which links to blogs on just about every theme possible."
		  }
		,
		  {
			"pattern": "Bot.AraTurka.com",
			"addition_date": "2018/06/27",
			"instances": [
			  "Bot.AraTurka.com/0.0.1"
			],
			"url": "http://www.araturka.com"
		  }
		,
		  {
			"pattern": "bot-pge.chlooe.com",
			"addition_date": "2018/06/27",
			"instances": [
			  "bot-pge.chlooe.com/1.0.0 (+http://www.chlooe.com/)"
			]
		  }
		,
		  {
			"pattern": "BoxcarBot",
			"addition_date": "2018/06/27",
			"instances": [
			  "Mozilla/5.0 (compatible; BoxcarBot/1.1; +awesome@boxcar.io)"
			],
			"url": "https://boxcar.io/"
		  }
		,
		  {
			"pattern": "BTWebClient",
			"addition_date": "2018/06/27",
			"instances": [
			  "BTWebClient/180B(9704)"
			],
			"url": "http://www.utorrent.com/",
			"description": "µTorrent BitTorrent Client"
		  }
		,
		  {
			"pattern": "ContextAd Bot",
			"addition_date": "2018/06/27",
			"instances": [
			  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0;.NET CLR 1.0.3705; ContextAd Bot 1.0)",
			  "ContextAd Bot 1.0"
			]
		  }
		,
		  {
			"pattern": "Digincore bot",
			"addition_date": "2018/06/27",
			"instances": [
			  "Mozilla/5.0 (compatible; Digincore bot; https://www.digincore.com/crawler.html for rules and instructions.)"
			],
			"url": "http://www.digincore.com/crawler.html"
		  }
		,
		  {
			"pattern": "Disqus",
			"addition_date": "2018/06/27",
			"instances": [
			  "Disqus/1.0"
			],
			"url": "https://disqus.com/",
			"description": "validate and quality check pages."
		  }
		,
		  {
			"pattern": "Feedly",
			"addition_date": "2018/06/27",
			"instances": [
			  "Feedly/1.0 (+http://www.feedly.com/fetcher.html; like FeedFetcher-Google)",
			  "FeedlyBot/1.0 (http://feedly.com)"
			],
			"url": "https://www.feedly.com/fetcher.html",
			"description": "Feedly Fetcher is how Feedly grabs RSS or Atom feeds when users choose to add them to their Feedly or any of the other applications built on top of the feedly cloud."
		  }
		,
		  {
			"pattern": "Fetch\\/",
			"addition_date": "2018/06/27",
			"instances": [
			  "Fetch/2.0a (CMS Detection/Web/SEO analysis tool, see http://guess.scritch.org)"
			]
		  }
		,
		  {
			"pattern": "Fever",
			"addition_date": "2018/06/27",
			"instances": [
			  "Fever/1.38 (Feed Parser; http://feedafever.com; Allow like Gecko)"
			],
			"url": "http://feedafever.com"
		  }
		,
		  {
			"pattern": "Flamingo_SearchEngine",
			"addition_date": "2018/06/27",
			"instances": [
			  "Flamingo_SearchEngine (+http://www.flamingosearch.com/bot)"
			]
		  }
		,
		  {
			"pattern": "FlipboardProxy",
			"addition_date": "2018/06/27",
			"instances": [
			  "Mozilla/5.0 (compatible; FlipboardProxy/1.1; +http://flipboard.com/browserproxy)",
			  "Mozilla/5.0 (compatible; FlipboardProxy/1.2; +http://flipboard.com/browserproxy)",
			  "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2) Gecko/20100115 Firefox/3.6 (FlipboardProxy/1.1; +http://flipboard.com/browserproxy)",
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:28.0) Gecko/20100101 Firefox/28.0 (FlipboardProxy/1.1; +http://flipboard.com/browserproxy)",
			  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:49.0) Gecko/20100101 Firefox/49.0 (FlipboardProxy/1.2; +http://flipboard.com/browserproxy)"
			],
			"url": "https://about.flipboard.com/browserproxy/",
			"description": "a proxy service to fetch, validate, and prepare certain elements of websites for presentation through the Flipboard Application"
		  }
		,
		  {
			"pattern": "g2reader-bot",
			"addition_date": "2018/06/27",
			"instances": [
			  "g2reader-bot/1.0 (+http://www.g2reader.com/)"
			],
			"url": "http://www.g2reader.com/"
		  }
		,
		  {
			"pattern": "G2 Web Services",
			"addition_date": "2019/03/01",
			"instances": [
			  "G2 Web Services/1.0 (built with StormCrawler Archetype 1.8; https://www.g2webservices.com/; developers@g2llc.com)"
			],
			"url": "https://www.g2webservices.com/"
		  }
		,
		  {
			"pattern": "imrbot",
			"addition_date": "2018/06/27",
			"instances": [
			  "Mozilla/5.0 (compatible; imrbot/1.10.8 +http://www.mignify.com)"
			],
			"url": "http://www.mignify.com"
		  }
		,
		  {
			"pattern": "K7MLWCBot",
			"addition_date": "2018/06/27",
			"instances": [
			  "K7MLWCBot/1.0 (+http://www.k7computing.com)"
			],
			"url": "http://www.k7computing.com",
			"description": "Virus scanner"
		  }
		,
		  {
			"pattern": "Kemvibot",
			"addition_date": "2018/06/27",
			"instances": [
			  "Kemvibot/1.0 (http://kemvi.com, marco@kemvi.com)"
			],
			"url": "http://kemvi.com"
		  }
		,
		  {
			"pattern": "Landau-Media-Spider",
			"addition_date": "2018/06/27",
			"instances": [
			  "Landau-Media-Spider/1.0(http://bots.landaumedia.de/bot.html)"
			],
			"url": "http://bots.landaumedia.de/bot.html"
		  }
		,
		  {
			"pattern": "linkapediabot",
			"addition_date": "2018/06/27",
			"instances": [
			  "linkapediabot (+http://www.linkapedia.com)"
			],
			"url": "http://www.linkapedia.com"
		  }
		,
		  {
			"pattern": "vkShare",
			"addition_date": "2018/07/02",
			"instances": [
			  "Mozilla/5.0 (compatible; vkShare; +http://vk.com/dev/Share)"
			],
			"url": "http://vk.com/dev/Share"
		  }
		,
		  {
			"pattern": "Siteimprove.com",
			"addition_date": "2018/06/22",
			"instances": [
			  "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0) LinkCheck by Siteimprove.com",
			  "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.0) Match by Siteimprove.com",
			  "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0) SiteCheck-sitecrawl by Siteimprove.com",
			  "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.0) LinkCheck by Siteimprove.com"
			]
		  }
		,
		  {
			"pattern": "BLEXBot\\/",
			"addition_date": "2018/07/07",
			"instances": [
			  "Mozilla/5.0 (compatible; BLEXBot/1.0; +http://webmeup-crawler.com/)"
			],
			"url": "http://webmeup-crawler.com"
		  }
		,
		  {
			"pattern": "DareBoost",
			"addition_date": "2018/07/07",
			"instances": [
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36 DareBoost"
			],
			"url": "https://www.dareboost.com/",
			"description": "Bot to test, Analyze and Optimize website"
		  }
		,
		  {
			"pattern": "ZuperlistBot\\/",
			"addition_date": "2018/07/07",
			"instances": [
			  "Mozilla/5.0 (compatible; ZuperlistBot/1.0)"
			]
		  }
		,
		  {
			"pattern": "Miniflux\\/",
			"addition_date": "2018/07/07",
			"instances": [
			  "Mozilla/5.0 (compatible; Miniflux/2.0.x-dev; +https://miniflux.net)",
			  "Mozilla/5.0 (compatible; Miniflux/2.0.3; +https://miniflux.net)",
			  "Mozilla/5.0 (compatible; Miniflux/2.0.7; +https://miniflux.net)",
			  "Mozilla/5.0 (compatible; Miniflux/2.0.10; +https://miniflux.net)",
			  "Mozilla/5.0 (compatibl$; Miniflux/2.0.x-dev; +https://miniflux.app)",
			  "Mozilla/5.0 (compatible; Miniflux/2.0.11; +https://miniflux.app)",
			  "Mozilla/5.0 (compatible; Miniflux/2.0.12; +https://miniflux.app)",
			  "Mozilla/5.0 (compatible; Miniflux/ae1dc1a; +https://miniflux.app)",
			  "Mozilla/5.0 (compatible; Miniflux/3b6e44c; +https://miniflux.app)"
			],
			"url": "https://miniflux.net",
			"description": "Miniflux is a minimalist and opinionated feed reader."
		  }
		,
		  {
			"pattern": "Feedspot",
			"addition_date": "2018/07/07",
			"instances": [
			  "Mozilla/5.0 (compatible; Feedspotbot/1.0; +http://www.feedspot.com/fs/bot)",
			  "Mozilla/5.0 (compatible; Feedspot/1.0 (+https://www.feedspot.com/fs/fetcher; like FeedFetcher-Google)"
			],
			"url": "http://www.feedspot.com/fs/bot"
		  }
		,
		  {
			"pattern": "Diffbot\\/",
			"addition_date": "2018/07/07",
			"instances": [
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1.2) Gecko/20090729 Firefox/3.5.2 (.NET CLR 3.5.30729; Diffbot/0.1; +http://www.diffbot.com)"
			],
			"url": "http://www.diffbot.com"
		  }
		,
		  {
			"pattern": "SEOkicks",
			"addition_date": "2018/08/22",
			"instances": [
			  "Mozilla/5.0 (compatible; SEOkicks; +https://www.seokicks.de/robot.html)"
			],
			"url": "https://www.seokicks.de/robot.html"
		  }
		,
		  {
			"pattern": "tracemyfile",
			"addition_date": "2018/08/23",
			"instances": [
			  "Mozilla/5.0 (compatible; tracemyfile/1.0; +bot@tracemyfile.com)"
			]
		  }
		,
		  {
			"pattern": "Nimbostratus-Bot",
			"addition_date": "2018/08/29",
			"instances": [
			  "Mozilla/5.0 (compatible; Nimbostratus-Bot/v1.3.2; http://cloudsystemnetworks.com)"
			]
		  }
		,
		  {
			"pattern": "zgrab",
			"addition_date": "2018/08/30",
			"instances": [
			  "Mozilla/5.0 zgrab/0.x"
			],
			"url": "https://zmap.io/"
		  }
		,
		  {
			"pattern": "PR-CY.RU",
			"addition_date": "2018/08/30",
			"instances": [
			  "Mozilla/5.0 (compatible; PR-CY.RU; + https://a.pr-cy.ru)"
			],
			"url": "https://a.pr-cy.ru/"
		  }
		,
		  {
			"pattern": "AdsTxtCrawler",
			"addition_date": "2018/08/30",
			"instances": [
			  "AdsTxtCrawler/1.0"
			]
		  },
		  {
			"pattern": "Datafeedwatch",
			"addition_date": "2018/09/05",
			"instances": [
			  "Datafeedwatch/2.1.x"
			],
			"url": "https://www.datafeedwatch.com/"
		  }
		,
		  {
			"pattern": "Zabbix",
			"addition_date": "2018/09/05",
			"instances": [
			  "Zabbix"
			],
			"url": "https://www.zabbix.com/documentation/3.4/manual/web_monitoring"
		  }
		,
		  {
			"pattern": "TangibleeBot",
			"addition_date": "2018/09/05",
			"instances": [
			  "TangibleeBot/1.0.0.0 (http://tangiblee.com/bot)"
			],
			"url": "http://tangiblee.com/bot"
		  }
		,
		  {
			"pattern": "google-xrawler",
			"addition_date": "2018/09/05",
			"instances": [
			  "google-xrawler"
			],
			"url": "https://webmasters.stackexchange.com/questions/105560/what-is-the-google-xrawler-user-agent-used-for"
		  }
		,
		  {
			"pattern": "axios",
			"addition_date": "2018/09/06",
			"instances": [
			  "axios/0.18.0",
			  "axios/0.19.0"
			],
			"url": "https://github.com/axios/axios"
		  }
		,
		  {
			"pattern": "Amazon CloudFront",
			"addition_date": "2018/09/07",
			"instances": [
			  "Amazon CloudFront"
			],
			"url": "https://aws.amazon.com/cloudfront/"
		  }
		,
		  {
			"pattern": "Pulsepoint",
			"addition_date": "2018/09/24",
			"instances": [
			  "Pulsepoint XT3 web scraper"
			]
		  }
		,
		  {
			"pattern": "CloudFlare-AlwaysOnline",
			"addition_date": "2018/09/27",
			"instances": [
			  "Mozilla/5.0 (compatible; CloudFlare-AlwaysOnline/1.0; +http://www.cloudflare.com/always-online) AppleWebKit/534.34",
			  "Mozilla/5.0 (compatible; CloudFlare-AlwaysOnline/1.0; +https://www.cloudflare.com/always-online) AppleWebKit/534.34"
			],
			"url" : "https://www.cloudflare.com/always-online/"
		  }
		,
		  {
			"pattern": "Google-Structured-Data-Testing-Tool",
			"addition_date": "2018/10/02",
			"instances": [
			  "Mozilla/5.0 (compatible; Google-Structured-Data-Testing-Tool +https://search.google.com/structured-data/testing-tool)",
			  "Mozilla/5.0 (compatible; Google-Structured-Data-Testing-Tool +http://developers.google.com/structured-data/testing-tool/)"
			],
			"url": "https://search.google.com/structured-data/testing-tool"
		  }
		,
		  {
			"pattern": "WordupInfoSearch",
			"addition_date": "2018/10/07",
			"instances": [
			  "WordupInfoSearch/1.0"
			]
		  }
		,
		  {
			"pattern": "WebDataStats",
			"addition_date": "2018/10/08",
			"instances": [
			  "Mozilla/5.0 (compatible; WebDataStats/1.0 ; +https://webdatastats.com/policy.html)"
			],
			"url": "https://webdatastats.com/"
		  }
		,
		  {
			"pattern": "HttpUrlConnection",
			"addition_date": "2018/10/08",
			"instances": [
			  "Jersey/2.25.1 (HttpUrlConnection 1.8.0_141)"
			]
		  }
		,
		  {
			"pattern": "Seekport Crawler",
			"addition_date": "2018/10/08",
			"instances": [
			  "Mozilla/5.0 (compatible; Seekport Crawler; http://seekport.com/)"
			],
			"url": "http://seekport.com/"
		  }
		,
		  {
			"pattern": "ZoomBot",
			"addition_date": "2018/10/10",
			"instances": [
			  "ZoomBot (Linkbot 1.0 http://suite.seozoom.it/bot.html)"
			],
			"url": "http://suite.seozoom.it/bot.html"
		  }
		,
		  {
			"pattern": "VelenPublicWebCrawler",
			"addition_date": "2018/10/09",
			"instances": [
			  "VelenPublicWebCrawler (velen.io)"
			]
		  }
		,
		  {
			"pattern": "MoodleBot",
			"addition_date": "2018/10/10",
			"instances": [
			  "MoodleBot/1.0"
			]
		  }
		,
		  {
			"pattern": "jpg-newsbot",
			"addition_date": "2018/10/10",
			"instances": [
			  "jpg-newsbot/2.0; (+https://vipnytt.no/bots/)"
			],
			"url": "https://vipnytt.no/bots/"
		  }
		,
		  {
			"pattern": "outbrain",
			"addition_date": "2018/10/14",
			"instances": [
			  "Mozilla/5.0 (Java) outbrain"
			],
			"url": "https://www.outbrain.com/help/advertisers/invalid-url/"
		  }
		,
		  {
			"pattern": "W3C_Validator",
			"addition_date": "2018/10/14",
			"instances": [
			  "W3C_Validator/1.3"
			],
			"url": "https://validator.w3.org/services"
		  }
		,
		  {
			"pattern": "Validator\\.nu",
			"addition_date": "2018/10/14",
			"instances": [
			  "Validator.nu/LV"
			],
			"url": "https://validator.w3.org/services"
		  }
		,
		  {
			"pattern": "W3C-checklink",
			"addition_date": "2018/10/14",
			"depends_on": ["libwww-perl"],
			"instances": [
			  "W3C-checklink/2.90 libwww-perl/5.64",
			  "W3C-checklink/3.6.2.3 libwww-perl/5.64",
			  "W3C-checklink/4.2 [4.20] libwww-perl/5.803",
			  "W3C-checklink/4.2.1 [4.21] libwww-perl/5.803",
			  "W3C-checklink/4.3 [4.42] libwww-perl/5.805",
			  "W3C-checklink/4.3 [4.42] libwww-perl/5.808",
			  "W3C-checklink/4.3 [4.42] libwww-perl/5.820",
			  "W3C-checklink/4.5 [4.154] libwww-perl/5.823",
			  "W3C-checklink/4.5 [4.160] libwww-perl/5.823"
			],
			"url": "https://validator.w3.org/services"
		  }
		,
		  {
			"pattern": "W3C-mobileOK",
			"addition_date": "2018/10/14",
			"instances": [
			  "W3C-mobileOK/DDC-1.0"
			],
			"url": "https://validator.w3.org/services"
		  }
		,
		  {
			"pattern": "W3C_I18n-Checker",
			"addition_date": "2018/10/14",
			"instances": [
			  "W3C_I18n-Checker/1.0"
			],
			"url": "https://validator.w3.org/services"
		  }
		,
		  {
			"pattern": "FeedValidator",
			"addition_date": "2018/10/14",
			"instances": [
			  "FeedValidator/1.3"
			],
			"url": "https://validator.w3.org/services"
		  }
		,
		  {
			"pattern": "W3C_CSS_Validator",
			"addition_date": "2018/10/14",
			"instances": [
			  "Jigsaw/2.3.0 W3C_CSS_Validator_JFouffa/2.0"
			],
			"url": "https://validator.w3.org/services"
		  }
		,
		  {
			"pattern": "W3C_Unicorn",
			"addition_date": "2018/10/14",
			"instances": [
			  "W3C_Unicorn/1.0"
			],
			"url": "https://validator.w3.org/services"
		  }
		,
		  {
			"pattern": "Google-PhysicalWeb",
			"addition_date": "2018/10/21",
			"instances": [
			  "Mozilla/5.0 (Google-PhysicalWeb)"
			]
		  }
		,
		  {
			"pattern": "Blackboard",
			"addition_date": "2018/10/28",
			"instances": [
			  "Blackboard Safeassign"
			],
			"url": "https://help.blackboard.com/Learn/Administrator/Hosting/Tools_Management/SafeAssign"
		  },
		  {
			"pattern": "ICBot\\/",
			"addition_date": "2018/10/23",
			"instances": [
			  "Mozilla/5.0 (compatible; ICBot/0.1; +https://ideasandcode.xyz"
			],
			"url": "https://ideasandcode.xyz"
		  },
		  {
			"pattern": "BazQux",
			"addition_date": "2018/10/23",
			"instances": [
			  "Mozilla/5.0 (compatible; BazQux/2.4; +https://bazqux.com/fetcher; 1 subscribers)"
			],
			"url": "https://bazqux.com/fetcher"
		  },
		  {
			"pattern": "Twingly",
			"addition_date": "2018/10/23",
			"instances": [
			  "Mozilla/5.0 (compatible; Twingly Recon; twingly.com)"
			],
			"url": "https://twingly.com"
		  },
		  {
			"pattern": "Rivva",
			"addition_date": "2018/10/23",
			"instances": [
			  "Mozilla/5.0 (compatible; Rivva; http://rivva.de)"
			],
			"url": "http://rivva.de"
		  },
		  {
			"pattern": "Experibot",
			"addition_date": "2018/11/03",
			"instances": [
			  "Experibot-v2 http://goo.gl/ZAr8wX",
			  "Experibot-v3 http://goo.gl/ZAr8wX"
			],
			"url": "https://amirkr.wixsite.com/experibot"
		  },
		  {
			"pattern": "awesomecrawler",
			"addition_date": "2018/11/24",
			"instances": [
			  "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.5 Safari/537.22 +awesomecrawler"
			]
		  },
		  {
			"pattern": "Dataprovider.com",
			"addition_date": "2018/11/24",
			"instances": [
			  "Mozilla/5.0 (compatible; Dataprovider.com)"
			],
			"url": "https://www.dataprovider.com/"
		  },
		  {
			"pattern": "GroupHigh\\/",
			"addition_date": "2018/11/24",
			"instances": [
			  "Mozilla/5.0 (compatible; GroupHigh/1.0; +http://www.grouphigh.com/"
			],
			"url": "http://www.grouphigh.com/"
		  },
		  {
			"pattern": "theoldreader.com",
			"addition_date": "2018/12/02",
			"instances": [
			  "Mozilla/5.0 (compatible; theoldreader.com)"
			],
			"url": "https://www.theoldreader.com/"
		  }
		,
		  {
			"pattern": "AnyEvent",
			"addition_date": "2018/12/07",
			"instances": [
			  "Mozilla/5.0 (compatible; U; AnyEvent-HTTP/2.24; +http://software.schmorp.de/pkg/AnyEvent)"
			],
			"url": "http://software.schmorp.de/pkg/AnyEvent.html"
		  }
		,
		  {
			"pattern": "Uptimebot\\.org",
			"addition_date": "2019/01/17",
			"instances": [
			  "Uptimebot.org - Free website monitoring"
			],
			"url": "http://uptimebot.org/"
		  }
		,
		  {
			"pattern": "Nmap Scripting Engine",
			"addition_date": "2019/02/04",
			"instances": [
			  "Mozilla/5.0 (compatible; Nmap Scripting Engine; https://nmap.org/book/nse.html)"
			],
			"url": "https://nmap.org/book/nse.html"
		  }
		,
		  {
			"pattern": "2ip.ru",
			"addition_date": "2019/02/12",
			"instances": [
			  "2ip.ru CMS Detector (https://2ip.ru/cms/)"
			],
			"url": "https://2ip.ru/cms/"
		  },
		  {
			"pattern": "Clickagy",
			"addition_date": "2019/02/19",
			"instances": [
			  "Clickagy Intelligence Bot v2"
			],
			"url": "https://www.clickagy.com"
		  },
		  {
			"pattern": "Caliperbot",
			"addition_date": "2019/03/02",
			"instances": [
			  "Caliperbot/1.0 (+http://www.conductor.com/caliperbot)"
			],
			"url": "http://www.conductor.com/caliperbot"
		  },
		  {
			"pattern": "MBCrawler",
			"addition_date": "2019/03/02",
			"instances": [
			  "MBCrawler/1.0 (https://monitorbacklinks.com)"
			],
			"url": "https://monitorbacklinks.com"
		  },
		  {
			"pattern": "online-webceo-bot",
			"addition_date": "2019/03/02",
			"instances": [
			  "Mozilla/5.0 (compatible; online-webceo-bot/1.0; +http://online.webceo.com)"
			],
			"url": "http://online.webceo.com"
		  },
		  {
			"pattern": "B2B Bot",
			"addition_date": "2019/03/02",
			"instances": [
			  "B2B Bot"
			]
		  },
		  {
			"pattern": "AddSearchBot",
			"addition_date": "2019/03/02",
			"instances": [
			  "Mozilla/5.0 (compatible; AddSearchBot/0.9; +http://www.addsearch.com/bot; info@addsearch.com)"
			],
			"url": "http://www.addsearch.com/bot"
		  },
		  {
			"pattern": "Google Favicon",
			"addition_date": "2019/03/14",
			"instances": [
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36 Google Favicon"
			]
		  },
		  {
			"pattern": "HubSpot",
			"addition_date": "2019/04/15",
			"instances": [
			  "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.116 Safari/537.36 HubSpot Webcrawler - web-crawlers@hubspot.com",
			  "Mozilla/5.0 (X11; Linux x86_64; HubSpot Single Page link check; web-crawlers+links@hubspot.com) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
			  "Mozilla/5.0 (compatible; HubSpot Crawler; web-crawlers@hubspot.com)",
			  "HubSpot Connect 2.0 (http://dev.hubspot.com/) - BizOpsCompanies-Tq2-BizCoDomainValidationAudit"
			]
		  },
		  {
			"pattern": "Chrome-Lighthouse",
			"addition_date": "2019/03/15",
			"instances": [
			  "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36(KHTML, like Gecko) Chrome/69.0.3464.0 Mobile Safari/537.36 Chrome-Lighthouse",
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36(KHTML, like Gecko) Chrome/69.0.3464.0 Safari/537.36 Chrome-Lighthouse",
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3694.0 Safari/537.36 Chrome-Lighthouse",
			  "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3694.0 Mobile Safari/537.36 Chrome-Lighthouse"
			],
			"url": "https://developers.google.com/speed/pagespeed/insights"
		  },
		  {
			"pattern": "HeadlessChrome",
			"url": "https://developers.google.com/web/updates/2017/04/headless-chrome",
			"addition_date": "2019/06/17",
			"instances": [
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/74.0.3729.169 Safari/537.36",
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/69.0.3494.0 Safari/537.36",
			  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/76.0.3803.0 Safari/537.36"
			]
		  },
		  {
			"pattern": "CheckMarkNetwork\\/",
			"addition_date": "2019/06/30",
			"instances": [
			  "CheckMarkNetwork/1.0 (+http://www.checkmarknetwork.com/spider.html)"
			],
			"url": "https://www.checkmarknetwork.com/"
		  },
		  {
			"pattern": "www\\.uptime\\.com",
			"addition_date": "2019/07/21",
			"instances": [
			  "Mozilla/5.0 (compatible; Uptimebot/1.0; +http://www.uptime.com/uptimebot)"
			],
			"url": "http://www.uptime.com/uptimebot"
		  }
		,
		  {
			"pattern": "Streamline3Bot\\/",
			"addition_date": "2019/07/21",
			"instances": [
			  "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 5.1) Streamline3Bot/1.0",
			  "Mozilla/5.0 (Windows NT 6.1; Win64; x64; +https://www.ubtsupport.com/legal/Streamline3Bot.php) Streamline3Bot/1.0"
			],
			"url": "https://www.ubtsupport.com/legal/Streamline3Bot.php"
		  }
		,
		  {
			"pattern": "serpstatbot\\/",
			"addition_date": "2019/07/25",
			"instances": [
			  "serpstatbot/1.0 (advanced backlink tracking bot; http://serpstatbot.com/; abuse@serpstatbot.com)",
			  "serpstatbot/1.0 (advanced backlink tracking bot; curl/7.58.0; http://serpstatbot.com/; abuse@serpstatbot.com)"
			],
			"url": "http://serpstatbot.com"
		  }
		,
		  {
			"pattern": "MixnodeCache\\/",
			"addition_date": "2019/08/04",
			"instances": [
			  "MixnodeCache/1.8(+https://cache.mixnode.com/)"
			],
			"url": "https://cache.mixnode.com/"
		  }
		,
		  {
			"pattern": "^curl",
			"addition_date": "2019/08/15",
			"instances": [
			  "curl",
			  "curl/7.29.0",
			  "curl/7.47.0",
			  "curl/7.54.0",
			  "curl/7.55.1",
			  "curl/7.64.0",
			  "curl/7.64.1",
			  "curl/7.65.3"
			],
			"url": "https://curl.haxx.se/"
		  }
		,
		  {
			"pattern": "SimpleScraper",
			"addition_date": "2019/08/16",
			"instances": [
			  "Mozilla/5.0 (compatible; SimpleScraper)"
			],
			"url": "https://github.com/ramonkcom/simple-scraper/"
		  }
		,
		  {
			"pattern": "RSSingBot",
			"addition_date": "2019/09/15",
			"instances": [
			  "RSSingBot (http://www.rssing.com)"
			],
			"url": "http://www.rssing.com"
		  }
		,
		  {
			"pattern": "Jooblebot",
			"addition_date": "2019/09/25",
			"instances": [
			  "Mozilla/5.0 (compatible; Jooblebot/2.0; Windows NT 6.1; WOW64; +http://jooble.org/jooble-bot) AppleWebKit/537.36 (KHTML, like Gecko) Safari/537.36"
			],
			"url": "http://jooble.org/jooble-bot"
		  }
		,
		  {
			"pattern": "fedoraplanet",
			"addition_date": "2019/09/28",
			"instances": [
			  "venus/fedoraplanet"
			],
			"url": "http://fedoraplanet.org/"
		  }
		,
		  {
			"pattern": "Friendica",
			"addition_date": "2019/09/28",
			"instances": [
			  "Friendica 'The Tazmans Flax-lily' 2019.01-1293; https://hoyer.xyz"
			],
			"url": "https://hoyer.xyz"
		  }
		,
		  {
			"pattern": "NextCloud",
			"addition_date": "2019/09/30",
			"instances": [
			  "NextCloud-News/1.0"
			],
			"url": "https://nextcloud.com/"
		  }
		,
		  {
			"pattern": "Tiny Tiny RSS",
			"addition_date": "2019/10/04",
			"instances": [
			  "Tiny Tiny RSS/1.15.3 (http://tt-rss.org/)",
			  "Tiny Tiny RSS/17.12 (a2d1fa5) (http://tt-rss.org/)",
			  "Tiny Tiny RSS/19.2 (b68db2d) (http://tt-rss.org/)",
			  "Tiny Tiny RSS/19.8 (http://tt-rss.org/)"
			],
			"url": "http://tt-rss.org/"
		  }
		,
		  {
			"pattern": "RegionStuttgartBot",
			"addition_date": "2019/10/17",
			"instances": [
			  "Mozilla/5.0 (compatible; RegionStuttgartBot/1.0; +http://it.region-stuttgart.de/competenzatlas/unternehmen-suchen/)"
			],
			"url": "http://it.region-stuttgart.de/competenzatlas/unternehmen-suchen/"
		  }
		,
		  {
			"pattern": "Bytespider",
			"addition_date": "2019/11/11",
			"instances": [
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.3754.1902 Mobile Safari/537.36; Bytespider",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.4454.1745 Mobile Safari/537.36; Bytespider",
			  "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.7597.1164 Mobile Safari/537.36; Bytespider;bytespider@bytedance.com",
				  "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2988.1545 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.4141.1682 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.3478.1649 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.5267.1259 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.7990.1979 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.2268.1523 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2576.1836 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.9681.1227 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.6023.1635 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.4944.1981 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.3613.1739 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.4022.1033 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.3248.1547 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.5527.1507 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.5216.1326 Mobile Safari/537.36; Bytespider",
				  "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.9038.1080 Mobile Safari/537.36; Bytespider"
		  	],
			"url": "https://stackoverflow.com/questions/57908900/what-is-the-bytespider-user-agent"
		  }
		,
		  {
			"pattern": "Datanyze",
			"addition_date": "2019/11/17",
			"instances": [
			  "Mozilla/5.0 (X11; Datanyze; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36"
			],
			"url": "https://www.datanyze.com/dnyzbot/"
		  }
		,
		  {
			"pattern": "Google-Site-Verification",
			"addition_date": "2019/12/11",
			"instances": [
			  "Mozilla/5.0 (compatible; Google-Site-Verification/1.0)"
			],
			"url": "https://support.google.com/webmasters/answer/9008080"
		  }
		,
		  {
			"pattern": "TrendsmapResolver",
			"addition_date": "2020/02/24",
			"instances": [
			  "Mozilla/5.0 (compatible; TrendsmapResolver/0.1)"
			],
			"url": "https://www.trendsmap.com/"
		  }
		,
		  {
			"pattern": "tweetedtimes",
			"addition_date": "2020/02/24",
			"instances": [
			  "Mozilla/5.0 (compatible; +http://tweetedtimes.com)"
			],
			"url": "https://tweetedtimes.com/"
		  }
		]
	`

	var patterns []CrawlerPattern

	err := json.Unmarshal([]byte(crawlers), &patterns)

	if err != nil {
		panic(err)
	}

	return patterns
}
