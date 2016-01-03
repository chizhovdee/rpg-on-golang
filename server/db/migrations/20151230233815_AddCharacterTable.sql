-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

create table characters (
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
  ep_updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
  hp_updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
  created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table characters;
