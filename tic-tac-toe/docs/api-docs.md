1. POST /games
    request:
    {
        "player1_name":
        "player2_name":
        "board_size"
    }

    response:
    {
        "status": 
        "message":
        "game_id": 
    }

2. POST /games/{game_id}/moves
    request:
    {
        "x":
        "y":
    }

    response:
    {
        "status":
        "message":
    }

3. GET /games/{game_id}
    request:
    {
        EMPTY
    }

    response:
    {
        "status":
        "message":
        "game": {
            "game_id":
            "players": [
                {
                    "name":
                    "symbol":
                },
                {
                    "name":
                    "symbol":
                }
            ],
            "board": [
                [],
                [],
                []
            ],
            "active_player":
            "game_state":
        }
    }