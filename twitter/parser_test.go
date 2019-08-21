package twitter

import "testing"

var jsonDataWithExtendedEntities = []byte(`
{
    "for_user_id": "1062458819794649088",
    "tweet_create_events": [
        {
            "created_at": "Wed May 08 20:11:27 +0000 2019",
            "id": 1126218073604796400,
            "id_str": "1126218073604796417",
            "text": "@whatanimalami what animal do you see? https://t.co/jfzBZULNMI",
            "display_text_range": [
                0,
                38
            ],
            "source": "<a href=\"http://twitter.com\" rel=\"nofollow\">Twitter Web Client</a>",
            "truncated": false,
            "in_reply_to_status_id": null,
            "in_reply_to_status_id_str": null,
            "in_reply_to_user_id": 1062458819794649100,
            "in_reply_to_user_id_str": "1062458819794649088",
            "in_reply_to_screen_name": "whatanimalami",
            "user": {
                "id": 20397392,
                "id_str": "20397392",
                "name": "James Mullen",
                "screen_name": "jamesnmullen",
                "location": "Manchester",
                "url": "https://blog.james-mullen.me.uk/",
                "description": null,
                "translator_type": "none",
                "protected": false,
                "verified": false,
                "followers_count": 94,
                "friends_count": 236,
                "listed_count": 0,
                "favourites_count": 747,
                "statuses_count": 573,
                "created_at": "Sun Feb 08 22:17:10 +0000 2009",
                "utc_offset": null,
                "time_zone": null,
                "geo_enabled": false,
                "lang": "en",
                "contributors_enabled": false,
                "is_translator": false,
                "profile_background_color": "C0DEED",
                "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
                "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
                "profile_background_tile": false,
                "profile_link_color": "1DA1F2",
                "profile_sidebar_border_color": "C0DEED",
                "profile_sidebar_fill_color": "DDEEF6",
                "profile_text_color": "333333",
                "profile_use_background_image": true,
                "profile_image_url": "http://pbs.twimg.com/profile_images/998941447943405568/4to703iG_normal.jpg",
                "profile_image_url_https": "https://pbs.twimg.com/profile_images/998941447943405568/4to703iG_normal.jpg",
                "profile_banner_url": "https://pbs.twimg.com/profile_banners/20397392/1542105926",
                "default_profile": true,
                "default_profile_image": false,
                "following": null,
                "follow_request_sent": null,
                "notifications": null
            },
            "geo": null,
            "coordinates": null,
            "place": null,
            "contributors": null,
            "is_quote_status": false,
            "quote_count": 0,
            "reply_count": 0,
            "retweet_count": 0,
            "favorite_count": 0,
            "entities": {
                "hashtags": [],
                "urls": [],
                "user_mentions": [
                    {
                        "screen_name": "whatanimalami",
                        "name": "What Animal Am I Bot",
                        "id": 1062458819794649100,
                        "id_str": "1062458819794649088",
                        "indices": [
                            0,
                            14
                        ]
                    }
                ],
                "symbols": [],
                "media": [
                    {
                        "id": 1126218015471734800,
                        "id_str": "1126218015471734784",
                        "indices": [
                            39,
                            62
                        ],
                        "media_url": "http://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg",
                        "media_url_https": "https://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg",
                        "url": "https://t.co/jfzBZULNMI",
                        "display_url": "pic.twitter.com/jfzBZULNMI",
                        "expanded_url": "https://twitter.com/jamesnmullen/status/1126218073604796417/photo/1",
                        "type": "photo",
                        "sizes": {
                            "thumb": {
                                "w": 150,
                                "h": 150,
                                "resize": "crop"
                            },
                            "medium": {
                                "w": 1200,
                                "h": 900,
                                "resize": "fit"
                            },
                            "large": {
                                "w": 1280,
                                "h": 960,
                                "resize": "fit"
                            },
                            "small": {
                                "w": 680,
                                "h": 510,
                                "resize": "fit"
                            }
                        }
                    }
                ]
            },
            "extended_entities": {
                "media": [
                    {
                        "id": 1126218015471734800,
                        "id_str": "1126218015471734784",
                        "indices": [
                            39,
                            62
                        ],
                        "media_url": "http://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg",
                        "media_url_https": "https://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg",
                        "url": "https://t.co/jfzBZULNMI",
                        "display_url": "pic.twitter.com/jfzBZULNMI",
                        "expanded_url": "https://twitter.com/jamesnmullen/status/1126218073604796417/photo/1",
                        "type": "photo",
                        "sizes": {
                            "thumb": {
                                "w": 150,
                                "h": 150,
                                "resize": "crop"
                            },
                            "medium": {
                                "w": 1200,
                                "h": 900,
                                "resize": "fit"
                            },
                            "large": {
                                "w": 1280,
                                "h": 960,
                                "resize": "fit"
                            },
                            "small": {
                                "w": 680,
                                "h": 510,
                                "resize": "fit"
                            }
                        }
                    }
                ]
            },
            "favorited": false,
            "retweeted": false,
            "possibly_sensitive": false,
            "filter_level": "low",
            "lang": "en",
            "timestamp_ms": "1557346287524"
        }
    ]
}`)

