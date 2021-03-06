create table type_test
(
    col_id              int primary key,
    col_bigint          bigint           default 0                     not null,
    col_bigint_n        bigint,
    col_binary          binary(255)      default 0                     not null,
    col_binary_n        binary(255),
    col_bit             bit              default 0                     not null,
    col_bit_n           bit,
    col_bool            bit              default 0                     not null,
    col_bool_n          bit,
    col_char            char(255)        default ''                    not null,
    col_char_n          char(255),
    col_date            date             default '1989-11-09'          not null,
    col_date_n          date,
    col_datetime        datetime         default '1989-11-09T00:00:00' not null,
    col_datetime_n      datetime,
    col_datetime2       datetime2        default '1989-11-09T00:00:00' not null,
    col_datetime2_n     datetime2,
    col_decimal         decimal          default 0                     not null,
    col_decimal_n       decimal,
    col_float           float            default 0                     not null,
    col_float_n         float,
    col_int             int              default 0                     not null,
    col_int_n           int,
    col_money           money            default 0                     not null,
    col_money_n         money,
    col_nchar           nchar(255)       default ''                    not null,
    col_nchar_n         nchar(255),
    col_nvarchar        nvarchar(255)    default ''                    not null,
    col_nvarchar_n      nvarchar(255),
    col_numeric         numeric          default 0                     not null,
    col_numeric_n       numeric,
    col_real            real             default 0                     not null,
    col_real_n          real,
    col_smalldatetime   smalldatetime    default '1989-11-09 00:00'    not null,
    col_smalldatetime_n smalldatetime,
    col_smallint        smallint         default 0                     not null,
    col_smallint_n      smallint,
    col_smallmoney      smallmoney       default 0                     not null,
    col_smallmoney_n    smallmoney,
    col_time            time             default '00:00:00'            not null,
    col_time_n          time,
    col_tinyint         tinyint          default '0'                   not null,
    col_tinyint_n       tinyint,
    col_uuid            uniqueidentifier default NEWID()               not null,
    col_uuid_n          uniqueidentifier,
    col_varbinary       varbinary(255)   default 0                     not null,
    col_varbinary_n     varbinary(255),
    col_varchar         varchar(255)     default ''                    not null,
    col_varchar_n       varchar(255)
)