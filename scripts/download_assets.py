#!/usr/bin/python3
# Created by James Raphael Tiovalen (2022)
# Usage: `mitmdump --quiet -nr <saved-flow-dump-file> -s scripts/download_assets.py`

import json
import os
import shutil
import time
from urllib import parse

import requests
from mitmproxy import ctx, http

MASTER_TABLE_FILENAME = "master_table.json"
MASTER_TABLE_URL = "msattack.snkplaymore.info/title/get_master_table"
STORAGE_URL = "strage.snkplaymore.info/snkp/msatk/prod"
STORAGE_PACK_URL = f"{STORAGE_URL}/pack"


class AssetDownloader:
    def response(self, flow: http.HTTPFlow) -> None:
        if STORAGE_PACK_URL in flow.request.pretty_url:
            filename = flow.request.pretty_url.split("/")[-1]
            version_number = flow.request.pretty_url.split("/")[-2]
            print(f"Downloading file {filename}...")
            # Handle arbitrary file size
            with requests.get(flow.request.pretty_url, stream=True) as r:
                r.raise_for_status()
                os.makedirs(f"data/pack/{version_number}", exist_ok=True)
                filepath = f"data/pack/{version_number}/{filename}"
                try:
                    with open(filepath, "xb+") as f:
                        print(f"Writing content of {filename}...")
                        shutil.copyfileobj(r.raw, f)
                except FileExistsError:
                    print(f"File {filename} already exists. Skipping...")

        elif STORAGE_URL in flow.request.pretty_url:
            filename = flow.request.pretty_url.split("/")[-1]
            version_number = flow.request.pretty_url.split("/")[-2]
            print(f"Downloading file {filename}...")
            # Handle arbitrary file size
            with requests.get(flow.request.pretty_url, stream=True) as r:
                r.raise_for_status()
                os.makedirs(f"data/{version_number}", exist_ok=True)
                filepath = f"data/{version_number}/{filename}"
                try:
                    with open(filepath, "xb+") as f:
                        print(f"Writing content of {filename}...")
                        shutil.copyfileobj(r.raw, f)
                except FileExistsError:
                    print(f"File {filename} already exists. Skipping...")

        elif MASTER_TABLE_URL in flow.request.pretty_url:
            table_key_values = parse.parse_qs(
                parse.urlparse(flow.request.pretty_url, allow_fragments=False).query
            )["table[]"]

            with open(f"data/{MASTER_TABLE_FILENAME}", "at+") as f:
                f.seek(0)
                try:
                    master_table = json.load(f)
                except json.decoder.JSONDecodeError:
                    master_table = {"table": []}

            master_table_values = master_table["table"]

            for table_key_value in table_key_values:
                # Sometimes the value has slashes for some reason, such as "event/.../..."
                table_key_actual_value = table_key_value.split("/")[-1]
                print(f"Getting master table value for {table_key_actual_value}...")
                with requests.get(
                    f"https://{MASTER_TABLE_URL}?table[]={table_key_value}"
                ) as r:
                    r.raise_for_status()
                    json_data = r.json()["table"][0]
                    if not any(
                        table_key_actual_value in pair for pair in master_table_values
                    ):
                        print(f"Adding {table_key_actual_value} to master table...")
                        master_table_values.append(json_data)
                    else:
                        print(f"Skipping {table_key_actual_value}...")

            with open(f"data/{MASTER_TABLE_FILENAME}", "wt+") as f:
                json.dump(master_table, f, indent=2)


addons = [AssetDownloader()]


if __name__ == "__main__":
    print("This script is not meant to be run directly.")
    print("Please use mitmproxy to run this script.")
    exit(1)
