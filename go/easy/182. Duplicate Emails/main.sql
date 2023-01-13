# sol2
select p.email
from person p
group by p.email
having count(p.email) > 1

# sol2
select distinct(p1.email)
from person p1
join person p2 on p1.id <> p2.id and p1.email = p2.email