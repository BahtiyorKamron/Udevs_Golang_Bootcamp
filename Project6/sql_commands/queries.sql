SELECT *
FROM users
WHERE password=crypt('mypassword18',password)



SELECT 
   b.id,
   b.title,
   b.created_at,
   b.body,
   u.name,
   u.lastname
FROM blogs b 
LEFT JOIN users u ON b.user_id=u.id



with old_data as (
                select
                    id,
                    username,
                    password,
                    name,
                    lastname
                from 
                    users
                where id = $1 
            )update users as u set
                username = (
                    case
                        when length($2) > 1 then $2
                        else o.username
                    end
                ),
                password = (
                    case
                        when length($3) > 1 then $3
                        else o.password
                    end
                ),
                name = (
                  case
                       when length($4) > 1 then $4
                       else o.name
                  end
                ),
                lastname = (
                   case 
                       when length($5) > 1 then $5
                       else o.lastname 
                )
            from old_data as o
            where u.id = $1
            returning u.*    

select * from users where username='james' and password=crypt('j@mshid2106', password)
