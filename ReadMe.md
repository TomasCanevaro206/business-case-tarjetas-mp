Postman:

GET localhost:8080/api/v1/cards/ -> (get all)

GET localhost:8080/api/v1/cards/1 -> (get by id)

POST localhost:8080/api/v1/cards/ -> (create)
          body:
        {
            "card_id": 6,
            "card_number": 987746,
            "card_type": "debit",
            "expiration_date": "4/8/2027",
            "card_state": "created",
            "timestamp_creation": "4/8/2021",
            "timestamp_modificaction": "17/6/2023"
        }