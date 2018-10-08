import requests
import json

resp = requests.post(
    "https://gql.twitch.tv/gql",
    json.dumps(
        {
            "operationName": "AnonFrontPage_TopChannels",
            "variables": {"platformType": "all", "isTagsExperiment": True},
            "extensions": {
                "persistedQuery": {
                    "version": 1,
                    "sha256Hash": "d94b2fd8ad1d2c2ea82c187d65ebf3810144b4436fbf2a1dc3af0983d9bd69e9",
                }
            },
        }
    ),
    headers={"Client-Id": "kimne78kx3ncx6brgo4mv6wki5h1ko"},
)

edges = json.loads(resp.content)["data"]["streams"]["edges"]
games = [(f["node"]["title"], f["node"]["viewersCount"]) for f in edges]

# games:
[
    ("Let us GAME", 78250),
    ("(REBROADCAST) Worlds Play-In Knockouts: Cloud9 vs. Gambit Esports", 36783),
    ("RuneFest 2018 - OSRS Reveals !schedule", 35042),
    (None, 25237),
    ("Front Page of TWITCH + Fortnite FALL SKIRMISH Training!", 22380),
    ("Reckful - 3v3 with barry and a german", 20399),
]
