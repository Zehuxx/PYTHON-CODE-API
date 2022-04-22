# Python code API

API to query python programs represented by nodes using [Drawflow library](https://github.com/jerosoler/Drawflow) and handled by [Dgraph](https://dgraph.io/).

## Run Dgraph

  docker run --rm -it -p 8000:8000 -p 8080:8080 -p 9080:9080 dgraph/standalone:v21.12.0
  
## Run project  
  
  go run main.go

## Endpoints

### Get Programs

  Return list of programs data.

* **URL**

  `GET` /api/programs/

* **Success Response:**

  * **Code:** 200 <br />
  * **Content:**   `{
                        "programs": [
                              {
                                  "programName": "program name",
                                  "uid": "0x1"
                              },
                              {
                                  "programName": "program name 2",
                                  "uid": "0x96"
                              }
                          ]
                      }`

 
* **Error Response:**

  * **Code:** 500 <br />
  * **Content:** `{ "Msg" : "Failed to get programs." }`

### Get Program By Uid

  Return program data filter by Uid.

* **URL**

  `GET` /api/program/{uid}/

* **Success Response:**

  * **Code:** 200 <br />
  * **Content:**   `{
    "programs": [
        {
            "programName": "print Hi!!! program.",
            "nodes": [
                {
                    "id": 1,
                    "name": "root",
                    "data": {},
                    "class": "node-root",
                    "html": "<div>\n\t\t\t\t\t\t\t\t\t <span>Root node<span/>\n\t\t\t\t\t\t\t\t </div>",
                    "typenode": false,
                    "inputs": {
                        "input_1": {
                            "connections": [
                                {
                                    "order": 1,
                                    "node": "6",
                                    "input": "output_1"
                                }
                            ]
                        }
                    },
                    "outputs": {},
                    "pos_x": 1113,
                    "pos_y": 396
                },
                {
                    "id": 6,
                    "name": "print",
                    "data": {
                        "msg": "Hi!!!"
                    },
                    "class": "node-print",
                    "html": "<div>\n <span>Print node<span/>\n <input type=\"text\" style=\"width:100%;\" df-msg>\n </div>",
                    "typenode": false,
                    "inputs": {},
                    "outputs": {
                        "output_1": {
                            "connections": [
                                {
                                    "node": "1",
                                    "output": "input_1"
                                }
                            ]
                        }
                    },
                    "pos_x": 820,
                    "pos_y": 432
                }
            ]
        }
    ]
}`

 
* **Error Response:**

  * **Code:** 400 <br />
  * **Content:** `{"Msg": "uid is required."}`

    OR
    
  * **Code:** 500 <br />
  * **Content:** `{"Msg": "Failed to get program."}`

### Save Program

  Save program data.

* **URL**

  `POST` /api/programs/

* **Data**
  
  * **Basic example of a program that does a print "Hi!!!".**

  * **Body:** `{
                    "programName": "print Hi!!! program.",
                    "nodes": [
                        {
                            "id": 1,
                            "name": "root",
                            "data": {},
                            "class": "node-root",
                            "html": "<div>\n\t\t\t\t\t\t\t\t\t  <span>Root node<span/>\n\t\t\t\t\t\t\t\t  </div>",
                            "typenode": false,
                            "inputs": {
                                "input_1": {
                                    "connections": [
                                        {
                                            "node": "6",
                                            "input": "output_1"
                                        }
                                    ]
                                }
                            },
                            "outputs": {},
                            "pos_x": 1113,
                            "pos_y": 396
                        },
                        {
                            "id": 6,
                            "name": "print",
                            "data": {
                                "msg": "Hi!!!"
                            },
                            "class": "node-print",
                            "html": "<div>\n                        <span>Print node<span/>\n                        <input type=\"text\" style=\"width:100%;\" df-msg>\n                    </div>",
                            "typenode": false,
                            "inputs": {},
                            "outputs": {
                                "output_1": {
                                    "connections": [
                                        {
                                            "node": "1",
                                            "output": "input_1"
                                        }
                                    ]
                                }
                            },
                            "pos_x": 820,
                            "pos_y": 432
                        }
                    ],
                      "uid": "_:program"
                    }`

* **Success Response:**

  * **Code:** 201 <br />

 
* **Error Response:**

  * **Code:** 500 <br />
  * **Content:** `{ Msg : "Failed to save program." }`
    
### Update Program

  update program data.

* **URL**

  `PUT` /api/programs/{uid}/

* **Data**
  
  * **Basic example of how to update a program.**

  * **Body:** `{
                "programName": "print Hi!!! program update.",
                "nodes": [
                    {
                        "id": 1,
                        "name": "root",
                        "data": {},
                        "class": "node-root",
                        "html": "<div>\n\t\t\t\t\t\t\t\t\t  <span>Root node<span/>\n\t\t\t\t\t\t\t\t  </div>",
                        "typenode": false,
                        "inputs": {
                            "input_1": {
                                "connections": [
                                    {
                                        "node": "6",
                                        "input": "output_1"
                                    }
                                ]
                            }
                        },
                        "outputs": {},
                        "pos_x": 1113,
                        "pos_y": 396
                    },
                    {
                        "id": 6,
                        "name": "print",
                        "data": {
                            "msg": "Hi!!! updated program"
                        },
                        "class": "node-print",
                        "html": "<div>\n                        <span>Print node<span/>\n                        <input type=\"text\" style=\"width:100%;\" df-msg>\n                    </div>",
                        "typenode": false,
                        "inputs": {},
                        "outputs": {
                            "output_1": {
                                "connections": [
                                    {
                                        "node": "1",
                                        "output": "input_1"
                                    }
                                ]
                            }
                        },
                        "pos_x": 820,
                        "pos_y": 432
                    }
                ],
                  "uid": "{uid}"
                }`

* **Success Response:**

  * **Code:** 204 <br />

 
* **Error Response:**

  * **Code:** 500 <br />
  * **Content:** `{ Msg : "Failed to update program." }`
    
    
 ### Delete Program

  Delete program data.

* **URL**

  `DELETE` /api/programs/{uid}/

* **Success Response:**

  * **Code:** 204 <br />

 
* **Error Response:**

  * **Code:** 500 <br />
  * **Content:** `{ Msg : "Failed to delete program." }`
