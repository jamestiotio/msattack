#!/usr/bin/python3
# Created by James Raphael Tiovalen (2022)
# Usage: `mitmdump --quiet -nr <saved-flow-dump-file> -s scripts/list_urls.py`

import json
import time

from mitmproxy import ctx, http


class Parser:
    def response(self, flow: http.HTTPFlow) -> None:
        print(f"{flow.request.pretty_url}\n")


addons = [Parser()]


if __name__ == "__main__":
    print("This script is not meant to be run directly.")
    print("Please use mitmproxy to run this script.")
    exit(1)
