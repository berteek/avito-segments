# avito-segments

[![Lint and Test](https://github.com/berteek/avito-segments/actions/workflows/lint_and_test.yml/badge.svg)](https://github.com/berteek/avito-segments/actions/workflows/lint_and_test.yml)

# Deployment
Clone repository. Run `make dub`. This will build a new image with docker-compose and run two containers. Run `make du` to run already built image.

# Database tables
Queries for creating tables are inside `sql/create_db.sql` file.

# Usage
### Create user: 
- url: localhost:8080/create_user
- body: **empty**

### Create segment:
- url: localhost:8080/create_seg
- body:
  
  ```
  {
    "slug": "AVITO_TEST_SEGMENT_TEST"
  }
  ```

### Delete segment:
- url: localhost:8080/delete_seg
- body:
  
  ```
  {
    "slug": "AVITO_TEST_SEGMENT_TEST"
  }
  ```

### Add and delete segments from user:
- url: localhost:8080/add_del_seg_from_user
- body:
  
  ```
  {
    "user_id": 1,
    "segments_to_add": [
        "AVITO_TEST_SEGMENT_ADD_0",
        "AVITO_TEST_SEGMENT_ADD_1",
        "AVITO_TEST_SEGMENT_ADD_2"
    ],
    "segments_to_delete": [
        "AVITO_TEST_SEGMENT_DEL_0",
        "AVITO_TEST_SEGMENT_DEL_1",
        "AVITO_TEST_SEGMENT_DEL_2"
    ]
  }
  ```

### Active segments for user:
- url: localhost:8080/active_seg_for_user
- body:
  
  ```
  {
    "user_id": 1
  }
  ```