var jsonDataWithEntities = []byte(`
{
    "for_user_id": "1062458819794649088",
    "tweet_create_events": [
        {
            "created_at": "Wed May 08 20:11:27 +0000 2019",
            "id": 1126218073604796400,
            "id_str": "1126218073604796417",
            "text": "@whatanimalami what animal do you see? https://t.co/jfzBZULNMI",
            "display_text_range": [
                0,
                38
            ],
            "source": "<a href=\"http://twitter.com\" rel=\"nofollow\">Twitter Web Client</a>",
            "truncated": false,
            "in_reply_to_status_id": null,
            "in_reply_to_status_id_str": null,
            "in_reply_to_user_id": 1062458819794649100,
            "in_reply_to_user_id_str": "1062458819794649088",
            "in_reply_to_screen_name": "whatanimalami",
            "user": {
                "id": 20397392,
                "id_str": "20397392",
                "name": "James Mullen",
                "screen_name": "jamesnmullen",
                "location": "Manchester",
                "url": "https://blog.james-mullen.me.uk/",
                "description": null,
                "translator_type": "none",
                "protected": false,
                "verified": false,
                "followers_count": 94,
                "friends_count": 236,
                "listed_count": 0,
                "favourites_count": 747,
                "statuses_count": 573,
                "created_at": "Sun Feb 08 22:17:10 +0000 2009",
                "utc_offset": null,
                "time_zone": null,
                "geo_enabled": false,
                "lang": "en",
                "contributors_enabled": false,
                "is_translator": false,
                "profile_background_color": "C0DEED",
                "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
                "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
                "profile_background_tile": false,
                "profile_link_color": "1DA1F2",
                "profile_sidebar_border_color": "C0DEED",
                "profile_sidebar_fill_color": "DDEEF6",
                "profile_text_color": "333333",
                "profile_use_background_image": true,
                "profile_image_url": "http://pbs.twimg.com/profile_images/998941447943405568/4to703iG_normal.jpg",
                "profile_image_url_https": "https://pbs.twimg.com/profile_images/998941447943405568/4to703iG_normal.jpg",
                "profile_banner_url": "https://pbs.twimg.com/profile_banners/20397392/1542105926",
                "default_profile": true,
                "default_profile_image": false,
                "following": null,
                "follow_request_sent": null,
                "notifications": null
            },
            "geo": null,
            "coordinates": null,
            "place": null,
            "contributors": null,
            "is_quote_status": false,
            "quote_count": 0,
            "reply_count": 0,
            "retweet_count": 0,
            "favorite_count": 0,
            "entities": {
                "hashtags": [],
                "urls": [
                    {
                        "url": "http://test.com/test.jpg"
                    }
                ],
                "user_mentions": [
                    {
                        "screen_name": "whatanimalami",
                        "name": "What Animal Am I Bot",
                        "id": 1062458819794649100,
                        "id_str": "1062458819794649088",
                        "indices": [
                            0,
                            14
                        ]
                    }
                ],
                "symbols": [],
                "media": []
            },
            "extended_entities": {
                "media": []
            },
            "favorited": false,
            "retweeted": false,
            "possibly_sensitive": false,
            "filter_level": "low",
            "lang": "en",
            "timestamp_ms": "1557346287524"
        }
    ]
}`)

