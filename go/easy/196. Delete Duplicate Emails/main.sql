delete p1.*
from Person p1
join Person p2 on p1.id <> p2.id and p1.Email = p2.email
where p1.id > p2.id