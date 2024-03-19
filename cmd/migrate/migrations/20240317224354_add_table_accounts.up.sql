create table if not exists accounts(
  id serial primary key,
	user_id uuid ,
	user_name text not null, 
  user_email text not null,
	created_at timestamp not null default now()
)