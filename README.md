```sql
CREATE KEYSPACE api WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'eu-west' : 1 };

CREATE TABLE supermarket (
  id UUID,
  name text,
  quantity int,
  PRIMARY KEY (id)
);
```

```shell
curl -X POST \
-H "Content-Type: application/x-www-form-urlencoded" \
-d 'name=Cheese&&quantity=3' \
"http://localhost:8080/items/new"
```