var jsonDataWithBothEntitiesAndExtendedEntities = []byte(`
{
    "for_user_id": "1062458819794649088",
    "tweet_create_events": [
        {
            "created_at": "Wed May 08 20:11:27 +0000 2019",
            "id": 1126218073604796400,
            "id_str": "1126218073604796417",
            "text": "@whatanimalami what animal do you see? https://t.co/jfzBZULNMI",
            "display_text_range": [
                0,
                38
            ],
            "source": "<a href=\"http://twitter.com\" rel=\"nofollow\">Twitter Web Client</a>",
            "truncated": false,
            "in_reply_to_status_id": null,
            "in_reply_to_status_id_str": null,
            "in_reply_to_user_id": 1062458819794649100,
            "in_reply_to_user_id_str": "1062458819794649088",
            "in_reply_to_screen_name": "whatanimalami",
            "user": {
                "id": 20397392,
                "id_str": "20397392",
                "name": "James Mullen",
                "screen_name": "jamesnmullen",
                "location": "Manchester",
                "url": "https://blog.james-mullen.me.uk/",
                "description": null,
                "translator_type": "none",
                "protected": false,
                "verified": false,
                "followers_count": 94,
                "friends_count": 236,
                "listed_count": 0,
                "favourites_count": 747,
                "statuses_count": 573,
                "created_at": "Sun Feb 08 22:17:10 +0000 2009",
                "utc_offset": null,
                "time_zone": null,
                "geo_enabled": false,
                "lang": "en",
                "contributors_enabled": false,
                "is_translator": false,
                "profile_background_color": "C0DEED",
                "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
                "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
                "profile_background_tile": false,
                "profile_link_color": "1DA1F2",
                "profile_sidebar_border_color": "C0DEED",
                "profile_sidebar_fill_color": "DDEEF6",
                "profile_text_color": "333333",
                "profile_use_background_image": true,
                "profile_image_url": "http://pbs.twimg.com/profile_images/998941447943405568/4to703iG_normal.jpg",
                "profile_image_url_https": "https://pbs.twimg.com/profile_images/998941447943405568/4to703iG_normal.jpg",
                "profile_banner_url": "https://pbs.twimg.com/profile_banners/20397392/1542105926",
                "default_profile": true,
                "default_profile_image": false,
                "following": null,
                "follow_request_sent": null,
                "notifications": null
            },
            "geo": null,
            "coordinates": null,
            "place": null,
            "contributors": null,
            "is_quote_status": false,
            "quote_count": 0,
            "reply_count": 0,
            "retweet_count": 0,
            "favorite_count": 0,
            "entities": {
                "hashtags": [],
                "urls": [
                    {
                        "url": "http://example.com/test2.jpg"
                    }
                ],
                "user_mentions": [
                    {
                        "screen_name": "whatanimalami",
                        "name": "What Animal Am I Bot",
                        "id": 1062458819794649100,
                        "id_str": "1062458819794649088",
                        "indices": [
                            0,
                            14
                        ]
                    }
                ],
                "symbols": [],
                "media": [
                    {
                        "id": 1126218015471734800,
                        "id_str": "1126218015471734784",
                        "indices": [
                            39,
                            62
                        ],
                        "media_url": "http://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg",
                        "media_url_https": "https://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg",
                        "url": "https://t.co/jfzBZULNMI",
                        "display_url": "pic.twitter.com/jfzBZULNMI",
                        "expanded_url": "https://twitter.com/jamesnmullen/status/1126218073604796417/photo/1",
                        "type": "photo",
                        "sizes": {
                            "thumb": {
                                "w": 150,
                                "h": 150,
                                "resize": "crop"
                            },
                            "medium": {
                                "w": 1200,
                                "h": 900,
                                "resize": "fit"
                            },
                            "large": {
                                "w": 1280,
                                "h": 960,
                                "resize": "fit"
                            },
                            "small": {
                                "w": 680,
                                "h": 510,
                                "resize": "fit"
                            }
                        }
                    }
                ]
            },
            "extended_entities": {
                "media": [
                    {
                        "id": 1126218015471734800,
                        "id_str": "1126218015471734784",
                        "indices": [
                            39,
                            62
                        ],
                        "media_url": "http://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg",
                        "media_url_https": "https://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg",
                        "url": "https://t.co/jfzBZULNMI",
                        "display_url": "pic.twitter.com/jfzBZULNMI",
                        "expanded_url": "https://twitter.com/jamesnmullen/status/1126218073604796417/photo/1",
                        "type": "photo",
                        "sizes": {
                            "thumb": {
                                "w": 150,
                                "h": 150,
                                "resize": "crop"
                            },
                            "medium": {
                                "w": 1200,
                                "h": 900,
                                "resize": "fit"
                            },
                            "large": {
                                "w": 1280,
                                "h": 960,
                                "resize": "fit"
                            },
                            "small": {
                                "w": 680,
                                "h": 510,
                                "resize": "fit"
                            }
                        }
                    }
                ]
            },
            "favorited": false,
            "retweeted": false,
            "possibly_sensitive": false,
            "filter_level": "low",
            "lang": "en",
            "timestamp_ms": "1557346287524"
        }
    ]
}`)

