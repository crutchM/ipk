create table chairs(
   id   serial
       constraint chairs_pk
           primary key,
   name varchar(255)
);

create table post(
                     id   serial
                         constraint post_pk
                             primary key,
 name varchar(255)
);

create table users(
                      id   serial
                          constraint users_pk
                              primary key,
    fullname varchar(255),
    login varchar(255),
    chair integer references chairs,
    post integer references post,
    password varchar(255)
);



create table test(
                     id   serial
                         constraint test_pk
                             primary key,
    name varchar(255)
);
create table block(
                      id   serial
                          constraint block_pk
                              primary key,
    name varchar(255)
);

create table testBlocks(
                           id   serial
                               constraint tb_pk
                                   primary key,
    test_id integer references test,
    block_id integer references block

);

create table question(
                         id   serial
                             constraint question_pk
                                 primary key,
    text varchar(255)
);

create table answer(
                       id   serial
                           constraint answer_pk
                               primary key,
    text varchar(255)
);

create table blockQuestions(
                               id   serial
                                   constraint bq_pk
                                       primary key,
    block_id integer references block,
    question_id integer references question
);

create table questionAnswers(
                                id   serial
                                    constraint qa_pk
                                        primary key,
    question_id integer references question,
    answer_id integer references answer
)


