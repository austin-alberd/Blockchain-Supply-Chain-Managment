# **API Documentation**
## Endpoints
```
/api/v1/
```

## Functions

### `POST` `/api/v1/createItem`
Creates a new item in the database
### Sample Request
``` JSON
{
    "id": "string",
    "name": "string",
    "description": "string",
    "quantity": "number"
}
```
### Responses
| HTTP Code | JSON Response | Meaning | 
| :--      | :----:        | --:      |
| ` 201 ` | `{"msg":"Successfully Created Object} ` | The Item was added successfully.|
| `400` | `{"msg: "Bad Request Object Not Created"}` | The object was not created due to an error in the request.|
| `500` | `{"msg":"Server Error Could Not Create Object", "errMSG": error message}` | The object was not created due to an error on the server.|