func TestTwitterParsingImageUrlInExtendedEntities(t *testing.T) {

	const expectedTweetID = 1126218073604796417
	const expectedScreenName = "jamesnmullen"
	const expectedImageURL = "http://pbs.twimg.com/media/D6EhUW_WkAAqUic.jpg"

	got := ParseTweetData(jsonDataWithExtendedEntities)

	if got.TweetID != expectedTweetID {
		t.Errorf("got %d, want %d", got.TweetID, expectedTweetID)
	}

	if got.ScreenName != expectedScreenName {
		t.Errorf("got %s, want %s", got.ScreenName, expectedScreenName)
	}

	if got.ImageURL != expectedImageURL {
		t.Errorf("got %s, want %s", got.ImageURL, expectedImageURL)
	}
}

func TestTwitterParsingImageUrlInEntities(t *testing.T) {

	const expectedTweetID = 1126218073604796417
	const expectedScreenName = "jamesnmullen"
	const expectedImageURL = "http://test.com/test.jpg"

	got := ParseTweetData(jsonDataWithEntities)

	if got.TweetID != expectedTweetID {
		t.Errorf("got %d, want %d", got.TweetID, expectedTweetID)
	}

	if got.ScreenName != expectedScreenName {
		t.Errorf("got %s, want %s", got.ScreenName, expectedScreenName)
	}

	if got.ImageURL != expectedImageURL {
		t.Errorf("got %s, want %s", got.ImageURL, expectedImageURL)
	}
}

func TestTwitterParsingImageUrlInEntitiesAndExtendedEntities(t *testing.T) {

	const expectedTweetID = 1126218073604796417
	const expectedScreenName = "jamesnmullen"
	const expectedImageURL = "http://example.com/test2.jpg"

	got := ParseTweetData(jsonDataWithBothEntitiesAndExtendedEntities)

	if got.TweetID != expectedTweetID {
		t.Errorf("got %d, want %d", got.TweetID, expectedTweetID)
	}

	if got.ScreenName != expectedScreenName {
		t.Errorf("got %s, want %s", got.ScreenName, expectedScreenName)
	}

	if got.ImageURL != expectedImageURL {
		t.Errorf("got %s, want %s", got.ImageURL, expectedImageURL)
	}
}
