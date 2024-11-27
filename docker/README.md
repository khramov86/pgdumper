# Тестовые данные

```sql
CREATE DATABASE db1;
CREATE DATABASE db2;
CREATE DATABASE db3;
```


```sh
\c db1
```
```sql
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    position VARCHAR(100),
    salary NUMERIC
);

INSERT INTO employees (name, position, salary) VALUES
('Alice', 'Engineer', 70000),
('Bob', 'Manager', 85000),
('Charlie', 'Analyst', 65000);
```

```sh
\c db2
```
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price NUMERIC,
    quantity INTEGER
);

INSERT INTO products (name, price, quantity) VALUES
('Laptop', 1200, 50),
('Smartphone', 800, 100),
('Tablet', 300, 75);
```

```sh
\c db3
```
```sql
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    phone VARCHAR(20)
);

INSERT INTO customers (name, email, phone) VALUES
('John Doe', 'john.doe@example.com', '555-1234'),
('Jane Smith', 'jane.smith@example.com', '555-5678'),
('Mike Johnson', 'mike.johnson@example.com', '555-9101');
```

