-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

create TABLE characters (
  id serial primary key,
  level integer not null default 1,
  energy integer not null default 0,
  ep integer not null default 0,
  health integer not null default 0,
  hp integer not null default 0,
  experience integer not null default 0,
  points integer not null default 0,
  basic_money integer not null default 0,
  vip_money integer not null default 0,
  ep_updated_at timestamptz not null DEFAULT CURRENT_TIMESTAMP,
  hp_updated_at timestamptz not null DEFAULT CURRENT_TIMESTAMP,
  created_at timestamptz not null DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz not null DEFAULT CURRENT_TIMESTAMP
);


create TABLE character_states (
  character_id integer references characters(id),
  quests jsonb
);


-- +goose StatementBegin
create or replace FUNCTION insert_character_state_func()
RETURNS trigger AS $first_trigger$
begin
  insert into character_states (character_id) values (new.id);
  return new;
end;
$first_trigger$ language plpgsql;
-- +goose StatementEnd


create TRIGGER insert_character_state after insert on characters
for each row
execute procedure insert_character_state_func();




-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop trigger if exists insert_character_state on characters;
drop function if exists insert_character_state_func();
drop table if exists character_states;
drop table if exists characters;


