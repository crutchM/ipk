create table chairs(
   id integer not null primary key,
   name varchar(255)
);

create table post(
 id integer not null primary key,
 name varchar(255)
);

create table users(
    id varchar(255) not null  primary key,
    fullname varchar(255),
    login varchar(255),
    chair integer references chairs,
    post integer references post,
    password varchar(255)
);



create table test(
    id integer not null primary key,
    name varchar(255)
);
create table block(
    id integer not null primary key,
    name varchar(255)
);

create table testBlocks(
    id integer not null primary key,
    test_id integer references test,
    block_id integer references block

);

create table question(
    id integer not null primary key,
    text varchar(255)
);

create table answer(
    id integer not null primary key,
    text varchar(255)
);

create table blockQuestions(
    id integer not null primary key,
    block_id integer references block,
    question_id integer references question
);

create table questionAnswers(
    id integer not null primary key,
    question_id integer references question,
    answer_id integer references answer
)


