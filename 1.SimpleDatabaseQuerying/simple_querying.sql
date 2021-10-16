select user1.id, user1.username, user2.username
from user user1
left join public.user user2
on user1.parent = user2.id;