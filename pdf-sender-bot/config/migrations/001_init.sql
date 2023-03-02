-- create database pdf_bot_db;

create table if not exists users (
  user_id                   bigint not null PRIMARY KEY,
  firstName                 text   not null,
  secondName                text   not null,
  thirdName                 text   not null,
  organization              text   not null,
  address                   text   not null,
  phone                     text   not null,
  OGRN                      text   not null,
  vehicleModel              text   not null,
  stateLicensePlate         text   not null,
  IDNumber                  text   not null,
  licenseRegistrationNumber text   not null,
  licenseSerial             text   not null,
  licenseNumber             text   not null,
  garageNumber              text   not null,
  personnelNumber           text   not null
);