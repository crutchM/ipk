create table chairs
(
    id   serial
        constraint chairs_pk
            primary key,
    name varchar(255)
);

create table post
(
    id   serial
        constraint post_pk
            primary key,
    name varchar(255)
);

create table users
(
    id       varchar(255)
        constraint users_pk primary key,
    fullname varchar(255),
    login    varchar(255),
    chair    integer references chairs,
    post     integer references post,
    password varchar(255)
);



create table test
(
    id   serial
        constraint test_pk
            primary key,
    name varchar(255)
);
create table block
(
    id   serial
        constraint block_pk
            primary key,
    name varchar(255)
);

create table testBlocks
(
    id       serial
        constraint tb_pk
            primary key,
    test_id  integer references test,
    block_id integer references block

);

create table question
(
    id   serial
        constraint question_pk
            primary key,
    number varchar(255),
    text varchar(255),
    answer integer
);


create table blockQuestions
(
    id          serial
        constraint bq_pk
            primary key,
    block_id    integer references block,
    question_id integer references question
);



create table expert
(
    id   serial
        constraint e_pk primary key,
    name varchar(255)
);

create table lessontype
(
    id   serial
        constraint lt_pk primary key,
    name varchar(255)
);



create table stat
(
    id         serial
        constraint s_pk primary key,
    userI      varchar(255) references users,
    post       integer references post,
    chair      integer references chairs,
    employment integer references lessontype,
    Expert     integer references expert,
    lessonDate date,
    anketDate  date
);


create table results
(
    id       serial
        constraint r_pk primary key,
    test     integer references stat,
    block    integer references block,
    question integer references question,
    answer   int
);




