CREATE TABLE IF NOT EXISTS users(
  id serial not null primary key,
  name varchar(128) not null ,
  lastname varchar(128) not null,
  username varchar(64) unique,
  password text not null
 );
 COMMENT ON TABLE users is 'foydalanuvchiu malumotlari saqlanadigan table'
 
 
 CREATE EXTENSION pgcrypto;
 
 
 INSERT INTO users(
    name,lastname,username,password
 ) VALUES (
   'Kamron',
   'Bahtiyorov',
   'kamron1804',
   crypt('mypassword18', gen_salt('bf', 8))
 );
  INSERT INTO users(
    name,lastname,username,password
 ) VALUES (
   'Jamshid',
   'Sobirov',
   'james',
   crypt('j@mshid2106', gen_salt('bf', 8))
 );
  INSERT INTO users(
    name,lastname,username,password
 ) VALUES (
   'Umidjon',
   'Pirmanov',
   'umid',
   crypt('umid2701', gen_salt('bf', 8))
 );
 
 CREATE TABLE IF NOT EXISTS blogs(
  id serial not null,
  user_id int not null references users(id),
  title text not null,
  body text not null,
  created_at timestamp,
  updated_at timestamp,
  status boolean not null
 );
   COMMENT ON TABLE blogs is 'Har bir foydalanuvchining shahsiy blog maqolalari saqlanadigan table'
   

 
 
 INSERT INTO blogs(
   user_id,title,body,created_at,status
 )VALUES(
  3,
  'G''azalkentga borganim haqida',
  'Not all PostgreSQL installations has the plpqsql language by default, this means you may have to call CREATE LANGUAGE plpgsql before creating the function, and afterwards have to remove the language again, to leave the database in the same state as it was before (but only if the database did not have the plpgsql language to begin with). See how the complexity grows?',
  Now(),
  true
 )
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
   
   
   
   
   
 
