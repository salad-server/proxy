# proxy
![go-version](https://img.shields.io/github/go-mod/go-version/salad-server/proxy) ![report-card](https://goreportcard.com/badge/github.com/salad-server/proxy) ![last-commit](https://img.shields.io/github/last-commit/salad-server/proxy)

Custom osu.ppy.sh proxy that allows you access certain routes that would require an account (such as osu!direct).

## Requirements
- **Go 1.18** is required.
- **make and git** are also required, but on most systems these are installed by default.
- **[upx](https://upx.github.io/)** is optional, but recommended as it will compress the proxy.

## Install
Clone the source:
```sh
$ git clone https://github.com/salad-server/proxy.git
$ cd proxy
```

Build the code and modify config:
```sh
$ make install
```

Configure your nginx:
```sh
$ sudo cp ext/nginx.conf /etc/nginx/conf.d/proxy.conf
```

Finally, run the proxy with:
```sh
$ ./proxy
```

You can use any process manager or something like [tmux](https://github.com/tmux/tmux) if you want to keep the proxy running in the background.

## osu!direct
If you want to use this proxy for osu!direct, you'll need to make a few changes to your nginx configs. This can vary between servers, but you need to redirect `/d/` and `/web/osu-search.php` to the proxy. Here's how mine looks:

```
server {
	listen 80;
	listen 443 ssl;
	server_name ~^(?:c[e4]?|osu|a|api)\.servername\.com$;

	... # location/ssl stuff

	# proxy inject
	location /d/ {
		rewrite ^(/d)(.*)$ https://proxy.servername.com/d/$2? permanent;
	}

	location /web/osu-search.php {
		rewrite (.*)$ https://proxy.servername.com/web/osu-search.php redirect;
	}
}
```

Assuming you have the proxy running on `proxy.servername.com`. If you had any issues setting this up you're welcome to open an issue. I'll get to it as soon as I can.
