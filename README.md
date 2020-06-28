# ed25519-client

A basic ed25519 signature verification app.
For detailed info see: https://github.com/spacemeshos/ed25519

## Building
```bash
go get
go make
```

## Running
1. Create a json file with the content of the copied signature from smapp.
For example:

```json
{
  "text": "Aviv Eyal avive@spacemesh.io",
  "signature": "0xb71ae47c5df50e8979996ebf0a4c8e9d6f8e60580a44859093f7bf5649bc32ba6d73c990d263978d79a57e13cf46606352cc285bb718c23aa75a6339ef45720e",
  "publicKey": "0x4cc083df81a82e1ef7d22b72614703fbcb451f7902a6c865bfb3309e9f311556"
}
```

2. Save it anywhere on your system.

3. Run the app with the first argument specifying the pass to the json file.
For example, for signed_msg.json in the same directory as the app:

```bash
ed25519-client signed_msg.json
```

4. Check the signature validity on the message in the app's console output. e.g.:
```bash
Valid signature. Message: Aviv Eyal avive@spacemesh.io. Account: 0x614703fbcb451f7902a6c865bfb3309e9f311556
```
