# Ben And Jerry's Ice-cream

## About

- This project was developed using Go and gRPC. 
- The goal of the project was to develop APIs for Ben And Jerry's Ice-cream.
- It comprises of 3 sub-systems : Ice-cream, User and Authorization

### Authorization Sub-system

- Authorization sub-system provides the following functionalities:
  - Login User
  - Verify User 
- Login user is used to login the user and generate a JWT Token.
- Verify User is used to verify the user using the JWT Token.

### User Sub-system:

- User sub-system provides the following functionalities:
  - Create User
  - Update User 
  - Delete User using Email Id
- Users cannot update or delete other users data.
- Authentication is required to perform update and delete operations.

### Ice-cream Sub-system:
- Ice-cream sub-system provides the following functionalities:
    - Create Icecream
    - Update Icecream
    - Delete Icecream using Product Id
    - Get Icecream using Product Id
    - Pagination on Icecream data
- Authentication is required to perform all the above operations.


## Project Structure
- The Project Structure uses the structure recommended by https://github.com/golang-standards/project-layout
- Each Sub-system has the following structure

```
                                         ┌──→ DB
Controller ──→ Service ──→ DataService ──├──→ Cache
                                         └──→ SearchEngine          
```
- DB can be used write to a persistant storage.

- Cache can be used to cache data for faster retrieval.

- SearchEngine can be used to filter and search data.

- For this project only DB has been used.

- Cache and SearchEngine can be integrated as and when required.


- The `internal` directory structure has been utilized to abstract and hide the DataServices from other Services.

- The DataService can only be accessed by its respective service; for example only UserService can access the UserDataService.

- This is done to avoid direct access of the DataServices by any other Service.

- If any sub-system needs data from any other sub-system it must use its Service.

- DB, Cache and SearchEngine are only accessible to the respective DataService.

- This is done to avoid direct access of DB, Cache or SearchEngine by the Service.


- All .proto files are available in `api/proto`.

- To .pb.go files are generated using the script at `/tool/generate-internal-proto.sh`.


## Implementation Details

**Database**
- MongoDB is used for back-end. 
- The ice-cream data is going to be read heavy.
- Create and update operations on ice-cream data will be very rare.  

