import requests
import json

resp = requests.post(
    "https://gql.twitch.tv/gql",
    json.dumps(
        {
            "operationName": "BrowsePage_AllDirectories",
            "variables": {
                "limit": 30,
                "directoryFilters": ["GAMES"],
                "isTagsExperiment": True,
                "tags": [],
            },
            "extensions": {
                "persistedQuery": {
                    "version": 1,
                    "sha256Hash": "75fb8eaa6e61d995a4d679dcb78b0d5e485778d1384a6232cba301418923d6b7",
                }
            },
        }
    ),
    headers={"Client-Id": "kimne78kx3ncx6brgo4mv6wki5h1ko"},
)

print()
edges = json.loads(resp.content)["data"]["directoriesWithTags"]["edges"]
games = [f["node"] for f in edges]
print(games)
# # games:
# [
#     ("Let us GAME", 78250),
#     ("(REBROADCAST) Worlds Play-In Knockouts: Cloud9 vs. Gambit Esports", 36783),
#     ("RuneFest 2018 - OSRS Reveals !schedule", 35042),
#     (None, 25237),
#     ("Front Page of TWITCH + Fortnite FALL SKIRMISH Training!", 22380),
#     ("Reckful - 3v3 with barry and a german", 20399),
# ]
