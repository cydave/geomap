# Geomap

Pipe a list of IP addresses into the binary to plot them on top of a card:

```bash
curl -s 'https://blocklists.0dave.ch/ssh.txt' | ./bin/geomap  ~/Downloads/GeoIP2-City.mmdb >/tmp/out.html && sleep 1 && firefox /tmp/out.html
```

You'll need a copy of MaxMind's GeoIP database.
