# Create table If Not Exists Employee (id int, name varchar(255), salary int, managerId int)

select e2.name as Employee
from employee e1
join employee e2 on e1.id = e2.managerId
where e1.salary < e2.salary