# SmartTriageCore

Provide the core component of the SmartTriage solution

## Compile on Makefile
#
This will compile the application and send a binary to the /bin directory

    make all


## Setup Docker
#
This will configure the containers and provide all the components for the application to run

    make all-docker



## Endpoints
#

- Method: GET   Endpoint: /v1/hmv/questions  
    - Response example
    ```json
    [
        {
            "_id": "6239e08b9053bc5a82662fe3",
            "answer": "",
            "description": "Você sente algum desconforto no peito/tórax?",
            "id": "1",
            "typeanswer": 0
        },
        {
            "_id": "6239e08b9053bc5a82662fe4",
            "answer": "",
            "description": "Essa dor, irradia-se para outras áreas do corpo?",
            "id": "2",
            "typeanswer": 0
        },
        {
            "_id": "6239e08b9053bc5a82662fe5",
            "answer": "",
            "description": "A área de irradiação é entre a mandíbula e o umbigo?",
            "id": "3",
            "typeanswer": 0
        },
        {
            "_id": "6239e08b9053bc5a82662fe6",
            "answer": "",
            "description": "A dor irradia-se para algum dos braços?",
            "id": "4",
            "typeanswer": 0
        },
        {
            "_id": "6239e08b9053bc5a82662fe7",
            "answer": "",
            "description": "Essa dor parece com pontadas?",
            "id": "5",
            "typeanswer": 0
        },
        {
            "_id": "6239e08b9053bc5a82662fe8",
            "answer": "",
            "description": "Essa dor, parece um aperto, opressão, pressão ou queimação?",
            "id": "6",
            "typeanswer": 0
        },
        {
            "_id": "6239e08b9053bc5a82662fe9",
            "answer": "",
            "description": "Você já está com essa dor a mais de 20 minutos?",
            "id": "7",
            "typeanswer": 0
        },
        {
            "_id": "6239e08b9053bc5a82662fea",
            "answer": "",
            "description": "Antes da dor, você fez exercícios físicos, uma refeição pesada ou passou alguma forte emoção?",
            "id": "8",
            "typeanswer": 0
        },
        {
            "_id": "6239e08b9053bc5a82662feb",
            "answer": "",
            "description": "Você sente algum desses sintomas? Suor excessivo, falta de ar, vômito, palpitações, palidez?",
            "id": "9",
            "typeanswer": 0
        }
    ]
    ```  

- Method: GET   Endpoint: /v1/hmv/questions/{cpf}
    - if status code 200, you get this response:
    ```json
    {
        "_id": "6239e0969053bc5a82662fef",
        "answers": [
            {
                "answer": 1,
                "id": "1"
            },
            {
                "answer": 1,
                "id": "2"
            },
            {
                "answer": 0,
                "id": "3"
            },
            {
                "answer": 1,
                "id": "4"
            },
            {
                "answer": 0,
                "id": "5"
            },
            {
                "answer": 1,
                "id": "6"
            },
            {
                "answer": 1,
                "id": "7"
            },
            {
                "answer": 1,
                "id": "8"
            },
            {
                "answer": 0,
                "id": "9"
            }
        ],
        "cpf": "B1ZRiqLF34rniGvNXuMntt_EPI0IErSIzY4v",
        "salt": "3316c31349471111f5ad75aec6f4dcfe031b8cac195b66ccc00d00fc2f8cd9b6"
    }
    ```
- Method: POST  Endpoint: /v1/hmv/questions/{cpf}
    - Response example
    ```json
    {
        "qrcode": "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABn0lEQVR42uyYvfHrIBDEl3GgkBJcCqVBaZSiEhQq0LBv9pDkr+f0P+ZGlxn9Eh/L7QKuuuqqn61I1RprBibOyOeSL2AGcFuRCjfgTrZzyReQ2ARQfSBLO5ccApElqA81h9UzUMHNVO0VMAnvfUD6JvvRgXNGhW1a+nZ/GWIjAydXtMN3MqxfjGdoQHvbbmRF0F9PbJpUfNXDzwMAcgNizVqfgUDOqTRnQO+DfDNspmD8t1HDAyT1qZ9Ns0+zlQ2uAP2ykcNALkil2dKTql0AURLGtCRS212zqfrVWD0AvQmLzqsd28BFydYfoD6QRSFhlqqjstCz7B0AuFdznO6bOrwKCc/G6QKwUYyJOq/a92KZ9t1xhgdkrIEW8MxxSlcA/AEwCWea4xTBNb8GvvGByAqbUTapkHjUOhJwcNUin6JtP6Ef0X1s4LhnWTxQgmeftx/XycGBx4sWdc8qex/cAftTDztw1MNYPQGAGYouYtaAz5c9JwDCnmlXJYZAb8ChatzshtKPsDvgnFH7U499ebuQjg9cddVVf17/AgAA//8RZkON4G3J3QAAAABJRU5ErkJggg==",
        "short_id": "6239E0969"
    }
    ```
- Method: PUT   Endpoint: /v1/hmv/questions/{cpf}
    - Respose example
    ```json
    {
        "qrcode": "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABn0lEQVR42uyYvfHrIBDEl3GgkBJcCqVBaZSiEhQq0LBv9pDkr+f0P+ZGlxn9Eh/L7QKuuuqqn61I1RprBibOyOeSL2AGcFuRCjfgTrZzyReQ2ARQfSBLO5ccApElqA81h9UzUMHNVO0VMAnvfUD6JvvRgXNGhW1a+nZ/GWIjAydXtMN3MqxfjGdoQHvbbmRF0F9PbJpUfNXDzwMAcgNizVqfgUDOqTRnQO+DfDNspmD8t1HDAyT1qZ9Ns0+zlQ2uAP2ykcNALkil2dKTql0AURLGtCRS212zqfrVWD0AvQmLzqsd28BFydYfoD6QRSFhlqqjstCz7B0AuFdznO6bOrwKCc/G6QKwUYyJOq/a92KZ9t1xhgdkrIEW8MxxSlcA/AEwCWea4xTBNb8GvvGByAqbUTapkHjUOhJwcNUin6JtP6Ef0X1s4LhnWTxQgmeftx/XycGBx4sWdc8qex/cAftTDztw1MNYPQGAGYouYtaAz5c9JwDCnmlXJYZAb8ChatzshtKPsDvgnFH7U499ebuQjg9cddVVf17/AgAA//8RZkON4G3J3QAAAABJRU5ErkJggg==",
        "short_id": "6239E0969"
    }
    ```

- Method: POST  Endpoint: /v1/hmv/questions/{cpf}/confirm
    - Return only status code 200 on success

## Directories :
#
- controllers 
    - Keep the roles responding to each endpoint
- database
    - Maintain database connection logic
- middleware
    - Keep json communication type
- mocks
    - Mock of questions collections
- models
    - Keep entities of collections
- routes
    - Maintain endpoint routes
- services
    - Other services required for application
