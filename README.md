# msattack

[Metal Slug Attack](https://game.snk-corp.co.jp/official/metalslug_attack/) Server Reimplementation

The purpose of this software is to preserve and archive a working version of Metal Slug Attack's backend server to prevent it from being lost and forgotten, as well as to retain the ability to usably play Metal Slug Attack which serves primarily to validate the accuracy of the implementation.

## Background Information

On October 12, 2022, SNK CORPORATION [announced](https://www.snk-corp.co.jp/us/press/2022/101201/) that this game's end of service is at January 12, 2023, 16:00 JST. Since the game version update on October 12, 2022, the game can no longer produce any more revenue for its developer/publisher, SNK CORPORATION, as medals will no longer be available for purchase. Furthermore, after January 12, 2023, users will have no way of playing the game since, presumably, all the servers would be shut down and taken offline. Hence, this game is considered as abandonware. This project was humbly started to offer an archived implementation of the backend server so as to keep the game available and accessible to everyone.

## Features

Coming soon!

## Instructions

1. Install [Git](https://git-scm.com/) and [Go](https://go.dev/) if you do not have them installed already.

2. Clone this Git repository to your local workspace:

   ```bash
   git clone https://github.com/jamestiotio/msattack
   ```

   If you have already cloned an older version of the Git repository, you can simply run the following command from the root of the repository to get the latest version:

   ```bash
   git pull
   ```

3. Connect your phone or your emulator to send Internet traffic via your PC.

4. Redirect traffic by installing and using your favorite proxy daemon ([mitmproxy](https://mitmproxy.org/), [Fiddler Classic](https://www.telerik.com/fiddler), [Burp Suite](https://portswigger.net/burp), [Wireshark](https://www.wireshark.org/), [Charles Proxy](https://www.charlesproxy.com/), etc.) Ensure that the certificate of your selected proxy software is installed as a trusted system-level root CA certificate (which might require rooting on Android). Since Metal Slug Attack implements [Certificate Pinning](https://en.wikipedia.org/wiki/HTTP_Public_Key_Pinning), you would also need to modify the APK/IPA respectively to disable Certificate Pinning and allow for HTTPS MITM (which might require jailbreaking on iOS). You can check [this section](https://docs.mitmproxy.org/stable/concepts-certificates/#certificate-pinning) to find a list of various tools that can be used to achieve this.

5. Once the proxy setup is done, you would need to obtain all of the required binary data. There are two main ways that you can achieve this:

   - If the original Metal Slug Attack servers are still up, you can do this by playing the game normally and intercepting all of the traffic between the Metal Slug Attack game and its server:

     ```bash
     mitmdump -k -w <saved-flow-dump-file> --set stream_large_bodies=1
     ```

     Then, you can download all of the required binary data by running the [`download_assets.py`](./scripts/download_assets.py) script from the root of this workspace:

     ```bash
     mitmdump --quiet -nr <saved-flow-dump-file> -s scripts/download_assets.py
     ```

     Ensure that nothing else on your machine is accessing the [`data`](./data/) directory (and any of its files/subdirectories) as the script would place all of the necessary files there. The script assumes that it is the only accessor of the `data` directory at the time it is being executed and any filesystem-related race conditions might mess up the final state.

     > While it might be possible to save the raw binary content of most of the files directly into the saved flow dump file without enabling streaming and process them later, it is highly improbable for this method to succeed for the master table. This is because the backend server is transmitting Base64-encoded binary data via JSON. Even with a very fast Internet speed via wired Ethernet, it is almost impossible to transmit all of the master table's contents back to the game process within the specified timeout requirements (since the bottleneck is likely to be somewhere else). Once a request has timed out, the game will attempt to retry the request. If there are 5 consecutive failed attempts for a particular request, the game will throw the [communication timeout error banner](./assets/timeout.jpg). Granted, this issue can be solved by increasing the timeout value in the game's APK/IPA or maybe even by physically/directly connecting to the Metal Slug Attack server, however ridiculous/impossible that sounds. Since both options require a lot of effort, it is much simpler/easier to just stream the responses first and re-request them later separately by using the [`requests`](https://requests.readthedocs.io/en/latest/) Python library/package.

   - If the original Metal Slug Attack servers are not up anymore, you would need to source this binary data from elsewhere. I am not including them in this repository due to 2 main reasons: to prevent copyright strikes or DMCA takedowns as the copyright for those binary assets are still owned by SNK CORPORATION, as well as to keep the size of this repository small and lean. I will keep a lookout on the possible ways to publicly serve this content without violating any copyright laws (feel free to inform me if you have any information about this by publicly raising an issue or privately emailing me).

6. You can now run the proxy service needed for the private server. If you are using `mitmproxy`, a script for `mitmproxy` has been provided [here](./scripts/proxy.py). Simply execute the following command:

   ```bash
   mitmdump -s scripts/proxy.py -k --set stream_large_bodies=1
   ```

7. Run the private server by executing this command from the terminal:

   ```bash
   go build -o msattack main.go
   chmod +x ./msattack
   ./msattack
   ```

   If you are on Windows, you would need to generate an executable file and double-click it to run it:

   ```bash
   go build -o msattack.exe main.go
   ```

8. Have fun and enjoy!

## Troubleshooting

### `exit status 259`

This error might be caused by your antivirus or/and firewall on a Windows machine. Ensure that your antivirus or/and firewall allow traffic to flow through the private server.

### Address already in use: bind

This error is derived from the server being unable to bind to a certain port. Ensure that there are no other processes using the same port numbers as the private server. If you are running on an operating system that restricts ports below `1024` to privileged users only (i.e., not Windows), choose a different port above `1024` for the private server and point the proxy server to that port.

## Legal Disclaimers

- METAL SLUG, METAL SLUG ATTACK and all related trademarks are the property of [SNK CORPORATION](https://www.snk-corp.co.jp/). This project is not affiliated, endorsed or supported in any way by [SNK CORPORATION](https://www.snk-corp.co.jp/). The use of information and software provided through this project may be used at your own risk. The information and software available through this project are provided as-is without any warranty or guarantee. By using this project you agree that: (1) We take no liability under any circumstance or legal theory for any software, error, omissions, loss of data or damage of any kind related to your use or exposure to any information provided through this project; (2) All software are made "AS AVAILABLE" and "AS IS" without any warranty or guarantee. All express and implied warranties are disclaimed. Some states do not allow limitations of incidental or consequential damages or on how long an implied warranty lasts, so the above may not apply to you.

- This project is non-commercial. The source code is available for free and always will be.

- This is a black-box re-implementation project. The code in this project was written by observing the game running and inspecting the behavior and data being transmitted between the game process and the backend server.

- If you want to contribute to this repository, your contribution must be either your own original code or open-source code with a clear acknowledgement of its origin. No code that was acquired through reverse engineering executable binaries or binary files will be accepted.

- No assets from the original game are included in this repository.
