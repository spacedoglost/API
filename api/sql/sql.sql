CREATE DATABASE IF NOT EXISTS anoteaqui;
USE anoteaqui;

DROP TABLE IF EXISTS usuarios

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null unique,
    CriadoEm timestamp default current_timestamp()
) ENGINE=INNODB;
